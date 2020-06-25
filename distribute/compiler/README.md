### 编译原理


[跟vczh看实例学编译原理——二：实现Tinymoe的词法分析](http://www.cppblog.com/vczh/archive/2014/03/02/206014.html)  
[编译原理 - 1 手撸状态机词法分析器](https://www.cnblogs.com/pointer-smq/p/4904531.html)
[从零开始写个编译器吧 - 词法分析器是一个状态机](https://segmentfault.com/a/1190000002478486)


编程语言组成:   
基本类型   
复合类型    
运算符    
控制语句  
函数   
绑定在类型上的函数   



词法分析：   
逐个读取字符，并标记状态，根据当前字符状态与上一个或下一个字符的状态，把一个大字符串分割成一个个token

语法分析：  
Abstract Syntax Tree   
将词法分析好的token，通过语法分析，把token分析成
可以被执行的对象。   
由大到小分别为：    
语句列表（statement list）  
语句（statement）    
表达式（expression）  

其中statement分类：  
Expression statement  
If statement  
For statement   
Return statement  
...

其中expression分类：
Value expression   
Identify expression（变量）   
Assign expression   
Binary expression（语法树）   
Function call expression   
...

Binary expression在执行之前要先转换成AST

归约，因式分解



















