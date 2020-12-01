# 按时间输出日志

年-月-日/时
2006-12-12/15.log

## 安装

```
go get -u github.com/masterZSH/log
```

## 使用

```go
    // 创建日志位置 
    l := log.New("/tmp/myProject")

    // info log 
    l.Info("foo")

    // error log
    l.Error("bar")
```
