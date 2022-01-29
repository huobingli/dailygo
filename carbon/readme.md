## repo
github.com/uniplaces/carbon

## 时间库
标准库time使用起来不太灵活，特别是日期时间的创建和运算。

## 时区

## 时间比较
标准库time可以使用time.Time对象的Before/After/Equal判断是否在另一个时间对象前，后，或相等。carbon库也可以使用上面的方法比较时间。除此之外，它还提供了多组方法，每个方法提供一个简短名，一个详细名：

- Eq/EqualTo：是否相等；
- Ne/NotEqualTo：是否不等；
- Gt/GreaterThan：是否在之后；
- Lt/LessThan：是否在之前；
- Lte/LessThanOrEqual：是否相同或在之前；
- Between：是否在两个时间之间。

判断当前时间是周几的方法：IsMonday/IsTuesday/.../IsSunday；
是否是工作日，周末，闰年，过去时间还是未来时间：IsWeekday/IsWeekend/IsLeapYear/IsPast/IsFuture。

## 高级特性
### 修饰器
就是对一些特定的时间操作，获取开始和结束时间。如当天、月、季度、年、十年、世纪、周的开始和结束时间，还能获得上一个周二、下一个周一、下一个工作日的时间等等