
// DO NOT EDIT THIS FILE - it is machine generated -*- c++ -*-

#ifndef __javax_swing_TransferHandler$TransferAction__
#define __javax_swing_TransferHandler$TransferAction__

#pragma interface

#include <javax/swing/AbstractAction.h>
extern "Java"
{
  namespace java
  {
    namespace awt
    {
      namespace datatransfer
      {
          class Clipboard;
      }
      namespace event
      {
          class ActionEvent;
      }
    }
  }
  namespace javax
  {
    namespace swing
    {
        class JComponent;
        class TransferHandler$TransferAction;
    }
  }
}

class javax::swing::TransferHandler$TransferAction : public ::javax::swing::AbstractAction
{

public:
  TransferHandler$TransferAction(::java::lang::String *);
  virtual void actionPerformed(::java::awt::event::ActionEvent *);
private:
  static ::java::awt::datatransfer::Clipboard * getClipboard(::javax::swing::JComponent *);
  ::java::lang::String * __attribute__((aligned(__alignof__( ::javax::swing::AbstractAction)))) command;
public:
  static ::java::lang::Class class$;
};

#endif // __javax_swing_TransferHandler$TransferAction__
