package normalize

import "testing"

func TestAlnum(t *testing.T) {
	result := Alnum(` @h3LLo !|   W0#Rld  `)
	expected := `h3llo w0rld`
	if result != expected {
		t.Errorf("result %q, expect %q", result, expected)
	}
}
