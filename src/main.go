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

func main() {

  //create a layer with 2 neurons and 'relu' activations 
  layer1 := bnn.NetworkLayer(2,"relu")
  layer2 := bnn.NetworkLayer(3,"sig")
  layers := [][]bnn.Neuron{layer1, layer2}
  fmt.Println(layers)

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
  weights, acc := bnn.TrainNetwork(nw_image, test, layers)
  fmt.Println(acc)
  
  //inference uses the updated weights, and finally returns an array with outputs 
  output := bnn.Inference(weights, input, layers)
  fmt.Println(output)
}
