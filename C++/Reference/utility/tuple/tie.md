# std::tie

## 头文件

- [`<tuple>`](https://en.cppreference.com/w/cpp/header/tuple)

## 概述

### 函数声明

| 函数原型 | SINCE | UNTIL |
| :- | :-: | :-: |
| `template <class... Types>`<br>`std::tuple<Types &...> tie(Types &... args) noexcept;` | ![](https://img.shields.io/badge/C%2B%2B-11-brightgreen) | ![](https://img.shields.io/badge/C%2B%2B-14-brightgreen) |
| `template <class... Types>`<br>`constexpr std::tuple<Types &...> tie(Types &... args) noexcept;` | ![](https://img.shields.io/badge/C%2B%2B-14-brightgreen) |  |

### 函数功能

&#160; &#160; &#160; &#160; 创建一个实参是左值引用或 std::ignore 实例对象的元组。

### 参数列表

- `args` 用于创建元组实例对象的零个或多个左值实参。

### 返回

&#160; &#160; &#160; &#160; 持有左值引用的元组对象。

## 可能的实现

```cpp
template <typename... Args>
constexpr  // since C++14
    std::tuple<Args &...>
    tie(Args &... args) noexcept {
    return {args...};
}
```

## 说明

&#160; &#160; &#160; &#160; `std::tie` 可用于解析一个 `std::pair` 对象，因为 `std::tuple` 拥有通过 `std::pair` 对象构造的转换构造函数。

```cpp
bool result;
std::tie(std::ignore, result) = set.insert(value);
```

## 示例代码

```cpp
#include <iostream>
#include <set>
#include <string>
#include <tuple>

struct S {
    int n;
    std::string s;
    float d;
    bool operator<(const S &rhs) const {
        // compares n to rhs.n,
        // then s to rhs.s,
        // then d to rhs.d
        return std::tie(n, s, d) < std::tie(rhs.n, rhs.s, rhs.d);
    }
};

int main() {
    std::set<S> set_of_s;  // S is LessThanComparable

    S value{42, "Test", 3.14};
    std::set<S>::iterator iter;
    bool inserted;

    // unpacks the return value of insert into iter and inserted
    std::tie(iter, inserted) = set_of_s.insert(value);

    if (inserted)
        std::cout << "Value was inserted successfully\n";
}
```

### 输出

```txt
Value was inserted successfully
```
