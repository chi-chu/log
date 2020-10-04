# Log  
[中文文档](./README_CN.md)

this is log package for golang  
it support stdout, file, mysql or other Database to write  
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
- **Stdout**  :  
this log will print in the shell
```go
    w, _ := stdout.New()
    log.New(w).SetFormat(log.FORMAT_JSON).SetReportCaller(true).SetLevel(log.DEBUG)

    // if you don`t like the colorful log, use this code
    // .SetColorTip(false)

    log.Info("info test %s %d", "hahahahhaha", 123)
    log.Warn("warn test %s", "hahahahhaha")
    log.Error("error test %s", "hahahahhaha")
```

- **File**  :  
this log will be written in file  
and the file can be rotated by minute/ hour/ day/ week/ month/ year 
```go
    w, err := file.New("./test/test.log")
    if err != nil {
        panic(err)
    }
    //if you set this, log file will be rotate  daily
    w.SetRotateFlag(true).SetRotatePlan(log.ROTATE_DAY)

        // you may define your own rotate filename, use this code
        // default rotate func will generate filename like
        //   logfile_20200101.log  logfile_20200102.log   logfile_20200103.log
        //.SetRotateFunc(func(s string) string{ return s})

    log.New(w).SetFormat(log.FORMAT_JSON).SetReportCaller(true).SetLevel(log.DEBUG)
    log.Info("info test %s %d", "hahahahhaha", 123)
    log.Warn("warn test %s", "hahahahhaha")
    log.Error("error test %s", "hahahahhaha")
```

- **Mysql** :  
```go
    w, err := mysql.New("./test/test.log")
    if err != nil {
        panic(err)
    }
    //if you set this, log file will be rotate  daily
    w.SetRotateFlag(true).SetRotatePlan(log.ROTATE_DAY)

        // you may define your own rotate filename, use this code
        // default rotate func will generate filename like
        //   logfile_20200101.log  logfile_20200102.log   logfile_20200103.log
        //.SetRotateFunc(func(s string) string{ return s})

    log.New(w).SetFormat(log.FORMAT_JSON).SetReportCaller(true).SetLevel(log.DEBUG)
    log.Info("info test %s %d", "hahahahhaha", 123)
    log.Warn("warn test %s", "hahahahhaha")
    log.Error("error test %s", "hahahahhaha")
```

- **Mongo**
```go
    w, err := mongo.New("./test/test.log")
    if err != nil {
        panic(err)
    }
    //if you set this, log file will be rotate  daily
    w.SetRotateFlag(true).SetRotatePlan(log.ROTATE_DAY)

        // you may define your own rotate filename, use this code
        // default rotate func will generate filename like
        //   logfile_20200101.log  logfile_20200102.log   logfile_20200103.log
        //.SetRotateFunc(func(s string) string{ return s})

    log.New(w).SetFormat(log.FORMAT_JSON).SetReportCaller(true).SetLevel(log.DEBUG)
    log.Info("info test %s %d", "hahahahhaha", 123)
    log.Warn("warn test %s", "hahahahhaha")
    log.Error("error test %s", "hahahahhaha")
```
other writer needs to be developing  
  
  
  

### Special function
- hook
```go
    // hook to add user-defined key value
    type hook struct {
    }

    func(h *hooktest) Set(e *log.Entry) {
    	e.Data["heelow"] = "just show your time"
    	e.Data["findasf"] = "123144"
    	e.Data["8888"] = "8888"
    }

    log.New(obj).SetHook(&hook{})
```

#### other
if you like this project star it  
if you have any problem  
please contract me at  544625106@qq.com