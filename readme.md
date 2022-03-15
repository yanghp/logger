# logger
logger 基于zap的封装

## Use Cases
```go
func func main() {
    log:= logger.New(logger.NewOptions())
    
    log.Debug("this is debug message")
    log.Info("this is info message")
    log.Warn("this is warn message")
    log.Error("this is error message")
}
```