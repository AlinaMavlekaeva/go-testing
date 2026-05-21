package normalize

import (
	"testing"
)

func TestClean(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		output string
	}{
		{"empty", "", ""},
		{"spaces", "  hello   ", "hello"},
		{"tabs", "  hello		hexlet 	", "hello hexlet"},
		{"case", "HellO", "hello"},
		{"normal", "normalized", "normalized"},
		{"symbols", " Hello, world!", "hello, world!"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := Clean(c.input)
			if got != c.output {
				t.Errorf("Clean(%s): got = %s, want = %s", c.input, got, c.output)
			}
		})
	}
}
