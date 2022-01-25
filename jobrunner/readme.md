## repo
github.com/bamzi/jobrunner

## 应用场景
我们在 Web 开发中时常会遇到这样的需求，执行一个操作之后，需要给用户一定形式的通知。例如，用户下单之后通过邮件发送电子发票，网上购票支付后通过短信发送车次信息。但是这类需求并不需要非常及时，如果放在请求流程中处理，会影响请求的响应时间。这类任务我们一般使用异步的方式来执行。jobrunner就是其中一个用来执行异步任务的 Go 语言库。得益于强大的cron库，再搭配jobrunner的任务状态监控，jobrunner非常易于使用。

## 使用pkg
https://pkg.go.dev/github.com/jordan-wright/email#section-readme


## 
``` go
// cron 类型job
func Schedule(spec string, job cron.Job) error

// 
func Every(duration time.Duration, job cron.Job)

// 立即运行job
func Now(job cron.Job)

// 经过duration时间 运行job
func In(duration time.Duration, job cron.Job)
```