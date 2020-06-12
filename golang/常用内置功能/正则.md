> 转：     
[go 正则-知乎](https://zhuanlan.zhihu.com/p/50398448)  
[官方文档](https://golang.org/pkg/regexp/)  
[Golang 中的正则表达式](https://www.cnblogs.com/golove/p/3269099.html)
```
分组：

        (子表达式)            被捕获的组，该组被编号 (子匹配)
        (?P<命名>子表达式)    被捕获的组，该组被编号且被命名 (子匹配)
        (?:子表达式)          非捕获的组 (子匹配)
        (?标记)               在组内设置标记，非捕获，标记影响当前组后的正则表达式
        (?标记:子表达式)      在组内设置标记，非捕获，标记影响当前组内的子表达式

        标记的语法是：
        xyz  (设置 xyz 标记)
        -xyz (清除 xyz 标记)
        xy-z (设置 xy 标记, 清除 z 标记)

        可以设置的标记有：
        i              不区分大小写 (默认为 false)
        m              多行模式：让 ^ 和 $ 匹配整个文本的开头和结尾，而非行首和行尾(默认为 false)
        s              让 . 匹配 \n (默认为 false)
        U              非贪婪模式：交换 x* 和 x*? 等的含义 (默认为 false)
```

### 1 概述
正则表达式，又称规则表达式，Regular Expression，在代码中常简写为 regex、regexp 或 RE。正则表达式是对字符串操作的一种逻辑公式，就是用事先定义好的一些特定字符、及这些特定字符的组合，组成一个“规则字符串”，这个“规则字符串”用来表达对字符串的一种过滤逻辑。 Perl 语言的正则表达式功能非常强大，很多语言设计正则式支持的时候都参考Perl的正则表达式。因此常用的表达式语法也是 Perl 兼容正则表达式。

Go 语言中使用包 regexp 提供对正则表达式的支持。本文说明 regexp 中常用的正则处理方法。

### 2 获取正则对象
通过编译正则表达式，可以得到正则操作对象，用于完成正则的相关处理： 函数：
```golang
regexp.Compile(expr string) (*Regexp, error)，用于编译一个正则表达式，如果成功返回 Regexp 对象。
regexp.MustCompile(str string) *Regexp，与 Compile 一致，差异是失败时会宕机。

reg,err := regexp.Compile(`\d+`)
reg := regexp.MustCompile(`\d+`)
```

### 3 匹配检测
函数:
```golang
func (re *Regexp) MatchString(s string) bool，测试字符串是否匹配正则。
func (re *Regexp) Match(b []byte) bool，检测字节切片是否匹配正则。
```
演示字符串的匹配：
```golang
text := "Hello Gopher，Hello 韩忠康"
reg := regexp.MustCompile(`\w+`)
fmt.Println(reg.MatchString(text))
// true
```
### 4 查找
函数： 
```golang
func (re *Regexp) FindString(s string) string，查找匹配模式的字符串，返回左侧第一个匹配的结果。 
func (re *Regexp) FindAllString(s string, n int) []string，用来查找匹配模式的字符串，返回多个匹配的结果，n 用于限定查找数量，-1不限制。 
func (re *Regexp) FindAll(b []byte, n int) [][]byte，用于在 []byte 中查找，返回 [][]byte。
```
匹配全部结果演示为：
```golang
text := "Hello Gopher，Hello 韩忠康"
reg := regexp.MustCompile(`\w+`)
fmt.Println(reg.FindAllString(text))
// [Hello Gopher Hello]
```

### 5 查找匹配位置
以下函数用于获取匹配正则子字符串的位置： 
```golang
func (re *Regexp) FindStringIndex(s string) (loc []int)，返回包含最左侧匹配结果的起止位置的切片。 
func (re *Regexp) FindIndex(b []byte) (loc []int)，返回包含最左侧匹配结果的起止位置的切片。 
func (re *Regexp) FindAllStringIndex(s string, n int) [][]int 会返回包含全部匹配结果的起止位置的切片的切片。
```
演示查找字符串最左侧匹配位置：
```golang
text := "Hello Gopher，Hello 韩忠康"
reg := regexp.MustCompile("llo")
fmt.Println(reg.FindStringIndex(text))
// [2 5]
```

### 6 查找匹配子模式
以下函数可以查找子模式，或查找子模式的位置： 
```golang
func (re *Regexp) FindStringSubmatch(s string) []string，查找字符串中最左侧子匹配结果 
func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string，查找字符串中全部匹配和子模式。 
func (re *Regexp) FindStringSubmatchIndex(s string) []int，查找字符串中包含最左侧子匹配结果的起止位置的切片
```
演示匹配全部子字符串如下：
```golang
re := regexp.MustCompile("Go(\w+)")
fmt.Println(re.FindAllStringSubmatch("Hello Gopher，Hello GoLang", -1))
// [["Gophoer" "phoer"], ["GoLang", "Lang"]]
```

### 7 替换
函数： 
```golang
func (re *Regexp) ReplaceAll(src, repl []byte) []byte //一致，针对的是 []byte。 替换时可以使用反向引用 $1，$2，来引用匹配的子模式内容。
func (re *Regexp) ReplaceAllFunc(src []byte, repl func([]byte) []byte) []byte
func (re *Regexp) ReplaceAllLiteral(src, repl []byte) []byte
func (re *Regexp) ReplaceAllLiteralString(src, repl string) string
func (re *Regexp) ReplaceAllString(src, repl string) string  // 将 src 中所有 re 的匹配结果都替换为 repl。
func (re *Regexp) ReplaceAllStringFunc(src string, repl func(string) string) string
```
demo
```golang
re := regexp.MustCompile("Go(\w+)")
fmt.Println(re.ReplaceAllString("Hello Gopher，Hello GoLang", "Hank$1"))
// Hello Hankpher，Hello HankLang
```

### 8 分割
函数： 
```golang
func (re *Regexp) Split(s string, n int) []string，使用正则分割字符串 s ，返回字符串切片。n 控制分割的片数，-1为不限制。
```
demo
```golang
reg := regexp.MustCompile("[\s,]")
fmt.Println(reg.Split("Hello Gopher,Hello GoLang", -1))
// [Hello Gopher Hello GoLang]
```
