package goson

import (
	"errors"
	"fmt"
)

// represents a json object which includes key/value pairs
// the key is string type, and the value is an JsonElement
type JsonObject struct {
	elements map[string]JsonElement
}

func NewJsonObject() *JsonObject {
	return &JsonObject{elements: make(map[string]JsonElement)}
}
func (jb *JsonObject) Put(name string, value interface{}) error {
	elem, err := toJsonElement(value)
	if err == nil {
		jb.elements[name] = elem
	}

	return err
}

func (jb *JsonObject) Add(value interface{}) error {
	return errors.New("JsonObject does not support Add() operation")
}

func (jb *JsonObject) Size() int {
	return len(jb.elements)
}

func (jb *JsonObject) GetAllPropName() ([]string, error) {
	r := make([]string, 0)
	for k, _ := range jb.elements {
		r = append(r, k)
	}
	return r, nil
}

func (jb *JsonObject) Get(index int) (JsonElement, error) {
	return nil, errors.New("JsonObject does not support Get() operation")
}

func (jb *JsonObject) Element(path ...string) (JsonElement, error) {
	if len(path) == 0 {
		return jb, nil
	}

	if elem, ok := jb.elements[path[0]]; ok {
		return elem.Element(path[1:]...)
	}
	return nil, errors.New("Not find element with path")
}

func (jb *JsonObject) Bool(path ...string) (bool, error) {
	elem, err := jb.Element(path...)
	if err != nil {
		return false, err
	}
	return elem.Bool()
}

func (jb *JsonObject) Float64(path ...string) (float64, error) {
	elem, err := jb.Element(path...)
	if err != nil {
		return 0, err
	}
	return elem.Float64()
}

func (jb *JsonObject) Float32(path ...string) (float32, error) {
	elem, err := jb.Element(path...)
	if err != nil {
		return 0, err
	}
	return elem.Float32()
}

func (jb *JsonObject) Int64(path ...string) (int64, error) {
	elem, err := jb.Element(path...)
	if err != nil {
		return 0, err
	}
	return elem.Int64()
}

func (jb *JsonObject) Int32(path ...string) (int32, error) {
	elem, err := jb.Element(path...)
	if err != nil {
		return 0, err
	}
	return elem.Int32()
}

func (jb *JsonObject) Int(path ...string) (int, error) {
	elem, err := jb.Element(path...)
	if err != nil {
		return 0, err
	}
	return elem.Int()
}

func (jb *JsonObject) Byte(path ...string) (byte, error) {
	elem, err := jb.Element(path...)
	if err != nil {
		return 0, err
	}
	return elem.Byte()
}

func (jb *JsonObject) String(path ...string) (string, error) {
	elem, err := jb.Element(path...)
	if err != nil {
		return "", err
	}
	return elem.String()
}

func (jb *JsonObject) GetByName(name string) (JsonElement, error) {
	value, ok := jb.elements[name]
	if !ok {
		return nil, fmt.Errorf("No such element %s", name)
	}
	return value, nil
}

func (jb *JsonObject) IsObject() bool {
	return true
}

func (jb *JsonObject) IsArray() bool {
	return false
}
func (jb *JsonObject) IsBool() bool {
	return false
}

func (jb *JsonObject) IsNumber() bool {
	return false
}

func (jb *JsonObject) IsString() bool {
	return false
}

func (jb *JsonObject) IsNull() bool {
	return false
}
