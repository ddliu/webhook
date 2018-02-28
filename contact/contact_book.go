package contact

type ContactBook struct {
	Contacts []*Contact
}

func (c *ContactBook) AddContact(c1 *Contact) {
	c.Contacts = append(c.Contacts, c1)
}

func (c *ContactBook) GetById(id string) *Contact {
	for _, v := range c.Contacts {
		if v.Id == id {
			return v
		}
	}

	return nil
}

func (c *ContactBook) GetByGroup(group string) []*Contact {
	var result []*Contact
	for _, v := range c.Contacts {
		if v.InGroup(group) {
			result = append(result, v)
		}
	}

	return result
}
