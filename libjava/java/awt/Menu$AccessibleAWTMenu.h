
// DO NOT EDIT THIS FILE - it is machine generated -*- c++ -*-

#ifndef __java_awt_Menu$AccessibleAWTMenu__
#define __java_awt_Menu$AccessibleAWTMenu__

#pragma interface

#include <java/awt/MenuItem$AccessibleAWTMenuItem.h>
extern "Java"
{
  namespace java
  {
    namespace awt
    {
        class Menu;
        class Menu$AccessibleAWTMenu;
    }
  }
  namespace javax
  {
    namespace accessibility
    {
        class AccessibleRole;
    }
  }
}

class java::awt::Menu$AccessibleAWTMenu : public ::java::awt::MenuItem$AccessibleAWTMenuItem
{

public: // actually protected
  Menu$AccessibleAWTMenu(::java::awt::Menu *);
public:
  virtual ::javax::accessibility::AccessibleRole * getAccessibleRole();
private:
  static const jlong serialVersionUID = 5228160894980069094LL;
public: // actually package-private
  ::java::awt::Menu * __attribute__((aligned(__alignof__( ::java::awt::MenuItem$AccessibleAWTMenuItem)))) this$0;
public:
  static ::java::lang::Class class$;
};

#endif // __java_awt_Menu$AccessibleAWTMenu__
