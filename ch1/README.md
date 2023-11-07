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