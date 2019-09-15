package goson

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// select JsonElement from current element
type ElementSelector interface {
	Search(root JsonElement, cur JsonElement) (JsonElement, error)
}

type PropNameSelector struct {
	// property name
	propName string
}

func NewPropNameSelector(propName string) *PropNameSelector {
	return &PropNameSelector{propName: propName}
}

func (ps *PropNameSelector) Search(root JsonElement, cur JsonElement) (JsonElement, error) {
	if strings.HasPrefix(ps.propName, ".") {
		array := NewJsonArray()
		searchElementsByPropName(ps.propName[1:], cur, func(elem JsonElement) {
			array.Add(elem)
		})
		if array.Size() > 1 {
			return array, nil
		} else if array.Size() == 1 {
			return array.Get(0)
		} else {
			return nil, errors.New("Fail to find the element by property name")
		}
	} else if cur.IsArray() {
		n := cur.Size()
		r := NewJsonArray()
		for i := 0; i < n; i++ {
			elem, err := cur.Get(i)
			if err != nil {
				continue
			}
			elem, err = elem.GetByName(ps.propName)
			if err == nil {
				r.Add(elem)
			}
		}
		if r.Size() > 0 {
			return r, nil
		}
		return nil, errors.New("Fail to find the element by property name")
	} else {
		return cur.GetByName(ps.propName)
	}
}

func searchElementsByPropName(propName string, cur JsonElement, elemFindCb func(JsonElement)) {
	if cur.IsObject() {
		elem, err := cur.GetByName(propName)
		if err == nil {
			elemFindCb(elem)
		} else {
			names, _ := cur.GetAllPropName()
			for _, name := range names {
				elem, _ := cur.GetByName(name)
				searchElementsByPropName(propName, elem, elemFindCb)
			}
		}
	} else if cur.IsArray() {
		n := cur.Size()
		for i := 0; i < n; i++ {
			elem, _ := cur.Get(i)
			searchElementsByPropName(propName, elem, elemFindCb)
		}
	}
}

type NthElementSelector struct {
	n int
}

func NewNthElementSelector(n int) *NthElementSelector {
	fmt.Printf("n=%d\n", n)
	return &NthElementSelector{n: n}
}

func (nes *NthElementSelector) Search(root JsonElement, cur JsonElement) (JsonElement, error) {
	if cur.IsArray() {
		fmt.Printf("cur.Size=%d\n", cur.Size())
		if nes.n >= 0 && cur.Size() > nes.n {
			return cur.Get(nes.n)
		} else if nes.n < 0 && cur.Size() > -1*nes.n {
			return cur.Get(cur.Size() + nes.n)
		} else {
			return nil, errors.New("Fail to get Nth element")
		}
	} else {
		return nil, errors.New("Can't get element by index from non-array element")
	}
}

type IndexesElementSelector struct {
	indexes []string
}

func NewIndexesElementSelector(indexes []string) *IndexesElementSelector {
	return &IndexesElementSelector{indexes: indexes}
}

func (ies *IndexesElementSelector) Search(root JsonElement, cur JsonElement) (JsonElement, error) {
	r := NewJsonArray()
	for _, index := range ies.indexes {
		if cur.IsArray() {
			i, err := strconv.Atoi(index)
			if err == nil {
				fmt.Printf("i=%d, size=%d\n", i, cur.Size())
				if i >= 0 && i < cur.Size() {
					elem, err := cur.Get(i)
					if err == nil {
						r.Add(elem)
					}
				} else if i < 0 && -1*i < cur.Size() {
					elem, err := cur.Get(cur.Size() + i)
					if err == nil {
						r.Add(elem)
					}
				}
			}
		} else if cur.IsObject() {
			elem, err := cur.GetByName(index)
			if err == nil {
				r.Add(elem)
			}
		}
	}

	if r.Size() > 0 {
		return r, nil
	}
	return nil, errors.New("not implement")
}

type ArraySliceElementSelector struct {
	start int
	end   int
	step  int
}

func NewArraySliceElementSelector(start int, end int, step int) *ArraySliceElementSelector {
	return &ArraySliceElementSelector{start: start, end: end, step: step}
}

func (ases *ArraySliceElementSelector) Search(root JsonElement, cur JsonElement) (JsonElement, error) {
	return nil, errors.New("not implement")
}

type FilterElementSelector struct {
	expression string
}

func NewFilterElementSelector(expression string) *FilterElementSelector {
	return &FilterElementSelector{expression: expression}
}

func (fes *FilterElementSelector) Search(root JsonElement, cur JsonElement) (JsonElement, error) {
	return nil, errors.New("not implement")
}

type ScriptElementSelector struct {
	expression string
}

func NewScriptElementSelector(expression string) *ScriptElementSelector {
	return &ScriptElementSelector{expression: expression}
}

func (ses *ScriptElementSelector) Search(root JsonElement, cur JsonElement) (JsonElement, error) {
	return nil, nil
}

type WildcastElementSelector struct {
}

func (wes *WildcastElementSelector) Search(root JsonElement, cur JsonElement) (JsonElement, error) {
	if cur.IsArray() || cur.IsObject() {
		return cur, nil
	}
	return nil, errors.New("not implement")
}

type JsonPathItem struct {
	selectors []ElementSelector
}

func NewJsonPathItem(item string) (*JsonPathItem, error) {
	fmt.Printf("%s\n", item)
	selectors := make([]ElementSelector, 0)
	//check if it is an array item
	if strings.HasSuffix(item, "]") {
		start := strings.Index(item, "[")
		end := strings.Index(item, "]")
		if start > 0 && end > start {
			selectors = append(selectors, NewPropNameSelector(item[0:start]))
		}
		//find more [ and ]
		for start > 0 && end > start {
			selector, err := createSubscriptionElementSelector(item[start+1 : end])
			if err != nil {
				return nil, err
			}
			selectors = append(selectors, selector)

			//ingore it currently
			start = nextCharPos(item, end+1, func(b byte) bool {
				return b == '['
			})
			end = nextCharPos(item, end+1, func(b byte) bool {
				return b == ']'
			})
		}
	} else {
		selectors = append(selectors, NewPropNameSelector(item))
	}
	if len(selectors) > 0 {
		return &JsonPathItem{selectors: selectors}, nil
	} else {
		return nil, fmt.Errorf("Invalid json path item:%s", item)
	}
}

func createSubscriptionElementSelector(subscription string) (ElementSelector, error) {
	fmt.Printf("subscription:%s\n", subscription)
	if subscription == "*" {
		return &WildcastElementSelector{}, nil
	}
	// syntax check: [index1,index2,…]
	indexes := strings.Split(subscription, ",")
	if len(indexes) > 1 {
		return NewIndexesElementSelector(indexes), nil
	}

	// syntax check: [start:end:step]
	slices := strings.Split(subscription, ":")
	if len(slices) > 1 {
		if len(slices) > 3 {
			return nil, errors.New("invalid array slice syntax")
		}
		start := 0
		end := -1
		step := 1
		var err error
		if len(slices[0]) > 0 {
			if start, err = strconv.Atoi(slices[0]); err != nil {
				return nil, err
			}
		}
		if len(slices[1]) > 0 {
			if end, err = strconv.Atoi(slices[1]); err != nil {
				return nil, err
			}
		}
		if len(slices) == 3 && len(slices[2]) > 0 {
			if step, err = strconv.Atoi(slices[2]); err != nil {
				return nil, err
			}
		}
		return NewArraySliceElementSelector(start, end, step), nil
	}

	// syntax check: [?(expression)]
	if strings.HasPrefix(subscription, "?(") && strings.HasSuffix(subscription, ")") {
		return NewFilterElementSelector(subscription[2 : len(subscription)-1]), nil
	}

	// syntax check: [(expression)]

	if strings.HasPrefix(subscription, "(") && strings.HasSuffix(subscription, ")") {
		return NewScriptElementSelector(subscription[1 : len(subscription)-1]), nil
	}

	// syntax check: [n]
	n, err := strconv.Atoi(subscription)
	if err == nil {
		return NewNthElementSelector(n), nil
	}
	return nil, errors.New("Invalid subscription syntax error")

}

// find element from by current field or filter or array indexes
func (jpi *JsonPathItem) FindElement(root JsonElement, cur JsonElement) (JsonElement, error) {
	var err error
	t := cur
	for _, selector := range jpi.selectors {
		t, err = selector.Search(root, t)
		if err != nil {
			return nil, err
		}
	}
	if t != nil {
		return t, nil
	} else {
		return nil, errors.New("Fail to find element by path")
	}
}

type JsonPath struct {
	items []*JsonPathItem
}

func nextCharPos(s string, start int, isChar func(byte) bool) int {
	n := len(s)
	for i := start; i < n; i++ {
		if isChar(s[i]) {
			return i
		}
	}
	return -1
}

func NewJsonPath(path string) (*JsonPath, error) {
	//json path must start from root
	if len(path) <= 0 || path[0] != '$' {
		return nil, errors.New("Invalid path")
	}
	start := 1
	n := len(path)
	elements := make([]string, 0)
	for start < n {
		switch path[start] {
		case '.':
			//find . or (
			end := nextCharPos(path, start+1, func(b byte) bool {
				return b == '.' || b == '['
			})
			//if fail to find the . or [
			if end == -1 {
				elements = append(elements, path[start+1:])
				start = n
				break
			}
			//if it is '['
			for end != -1 && path[end] == '[' {
				//try to find ']'
				end = nextCharPos(path, end+1, func(b byte) bool {
					return b == ']'
				})
				if end == -1 {
					return nil, errors.New("Invalid path")
				}
				end = nextCharPos(path, end+1, func(b byte) bool {
					return b == '.' || b == '['
				})
			}
			if end == start+1 {
				end = nextCharPos(path, end+1, func(b byte) bool {
					return b == '.'
				})
			}
			if end == -1 {
				elements = append(elements, path[start+1:])
				start = n
			} else {
				elements = append(elements, path[start+1:end])
				start = end
			}
		case '[':
			end := nextCharPos(path, start+1, func(b byte) bool {
				return b == ']'
			})
			elem := ""
			if end == -1 {
				elem = path[start+1:]
			} else {
				elem = path[start+1 : end]
			}
			//if start with char '
			if strings.HasPrefix(elem, "'") {
				if !strings.HasSuffix(elem, "'") {
					return nil, errors.New("Invalid path")
				}
				elements = append(elements, elem[1:len(elem)-1])
			} else if len(elements) <= 0 {
				return nil, errors.New("Invalid path")
			} else { //merge to previous element
				elements[len(elements)-1] = fmt.Sprintf("%s[%s]", elements[len(elements)-1], elem)
			}
			if end == -1 {
				break
			}
			start = end + 1
		default:
			return nil, errors.New("Invalid path")
		}
	}
	n = len(elements)
	items := make([]*JsonPathItem, 0)
	for i := 0; i < n; i++ {
		elem := elements[i]
		if i+1 < n && elements[i+1][0] == '[' {
			i++
			elem = fmt.Sprintf("%s%s", elem, elements[i+1])
		}
		item, err := NewJsonPathItem(elem)
		if err != nil {
			return nil, errors.New("Invalid path format")
		}
		items = append(items, item)
	}
	return &JsonPath{items}, nil
}

// find the json element with json path
func (jp *JsonPath) FindElement(root JsonElement) (JsonElement, error) {
	curElement := root
	var err error

	for _, item := range jp.items {
		curElement, err = item.FindElement(root, curElement)
		if err != nil {
			return nil, err
		}
		if curElement == nil {
			return nil, errors.New("Fail to find element")
		}
	}
	return curElement, nil
}
