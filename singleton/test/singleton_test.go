package singleton

import (
	"testing"

	"github.com/applepine1125/go-design-pattern/singleton"
)

func TestSingleton(t *testing.T) {
	obj1 := singleton.GetInstance()
	obj2 := singleton.GetInstance()
	if obj1 != obj2 {
		t.Fatalf("expect obj1 == obj2, but obj1 != obj2")
	}
}
