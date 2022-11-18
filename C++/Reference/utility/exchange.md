# std::exchange

## 头文件

- [`<utility>`](https://en.cppreference.com/w/cpp/header/utility)

## 概述

### 函数声明

| 函数原型 | SINCE | UNTIL |
| :- | :-: | :-: |
| `template<class T, class U = T>`<br>`T exchange(T &obj, U &&new_value);` | ![](https://img.shields.io/badge/C%2B%2B-14-brightgreen) | ![](https://img.shields.io/badge/C%2B%2B-20-brightgreen) |
| `template<class T, class U = T>`<br>`constexpr T exchange(T &obj, U &&new_value);` | ![](https://img.shields.io/badge/C%2B%2B-20-brightgreen) | ![](https://img.shields.io/badge/C%2B%2B-23-brightgreen) |
| `template<class T, class U = T>`<br>`constexpr T exchange(T &obj, U &&new_value) noexcept(/* see below */);` | ![](https://img.shields.io/badge/C%2B%2B-23-brightgreen) | |

### 函数功能

&#160; &#160; &#160; &#160; 使用 `new_value` 的值来替代 `obj` 的值，并将 `obj` 的旧值作为返回值返回。

### 参数列表

- `obj` 待替换的变量。

- `new_value` 需要替换成的新值。

### 返回

&#160; &#160; &#160; &#160; `obj` 的旧值。

### 异常

| | SINCE | UNTIL |
| :- | :-: | :-: |
| 无 | | ![](https://img.shields.io/badge/C%2B%2B-23-brightgreen) |
| `noexcept(std::is_nothrow_move_constructible_v<T> && std::is_nothrow_assignable_v<T&, U>)` | ![](https://img.shields.io/badge/C%2B%2B-23-brightgreen) | |

### 可用于特征测试的宏

- [`__cpp_lib_exchange_function`](https://en.cppreference.com/w/cpp/feature_test#Library_features)


## 可能的实现

```cpp
template<class T, class U = T>
constexpr // since C++20
T exchange(T &obj, U &&new_value) noexcept( // since C++23
        std::is_nothrow_move_constructible<T>::value &&
        std::is_nothrow_assignable<T&, U>::value) {
    T old_value = std::move(obj);
    obj = std::forward<U>(new_value);
    return old_value;
}
```

## 说明

&#160; &#160; &#160; &#160; `std::exchange` 可用于实现**移动赋值运算符重载**和**移动构造函数**：

```cpp
struct S {
    int n;

    S(S &&other) noexcept : n{std::exchange(other.n, 0)} { }

    S &operator=(S &&other) noexcept {
        if (this != &other)
            // 移动 n 值，并将 other.n 置为 0
            n = std::exchange(other.n, 0);
        return *this;
    }
};
```

## 示例代码

```cpp
#include <iostream>
#include <iterator>
#include <utility>
#include <vector>

class stream {
public:
    using flags_type = int;

public:
    flags_type flags() const { return flags_; }

    // 使用 newf 替换 flags_，并将 flags_ 的旧值返回。
    flags_type flags(flags_type newf) { return std::exchange(flags_, newf); }

private:
    flags_type flags_ = 0;
};

void f() {
    std::cout << "f()";
}

int main() {
    stream s;

    std::cout << s.flags() << '\n';
    std::cout << s.flags(12) << '\n';
    std::cout << s.flags() << "\n\n";

    std::vector<int> v;

    // 由于模板的第二个参数已经指定了默认类型，所以使用初始列表作为其参数是合理的。
    // 这个表达式的效果等同于 std::exchange(v, std::vector<int>{1, 2, 3, 4})
    std::exchange(v, {1, 2, 3, 4});

    std::copy(begin(v), end(v), std::ostream_iterator<int>(std::cout, ", "));

    std::cout << "\n\n";

    void (*fun)();

    // 模板参数的默认值同样适用于普通函数作为实参的场景。
    // 这个表达式的效果等同于 std::exchange(fun, static_cast<void(*)()>(f))
    std::exchange(fun, f);
    fun();

    std::cout << "\n\nFibonacci sequence: ";
    for (int a{0}, b{1}; a < 100; a = std::exchange(b, a + b))
        std::cout << a << ", ";
    std::cout << "...\n";
}
```

### 输出

```txt
0
0
12
 
1, 2, 3, 4, 
 
f()
 
Fibonacci sequence: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
```


