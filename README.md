# 按时间 分割日志文件

baseDir 基础目录

baseDir/年-月-日/时.log


年-月-日/时
2006-12-12/15.log

## 安装

```
go get -u github.com/masterZSH/log
```

## 使用

```go
    // 创建日志位置 
    l := log.New("/data/log/services/auth")

    // info log 
    l.Info("foo")

    // error log
    l.Error("bar")
```


// 2020-12-12日12点的日志文件
/data/log/services/auth/2020-12-12/12.log

读取文件数据到influxdb
