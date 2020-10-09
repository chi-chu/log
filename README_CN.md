# Log  
![License](https://img.shields.io/:license-apache-blue.svg)  ![Build Status](https://travis-ci.org/chi-chu/log.svg?branch=master)

[English Document](./README.md)

致力于构建简单易用的log日志包  
它提供了 标准输出，写文件，写mysql，写入Mongo，写入Elastic search  
并且在写入文件、数据库或者其他 提供了日志滚动  
只需要简单的设置rotate就可以， 默认将用时间以天为单位滚动  
就像如下：
```go
|-- LogDir | Database
|   |-- logfile_2020101000.log  | table
|   |-- logfile_2020101100.log
|   |-- logfile_2020101200.log
|   |-- logfile_2020101300.log 
|   |-- logfile_2020101400.log
```
**__特别的__** 能提供彩色日志，使你更快捷的定位和注意到bug 

![Image text](example.png)  

特别醒目有木有！！！！！！！  
ERROR一览无余有木有！！！！！！  

下载本包:
```bash
go get github.com/chi-chu/log
```
导入项目:
```go
import "github.com/chi-chu/log"
```

## 使用示例
- **Stdout** (Default) :  
日志将会标准输出
```go
    //如果你要进行log日志设置请使用下列代码
    //log.Opt(
    //  log.SetLevel(define.DEBUG),
    //  log.SetReportCaller(false),
    //  log.SetFormat(log.FORMAT_JSON),
    //  )

    log.Info("info test %s %d", "hahahahhaha", 123)
    log.Warn("warn test %s", "hahahahhaha")
    log.Error("error test %s", "hahahahhaha")
```

- **File**  :  
日志将会写入文件
```go
    w, err := file.New("./LogDir/logfile.log")
    if err != nil {
        panic(err)
    }
    log.Opt(
        log.SetWriterAndRotate(w, true, log.ROTATE_DAY),
        //如果你需要对log进行level设置等等，使用如下代码
        //log.SetLevel(define.DEBUG),
        //log.SetReportCaller(true),
        //log.SetFormat(log.FORMAT_JSON),
    )
    log.Info("info test %s %d", "hahahahhaha", 123)
```

- **Mysql** :  
    - 默认使用 [Archive 引擎](https://dev.mysql.com/doc/index-archive.html) 来创建表
    - **__注意：__** 目前使用 [gorm](https://github.com/go-gorm/gorm) 驱动来写入mysql数据库  
    你应该创建日志相对应的struct结构，gorm标签应该置前，如下示例  
    标签应使用书写 : gorm:"size:128",json:"func"   
    这种情况将报错 : ~~json:"func",gorm:"size:128"~~
```go
    type LogModel struct {
    	ID        	uint            `gorm:"primaryKey",json:"-"`
    	Func		string          `gorm:"size:128",json:"func"`
    	Line		string          `gorm:"size:64",json:"line"`
    	File		string          `gorm:"size:256",json:"file"`
    	Level		string          `gorm:"size:4",json:"level"`
    	Time		string          `json:"time"`
    	Msg	        string          `json:"msg"`
    }
    cf := &mysql.Config{"root", "123456", "127.0.0.1", 3306, "log"}
    dsn := cf.String()
    //或者你可以直接使用dsn连接
    //  root:password@tcp(127.0.0.1:3306)/database?charset=utf8mb4&parseTime=True&loc=Local
    w, err := mysql.New(dsn, "log", LogModel{})
    if err != nil {
        panic(err)
    }
    log.Opt(
        log.SetWriterAndRotate(w, true, log.ROTATE_DAY),
        )
    log.Debug("info test %s %d", "hahahahhaha", 123)
```

- **Mongo**  
    需要MongoDB 2.6 或者更高的版本
```go
    a := mongo.Config{"root","123456","localhost",0,"log"}
    // 也可以直接使用dns连接
    //dns := "mongodb://localhost:27017 
    //dns := "mongodb://root:123456@localhost:27017/log?authSource=log"
    w, err := mongo.New(a.String(), "log","testlog")
    if err != nil {
        panic(err)
    }
    log.Opt(
        log.SetLevel(define.DEBUG),
        log.SetWriterAndRotate(w, true, log.ROTATE_MINITE),
        )
    log.Debug("debug test %s %d", "hahahahhaha", 123)
    log.Info("info test %s", "hahahahhaha")
```
- **ElasticSearch**  
需要 ElasticSearch 7.x
批量写入的api正在研究中。。。
```go
   w, err := elasticsearch.New([]string{"http://127.0.0.1:9200", "http://127.0.0.2:9200"}, "log",
        //elasticsearch.SetReplicas(4),
        elasticsearch.SetShards(3),
        )
    if err != nil {
        panic(err)
    }
    log.Opt(
        log.SetWriterAndRotate(w, false, log.ROTATE_DAY),
    )
    log.Debug("debug test %s %d", "hahahahhaha", 123)
    log.Info("info test %s", "hahahahhaha")
    log.Warn("warn test %s", "hahahahhaha")
```
其他日志驱动正在集成中。。。。
  
  
  

### 钩子
- hook
```go
    // 注入自定义数据
    type hook struct {
    }

    func(h *hook) Set(e *log.Entry) {
    	e.Data["heelow"] = "just show your time"
    	e.Data["findasf"] = "123144"
    	e.Data["8888"] = "8888"
    }

    log.Opt(log.SetHook(&hook{})
```

#### 其他  
如果有任何问题，请联系我 544625106@qq.com  
欢迎加入一起开发。