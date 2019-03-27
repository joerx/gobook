package bitmask

import (
  "testing"
)

func TestBitmask(t *testing.T) {
  tests := []struct {
    bitmask string
    want int
  } {
    {"777", 0777},
    {"775", 0775},
    {"600", 0600},
    {"644", 0644},
  }

  for _, test := range tests {
    if got := Bitmask(test.bitmask); got != test.want {
      t.Errorf("Bitmask(%s), got %o, want %o", test.bitmask, got, test.want)
    }
  }
}

func TestIsSet(t *testing.T) {
  tests := []struct {
    bitmask int
    role byte
    flag byte
    want bool
  } {
    {0777, User, Read, true},
    {0770, Other, Read, false},
    {0770, Other, Write, false},
    {0770, Other, Exec, false},
    {0644, User, Exec, false},
    {0644, User, Write, true},
    {0644, Group, Read, true},
    {0644, Group, Write, false},
  }

  for _, test := range tests {
    if got := IsSet(test.bitmask, test.role, test.flag); got != test.want {
      t.Errorf("IsSet(%o, %d, %d), got %t, want %t", test.bitmask, test.role, test.flag, got, test.want)
    }
  }
}
