package simplefactory

import (
	"fmt"
)

//API interface
type API interface {
	Say(name string) string
}

//hiAPI is one of API implement
type hiAPI struct{}

func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi,%s", name)
}

//helloAPI is one of API implement
type helloAPI struct{}

func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello,%s", name)
}

//NewAPI return api instance by type
func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}

	return nil
}
