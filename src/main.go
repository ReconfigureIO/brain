package main

import "fmt"
import "bnn"

func main() {
  layer1 := bnn.NetworkLayer(2,"relu")
  fmt.Println(layer1[1])
}
