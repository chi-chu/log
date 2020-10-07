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
)

type printer struct {
	client				*elastic.Client
	indexName			string
	docSlice			[]*entry.Entry
	sendGorouStart		bool
}

func New(host []string, indexName string, option ...Option) (*printer, error) {
	if indexName == "" {
		return nil, errors.New("index name can`t be nil")
	}
	opt(option...)
	client , err := elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	return &printer{client:client, indexName:indexName, docSlice:make([]*entry.Entry, 0, 64)}, nil
}

func (p *printer) Print(e *entry.Entry) {
	if defaultConfig.bulkInsert {
		p.docSlice = append(p.docSlice, e)
	} else {
		a := elastic.NewBulkIndexRequest().Index(p.indexName).Doc(e.Data)
		rsp ,err := p.client.Bulk().Add(a).Do(context.Background())
		if err != nil {
			fmt.Println("[log] insert into es err :", err)
			return
		}
		fmt.Println(rsp)
	}
}

func (p *printer) Rotate(b bool) error {
	index := p.indexName
	if b {
		index = p.indexName + "_" + time.Now().Format("200601021504")
	}
	exist, err := p.client.IndexExists(index).Do(context.Background())
	if err != nil {
		return err
	}
	if !exist {
		logArr := p.docSlice
		p.docSlice = make([]*entry.Entry, 0, 64)
		req := elastic.NewBulkIndexRequest().Index(p.indexName)
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
						errString += r.Error.Error() + "  \n"
					}
				}
			}
			return errors.New(errString)
		}
		result, err := p.client.CreateIndex(index).BodyJson(newMap()).Do(context.Background())
		if err != nil {
			return err
		}
		if !result.Acknowledged || !result.ShardsAcknowledged {
			fmt.Printf("[log] Elastic search create doc: %s Warning Acknowledged:%v, ShardsAcknowledged:%v\n",
				index, result.Acknowledged, result.ShardsAcknowledged)
		}
	}
	p.indexName = index
	if !p.sendGorouStart {
		go p.sendLog(BULK_WAIT_TIME)
		p.sendGorouStart = true
	}
	return nil
}

func (p *printer) Exit() {
	//if bulk insert flush the cache
	p.sendLog(0)
	p.client.Stop()
}

func (p *printer) sendLog(t int) {
	for {
		time.Sleep(time.Duration(t)*time.Second)
		req := elastic.NewBulkIndexRequest().Index(p.indexName)
		for _, log := range p.docSlice {
			req = req.Doc(log)
		}
		rsp ,err := p.client.Bulk().Add(req).Do(context.Background())
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
						errString += r.Error.Error() + "  \n"
					}
				}
			}
			fmt.Println(errString)
		}
	}
}