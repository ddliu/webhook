package context

import (
	"encoding/json"
	"errors"
	"strings"
)

// Context container
type Context struct {
	data interface{}
}

func (c *Context) GetValue(path string) (interface{}, error) {
	if path == "" || path == "." {
		return c.data, nil
	}

	parts := strings.Split(path, ".")
	v := c.data
	for _, part := range parts {
		if part == "" {
			continue
		}
		m, ok := c.data.(map[string]interface{})
		if !ok {
			return nil, errors.New("Invalid data type")
		}

		v, ok = m[part]
		if !ok {
			return nil, errors.New("Path does not exist")
		}
	}

	return v, nil
}

func (c *Context) SetValue(path string, value interface{}) {
	if path == "" || path == "." {
		c.data = value
		return
	}

	parts := strings.Split(path, ".")
	c.data = setValueRecursive(c.data, parts, value)
}

func (c *Context) GetContext(path string) (*Context, error) {
	v, err := c.GetValue(path)
	if err != nil {
		return nil, err
	}

	return &Context{
		data: v,
	}, nil
}

func (c *Context) Unmarshal(i interface{}) error {
	b, err := json.Marshal(c.data)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, i)
}

func (c *Context) Exist(path string) bool {
	_, err := c.GetValue(path)
	return err == nil
}

func setValueRecursive(data interface{}, path []string, value interface{}) map[string]interface{} {
	current := path[0]
	dataMap, ok := data.(map[string]interface{})
	if !ok {
		dataMap = make(map[string]interface{})
	}
	if len(path) == 1 {
		dataMap[current] = value
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
