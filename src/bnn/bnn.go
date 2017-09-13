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

package bnn

import (
    "os"
//    "math/rand"
)

//essentially a link connecting neurons 
type Synapse struct {
    //weight associated with the synapse
    Weight      float32
    //no of the input/output neuron
    In, Out     int
}

//FIXME calculate deltas locally per neuron for BP
type Neuron struct {
    //activation function
    Activation  string
    //no of inputs and outputs per neuron
    Inps, Outs  []int
    //for calculating deltas
    DeltaTemp   float32
    //neuron's output
    OutVal      float32
}

//inference takes an input image and uses the weights from training  
//FIXME add bias
//FIXME pass array of layers  
//func Inference(weights []float32, input []byte, layers [][]Neuron) []byte{
// output = input * layers * weights
// return output
//}

//trains the network of layers based on the input batches
//compares the output based on the test in the dataset 
//FIXME add bias and weight distributions as input 
//FIXME pass a pointer to the network
func TrainNetwork(image []byte, test []byte, network [][]Neuron) ([][]float32, float32){

 var accuracy float32
 var weights [][]float32

}

//reshapes images based on the resize factors should support:
//padding, flipping, rotation, transpose, etc.
//FIXME resize to be implemented as a struct wrt alignment factors
//FIXME implement it as a separate package 
func ReshapeImage(image []byte) []byte{
 return image
}

//reads in images located in 'path' and returns an array 
func ReadImage(path string) []byte{

   //open the image file
   f, err := os.Open(path)
   if err != nil {
        panic(err)
   }

   //get the file status
   fi, err := f.Stat()
   if err != nil {
       panic(err)
   }

   //create an arraye of size 'image'
   arr := make([]byte, fi.Size())
   f.Read(arr)

   f.Close()
   return arr
}

//constructs a layer of neurons with arbitrary 'size' and 'activation' functions
func NetworkLayer(size int, act string) []Neuron{

  layer := make([]Neuron, size)

  //init the array
  for i, _:= range layer {

    layer[i].act = act
    layer[i].inps = 0
    layer[i].outs = 0
  }

  return layer
}
