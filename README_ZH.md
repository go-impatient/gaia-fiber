## 运行
```
go run main.go server start -config ./../../config/config.json
```
## 代码结构

```bash

```

### ginhttp
```
s := ginhttp.NewServer(ginhttp.Addr(":4000))
s.AddBeforeServerStartFunc(bs.InitPprof(), bs.InitExpvar())
s.AddAfterServerStopFunc(bs.CloseLogger())
s.Serve();
```