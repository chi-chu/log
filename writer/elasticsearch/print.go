package elasticsearch

import (
	"context"
	"errors"
	"fmt"
	"github.com/chi-chu/log/entry"
	"github.com/olivere/elastic/v7"
	"time"
)

const (
	BULK_WAIT_TIME			= 5
	default_cache_size		= 64
)

type printer struct {
	client				*elastic.Client
	indexName			string
	rotateIndexName		string
	docSlice			[]map[string]string
	sendGorouStart		bool
}

func New(host []string, indexName string, option ...Option) (*printer, error) {
	if indexName == "" {
		return nil, errors.New("index name can`t be nil")
	}
	opt(option...)
	client , err := elastic.NewClient(elastic.SetURL(host...), elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	return &printer{client:client, indexName:indexName, docSlice:make([]map[string]string, 0, default_cache_size)}, nil
}

func (p *printer) Print(e *entry.Entry) {
	if defaultConfig.bulkInsert {
		obj := make(map[string]string)
		for k,v := range e.Data {
			obj[k] = v
		}
		p.docSlice = append(p.docSlice, obj)
	} else {
		//a := elastic.NewBulkIndexRequest().Index(p.rotateIndexName).Doc(e.Data)
		//rsp ,err := p.client.Bulk().Add(a).Do(context.Background())
		_ , err := p.client.Index().Index(p.rotateIndexName).BodyJson(e.Data).Do(context.Background())
		if err != nil {
			fmt.Println("[log] insert into es err :", err)
			return
		}
	}
}

func (p *printer) Rotate(b bool) error {
	oldIndex := p.rotateIndexName
	index := p.indexName
	if b {
		index = p.indexName + "_" + time.Now().Format("200601021504")
	}
	exist, err := p.client.IndexExists(index).Do(context.Background())
	if err != nil {
		return err
	}
	if !exist {
		result, err := p.client.CreateIndex(index).BodyJson(newMap()).Do(context.Background())
		if err != nil {
			return err
		}
		if !result.Acknowledged || !result.ShardsAcknowledged {
			fmt.Printf("[log] Elastic search create doc: %s Warning Acknowledged:%v, ShardsAcknowledged:%v\n",
				index, result.Acknowledged, result.ShardsAcknowledged)
		}
		if oldIndex == "" {
			oldIndex = index
		}
		logArr := p.docSlice
		p.docSlice = make([]map[string]string, 0, default_cache_size)
		if len(logArr) > 0 {
			req := elastic.NewBulkIndexRequest().Index(oldIndex)
			for _, log := range logArr {
				req = req.Doc(log)
			}
			rsp ,err := p.client.Bulk().Add(req).Do(context.Background())
			if err != nil {
				return err
			}
			if rsp.Errors {
				errString := "[log] bulk insert err:  \n"
				for _,item := range rsp.Items {
					if item == nil {
						continue
					}
					for _, r := range item {
						if r != nil {
							errString += r.Error.Reason + "  \n"
						}
					}
				}
				return errors.New(errString)
			}
		}
	}
	p.rotateIndexName = index
	if !p.sendGorouStart && defaultConfig.bulkInsert {
		go p.sendLog()
		p.sendGorouStart = true
	}
	return nil
}

func (p *printer) Exit() {
	//if bulk insert flush the cache
	p.sendOnce()
	p.client.Stop()
}

func (p *printer) sendLog() {
	for {
		time.Sleep(BULK_WAIT_TIME*time.Second)
		p.sendOnce()
	}
}

func (p *printer) sendOnce() {
	if len(p.docSlice) < 1 {
		return
	}
	//something wrong
	reqArr := make([]elastic.BulkableRequest, len(p.docSlice))
	for _, log := range p.docSlice {
		reqArr = append(reqArr, elastic.NewBulkIndexRequest().Index(p.rotateIndexName).Doc(log))
	}
	fmt.Println(reqArr)
	fmt.Println("  ")
	p.docSlice = make([]map[string]string, 0, default_cache_size)
	rsp ,err := elastic.NewBulkService(p.client).Add(reqArr...).Do(context.Background())
	if err != nil {
		fmt.Println("[log] bulk insert error ", err)
		return
	}
	if rsp.Errors {
		errString := "[log] bulk insert err:  \n"
		for _,item := range rsp.Items {
			if item == nil {
				continue
			}
			for _, r := range item {
				if r != nil {
					errString += r.Error.Reason + "  \n"
				}
			}
		}
		fmt.Println(errString)
	}
}