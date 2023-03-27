> 参考: https://www.cnblogs.com/CPlusPlusZhangChi/p/3153260.html

## 换行符'\n'和回车符'\r'的区别

\r 是回到本行开始输出; \n 是转到下一行继续输出

```cpp
#include <iostream>
using std::cout;
using std::endl;
int main()
{
    cout<<"Hello"<<'\n'<<"World"<<endl;
    cout<<"Hello"<<'\r'<<"World"<<endl;
    //system("pause");
    return 0;
}
```
输出:
```
changlie@as:~/ws/tmp$ ./hey
Hello
World
World

```

