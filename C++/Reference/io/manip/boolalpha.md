# std::boolalpha, std::noboolalpha

## 头文件

- [`<ios>`](https://en.cppreference.com/w/cpp/header/ios)

## 概述

### 函数原型

`std::ios_base &boolalpha(std::ios_base &str);`

`std::ios_base &noboolalpha(std::ios_base &str);`

### 函数功能

&#160; &#160; &#160; &#160; 激活 / 取消激活输入输出流的 boolalpha 标识。

### 参数列表

- str I/O 流对象的引用

### 返回

&#160; &#160; &#160; &#160; 经过逻辑执行处理后的 str(I/O 流对象的引用)。

## 示例代码

```cpp
#include <iostream>
#include <locale>
#include <sstream>
int main() {
    // boolalpha output
    std::cout << std::boolalpha << "boolalpha true: " << true << '\n' << "boolalpha false: " << false << '\n';
    std::cout << std::noboolalpha << "noboolalpha true: " << true << '\n' << "noboolalpha false: " << false << '\n';
    // boolalpha parse
    bool b1, b2;
    std::istringstream is("true false");
    is >> std::boolalpha >> b1 >> b2;
    std::cout << '\"' << is.str() << "\" parsed as " << b1 << ' ' << b2 << '\n';
}
```

### 输出

```txt
boolalpha true: true
boolalpha false: false
noboolalpha true: 1
noboolalpha false: 0
"true false" parsed as 1 0
```
