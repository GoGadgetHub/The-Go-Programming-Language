## helloworld
1. Go 是编译型语言
2. Go 原生支持 Unicode，所以可以处理所有国家的语言
3. `go build xxx` 生成一个二进制文件
4. Go 语言使用包来组织，包里包含了单个或多个`.go`文件，使用 `package xxx` 指明该文件属于哪个包
5. `main` 函数很特殊，是每个程序的开始位置
6. `import` 导入包，且必须在 `package`声明之后，严格要求防止引用不需要的包
7. 函数，变量，常量，类型声明（func, var, const, type）
8. 函数：`func funcName(arguments...) (res...) {func body}`
9. 不使用分号，除非多个语句或声明在同一行
10. 因为内部会把跟在特定符号后面的换行符转换为分号，所以换行会影响`Go`代码的解析，比如 `x+y`，换行符只能在加号后面不能在前面
11. 了解下 gofmt，goimports

## echo1, echo2, echo3
1. os包和操作系统打交道
2. os.Args是命令行参数的变量访问，os.Args是一个字符串slice
3. slice是一个动态容量的顺序数组s，s[i]获取值，s[m:n]获取区间的所有值，len(s)返回长度
4. 变量可以在声明的时候初始化，如果变量没有初始化，它将隐式初始化为这个类的空值
5. Go中i++，i--是语句，不是表达式，所以不能使用j = i++，并且仅支持后缀，所以，--i也不行
6. for是Go里面唯一的循环语句
```go
// for 循环
// 执行顺序，initialization condition 代码 post condition 代码
for initialization; condition; post{ 代码 }

// while
for condition {}
```
7. 使用 `_` 表示空标识符，因为Go中不能有没有使用过的变量，所以使用一个占位符

### dup1
1. map存储一个键/值对集合，键可以是其值能够进行相等（==）比较的任意类型，字符串最常见，值可以是任意类型
2. make可用来新建map
3. 当键在map中不存在时，第一次出现时，会初始化为其值类型的初始值，比如 int初始化为0
4. map里面的键迭代顺序不是固定的，通常是随机的，每次运行都不一致，防止程序依赖某种特定的序列，此处不对排序做保证
5. 在`fmt.Printf()`打印一个格式化输出。第一个参数是格式化指示字符串，由它指定其它参数如何格式化。每一个参数的格式是一个转义字符，一个百分号加一个字符
   - %d 十进制
   - %x，%o，%e 16，8，2进制
   - %f，%g，%e 浮点数
   - %t 布尔
   - %c 字符（Unicode）
   - %s 字符串
   - %v 任何值类型
   - 。。。
6. ln结尾的打印函数，如Println，内部使用%v的方式来格式化参数，且末尾加了换行符
7. 函数和其他包级别的实体可以任意次序声明
8. map是一个使用make创建的引用。当一个map被传递给一个函数时，函数接收到这个引用的副本，当被调用函数中对map数据结构进行改变，对map本身也会改变
9. String.Split 可以把一个字符串按照规定字符分割成一个 slice
10. const声明用来给常量命名，常量是其值在编译期间固定的值
11. const声明可以出现在包级别和函数内，常量必须是字符串，布尔值，数字

### fetch
1. [ReadAll和Copy的区别](https://juejin.cn/spost/7300812626278563859)
2. `goroutine`是一个并发执行函数。
3. 通道是一种允许某一例程向另一例程传递指定类型的值的通信机制。
4. main 函数在一个 goroutine 中执行，go语句可以创建额外的 goroutine
5. 当一个 goroutine 试图在一个通道上进行发送或接收操作时，它会阻塞，直到另一个 goroutine 试图进行接收或发送操作才传递值
6. 两个并发请求同时更新某值时，会产生一个竞态bug，为了避免必须确保同一时间只有一个修改操作（加锁）
7. go 允许一个简单语句（如一个局部变量声明）跟在if条件前面，这在错误处理的时候特别有用
8. go中提供了指针，它的值是变量的地址，使用&操作符可以获取变量的地址，使用*操作符可以获取指针引用的变量的值，但指针不支持算术运算