golang内置类型和函数
内置类型
值类型：
bool
int(32 or 64), int8, int16, int32, int64
uint(32 or 64), uint8(byte), uint16, uint32, uint64
float32, float64
string
complex64, complex128
array    -- 固定长度的数组

引用类型：(指针类型)
slice   -- 序列数组(最常用)
map     -- 映射
chan    -- 管道

内置函数
append  	-- 把东西增加到slice里面,返回修改后的slice
close   	-- 关闭channel
delete    	-- 从map中删除key对应的value
panic    	-- 停止常规的goroutine
recover 	-- 允许程序定义goroutine的panic动作
imag    	-- 返回complex的实部
real    	-- 返回complex的虚部
make    	-- 返回Type本身(只能应用于slice, map, channel)
new        	-- 返回指向Type的指针
cap        	-- 容量，容积capacity
copy    	-- 复制slice，返回复制的数目
len        	-- 返回长度

内置接口error

type error interface {        //只要实现了Error()函数，返回值为String的都实现了err接口

        Error()    String

}