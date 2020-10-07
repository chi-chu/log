# Log  
![License](https://img.shields.io/:license-apache-blue.svg)  ![Build Status](https://travis-ci.org/chi-chu/log.svg?branch=master)

[中文文档](./README_CN.md)

this is log package for golang  
it support stdout, file, mysql or other Database to write  
and auto rotate like this, default daily
```go
|-- LogDir | Database  
|   |-- logfile_2020101000.log | table
|   |-- logfile_2020101100.log
|   |-- logfile_2020101200.log
|   |-- logfile_2020101300.log 
|   |-- logfile_2020101400.log
```
**_SPECIAL_** it has some cool stdout like this below 

![Image text](example.png)  

it`s so striking, emmmm, easy to debug your code  
find the **ERROR** as soon as possible

To download the package, run:
```bash
go get github.com/chi-chu/log
```
Import it in your program as:
```go
import "github.com/chi-chu/log"
```

## Usage
- **Stdout** (Default) :  
this log will print in the shell
```go
    //this code will set the option if you like
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
this log will be written in file  
and the file can be rotated by minute/ hour/ day/ week/ month/ year 
```go
    w, err := file.New("./LogDir/logfile.log")
    if err != nil {
        panic(err)
    }
    log.Opt(
        log.SetWriterAndRotate(w, true, log.ROTATE_DAY),
    )
    log.Info("info test %s %d", "hahahahhaha", 123)
```

- **Mysql** :  
    - it will Default use [Archive Engine](https://dev.mysql.com/doc/index-archive.html) to create table
    - **__Attention__** it ues [gorm](https://github.com/go-gorm/gorm) to drive database  
    you should give struct detail like this 
    
    remember to use: gorm:"size:128",json:"func"   
    do not to use: ~~json:"func",gorm:"size:128"~~
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
    //or you can direct use the dsn url like
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
```go
   
```

- **ElasticSearch**
```go
   
```
other writer needs to be developing  
  
  
  

### Special function
- hook
```go
    // hook to add user-defined key value
    type hook struct {
    }

    func(h *hook) Set(e *log.Entry) {
    	e.Data["heelow"] = "just show your time"
    	e.Data["findasf"] = "123144"
    	e.Data["8888"] = "8888"
    }

    log.Opt(log.SetHook(&hook{})
```

#### other
if you like this project star it  
if you have any problem  
please contract me at  544625106@qq.com  
welcome to join and perfect it