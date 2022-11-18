# std::exchange

## 头文件

- [`<utility>`](https://en.cppreference.com/w/cpp/header/utility)

## 说明

### 函数声明

| 函数原型 | SINCE | UNTIL |
| :- | :-: | :-: |
| `template<class T, class U = T>`<br>`T exchange(T &obj, U &&new_value);` | ![](https://img.shields.io/badge/C%2B%2B-14-brightgreen) | ![](https://img.shields.io/badge/C%2B%2B-20-brightgreen) |
| `template<class T, class U = T>`<br>`constexpr T exchange(T &obj, U &&new_value);` | ![](https://img.shields.io/badge/C%2B%2B-20-brightgreen) | ![](https://img.shields.io/badge/C%2B%2B-23-brightgreen) |
| `template<class T, class U = T>`<br>`constexpr T exchange(T &obj, U &&new_value) noexcept(/* see below */);` | ![](https://img.shields.io/badge/C%2B%2B-23-brightgreen) | |

### 概述

&#160; &#160; &#160; &#160; 使用 `new_value` 的值来替代 `obj` 的值，并将 `obj` 的旧值作为返回值返回。

### 参数列表

- `obj` 待替换的变量。

- `new_value` 需要替换成的新值。
