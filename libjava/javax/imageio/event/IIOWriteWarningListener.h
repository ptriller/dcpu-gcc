
// DO NOT EDIT THIS FILE - it is machine generated -*- c++ -*-

#ifndef __javax_imageio_event_IIOWriteWarningListener__
#define __javax_imageio_event_IIOWriteWarningListener__

#pragma interface

#include <java/lang/Object.h>
extern "Java"
{
  namespace javax
  {
    namespace imageio
    {
        class ImageWriter;
      namespace event
      {
          class IIOWriteWarningListener;
      }
    }
  }
}

class javax::imageio::event::IIOWriteWarningListener : public ::java::lang::Object
{

public:
  virtual void warningOccurred(::javax::imageio::ImageWriter *, jint, ::java::lang::String *) = 0;
  static ::java::lang::Class class$;
} __attribute__ ((java_interface));

#endif // __javax_imageio_event_IIOWriteWarningListener__
