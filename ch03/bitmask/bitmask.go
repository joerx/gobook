package bitmask

import (
  "strconv"
)

const Read byte = 2
const Write byte = 1
const Exec byte = 0

const User byte = 6
const Group byte = 3
const Other byte = 0

func Bitmask(mask string) int {
  val, err := strconv.ParseInt(mask, 8, 16)
  if err != nil {
    panic("Invalid bitmask")
  }
  return int(val)
}

func IsSet(mask int, role byte, flag byte) bool {
  return mask&(1<<(role+flag)) > 0
}
