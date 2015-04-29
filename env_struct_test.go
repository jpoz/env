package env_struct

import (
	"os"
	"testing"
)

type Foo struct {
	Boo string `env:"BOO"`
	Loo string `env:"ROO"`
	Doo string `expand:"${POO}.com"`
	Moo string
}

func TestDecode(t *testing.T) {
	foo := &Foo{
		Boo: "0",
	}
	Decode(foo)

	if want, got := "0", foo.Boo; want != got {
		t.Errorf("Wanted %q, got %q", want, got)
	}
	if want, got := "", foo.Loo; want != got {
		t.Errorf("Wanted %q, got %q", want, got)
	}

	os.Setenv("BOO", "1")
	os.Setenv("ROO", "2")
	os.Setenv("POO", "3")
	os.Setenv("MOO", "4")

	Decode(foo)
	if want, got := "1", foo.Boo; want != got {
		t.Errorf("Boo should have been %q, got %q", want, got)
	}
	if want, got := "2", foo.Loo; want != got {
		t.Errorf("Loo should have been %q, got %q", want, got)
	}
	if want, got := "3.com", foo.Doo; want != got {
		t.Errorf("Doo should have been %q, got %q", want, got)
	}
	if want, got := "4", foo.Moo; want != got {
		t.Errorf("Moo should have been %q, got %q", want, got)
	}
}
