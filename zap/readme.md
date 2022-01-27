## repo
go.uber.org/zap

## 日志库
在热点函数中记录日志对日志库的执行性能有较高的要求，不能影响正常逻辑的执行时间。uber开源的日志库zap，对性能和内存分配做了极致的优化。

## 日志级别
Debug/Info/Error/Warn

## 创建logger方法
``` go
// zap使用预定义设置
zap.NewExample()            // 适用于测试
zap.NewDevelopment()        // 适用于开发
zap.NewProduction()         // 适用于生产

// zap定制化
zap.New()
```

## zap.Type / zap.Typep / zap.Types
Type为bool/int/uint/float64/complex64/time.Time/time.Duration/error等   
Typep 上述类型指针
Types 上述类型切片

## 代码行数
``` go
zap.AddCaller()
```

## 打印堆栈
``` go
zap.AddStackTrace(lvl zapcore.LevelEnabler)
```