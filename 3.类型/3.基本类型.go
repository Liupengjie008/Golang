Golang 更明确的数字类型命名，支持 Unicode，支持常用数据结构。

类型            度      默认值    说明
bool            1       false     
byte            1       0         uint8
rune            4       0         Unicode Code Point, int32
int, uint       4或8    0         32 或 64 位
int8, uint8     1       0         -128 ~ 127, 0 ~ 255
int16, uint16   2       0         -32768 ~ 32767, 0 ~ 65535
int32, uint32   4       0         -21亿~ 21亿, 0 ~ 42亿
int64, uint64   8       0
float32         4       0.0
float64         8       0.0
complex64       8
complex128      16
uintptr         4或8              以存储指针的 uint32 或 uint64 整数
array                             值类型
struct                            值类型
string                  ""        UTF-8 字符串
slice                   nil       引用类型
map                     nil       引用类型
channel                 nil       引用类型
interface               nil       接口
function                nil       函数
                        

支持八进制、 六进制，以及科学记数法。标准库 math 定义了各数字类型取值范围。 

a, b, c, d := 071, 0x1F, 1e9, math.MinInt16

空指针值 nil，而非C/C++ NULL。

基本类型应用：

bool:只能存true和false

uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 的别名 ：字符类型

rune // int32 的别名
     // 代表一个Unicode码，用UTF-8 进行编码。
     这个类型在什么时候使用呢？
     例如需要遍 历字符串中的字符。可以循环每个字节（仅在使用US ASCII 编码字符串时与字符等价， 而它们在Go 中不存在！）。因此为了获得实际的字符，需要使用rune 类型。

string rune byte 的关系

在Go当中 string底层是用byte数组存的，并且是不可以改变的。

例如 s:="Go编程" fmt.Println(len(s)) 输出结果应该是8因为中文字符是用3个字节存的。

len(string(rune('编')))的结果是3

如果想要获得我们想要的情况的话，需要先转换为rune切片再使用内置的len函数

fmt.Println(len([]rune(s)))

结果就是4了。

所以用string存储unicode的话，如果有中文，按下标是访问不到的，因为你只能得到一个byte。 要想访问中文的话，还是要用rune切片，这样就能按下表访问。

float32 float64

complex64 complex128
与导入语句一样，变量的定义“打包”在一个语法块中。

int，uint 和 uintptr 类型在32位的系统上一般是32位，而在64位系统上是64位。当你需要使用一个整数类型时，你应该首选 int，仅当有特别的理由才使用定长整数类型或者无符号整数类型。


值类型：变量直接存储值，内存通常在栈中分配。
        获取变量的地址，用&，比如： var a int, 获取a的地址：&a
        基本数据类型int、float、bool、string以及数组和struct。