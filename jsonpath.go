package goson
import (
    "errors"
    "fmt"
    "strings"
)
type JsonPathFilter interface {
}

type JsonPathArray interface {
}
type JsonPathItem struct {
    element string
    index int
    recursive bool
}

func NewJsonPathItem( item string ) (*JsonPathItem, error ){
    if strings.HasSuffix( item, "]" ) {
        start := strings.Index( item, "[" )
        end := strings.Index( item, "]" )
        pathItem := &JsonPathItem{ element: "", index: -1, recursive: false }
        if start > 0 && end > start {
            pathItem.element = item[0:start]
            if strings.HasPrefix( pathItem.element, "." ) {
                pathItem.element = pathItem.element[1:]
                pathItem.recursive = true
            }
        }
        for start > 0 && end > start {
            //ingore it currently
            start = nextCharPos( item, end + 1, func( b byte) bool {
                return b == '['
            })
            end = nextCharPos( item, end + 1, func( b byte) bool {
                return b == ']'
            })
        }
        return pathItem, nil
    } else if strings.HasPrefix( item, "." ) {
        return &JsonPathItem{ element: item[1:], index: -1, recursive: true }, nil
    } else {
        return  &JsonPathItem{ element: item[1:], index: -1, recursive: false }, nil
    }
}

func (jpi *JsonPathItem)Element() string {
    return jpi.element
}

func (jpi *JsonPathItem)ArrayIndex()int {
    return jpi.index
}

func (jpi *JsonPathItem)IsRecursive() bool {
    return jpi.recursive
}

func (jpi *JsonPathItem)String() string{
    if jpi.recursive {
        return fmt.Sprintf( ".%s", jpi.element )
    } else {
        return jpi.element
    }
}

// find element from by current field or filter or array indexes
func (jpi *JsonPathItem)FindElement(root JsonElement, cur JsonElement) (JsonElement, error ) {
    return nil, nil
}

type JsonPath struct {
    items[]*JsonPathItem
}


func nextCharPos( s string, start int, isChar func( byte ) bool ) int {
    n := len( s )
    for i := start; i < n; i++ {
        if isChar( s[i] ) {
            return i
        }
    }
    return -1
}

func NewJsonPath( path string ) (*JsonPath, error ){
    //json path must start from root
    if len(path) <= 0 || path[0] != '$' {
        return nil, errors.New( "Invalid path" )
    }
    start := 1
    n := len( path )
    elements := make( []string, 0 )
    for start < n {
        switch( path[start] ) {
            case '.':
                //find . or (
                end := nextCharPos( path, start + 1, func( b byte) bool {
                    return b == '.' || b == '['
                })
                //if fail to find the . or [
                if end == -1 {
                    elements = append( elements, path[ start + 1:] )
                    start = n
                    break
                }
                //if it is '['
                for end != -1 && path[end] == '[' {
                    //try to find ']'
                    end = nextCharPos( path, end + 1, func( b byte) bool {
                        return b == ']'
                    })
                    if end == -1 {
                        return nil, errors.New( "Invalid path")
                    }
                    end = nextCharPos( path, end + 1, func( b byte) bool {
                        return b == '.' || b == '['
                    })
                }
                if end == start + 1 {
                    end = nextCharPos( path, end + 1, func( b byte) bool {
                        return b == '.'
                    })
                }
                if end == -1 {
                    elements = append( elements, path[ start + 1:] )
                    start = n
                } else {
                    elements = append( elements, path[start+ 1: end ] )
                    start = end
                }
            case '[':
                end := nextCharPos( path, start + 1, func( b byte) bool {
                    return b == ']'
                })
                elem := ""
                if end == -1 {
                    elem = path[ start + 1:]
                } else {
                    elem = path[start+ 1: end ]
                }
                //if start with char '
                if strings.HasPrefix( elem, "'" ) {
                    if !strings.HasSuffix( elem, "'" ) {
                        return nil, errors.New( "Invalid path" )
                    }
                    elements = append( elements, elem[1: len( elem ) - 1] )
                } else if len( elements ) <= 0 {
                    return nil, errors.New( "Invalid path" )
                } else {//merge to previous element
                    elements[ len( elements ) - 1] = fmt.Sprintf( "%s[%s]", elements[ len( elements ) - 1 ], elem )
                }
                if end == -1 {
                    break
                }
                start = end+1
            default:
                return nil, errors.New( "Invalid path")
        }
    }
    n = len( elements )
    items := make( []*JsonPathItem, 0 )
    for i := 0; i < n; i++ {
        elem := elements[i]
        if i + 1 < n && elements[i+1][0] == '[' {
            i ++
            elem = fmt.Sprintf("%s%s", elem, elements[i+1] )
        }
        item, err := NewJsonPathItem( elem )
        if err != nil {
            return nil, errors.New( "Invalid path format" )
        }
        items = append( items, item )
    }
    return &JsonPath{ items }, nil
}

// find the json element with json path
func (jp *JsonPath)FindElement( root JsonElement) ( JsonElement, error ) {
    curElement := root

    for _, item := range jp.items {
        curElement, err := item.FindElement( root, curElement )
        if err != nil {
            return nil, err
        }
        if curElement == nil {
            return nil, errors.New( "Fail to find element" )
        }
    }
    return curElement, nil
}

