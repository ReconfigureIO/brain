/*
Copyright 2017 Reconfigure.io Ltd. All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import "fmt"
import "bnn"
//import "math/rand"

import "github.com/sjwhitworth/golearn/base"


const INP_LAYER_SIZE int = 2
const HID_LAYER_SIZE int = 3
const OUT_LAYER_SIZE int = 2

func main() {

   //read data from dataset 
   rawData, err := base.ParseCSVToInstances("datasets/mnist_test.csv", false)
   if err != nil {
      panic(err)
   }
   fmt.Println(rawData)

  //cast rawdate to input var
  var input [][]float32

  //build a network with 3 layers of input, hidden, and output
  layer_in := bnn.NetworkLayer(INP_LAYER_SIZE,"sig")
  layer_hidden := bnn.NetworkLayer(HID_LAYER_SIZE,"sig")
  layer_out := bnn.NetworkLayer(OUT_LAYER_SIZE,"relu")

  network := [][]bnn.Neuron{layer_in, layer_hidden, layer_out}
  fmt.Println(network)

  //load image 
  image := bnn.ReadImage("dataset")
  fmt.Println(image)

  //load validations 
  test := bnn.ReadImage("dataset")
  fmt.Println(test)

  //reshape image 
  nw_image:= bnn.ReshapeImage(image)
  fmt.Println(nw_image)

  //train network and return accuracy
  //FIXME add initial weight and bias distribution
  weights, acc := bnn.TrainNetwork(nw_image, test, network)
  fmt.Println(acc, weights)

  //inference uses the updated weights, and finally returns an array with outputs 
  output := bnn.Inference(weights, input, network)
  fmt.Println(output)
}
