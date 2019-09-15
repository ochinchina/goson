package goson

import (
	"encoding/json"
	"fmt"
)

// json element interface, represents one of following:
// - JsonObject, with key (string type) and value (JsonObject, JsonArray or json primitive)
// - JsonArray, json array and start from index 0
// - JsonString, a primitive with string type
// - JsonNumber, a primitive with json number type
// - JsonBool, a primitive with json bool
// - JsonNull, a null value
// all above objects implement this JsonElement interface
type JsonElement interface {
	// return True if this json element is an object
	IsObject() bool

	//return true if this json element is an array
	IsArray() bool

	//return true if this json element is a bool
	IsBool() bool

	//return true if this json element is number
	IsNumber() bool

	//return true if this json element is a string
	IsString() bool

	//return true if this json element is null
	IsNull() bool

	// only valid for JsonObject
	// add the name/value pair to the JsonObject
	// the value can be one of:
	// - JsonElement
	// - string
	// - json.Number
	// - float
	// - bool
	// - nil
	Put(name string, value interface{}) error

	// Only valid for JsonArray
	// Add the value to the json array
	// the value can be one of:
	// - JsonElement
	// - string
	// - json.Number
	// - float
	// - bool
	// - nil
	Add(value interface{}) error

	// get the size of JsonObject or JsonArray
	// return >= 0 if it is JsonObject or JsonArray, otherwise return -1
	Size() int

	// Only valid for JsonObject
	// Get all the property names
	GetAllPropName() ([]string, error)

	// Only valid for JsonArray
	// Get json element by index
	Get(index int) (JsonElement, error)

	// Only valid for JsonObject
	// Get element by property name
	GetByName(name string) (JsonElement, error)

	//find json element by path
	Element(path ...string) (JsonElement, error)

	//get the bool value by path
	Bool(path ...string) (bool, error)

	//get float64 value by path
	Float64(path ...string) (float64, error)

	//get float32 value by path
	Float32(path ...string) (float32, error)

	// get int64 by path
	Int64(path ...string) (int64, error)

	// get int32 by path
	Int32(path ...string) (int32, error)

	// get int by path
	Int(path ...string) (int, error)

	// get byte by path
	Byte(path ...string) (byte, error)

	//find json element by path and return its string value
	//every kind of JsonElement have their string represents
	String(path ...string) (string, error)
}

func toJsonElement(value interface{}) (JsonElement, error) {
	if value == nil {
		return NewJsonNull(), nil
	}
	switch value.(type) {
	case string:
		s, _ := value.(string)
		return NewJsonString(s), nil
	case JsonElement:
		v, _ := value.(JsonElement)
		return v, nil
	case json.Number:
		n, _ := value.(json.Number)
		return NewJsonNumber(n), nil
	case float64:
		v, _ := value.(float64)
		return NewJsonNumber(json.Number(fmt.Sprintf("%f", v))), nil
	case float32:
		v, _ := value.(float32)
		return NewJsonNumber(json.Number(fmt.Sprintf("%f", v))), nil
	case int64:
		v, _ := value.(int64)
		return NewJsonNumber(json.Number(fmt.Sprintf("%d", v))), nil
	case int32:
		v, _ := value.(int32)
		return NewJsonNumber(json.Number(fmt.Sprintf("%d", v))), nil
	case int:
		v, _ := value.(int)
		return NewJsonNumber(json.Number(fmt.Sprintf("%d", v))), nil
	case byte:
		v, _ := value.(byte)
		return NewJsonNumber(json.Number(fmt.Sprintf("%d", v))), nil
	case bool:
		v, _ := value.(bool)
		return NewJsonBool(v), nil
	}

	return nil, fmt.Errorf("%v can't be converted to JsonElement", value)
}
