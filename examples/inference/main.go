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

import (
	"github.com/reconfigureio/brain/bnn"
	"github.com/reconfigureio/brain/utils"
)


const INP_LAYER_SIZE int = 2
const HID_LAYER_SIZE int = 3
const OUT_LAYER_SIZE int = 2

func TOP(
	// The first set of arguments will be the ports for interacting with host 
	output float32,
	// The second set of arguments will be the ports for interacting with memory
	memReadAddr chan<- axiprotocol.Addr,
	memReadData <-chan axiprotocol.ReadData,

	memWriteAddr chan<- axiprotocol.Addr,
	memWriteData chan<- axiprotocol.WriteData,
	memWriteResp <-chan axiprotocol.WriteResp){

	//cast rawdate to input vars
	training_data := [][]float32{
    		{0, 0},
    		{0, 1},
		{1, 0},
		{1, 1}
	}
	target_data := []float32{
    		{0},
    		{1},
		{1},
		{0}
	}
	test_data := [][]float32{
    		{0, 1},
    		{1, 1},
		{1, 0},
		{1, 1}
	}

	//build a network with 3 layers of input, hidden, and output
	layer_in := bnn.NetworkLayer(INP_LAYER_SIZE,"relu")
	layer_hidden := bnn.NetworkLayer(HID_LAYER_SIZE,"relu")
	layer_out := bnn.NetworkLayer(OUT_LAYER_SIZE,"sig")

	network := [][]bnn.Neuron{layer_in, layer_hidden, layer_out}

	//train network and return accuracy
	//FIXME add initial weight and bias distribution
	weights, acc := bnn.TrainNetwork(training_data, target_data, network)

	//inference uses the updated weights, and finally returns an array with outputs 
	output := bnn.Inference(weights, test_data, network)

}
