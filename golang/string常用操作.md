```golang
import (
	"strconv"
	"fmt"
	"reflect"
	"strings"
)

func main() {
	// 类型转换
	//  1.转int
	i, _ := strconv.Atoi("100");
	fmt.Println(i, reflect.TypeOf(i))
	//  2.转flocat64
	f64, _ := strconv.ParseFloat("124.89", 10)
	fmt.Println(f64, reflect.TypeOf(f64))


	// 字符串比较
	fmt.Println("string compare: ", strings.Compare("to", "and")==0)
	fmt.Println("string compare: ", strings.Compare("我去", "我去")==0)


	// 检索子字符串的位置, 不存在时返回-1
	fmt.Println("get index of substring", strings.Index("hello,world", "l"))
	fmt.Println("get index of substring", strings.Index("hello,world", "txt"))
	// lastIndex
	fmt.Println("get index of substring", strings.LastIndex("hello,world", "l"))


	// 统计子字符串出现的次数
	fmt.Println("substring count: ", strings.Count("hello world", "l"))
	fmt.Println("substring count: ", strings.Count("yellow", ""))


	// 重复拼接指定字符串n次,并返回
	fmt.Println("repeat string: ", strings.Repeat("key! ", 3))


	// 字符串替换
	fmt.Println("replace: ", strings.Replace("com/changlie/it/Server", "/", ".", -1))


	// 字符串修剪(trim)
	fmt.Println("trim: ", strings.Trim("//www.baidu.com/", "/"))
	fmt.Println("trim: ", strings.TrimLeft("//www.baidu.com/", "/"))
	fmt.Println("trim: ", strings.TrimRight("//www.baidu.com/", "/"))
	fmt.Println("trimSpace: ", "--"+strings.TrimSpace("  double kill  ")+"--")


	// 字符串大小写
	fmt.Println("Letter LowerUpper:",strings.Title("fly in the sky"))
	fmt.Println("Letter LowerUpper:",strings.Title("changlie"))
	fmt.Println("Letter LowerUpper:",strings.ToLower("DEPARTMENT"))
	fmt.Println("Letter LowerUpper:",strings.ToUpper("hero"))


	// 前缀,后缀判断
	fmt.Println("suffix, prefix: ", strings.HasPrefix("#smoke", "#"))
	fmt.Println("suffix, prefix: ", strings.HasSuffix("#smoke", "kg"))


	// 判断是否存在某个字符或者子串
	fmt.Println("Contains: ", strings.Contains("it's good day!", "day"))
	fmt.Println("Contains: ", strings.Contains("it's good day!", "day1"))
	fmt.Println("ContainsAny:", strings.ContainsAny("fight with the punch!", "hand"))
	fmt.Println("ContainsAny:", strings.ContainsAny("fight with the punch!", "xyz"))
	fmt.Println("ContainsRune:", strings.ContainsRune("天空之城", '天'))



	// 字符串分割
	fmt.Println("splitBySpaceChar:", strings.Join(strings.Fields("a tome   world cat"), "-=>  "))
	fmt.Println("splitBySubstring:", strings.Join(strings.Split("hell#fate#Heaven", "#"), ","))
	fmt.Println("SplitNBySubstring:", strings.Join(strings.SplitN("hell#fate#Heaven", "#", 2), ","))
	fmt.Println("splitAfter:", strings.Join(strings.SplitAfter("a+b+c", "+"), "<"))
	fmt.Println("SplitAfterN:", strings.Join(strings.SplitAfterN("11<22<33", "<", 2), "@"))

}
```
> 参考:

[golang 字符串常用操作](https://studygolang.com/articles/11388)

[strings 字符串操作](https://blog.csdn.net/weixin_41036574/article/details/97611647)
