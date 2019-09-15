package goson

import (
	"fmt"
	"testing"
)

func Float64Equal(f1 float64, f2 float64, precision float64) bool {
	if f1 > f2 {
		return f1 <= f2+precision
	} else {
		return f2 <= f1+precision
	}
}

func Float32Equal(f1 float32, f2 float32, precision float32) bool {
	if f1 > f2 {
		return f1 <= f2+precision
	} else {
		return f2 <= f1+precision
	}

}

func TestValidJsonParse(t *testing.T) {
	s := `{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}`

	_, err := ParseJson(s)
	if err != nil {
		t.Fail()
	}
}

func TestInvalidJsonParse(t *testing.T) {
	s := `{"ID":1,"Name":"Reds","Colors":"Crimson","Red","Ruby","Maroon"]}`

	_, err := ParseJson(s)
	if err == nil {
		t.Fail()
	}
}

func TestFindElementByPath(t *testing.T) {
	s := `{
    "tool": 
    {
        "jsonpath": 
        {
            "creator": 
            {
                "name": "Jayway Inc.",
                "location": 
                [
                    "Malmo",
                    "San Francisco",
                    "Helsingborg"
                ]
            }
        }
    },
 
    "book": 
    [
        {
            "title": "Beginning JSON",
            "price": 49.99
        },
 
        {
            "title": "JSON at Work",
            "price": 29.99
        }
    ]
}`
	elem, err := ParseJson(s)
	if err != nil {
		t.Fail()
	}
	name, err := elem.Element("tool", "jsonpath", "creator", "location", "2")
	if err != nil || !name.IsString() {
		t.Fail()
	}
	v, err := name.String()
	if err != nil || v != "Helsingborg" {
		t.Fail()
	}

}

func TestGetFloat64ByPath(t *testing.T) {
	s := `{
    "tool":
    {
        "jsonpath":
        {
            "creator":
            {
                "name": "Jayway Inc.",
                "location":
                [
                    "Malmo",
                    "San Francisco",
                    "Helsingborg"
                ]
            }
        }
    },

    "book":
    [
        {
            "title": "Beginning JSON",
            "price": 49.99
        },

        {
            "title": "JSON at Work",
            "price": 29.99
        }
    ]
}`
	elem, err := ParseJson(s)

	if err != nil {
		t.Fail()
	}
	price, err := elem.Float64("book", "1", "price")
	fmt.Printf("%f\n", price)
	if err != nil || !Float64Equal(price, 29.99, 0.001) {
		t.Fail()
	}
}

func TestGetFloat32ByPath(t *testing.T) {
	s := `{
    "tool":
    {
        "jsonpath":
        {
            "creator":
            {
                "name": "Jayway Inc.",
                "location":
                [
                    "Malmo",
                    "San Francisco",
                    "Helsingborg"
                ]
            }
        }
    },

    "book":
    [
        {
            "title": "Beginning JSON",
            "price": 49.99
        },

        {
            "title": "JSON at Work",
            "price": 29.99
        }
    ]
}`
	elem, err := ParseJson(s)

	if err != nil {
		t.Fail()
	}
	price, err := elem.Float32("book", "0", "price")
	if err != nil || !Float32Equal(price, 49.99, 0.001) {
		t.Fail()
	}
}
