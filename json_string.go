package goson

import (
    "errors"
    "strconv"
)
type JsonString struct {
    value string
}

func NewJsonString(value string) *JsonString {
    return &JsonString{value: value}
}

func (js *JsonString) Get() string {
    return js.value
}

func (js *JsonString)Element( path...string)(JsonElement, error ) {
    if len( path ) == 0 {
        return js, nil
    }
    return nil, errors.New( "Not find element with path" )
}

func (js *JsonString)Bool(path...string)(bool, error ) {
    if len( path ) == 0 {
        return strconv.ParseBool( js.value )
    }
    return false, errors.New( "Not find element with path" )
}

func (js *JsonString)Float64(path...string)(float64, error ) {
    if len( path ) == 0 {
        return strconv.ParseFloat( js.value, 64 )
    }
    return 0, errors.New( "Not find element with path" )
}

func (js *JsonString)Float32(path...string)(float32, error ) {
    if len( path ) == 0 {
        v, err := strconv.ParseFloat( js.value, 32 )
        return float32(v), err
    }
    return 0, errors.New( "Not find element with path" )
}
func (js *JsonString)Int64(path...string)(int64, error ) {
    if len( path ) == 0 {
        return strconv.ParseInt( js.value, 0, 64 )
    }
    return 0, errors.New( "Not find element with path" )
}

func (js *JsonString)Int32(path...string)(int32, error ) {
    if len( path ) == 0 {
        v, err := strconv.ParseInt( js.value, 0, 32 )
        return int32(v), err
    }
    return 0, errors.New( "Not find element with path" )
}

func (js *JsonString)Int(path...string)(int, error ) {
    if len( path ) == 0 {
        v, err := strconv.ParseInt( js.value, 0, 64 )
        return int(v), err
    }
    return 0, errors.New( "Not find element with path" )
}

func (js *JsonString)Byte(path...string)(byte, error ) {
    if len( path ) == 0 {
        v, err := strconv.ParseInt( js.value, 0, 8 )
        return byte(v), err
    }
    return 0, errors.New( "Not find element with path" )
}

func (js *JsonString)String(path...string)(string, error ) {
    if len( path ) == 0 {
        return js.value, nil
    }
    return "", errors.New( "Not find element with path" )
}


func (js *JsonString) IsObject() bool {
    return false
}
func (js *JsonString) IsArray() bool {
    return false
}

func (js *JsonString) IsBool() bool {
    return false
}

func (js *JsonString) IsNumber() bool {
    return false
}
func (js *JsonString) IsString() bool {
    return true
}

func (js *JsonString) IsNull() bool {
    return false
}

