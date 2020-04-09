java char 默认使用utf-16进行编码
```java
public class CharTest1 {
    public static void main(String[] args) {
        String s = "天上𥊍𪚥";
        System.out.println("len: "+ s.length());
        for(int i=0; i<s.length(); i++){
            System.out.println(i+": "+s.charAt(i));
        }
        System.out.println("source: "+ s);
    }
}
```
输出:
```
len: 6
0: 天
1: 上
2: ?
3: ?
4: ?
5: ?
source: 天上𥊍𪚥
```

> 参考:
https://xiaogd.net/java-%E8%AF%AD%E8%A8%80%E4%B8%AD%E4%B8%80%E4%B8%AA%E5%AD%97%E7%AC%A6%E5%8D%A0%E5%87%A0%E4%B8%AA%E5%AD%97%E8%8A%82%EF%BC%9F/
https://blog.csdn.net/jiangfuqiang/article/details/17222255
