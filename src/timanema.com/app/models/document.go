package models

import (
  "github.com/robfig/revel"
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
  "errors"
  "reflect"
)

type Document struct {
  D           interface{}
  LastError   error
}

func (doc *Document) Id() string {
  return reflect.ValueOf(doc.D).Elem().FieldByName("Id").String()
}

func (doc *Document) IsNew() bool {
  return !doc.IsPersisted()
}

func (doc *Document) IsPersisted() bool {
  return bson.ObjectId(doc.Id()).Valid()
}

func (doc *Document) Validate(v *revel.Validation) {
  reflect.ValueOf(doc.D).MethodByName("Validate").Call([]reflect.Value{reflect.ValueOf(v)})
}

func (doc *Document) Get(field_name string) (val interface {}) {
  return reflect.ValueOf(doc.D).Elem().FieldByName(field_name).Interface()
}

func (doc *Document) Set(field_name string, v interface{}) {
  reflect.ValueOf(doc.D).Elem().FieldByName(field_name).Set(reflect.Value(reflect.ValueOf(v)))
}

func (doc *Document) Save() bool {
  collection_name := collection_name_from(doc.D)
  err :=  with_collection(collection_name, func(c *mgo.Collection) (err error) {
    if doc.IsPersisted() {
      err = c.UpdateId(doc.Id(), doc.D)
    }else{
      err = c.Insert(doc.D)
    }
    doc.LastError = err
    return
  })
  return err == nil
}

func (doc *Document) Update(changes interface{}) bool {
  collection_name := collection_name_from(doc.D)
  err := with_collection(collection_name, func(c *mgo.Collection) (err error) {
    if doc.IsPersisted() {
      err = c.UpdateId(doc.Id(), changes)
    }else{
      err = errors.New("Document is not persisted, Please use Save instead of Update")
    }
    doc.LastError = err
    return
  })
  return err == nil
}
