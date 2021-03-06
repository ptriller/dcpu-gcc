// PR c++/51433
// { dg-options -std=c++0x }

constexpr int f();
constexpr int g() { return f(); }
extern const int n = g();	// dynamic initialization
constexpr int f() { return 42; }
extern const int m = g();
static_assert(m == 42, "m == 42");
