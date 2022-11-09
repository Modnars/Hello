# 协程

- C++20

## Coroutine Handles

```cpp
#include <concepts>
#include <coroutine>
#include <exception>
#include <iostream>

struct ReturnObject {
    struct promise_type {
        ReturnObject get_return_object() { return {}; }
        std::suspend_never initial_suspend() { return {}; }
        std::suspend_never final_suspend() noexcept { return {}; }
        void unhandled_exception() { }
    };
};

struct Awaiter {
    std::coroutine_handle<> *hp_;
    constexpr bool await_ready() const noexcept { return false; }
    void await_suspend(std::coroutine_handle<> h) { *hp_ = h; }
    constexpr void await_resume() const noexcept { }
};

ReturnObject counter(std::coroutine_handle<> *continuation_out) {
    Awaiter a{continuation_out};
    for (unsigned i = 0;; ++i) {
        co_await a;
        std::cout << "counter: " << i << std::endl;
    }
}

void main1() {
    std::coroutine_handle<> h;
    counter(&h);
    for (int i = 0; i < 3; ++i) {
        std::cout << "In main1 function\n";
        h();
    }
    h.destroy();
}
```

&#160; &#160; &#160; &#160; C++20 使用 `co_await` 将当前协程挂起时，当前的状态、变量等数据都会绑定在堆中，并创建了一个可调用对象用于后续唤醒当前协程以恢复执行。这个可调用对象的类型，就是 `std::coroutine_handle<>`。

&#160; &#160; &#160; &#160; 当执行 `co_await a;` 时，需要 `a` 的类型必须支持一些指定的方法，而更常见的场景是，需要将 `a` 视为一个 **awaitable** 的对象或一个 **awaiter**。

&#160; &#160; &#160; &#160; `await_ready` 接口实际属于一个优化接口，每次执行 `co_await` 时，先执行此接口来判断是否要将当前协程挂起。当然，也可以在 `await_suspend` 中补充逻辑达到相同的优化目的，但需要注意的是，在调用 `await_suspend` 之前，编译器就已经需要将协程状态等数据绑定到可通过 coroutine_handle 访问的堆对象了，而这一操作本身就属于是「重量级」操作了。

&#160; &#160; &#160; &#160; 由此可知，`std::await_always::await_ready` 总是会返回 `false`，而 `std::await_never::await_ready` 则总是会返回 `true`，而这两个类型的其他接口内都没有任何代码逻辑。

## The coroutine return object

```cpp

struct ReturnObject2 {
    struct promise_type {
        ReturnObject2 get_return_object() {
            // Uses C++20 designated initializer syntax
            return {.h_ = std::coroutine_handle<promise_type>::from_promise(*this)};
        }
        std::suspend_never initial_suspend() { return {}; }
        std::suspend_never final_suspend() noexcept { return {}; }
        void unhandled_exception() { }
    };

    std::coroutine_handle<promise_type> h_;
    operator std::coroutine_handle<promise_type>() const { return h_; }
    // A coroutine_handle<promise_type> converts to coroutine_handle<>
    operator std::coroutine_handle<>() const { return h_; }
};

ReturnObject2 counter2() {
    for (unsigned i = 0;; ++i) {
        co_await std::suspend_always{};
        std::cout << "counter2: " << i << std::endl;
    }
}

void main2() {
    std::coroutine_handle<> h = counter2();
    for (int i = 0; i < 3; ++i) {
        std::cout << "In main2 function\n";
        h();
    }
    h.destroy();
}
```

&#160; &#160; &#160; &#160; 在 `promise_type` 定义中，可以使用静态方法 `coroutine_handle::from_promise` 来获取协程处理程序（coroutine handle）。

&#160; &#160; &#160; &#160; 与示例 1 的代码不同，示例 2 中不再需要使用 awaiter （await_suspend）来保存 coroutine_handle 了，所以这里可以直接使用 `co_wait std::suspend_always{}` 来实现相同的挂起协程的效果。

&#160; &#160; &#160; &#160; 此外，尽管在 main2 函数的第一行，通过赋值操作会使得 counter2 返回的协程对象离开函数体被“销毁”，但我们可以将 coroutine_handle 理解为指针，也就是说尽管旧的指针对象已经被销毁了，但我们依然用 h 缓存了指针对象的值，我们依旧可以通过 h 来访问其指向的协程对象，因为这个协程对象本身是不会被销毁掉的。这也就是说，调用 counter2 的同时，需要显示调用 `h.destroy()` 来释放掉这段协程对象空间，否则会存在内存泄漏问题。值得一提的是，如果这里没有使用 h 来“接收” counter2 的返回值，那么就已经存在内存泄漏的问题了。
