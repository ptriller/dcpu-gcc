// PR c++/41090
// { dg-do run }
// { dg-options "" }

int i;
struct C
{
  C();
};

C::C()	// { dg-bogus "can never be copied" "" { xfail { { *-apple-darwin* } || { hppa*-*-hpux* && { ! lp64 } } } } }
{
  static void *labelref = &&label;
  goto *labelref;
 label: i = 1;
}

int main()
{
  C c;
  return (i != 1);
}
