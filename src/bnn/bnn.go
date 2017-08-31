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
)


type neuron struct {
    //activation function
    act string
    //no of inputs and outputs per neuron
    inps, outs int
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
func NetworkLayer(size int, act string) []neuron{

  layer := make([]neuron, size)

  //init the array
  for i, _:= range layer {

    layer[i].act = act
    layer[i].inps = 0
    layer[i].outs = 0
  }

  return layer
}


