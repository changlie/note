### java 爬虫技术：
HttpUnit -- 相当于一个没有UI的浏览器

Jsoup -- 一款比较好的Java版HTML解析器

Selenium java -- java webdriver.

Crawler4j -- java open source web crawler

### -----------
## httpclient：支持HTTP协议的客户端编程工具包. 
> 1）封装实现了所有HTTP的方法，如GET，POST，PUT，HEAD   
2）支持redirect，会话保持   
3）支持文件上传 

## httpunit: HTTP请求的测试辅助工具，能处理web测试的需求 .
>相比于HttpClient，不同之处在于：    
1）HttpUnit能对HTTP返回的结果页进行解析，比如DOM元素定位  
2）HttpUnit能自己启动一个servlet来运行被测服务

## htmlunit: HtmlUnit相比HttpUnit功能更加强大，
> HtmlUnit是Junit的扩展测试框架之一，开发者可以使用其提供的API对页面的元素进行操作。
页面的各种元素都可以被当作对象进行调用，对JavaScript的支持也比较好。

## JWebUnit以HttpUnit和JUnit为基础的一个web测试工具。
> 可以用来验证链接跳转、表单输入和提交、表格内容以及Web应用程序特性的正确性。
相比于HtmlUnit，JWebUnit封装的更友好，编写case也会更加简单
