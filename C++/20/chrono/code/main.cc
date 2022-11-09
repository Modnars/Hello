/*
 * @Author: modnarshen
 * @Date: 2022/09/27 16:55:20
 * @LastEditors: modnarshen
 * @Description: std::chrono
 */
#include <chrono>
#include <iostream>

int main(int argc, char *argv[]) {
    // 时间点
    auto now = std::chrono::system_clock::now();
    // 时长
    std::chrono::duration<int64_t, std::nano> ss = now.time_since_epoch();
    // 时长
    auto hh = std::chrono::floor<std::chrono::days>(ss);
    // 时间点
    std::chrono::time_point<std::chrono::system_clock, std::chrono::days> tp =
        std::chrono::floor<std::chrono::days>(now);
    // hh_mm_ss 使用的是时长
    std::chrono::hh_mm_ss<std::chrono::microseconds> tod{
        std::chrono::duration_cast<std::chrono::microseconds>(ss - hh)};
    // year_month_day 使用的是时间点
    const std::chrono::year_month_day ymd(tp);
    std::cout << static_cast<int>(ymd.year()) << static_cast<unsigned>(ymd.month()) << static_cast<unsigned>(ymd.day())
              << " " << tod.hours().count() << ":" << tod.minutes().count() << ":" << tod.seconds().count() << "."
              << tod.subseconds().count() << std::endl;
    return 0;
}
