## repo
github.com/rs/zerolog/log

## 日志
zerolog只专注于记录 JSON 格式的日志，号称 0 内存分配！

## 日志输出
Msg()
Send()

## 链式调用
``` go
func main() {
  log.Debug().
    Str("Scale", "833 cents").
    Float64("Interval", 833.09).
    Msg("Fibonacci is everywhere")

  log.Debug().
    Str("Name", "Tom").
    Send()
}
```

## 日志级别
panic/fatal/error/warn/info/debug/trace

调用SetGlobalLevel()设置全局Logger的日志级别。

## 美化输出
ConsoleWriter 不过性能不够理想，适合开发环境

## 输出行号