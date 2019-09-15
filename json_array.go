package goson

import (
	"fmt"
	"strconv"
)

// represents a json array object. The element of array is another
// JsonElement, which can be accessed with index (start from 0 )
type JsonArray struct {
	elements []JsonElement
}

func NewJsonArray() *JsonArray {
	return &JsonArray{elements: make([]JsonElement, 0)}
}

func (ja *JsonArray) Put(name string, value interface{}) error {
	return fmt.Errorf("json array does not support Put() operation")
}

func (ja *JsonArray) Add(value interface{}) error {
	elem, err := toJsonElement(value)
	if err == nil {
		ja.elements = append(ja.elements, elem)
	}
	return err
}

func (ja *JsonArray) Size() int {
	return len(ja.elements)
}

func (ja *JsonArray) GetAllPropName() ([]string, error) {
	return nil, fmt.Errorf("json array does not support GetAllPropName() operation")
}
func (ja *JsonArray) Get(index int) (JsonElement, error) {
	if index < 0 || index >= len(ja.elements) {
		return nil, fmt.Errorf("%d is out of index", index)
	}
	return ja.elements[index], nil
}

func (ja *JsonArray) GetByName(name string) (JsonElement, error) {
	return nil, fmt.Errorf("json array does not support GetByName() operation")
}

func (ja *JsonArray) Element(path ...string) (JsonElement, error) {
	if len(path) == 0 {
		return ja, nil
	}
	index, err := strconv.Atoi(path[0])
	if err != nil {
		return nil, err
	}
	if index >= 0 && index < len(ja.elements) {
		return ja.elements[index].Element(path[1:]...)
	}
	return nil, fmt.Errorf("%d is out of index", index)
}

// get the element by path as bool type
func (ja *JsonArray) Bool(path ...string) (bool, error) {
	elem, err := ja.Element(path...)

	if err != nil {
		return false, err
	}

	return elem.Bool()
}

func (ja *JsonArray) Float64(path ...string) (float64, error) {
	elem, err := ja.Element(path...)

	if err != nil {
		return 0, err
	}

	return elem.Float64()
}

func (ja *JsonArray) Float32(path ...string) (float32, error) {
	elem, err := ja.Element(path...)

	if err != nil {
		return 0, err
	}

	return elem.Float32()
}

func (ja *JsonArray) Int64(path ...string) (int64, error) {
	elem, err := ja.Element(path...)

	if err != nil {
		return 0, err
	}

	return elem.Int64()
}
func (ja *JsonArray) Int32(path ...string) (int32, error) {
	elem, err := ja.Element(path...)

	if err != nil {
		return 0, err
	}

	return elem.Int32()
}

func (ja *JsonArray) Int(path ...string) (int, error) {
	elem, err := ja.Element(path...)

	if err != nil {
		return 0, err
	}

	return elem.Int()
}

func (ja *JsonArray) Byte(path ...string) (byte, error) {
	elem, err := ja.Element(path...)

	if err != nil {
		return 0, err
	}

	return elem.Byte()
}

func (ja *JsonArray) String(path ...string) (string, error) {
	elem, err := ja.Element(path...)

	if err != nil {
		return "", err
	}

	return elem.String()
}

func (ja *JsonArray) IsObject() bool {
	return false
}

func (ja *JsonArray) IsArray() bool {
	return true
}
func (ja *JsonArray) IsBool() bool {
	return false
}

func (ja *JsonArray) IsNumber() bool {
	return false
}

func (ja *JsonArray) IsString() bool {
	return false
}

func (ja *JsonArray) IsNull() bool {
	return false
}
