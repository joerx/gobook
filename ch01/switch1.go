package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  heads, tails, edge := 0, 0, 0
  throws := 10
  r := rand.New(rand.NewSource(time.Now().UnixNano()))

  for i := 0; i < throws; i++ {
    result := coinflip(r)
    fmt.Println(result)

    switch result {
    case "heads":
      heads++;
    case "tails":
      tails++;
    default:
      edge++;
    }
  }

  fmt.Printf("heads: %d, tails: %d, edge: %d\n", heads, tails, edge)
}

func coinflip(r *rand.Rand) string {
  val := r.Float32()
  switch {
  case val > 0.5:
    return "heads"
  case val < 0.5:
    return "tails"
  default:
    return "edge"
  }
}
