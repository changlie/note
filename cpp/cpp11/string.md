### string 相关

```cpp
#include <iostream>
#include <string>
#include <stdio.h>
#include <codecvt>
#include <locale>
using namespace std;

int main (){
    string s5 = "hiya";//拷贝初始化
    s5.append(" go");
    cout << s5 << '\n';

    // string 转 int.
    string tel = "110";
    int a = stoi(tel)+1;
    cout << a << endl;

    // 处理utf-8
    u32string input = U"规定998,死";

    wstring_convert<codecvt_utf8<char32_t>, char32_t> converter;

    int count = 0;
    for(char32_t c : input){
       count++;
       std::cout << converter.to_bytes(c) << std::endl;
    }
    printf("count: %d \n", count);
}
```
