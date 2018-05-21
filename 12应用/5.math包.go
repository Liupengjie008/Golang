import "math"

math包提供了基本的数学常数和数学函数。


func Ceil(x float64) float64
返回不小于x的最小整数（的浮点值），特例如下：
Ceil(±0) = ±0
Ceil(±Inf) = ±Inf
Ceil(NaN) = NaN

func Floor(x float64) float64
返回不大于x的最小整数（的浮点值），特例如下：

Floor(±0) = ±0
Floor(±Inf) = ±Inf
Floor(NaN) = NaN

func Trunc(x float64) float64
返回x的整数部分（的浮点值）。特例如下：

Trunc(±0) = ±0
Trunc(±Inf) = ±Inf
Trunc(NaN) = NaN

func Modf(f float64) (int float64, frac float64)
返回f的整数部分和小数部分，结果的正负号和都x相同；特例如下：

Modf(±Inf) = ±Inf, NaN
Modf(NaN) = NaN, NaN

func Abs(x float64) float64
返回x的绝对值；特例如下：

Abs(±Inf) = +Inf
Abs(NaN) = NaN


func Max(x, y float64) float64
返回x和y中最大值，特例如下：

Max(x, +Inf) = Max(+Inf, x) = +Inf
Max(x, NaN) = Max(NaN, x) = NaN
Max(+0, ±0) = Max(±0, +0) = +0
Max(-0, -0) = -0

func Min(x, y float64) float64
返回x和y中最小值，特例如下：

Min(x, -Inf) = Min(-Inf, x) = -Inf
Min(x, NaN) = Min(NaN, x) = NaN
Min(-0, ±0) = Min(±0, -0) = -0

func Sqrt(x float64) float64
返回x的二次方根，特例如下：

Sqrt(+Inf) = +Inf
Sqrt(±0) = ±0
Sqrt(x < 0) = NaN
Sqrt(NaN) = NaN

func Cbrt(x float64) float64
返回x的三次方根，特例如下：

Cbrt(±0) = ±0
Cbrt(±Inf) = ±Inf
Cbrt(NaN) = NaN