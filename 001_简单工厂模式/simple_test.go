package simplefactory

import (
	"testing"
)

//测试 简单工厂模式下的 hiAPI
func TestType1(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("jack")
	if s != "Hi,jack" {
		t.Fatal("Type1 test fail")
	}
}

//测试 简单工厂模式下的 helloAPI
func TestType2(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("jack")
	if s != "Hello,jack" {
		t.Fatal("Type2 test fail")
	}
}
