package tpl

type Vars struct {
	v map[string]interface{}
}

func (v *Vars) Set(path string, value interface{}) error {
	return nil
}

func (v *Vars) SetJson(path, j string) error {
	return nil
}

func (v *Vars) Get(path string) interface{} {
	return nil
}

func (v *Vars) GetString(path string) string {
	return ""
}

func (v *Vars) Exist(path string) bool {
	return false
}
