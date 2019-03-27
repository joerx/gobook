package bitmask

import (
  "strconv"
  "strings"
)

const Read byte = 2
const Write byte = 1
const Exec byte = 0

const User byte = 6
const Group byte = 3
const Other byte = 0

func Bitmask(mask string) uint16 {
  arr := strings.Split(mask, "")
  res := uint16(0)

  for i := 0; i < 3; i++ {
    val, err := strconv.ParseInt(arr[i], 10, 8)
    if err != nil {
      return 0 // FIXME: handle this shit (return error)
    }
    res |= uint16(val) << byte((2-i)*3)
  }

  return res
}

func IsSet(mask uint16, role byte, flag byte) bool {
  return mask&(1<<(role+flag)) > 0
}
