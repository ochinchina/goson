package goson

import (
	"errors"
)

type JsonBool struct {
	value bool
}

func NewJsonBool(value bool) *JsonBool {
	return &JsonBool{value: value}
}

func (jb *JsonBool) Put(name string, value interface{}) error {
	return errors.New("JsonBool does not support Put() operation")
}

func (jb *JsonBool) Add(value interface{}) error {
	return errors.New("JsonBool does not support Add() operation")
}

func (jb *JsonBool) Size() int {
	return -1
}

func (jb *JsonBool) GetAllPropName() ([]string, error) {
	return nil, errors.New("JsonBool does not support GetAllPropName() operation")
}

func (jb *JsonBool) Get(index int) (JsonElement, error) {
	return nil, errors.New("JsonBool does not support Get() operation")
}

func (jb *JsonBool) GetByName(name string) (JsonElement, error) {
	return nil, errors.New("JsonBool does not support GetByName() operation")
}

func (jb *JsonBool) Element(path ...string) (JsonElement, error) {
	if len(path) == 0 {
		return jb, nil
	}
	return nil, errors.New("Not find element with path")
}

func (jb *JsonBool) Bool(path ...string) (bool, error) {
	return jb.value, nil
}

func (jb *JsonBool) Float64(path ...string) (float64, error) {
	return 0, errors.New("Not float type")
}

func (jb *JsonBool) Float32(path ...string) (float32, error) {
	return 0, errors.New("Not float type")
}

func (jb *JsonBool) Int64(path ...string) (int64, error) {
	return 0, errors.New("Not int type")
}

func (jb *JsonBool) Int32(path ...string) (int32, error) {
	return 0, errors.New("Not int type")
}

func (jb *JsonBool) Int(path ...string) (int, error) {
	return 0, errors.New("Not int type")
}

func (jb *JsonBool) Byte(path ...string) (byte, error) {
	return 0, errors.New("Not int type")
}

func (jb *JsonBool) String(path ...string) (string, error) {
	if len(path) == 0 {
		if jb.value {
			return "true", nil
		} else {
			return "false", nil
		}
	}
	return "", errors.New("Not String type")
}

func (jb *JsonBool) IsObject() bool {
	return false
}

func (jb *JsonBool) IsArray() bool {
	return false
}

func (jb *JsonBool) IsBool() bool {
	return true
}

func (jb *JsonBool) IsNumber() bool {
	return false
}

func (jb *JsonBool) IsString() bool {
	return false
}
func (jb *JsonBool) IsNull() bool {
	return false
}
