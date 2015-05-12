package decoder

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

type Bar struct {
	Tar string
	Foo Foo
}

type Muc struct {
	Lup   int   `env:"LUP"`
	Lup32 int32 `env:"LUP"`
	Lup64 int64 `env:"LUP"`
	Tup   bool  `env:"PUP"`
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

	os.Setenv("TAR", "5")
	bar := &Bar{}
	Decode(bar)
	if want, got := "1", bar.Foo.Boo; want != got {
		t.Errorf("bar.Foo.Boo should have been %q, got %q", want, got)
	}
	if want, got := "2", bar.Foo.Loo; want != got {
		t.Errorf("bar.Foo.Loo should have been %q, got %q", want, got)
	}
	if want, got := "3.com", bar.Foo.Doo; want != got {
		t.Errorf("bar.Foo.Doo should have been %q, got %q", want, got)
	}
	if want, got := "4", bar.Foo.Moo; want != got {
		t.Errorf("bar.Foo.Moo should have been %q, got %q", want, got)
	}
	if want, got := "5", bar.Tar; want != got {
		t.Errorf("bar.Tar should have been %q, got %q", want, got)
	}

	os.Setenv("LUP", "800")
	os.Setenv("PUP", "true")
	muc := &Muc{}
	Decode(muc)
	if want, got := int(800), muc.Lup; want != got {
		t.Errorf("muc.Lup should have been %d, got %d", want, got)
	}
	if want, got := int32(800), muc.Lup32; want != got {
		t.Errorf("muc.Lup should have been %d, got %d", want, got)
	}
	if want, got := int64(800), muc.Lup64; want != got {
		t.Errorf("muc.Lup should have been %d, got %d", want, got)
	}
	if want, got := true, muc.Tup; want != got {
		t.Errorf("muc.Tup should have been %v, got %v", want, got)
	}

	os.Setenv("PUP", "NOTABOOL")
	err := Decode(muc)
	if err == nil {
		t.Error("Bool parse of should of raised error but did not")
	}
}
