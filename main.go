package main

import "fmt"
import "gorand/rand"

func main() {
  rnd := rand.Rand{}

  rnd.Powerup()
  
  b := 3
  for i:=0; i<1000; i++{
    b = rnd.GetBit()
    fmt.Println(b)
  }
  rnd.Shutdown()
}

