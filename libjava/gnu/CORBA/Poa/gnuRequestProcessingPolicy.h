
// DO NOT EDIT THIS FILE - it is machine generated -*- c++ -*-

#ifndef __gnu_CORBA_Poa_gnuRequestProcessingPolicy__
#define __gnu_CORBA_Poa_gnuRequestProcessingPolicy__

#pragma interface

#include <gnu/CORBA/_PolicyImplBase.h>
extern "Java"
{
  namespace gnu
  {
    namespace CORBA
    {
      namespace Poa
      {
          class gnuRequestProcessingPolicy;
      }
    }
  }
  namespace org
  {
    namespace omg
    {
      namespace PortableServer
      {
          class RequestProcessingPolicyValue;
      }
    }
  }
}

class gnu::CORBA::Poa::gnuRequestProcessingPolicy : public ::gnu::CORBA::_PolicyImplBase
{

public:
  gnuRequestProcessingPolicy(::org::omg::PortableServer::RequestProcessingPolicyValue *);
  virtual ::org::omg::PortableServer::RequestProcessingPolicyValue * value();
private:
  static const jlong serialVersionUID = 1LL;
public:
  static ::java::lang::Class class$;
};

#endif // __gnu_CORBA_Poa_gnuRequestProcessingPolicy__
