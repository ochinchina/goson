package goson

// json element interface, represents one of following:
// - JsonObject, with key (string type) and value (JsonObject, JsonArray or json primitive)
// - JsonArray, json array and start from index 0
// - JsonString, a primitive with string type
// - JsonNumber, a primitive with json number type
// - JsonBool, a primitive with json bool
// - JsonNull, a null value
// all above objects implement this JsonElement
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

    //find json element by path
    Element( path ... string )( JsonElement, error )

    //get the bool value by path
    Bool( path...string)(bool, error)

    //get float64 value by path
    Float64( path...string)(float64, error )


    //get float32 value by path
    Float32( path...string)(float32, error )

    // get int64 by path
    Int64( path...string)(int64, error )

    // get int32 by path
    Int32( path...string)(int32, error )

    // get int by path
    Int( path...string)(int, error )

    // get byte by path
    Byte( path...string)(byte, error )

    //find json elemen by path and return its string value
    //every kind of JsonElement have their string represents
    String( path...string)(string, error )
}

