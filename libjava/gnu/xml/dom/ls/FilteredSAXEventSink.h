
// DO NOT EDIT THIS FILE - it is machine generated -*- c++ -*-

#ifndef __gnu_xml_dom_ls_FilteredSAXEventSink__
#define __gnu_xml_dom_ls_FilteredSAXEventSink__

#pragma interface

#include <gnu/xml/dom/ls/SAXEventSink.h>
#include <gcj/array.h>

extern "Java"
{
  namespace gnu
  {
    namespace xml
    {
      namespace dom
      {
        namespace ls
        {
            class FilteredSAXEventSink;
        }
      }
    }
  }
  namespace org
  {
    namespace w3c
    {
      namespace dom
      {
          class Attr;
          class Node;
        namespace ls
        {
            class LSParserFilter;
        }
      }
    }
    namespace xml
    {
      namespace sax
      {
          class Attributes;
      }
    }
  }
}

class gnu::xml::dom::ls::FilteredSAXEventSink : public ::gnu::xml::dom::ls::SAXEventSink
{

public: // actually package-private
  FilteredSAXEventSink(::org::w3c::dom::ls::LSParserFilter *);
public:
  virtual void startDocument();
  virtual void endDocument();
  virtual void startElement(::java::lang::String *, ::java::lang::String *, ::java::lang::String *, ::org::xml::sax::Attributes *);
public: // actually protected
  virtual ::org::w3c::dom::Attr * createAttr(::org::xml::sax::Attributes *, jint);
public:
  virtual void endElement(::java::lang::String *, ::java::lang::String *, ::java::lang::String *);
  virtual void characters(JArray< jchar > *, jint, jint);
  virtual void processingInstruction(::java::lang::String *, ::java::lang::String *);
  virtual void startDTD(::java::lang::String *, ::java::lang::String *, ::java::lang::String *);
  virtual void endDTD();
  virtual void comment(JArray< jchar > *, jint, jint);
public: // actually package-private
  virtual jshort getDecision(::org::w3c::dom::Node *, jboolean);
  ::org::w3c::dom::ls::LSParserFilter * __attribute__((aligned(__alignof__( ::gnu::xml::dom::ls::SAXEventSink)))) filter;
  jint whatToShow;
  ::java::util::LinkedList * nodes;
  ::java::util::LinkedList * decisions;
  jboolean rejecting;
public:
  static ::java::lang::Class class$;
};

#endif // __gnu_xml_dom_ls_FilteredSAXEventSink__
