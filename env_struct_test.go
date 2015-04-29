package env_struct

import (
	"os"
	"testing"
)

type Foo struct {
	Boo string `env:"BOO" default:"Boo"`
	Loo string `env:"LOO"`
	Moo string
}

func TestDecode(t *testing.T) {
	foo := &Foo{
		Boo: "taco",
	}

	os.Setenv("LOO", "1")

	Decode(foo)
	if want, got := "Boo", foo.Boo; want != got {
		t.Errorf("Wanted %q, got %q", want, got)
	}

	if want, got := "1", foo.Loo; want != got {
		t.Errorf("Wanted %q, got %q", want, got)
	}
}
