package goson

import (
	"errors"
)

type JsonNull struct {
}

func NewJsonNull() *JsonNull {
	return &JsonNull{}
}

func (jn *JsonNull) Put(name string, value interface{}) error {
	return errors.New("JsonNull does not support Put() operation")
}

func (jn *JsonNull) Add(value interface{}) error {
	return errors.New("JsonNull does not support Add() operation")
}

func (jn *JsonNull) Size() int {
	return -1
}

func (jn *JsonNull) GetAllPropName() ([]string, error) {
	return nil, errors.New("JsonNull does not support GetAllPropName() operation")
}

func (jn *JsonNull) Get(index int) (JsonElement, error) {
	return nil, errors.New("JsonNull does not support Get() operation")
}

func (jn *JsonNull) GetByName(name string) (JsonElement, error) {
	return nil, errors.New("JsonNull does not support GetByName() operation")
}

func (jn *JsonNull) Element(path ...string) (JsonElement, error) {
	if len(path) == 0 {
		return jn, nil
	}
	return nil, errors.New("Not find element with path")
}

func (jn *JsonNull) Bool(path ...string) (bool, error) {
	return false, errors.New("Not a bool type")
}

func (jn *JsonNull) Float64(path ...string) (float64, error) {
	return 0, errors.New("Not a float type")
}

func (jn *JsonNull) Float32(path ...string) (float32, error) {
	return 0, errors.New("Not a float type")
}

func (jn *JsonNull) Int64(path ...string) (int64, error) {
	return 0, errors.New("Not an integer type")
}

func (jn *JsonNull) Int32(path ...string) (int32, error) {
	return 0, errors.New("Not an integer type")
}

func (jn *JsonNull) Int(path ...string) (int, error) {
	return 0, errors.New("Not an integer type")
}

func (jn *JsonNull) Byte(path ...string) (byte, error) {
	return 0, errors.New("Not an integer type")
}

func (jn *JsonNull) String(path ...string) (string, error) {
	if len(path) == 0 {
		return "null", nil
	}
	return "", errors.New("Not a string type")
}

func (jn *JsonNull) IsObject() bool {
	return false
}

func (jn *JsonNull) IsArray() bool {
	return false
}

func (jn *JsonNull) IsBool() bool {
	return false
}

func (jn *JsonNull) IsNumber() bool {
	return false
}

func (jn *JsonNull) IsString() bool {
	return false
}

func (jn *JsonNull) IsNull() bool {
	return true
}
