
// DO NOT EDIT THIS FILE - it is machine generated -*- c++ -*-

#ifndef __gnu_javax_naming_jndi_url_rmi_rmiURLContextFactory__
#define __gnu_javax_naming_jndi_url_rmi_rmiURLContextFactory__

#pragma interface

#include <java/lang/Object.h>
extern "Java"
{
  namespace gnu
  {
    namespace javax
    {
      namespace naming
      {
        namespace jndi
        {
          namespace url
          {
            namespace rmi
            {
                class rmiURLContextFactory;
            }
          }
        }
      }
    }
  }
  namespace javax
  {
    namespace naming
    {
        class Context;
        class Name;
    }
  }
}

class gnu::javax::naming::jndi::url::rmi::rmiURLContextFactory : public ::java::lang::Object
{

public:
  rmiURLContextFactory();
  virtual ::java::lang::Object * getObjectInstance(::java::lang::Object *, ::javax::naming::Name *, ::javax::naming::Context *, ::java::util::Hashtable *);
  static ::java::lang::Class class$;
};

#endif // __gnu_javax_naming_jndi_url_rmi_rmiURLContextFactory__
