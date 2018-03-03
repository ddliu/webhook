package context

import (
	"github.com/spf13/cast"
	"testing"
)

func TestContext(t *testing.T) {
	c := New(nil)
	c.SetValue("a.b.c.d", 99)
	if cast.ToInt(c.GetValue("a.b.c.d")) != 99 {
		t.Error("Get value error")
	}

	c.SetValue("a1.b.c.d", map[string]string{
		"k1": "v1",
		"k2": "v2",
	})

	if cast.ToString(c.GetValue("a1.b.c.d.k1")) != "v1" {
		t.Error("Get value error")
	}
}

func TestTemplate(t *testing.T) {
	tpl := `Author: ${author.name}; License: ${license}; `
	expected := `Author: Dong; License: MIT; `

	c := New(nil)
	c.SetValue("author", map[string]string{
		"name":  "Dong",
		"email": "test@example.com",
	})
	c.SetValue("license", "MIT")

	parsed := c.Tpl(tpl)
	if parsed != expected {
		t.Error("Tpl error: " + parsed)
	}
}
