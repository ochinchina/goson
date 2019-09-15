package goson

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func ParseJson(data interface{}) (JsonElement, error) {
	switch data.(type) {
	case string:
		s, _ := data.(string)
		if _, err := os.Stat(s); err == nil {
			f, err := os.Open(s)
			if err == nil {
				defer f.Close()
				return parseJson(f)
			} else {
				return nil, fmt.Errorf("Fail to read file %s", s)
			}
		} else {
			return parseJson(strings.NewReader(s))
		}
	case []byte:
		b, _ := data.([]byte)
		return parseJson(bytes.NewBuffer(b))
	case io.Reader:
		reader, _ := data.(io.Reader)
		return parseJson(reader)
	}
	return nil, errors.New("Don't know how to parse data")
}

func parseJson(reader io.Reader) (JsonElement, error) {

	decoder := json.NewDecoder(reader)
	decoder.UseNumber()
	token, err := decoder.Token()
	if err != nil {
		return nil, err
	}
	if token == nil {
		return NewJsonNull(), nil
	}
	if s, ok := token.(string); ok {
		return NewJsonString(s), nil
	}
	if b, ok := token.(bool); ok {
		return NewJsonBool(b), nil
	}
	if n, ok := token.(json.Number); ok {
		return NewJsonNumber(n), nil
	}
	if delim, ok := token.(json.Delim); ok {
		switch delim.String() {
		case "[":
			array := NewJsonArray()
			err := parseJsonArray(decoder, array)
			if err != nil {
				return nil, err
			}
			return array, nil
		case "{":
			obj := NewJsonObject()
			err := parseJsonObject(decoder, obj)
			if err != nil {
				return nil, err
			}
			return obj, nil
		default:
			return nil, fmt.Errorf("Invalid json format with delimeter %s", delim)
		}
	}

	return nil, fmt.Errorf("Invalid json format")
}

func parseJsonArray(decoder *json.Decoder, array *JsonArray) error {
	for {
		token, err := decoder.Token()
		if err != nil {
			return err
		}
		if delim, ok := token.(json.Delim); ok && delim.String() == "]" {
			return nil
		}
		if v, err := toPrimitive(token); err == nil {
			array.Add(v)
			continue
		}
		delim, ok := token.(json.Delim)
		if ok {
			switch delim.String() {
			case "{":
				jObj := NewJsonObject()
				err := parseJsonObject(decoder, jObj)
				if err == nil {
					array.Add(jObj)
				} else {
					return err
				}
			case "]":
				return nil
			case "[":
				ja := NewJsonArray()
				err := parseJsonArray(decoder, ja)
				if err == nil {
					array.Add(ja)
				} else {
					return err
				}
			case "}":
				return fmt.Errorf("invalid json format with delimiter }")

			}
		}
	}
	return fmt.Errorf("invalid json format")
}
func toPrimitive(token json.Token) (JsonElement, error) {
	if token == nil {
		return NewJsonNull(), nil
	}
	if s, ok := token.(string); ok {
		return NewJsonString(s), nil
	}

	if b, ok := token.(bool); ok {
		return NewJsonBool(b), nil
	}
	if n, ok := token.(json.Number); ok {
		return NewJsonNumber(n), nil
	}
	return nil, fmt.Errorf("Not a primitive value")

}
func parseJsonObject(decoder *json.Decoder, obj *JsonObject) error {
	for {
		token, err := decoder.Token()
		if err != nil {
			return err
		}
		if delim, ok := token.(json.Delim); ok && delim.String() == "}" {
			return nil
		}
		field, ok := token.(string)
		if !ok {
			return fmt.Errorf("invalid json format: key should be string")
		}
		token, err = decoder.Token()
		if err != nil {
			return err
		}
		if v, err := toPrimitive(token); err == nil {
			obj.Put(field, v)
			continue
		}

		delim, ok := token.(json.Delim)
		if ok {
			switch delim.String() {
			case "{":
				jObj := NewJsonObject()
				if err = parseJsonObject(decoder, jObj); err != nil {
					return err
				}
				obj.Put(field, jObj)
			case "[":
				ja := NewJsonArray()
				if err = parseJsonArray(decoder, ja); err != nil {
					return err
				}
				obj.Put(field, ja)
			case "}":
				return nil
			case "]":
				return fmt.Errorf("invalid json format with delimeter ]")
			}
		}
	}
	return fmt.Errorf("invalid json object format")
}
