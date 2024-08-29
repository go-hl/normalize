package normalize

import "testing"

func TestAlnumRhuanPatriky(t *testing.T) {
	result := Alnum(`  r!Hu4n$ % Pa(TR!)1ky`)
	expected := `rhu4n patr1ky`
	if result != expected {
		t.Errorf("result %q, expect %q", result, expected)
	}
}

func TestAlnumHelloWorld(t *testing.T) {
	result := Alnum(` @h3LLo !|   W0#Rld  `)
	expected := `h3llo w0rld`
	if result != expected {
		t.Errorf("result %q, expect %q", result, expected)
	}
}
