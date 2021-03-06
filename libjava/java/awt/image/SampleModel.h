
// DO NOT EDIT THIS FILE - it is machine generated -*- c++ -*-

#ifndef __java_awt_image_SampleModel__
#define __java_awt_image_SampleModel__

#pragma interface

#include <java/lang/Object.h>
#include <gcj/array.h>

extern "Java"
{
  namespace java
  {
    namespace awt
    {
      namespace image
      {
          class DataBuffer;
          class SampleModel;
      }
    }
  }
}

class java::awt::image::SampleModel : public ::java::lang::Object
{

public:
  SampleModel(jint, jint, jint, jint);
  virtual jint getWidth();
  virtual jint getHeight();
  virtual jint getNumBands();
  virtual jint getNumDataElements() = 0;
  virtual jint getDataType();
  virtual jint getTransferType();
  virtual JArray< jint > * getPixel(jint, jint, JArray< jint > *, ::java::awt::image::DataBuffer *);
  virtual ::java::lang::Object * getDataElements(jint, jint, ::java::lang::Object *, ::java::awt::image::DataBuffer *) = 0;
  virtual ::java::lang::Object * getDataElements(jint, jint, jint, jint, ::java::lang::Object *, ::java::awt::image::DataBuffer *);
  virtual void setDataElements(jint, jint, ::java::lang::Object *, ::java::awt::image::DataBuffer *) = 0;
  virtual void setDataElements(jint, jint, jint, jint, ::java::lang::Object *, ::java::awt::image::DataBuffer *);
  virtual JArray< jfloat > * getPixel(jint, jint, JArray< jfloat > *, ::java::awt::image::DataBuffer *);
  virtual JArray< jdouble > * getPixel(jint, jint, JArray< jdouble > *, ::java::awt::image::DataBuffer *);
  virtual JArray< jint > * getPixels(jint, jint, jint, jint, JArray< jint > *, ::java::awt::image::DataBuffer *);
  virtual JArray< jfloat > * getPixels(jint, jint, jint, jint, JArray< jfloat > *, ::java::awt::image::DataBuffer *);
  virtual JArray< jdouble > * getPixels(jint, jint, jint, jint, JArray< jdouble > *, ::java::awt::image::DataBuffer *);
  virtual jint getSample(jint, jint, jint, ::java::awt::image::DataBuffer *) = 0;
  virtual jfloat getSampleFloat(jint, jint, jint, ::java::awt::image::DataBuffer *);
  virtual jdouble getSampleDouble(jint, jint, jint, ::java::awt::image::DataBuffer *);
  virtual JArray< jint > * getSamples(jint, jint, jint, jint, jint, JArray< jint > *, ::java::awt::image::DataBuffer *);
  virtual JArray< jfloat > * getSamples(jint, jint, jint, jint, jint, JArray< jfloat > *, ::java::awt::image::DataBuffer *);
  virtual JArray< jdouble > * getSamples(jint, jint, jint, jint, jint, JArray< jdouble > *, ::java::awt::image::DataBuffer *);
  virtual void setPixel(jint, jint, JArray< jint > *, ::java::awt::image::DataBuffer *);
  virtual void setPixel(jint, jint, JArray< jfloat > *, ::java::awt::image::DataBuffer *);
  virtual void setPixel(jint, jint, JArray< jdouble > *, ::java::awt::image::DataBuffer *);
  virtual void setPixels(jint, jint, jint, jint, JArray< jint > *, ::java::awt::image::DataBuffer *);
  virtual void setPixels(jint, jint, jint, jint, JArray< jfloat > *, ::java::awt::image::DataBuffer *);
  virtual void setPixels(jint, jint, jint, jint, JArray< jdouble > *, ::java::awt::image::DataBuffer *);
  virtual void setSample(jint, jint, jint, jint, ::java::awt::image::DataBuffer *) = 0;
  virtual void setSample(jint, jint, jint, jfloat, ::java::awt::image::DataBuffer *);
  virtual void setSample(jint, jint, jint, jdouble, ::java::awt::image::DataBuffer *);
  virtual void setSamples(jint, jint, jint, jint, jint, JArray< jint > *, ::java::awt::image::DataBuffer *);
  virtual void setSamples(jint, jint, jint, jint, jint, JArray< jfloat > *, ::java::awt::image::DataBuffer *);
  virtual void setSamples(jint, jint, jint, jint, jint, JArray< jdouble > *, ::java::awt::image::DataBuffer *);
  virtual ::java::awt::image::SampleModel * createCompatibleSampleModel(jint, jint) = 0;
  virtual ::java::awt::image::SampleModel * createSubsetSampleModel(JArray< jint > *) = 0;
  virtual ::java::awt::image::DataBuffer * createDataBuffer() = 0;
  virtual JArray< jint > * getSampleSize() = 0;
  virtual jint getSampleSize(jint) = 0;
public: // actually protected
  jint __attribute__((aligned(__alignof__( ::java::lang::Object)))) width;
  jint height;
  jint numBands;
  jint dataType;
public:
  static ::java::lang::Class class$;
};

#endif // __java_awt_image_SampleModel__
