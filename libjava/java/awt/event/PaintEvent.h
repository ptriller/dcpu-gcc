
// DO NOT EDIT THIS FILE - it is machine generated -*- c++ -*-

#ifndef __java_awt_event_PaintEvent__
#define __java_awt_event_PaintEvent__

#pragma interface

#include <java/awt/event/ComponentEvent.h>
extern "Java"
{
  namespace java
  {
    namespace awt
    {
        class Component;
        class Rectangle;
      namespace event
      {
          class PaintEvent;
      }
    }
  }
}

class java::awt::event::PaintEvent : public ::java::awt::event::ComponentEvent
{

public:
  PaintEvent(::java::awt::Component *, jint, ::java::awt::Rectangle *);
  virtual ::java::awt::Rectangle * getUpdateRect();
  virtual void setUpdateRect(::java::awt::Rectangle *);
  virtual ::java::lang::String * paramString();
private:
  static const jlong serialVersionUID = 1267492026433337593LL;
public:
  static const jint PAINT_FIRST = 800;
  static const jint PAINT_LAST = 801;
  static const jint PAINT = 800;
  static const jint UPDATE = 801;
private:
  ::java::awt::Rectangle * __attribute__((aligned(__alignof__( ::java::awt::event::ComponentEvent)))) updateRect;
public:
  static ::java::lang::Class class$;
};

#endif // __java_awt_event_PaintEvent__
