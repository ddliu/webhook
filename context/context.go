package context

import (
	"encoding/json"
	"errors"
	"github.com/spf13/cast"
	"reflect"
	"regexp"
	"strings"
)

func New(data interface{}) *Context {
	return &Context{
		data: data,
	}
}

// Context container
type Context struct {
	data interface{}
}

func (c *Context) GetValueE(path string) (interface{}, error) {
	if path == "" || path == "." {
		return c.data, nil
	}

	parts := strings.Split(path, ".")
	v := c.data
	for _, part := range parts {
		if part == "" {
			continue
		}

		m, ok := toMap(v)
		if !ok {
			return nil, errors.New("Invalid data type")
		}

		vv, ok := m[part]
		if !ok {
			return nil, errors.New("Path does not exist")
		}

		v = vv
	}

	return v, nil
}

func (c *Context) GetValue(path string) interface{} {
	v, err := c.GetValueE(path)
	if err != nil {
		return nil
	}

	return v
}

func (c *Context) SetValue(path string, value interface{}) {
	if path == "" || path == "." {
		c.data = value
		return
	}

	parts := strings.Split(path, ".")
	c.data = setValueRecursive(c.data, parts, value)
}

func (c *Context) GetContextE(path string) (*Context, error) {
	v, err := c.GetValueE(path)
	if err != nil {
		return nil, err
	}

	return &Context{
		data: v,
	}, nil
}

func (c *Context) GetContext(path string) *Context {
	return &Context{
		data: c.GetValue(path),
	}
}

func (c *Context) Unmarshal(i interface{}) error {
	b, err := json.Marshal(c.data)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, i)
}

func (c *Context) Exist(path string) bool {
	_, err := c.GetValueE(path)
	return err == nil
}

func setValueRecursive(data interface{}, path []string, value interface{}) map[string]interface{} {
	current := path[0]
	dataMap, ok := toMap(data)
	if !ok || dataMap == nil {
		dataMap = make(map[string]interface{})
	}
	if len(path) == 1 {
		if valueAsContext, ok := value.(*Context); ok {
			if valueAsContext == nil {
				dataMap[current] = nil
			} else {
				dataMap[current] = valueAsContext.GetValue(".")
			}
		} else {
			dataMap[current] = value
		}
	} else {
		nextData, ok := dataMap[current]
		var nextDataMap map[string]interface{}
		if !ok {
			nextData = make(map[string]interface{})
		} else {
			nextDataMap, ok = nextData.(map[string]interface{})
			if !ok {
				nextData = make(map[string]interface{})
			}
		}

		dataMap[current] = setValueRecursive(nextDataMap, path[1:], value)
	}

	return dataMap
}

func (c *Context) Tpl(tpl string) string {
	re := regexp.MustCompile(`\$\{[a-zA-Z0-9\._]+\}`)
	return re.ReplaceAllStringFunc(tpl, c.rep)
}

func (c *Context) rep(str string) string {
	str = str[2 : len(str)-1]
	return cast.ToString(c.GetValue(str))
}

func toMap(data interface{}) (map[string]interface{}, bool) {
	ref := reflect.ValueOf(data)
	if ref.Kind() != reflect.Map {
		return nil, false
	}

	result := make(map[string]interface{})

	for _, k := range ref.MapKeys() {
		result[k.String()] = ref.MapIndex(k).Interface()
	}

	return result, true
}
