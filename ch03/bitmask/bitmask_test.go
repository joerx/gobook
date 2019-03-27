package bitmask

import (
  "testing"
  "strconv"
)

func TestBitmask(t *testing.T) {
  tests := []struct {
    bitmask string
    want string
  } {
    {"777", "111111111"},
    {"775", "111111101"},
    {"600", "110000000"},
    {"644", "110100100"},
  }

  for _, test := range tests {
    if got := Bitmask(test.bitmask); strconv.FormatInt(int64(got), 2) != test.want {
      t.Errorf("Bitmask(%s), got %09b, want %s", test.bitmask, got, test.want)
    }
  }
}

func TestIsSet(t *testing.T) {
  tests := []struct {
    bitmask string
    role byte
    flag byte
    want bool
  } {
    {"777", User, Read, true},
    {"770", Other, Read, false},
    {"770", Other, Write, false},
    {"770", Other, Exec, false},
    {"644", User, Exec, false},
    {"644", User, Write, true},
    {"644", Group, Read, true},
    {"644", Group, Write, false},
  }

  for _, test := range tests {
    if got := IsSet(Bitmask(test.bitmask), test.role, test.flag); got != test.want {
      t.Errorf("IsSet(%q, %d, %d), got %t, want %t", test.bitmask, test.role, test.flag, got, test.want)
    }
  }
}
