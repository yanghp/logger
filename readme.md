# logger
logger 基于zap的封装

## Use Cases
```go
package main

import "github.com/yanghp/logger"

func main() {
    log:= logger.New(logger.NewOptions())
    
    log.Debug("this is debug message")
    log.Info("this is info message")
    log.Warn("this is warn message")
    log.Error("this is error message")
}

```