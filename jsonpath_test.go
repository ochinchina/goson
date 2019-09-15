package goson

import (
	"fmt"
	"testing"
)

const (
	TESTING_JSON = `{
  "store": {
    "book": [
      {
        "category": "reference",
        "author": "Nigel Rees",
        "title": "Sayings of the Century",
        "price": 8.95
      },
      {
        "category": "fiction",
        "author": "Herman Melville",
        "title": "Moby Dick",
        "isbn": "0-553-21311-3",
        "price": 8.99
      },
      {
        "category": "fiction",
        "author": "J.R.R. Tolkien",
        "title": "The Lord of the Rings",
        "isbn": "0-395-19395-8",
        "price": 22.99
      }
    ],
    "bicycle": {
      "color": "red",
      "price": 19.95
    }
  },
  "expensive": 10
}`
)

func TestParseJsonPathOk(t *testing.T) {
	NewJsonPath("$.store.book[?(@.price < 10)][0]..test")
	NewJsonPath("$['store']['book'][?(@.price < 10)][0]..test")
	path, err := NewJsonPath("$..book[*].title")
	if err != nil {
		t.Fail()
	}
	fmt.Printf("%v\n", path)

}

func TestFindByPropName(t *testing.T) {
	path, _ := NewJsonPath("$.store.bicycle.color")
	elem, _ := ParseJson(TESTING_JSON)
	r, err := path.FindElement(elem)
	if err != nil {
		t.Fail()
	}
	fmt.Printf("r=%v\n", r)
	if !r.IsString() {
		t.Fail()
	}
	s, err := r.String()
	if err != nil || s != "red" {
		t.Fail()
	}
}

func TestFindBySingleIndex(t *testing.T) {
	path, _ := NewJsonPath("$..book[2]")
	elem, _ := ParseJson(TESTING_JSON)

	r, err := path.FindElement(elem)
	if err != nil {
		t.Fail()
	}
	fmt.Printf("r=%v, err = %v\n", r, err)
}

func TestFindByWildcast(t *testing.T) {
	path, _ := NewJsonPath("$..book[*].price")
	elem, _ := ParseJson(TESTING_JSON)

	r, err := path.FindElement(elem)
	if err != nil {
		t.Fail()
	}
	fmt.Printf("r=%v, err = %v\n", r, err)
}

func TestFindByIndexes(t *testing.T) {
	path, _ := NewJsonPath("$..book[1,2].price")
	elem, _ := ParseJson(TESTING_JSON)

	r, err := path.FindElement(elem)
	if err != nil {
		t.Fail()
	}
	if r.Size() != 2 {
		t.Fail()
	}
	item, err := r.Get(0)
	if err != nil || !item.IsNumber() {
		t.Fail()
	}
	n1, err := item.Float64()
	if err != nil || n1 != 8.99 {
		t.Fail()
	}

}
