package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEcho(t *testing.T) {
	tests := []struct{
		newline bool
		sep string
		args []string
		want string
	} {
		{true, "", []string{}, "\n"},
		{false, "", []string{}, ""},
		{false, "", []string{"hello", "world"}, "helloworld"},
		{false, ",", []string{"hello", "world"}, "hello,world"},
		{false, " ", []string{"hello", "world"}, "hello world"},
		{false, "\n", []string{"hello", "world"}, "hello\nworld"},
		{true, "\n", []string{"hello", "world"}, "hello\nworld\n"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("echo(%t, %q, %q)", test.newline, test.sep, test.args)
		out = new(bytes.Buffer)

		if err := echo(test.newline, test.sep, test.args); err != nil {
			t.Errorf("%s failed: %v", descr, err)
		}

		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
