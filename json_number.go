package goson

import (
    "encoding/json"
    "errors"
)
type JsonNumber struct {
    value json.Number
}

func NewJsonNumber(value json.Number) *JsonNumber {
    return &JsonNumber{value: value}
}

func (jn *JsonNumber) Bool( path...string ) (bool, error) {
    if len( path ) == 0 {
        v, err := jn.Int()
        return v != 0, err
    }
    return false, errors.New( "Not find element with path" )
}

func (jn *JsonNumber) Float64( path...string ) (float64, error) {
    if len( path ) == 0 {
        return jn.value.Float64()
    }
    return 0, errors.New( "Not find element with path" )
}

func (jn *JsonNumber) Float32( path...string ) (float32, error) {
    if len( path ) == 0 {
        v, err := jn.value.Float64()
        return float32( v ), err
    }
    return 0, errors.New( "Not find element with path" )
}


func (jn *JsonNumber)Element( path...string)(JsonElement, error ) {
    if len( path ) == 0 {
        return jn, nil
    }
    return nil, errors.New( "Not find element with path" )
}

func (jn *JsonNumber)Int64(path...string)(int64, error ) {
    if len( path ) == 0 {
        return jn.value.Int64()
    }
    return 0, errors.New( "Not find element with path" )
}

func (jn *JsonNumber)Int32(path...string)(int32, error ) {
    if len( path ) == 0 {
        v, err := jn.value.Int64()
        return int32(v), err
    }
    return 0, errors.New( "Not find element with path" )
}

func (jn *JsonNumber)Int(path...string)(int, error ) {
    if len( path ) == 0 {
        v, err := jn.value.Int64()
        return int(v), err
    }
    return 0, errors.New( "Not find element with path" )
}

func (jn *JsonNumber)Byte(path...string)(byte, error ) {
    if len( path ) == 0 {
        v, err := jn.value.Int64()
        return byte(v), err
    }
    return 0, errors.New( "Not find element with path" )
}


func (jn *JsonNumber)String(path...string)(string, error ) {
    if len( path ) == 0 {
        return jn.value.String(), nil
    }
    return "", errors.New( "Not find element with path" )
}

func (jn *JsonNumber) IsObject() bool {
    return false
}
func (jn *JsonNumber) IsArray() bool {
    return false
}

func (jn *JsonNumber) IsBool() bool {
    return false
}

func (jn *JsonNumber) IsNumber() bool {
    return true
}
func (jn *JsonNumber) IsString() bool {
    return false
}

func (jn *JsonNumber) IsNull() bool {
    return false
}


