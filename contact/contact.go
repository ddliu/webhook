package contact

type Contact struct {
	Id     string
	Groups []string

	// Name   string
	// Email  string
	Properties map[string]interface{}
}

func (c *Contact) InGroup(group string) bool {
	for _, v := range c.Groups {
		if v == group {
			return true
		}
	}

	return false
}

func (c *Contact) GetProperty(name string) interface{} {
	return c.Properties[name]
}
