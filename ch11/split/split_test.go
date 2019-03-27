package split

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	var words []string

	tests := []struct {
		want int
		s    string
		sep  string
	}{
		{3, ":", "a:b:c"},
		{1, ":", ""},
		{1, ":", "a"},
	}

	for _, test := range tests {
		words = strings.Split(test.sep, test.s)
		if got, want := len(words), test.want; got != want {
			t.Errorf("Split(%q, %q), want %d words, got %d", test.s, test.sep, test.want, got)
		}
	}
}
