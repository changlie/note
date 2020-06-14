> 转: [浅显易懂的GCC使用教程——初级篇](https://blog.csdn.net/qq_42475711/article/details/85224010)

> 本文gcc版本:  gcc (Ubuntu 5.4.0-6ubuntu1~16.04.10) 5.4.0 20160609

### 什么是gcc，它能干什么？
  GCC(GNU Compiler Collection)即GNU编译器套件，属于一种编程语言编译器，其原名为GCC（GNU C Compiler）即GNU c语言编译器，虽然缩写一样但是功能上区别很大。GCC的初衷是为GNU操作系统专门编写的一款编译器，原本的GNU是专用于编译C代码，现如今已扩展为可以编译C、C++、Java、Objective-C等多种编程语言的编译器集合了。这篇文章主要介绍gcc或g++的使用。

### GCC、gcc、g++三者有何关系？
  gcc（GUN C Compiler）是GCC中的c编译器，而g++（GUN C++ Compiler）是GCC中的c++编译器。
  gcc和g++两者都可以编译c和cpp文件，但存在差异。gcc在编译cpp时语法按照c来编译但默认不能链接到c++的库（gcc默认链接c库，g++默认链接c++库）。g++编译.c和.cpp文件都统一按cpp的语法规则来编译。所以一般编译c用gcc，编译c++用g++。后文有时间会继续深入探讨区别。

用gcc在Windows上编译*.c文件并非直接生成*.exe文件（Linux上为*.out），中间还经历了预处理、编译和汇编几个过程
```
*.c -(预处理)-> *.i -(编译)-> *.s -(汇编)-> *.o -(链接)-> *.exe/*.out
```
- 预处理: 把源文件所有资源合成一个大文件
- 编译: 进行语法检查, 并生产汇编代码
- 汇编: 生成二进制代码
- 链接: 链接就是将汇编生成的OBJ文件、系统库的OBJ文件、库文件链接起来，最终生成可以在特定平台运行的可执行程序
> 函数库一般分为静态库和动态库两种。静态库是指编译链接时，把库文件的代码全部加入到可执行文件中，因此生成的文件比较大，但在运行时也就不再需要库文件 了。其后缀名一般为”.a”。动态库与之相反，在编译链接时并没有把库文件的代码加入到可执行文件中，而是在程序执行时由运行时链接文件加载库，这样可以 节省系统的开销。动态库一般后缀名为”.so”，如前面所述的libc.so.6就是动态库。gcc在编译时默认使用动态库。

### 常用参数
- `-E` 只进行预处理,   
Preprocess only; do not compile, assemble or link   
- `-S` 只进行预处理, 编译   
Compile only; do not assemble or link   
- `-c` 只进行预处理, 编译, 汇编
Compile and assemble, but do not link   
- `-o <file>`  把输出保存到指定文件   
Place the output into <file>


### 实例
实例1,2 源码
```cpp
#include<stdio.h>

int main(void){
    printf("hey g++\n");
}
```

1. 直接生成可执行文件
```bash
changlie@as:~/ws/cpp/cpp11demo
$ gcc demo.cpp
changlie@as:~/ws/cpp/cpp11demo
$ ./a.out
hey g++
```

2. 用`-o`指定可执行文件的名字
```bash
changlie@as:~/ws/cpp/cpp11demo
$ gcc demo.cpp -o demo
changlie@as:~/ws/cpp/cpp11demo
$ ./demo
hey g++
```

3. 加载自定义库文件
源码:
demo.cpp
```cpp
#include <stdio.h>
#include "share.h"

int main(){
    int a = N;   //宏常量
    int b = 2;
    int c = 0;

    c = a + b;
    printf("%d\n", c);

    CODE   //宏替换代码段

	  DoNothing();
}
```
share.h
```cpp
#define N 1

#define CODE  if(c > 2)                      \
                {                            \
                    printf("c > 2\n");       \
                }

void DoNothing(void);          //函数声明(该函数未被调用)
```
share.cpp
```cpp
#include "share.h"

void DoNothing(void){
    ;
    return ;
}
```
(1) 先汇编生成二进制文件, 再生成可执行文件
```bash
changlie@as:~/ws/cpp/cpp11demo
$ gcc demo.cpp -c -o demo.o
changlie@as:~/ws/cpp/cpp11demo
$ gcc demo.o -o demo
demo.o：在函数‘main’中：
demo.cpp:(.text+0x4d)：对‘DoNothing()’未定义的引用
collect2: error: ld returned 1 exit status
```
由于编译器找不到我们的自定义库文件中函数的实现, 故生成可执行文件时报错了

(2) 基于(1), 把share.cpp封装成静态库供其调用
```bash
changlie@as:~/ws/cpp/cpp11demo
$ gcc share.cpp -c -o share.o
changlie@as:~/ws/cpp/cpp11demo
$ ar -rcs libShare.a share.o
changlie@as:~/ws/cpp/cpp11demo
$ gcc demo.o libShare.a -o demo
changlie@as:~/ws/cpp/cpp11demo
$ ./demo
3
c > 2
```

(3) 使用自定义库文件, 一步到位.
```bash
changlie@as:~/ws/cpp/cpp11demo
$ gcc demo.cpp -include share.h share.cpp -o demo
changlie@as:~/ws/cpp/cpp11demo
$ ./demo
4
c: 4, a: 2
I am a string that come from outlog func
```
(3)源码:
demo.cpp
```cpp
#include <stdio.h>
#include "share.h"

int main() {
    int a = N;   //宏常量
    int b = 2;
    int c = 0;

    c = a + b;
    printf("%d\n", c);

    CODE   //宏替换代码段

	  DoNothing();

    char* msg = outlog();
    printf("%s\n", msg);
}
```
share.h
```cpp
#define N 2

#define CODE  if(c > 2)                      \
                {                            \
                    printf("c: %d, a: %d\n", c, a);       \
                }

void DoNothing(void);          //函数声明(该函数未被调用)

char* outlog();
```
share.cpp
```cpp
#include "share.h"

void DoNothing(void) {
    ;
    return ;
}


char* outlog() {
  char* msg = (char*) "I am a string that come from outlog func";
  return msg;
}
```
