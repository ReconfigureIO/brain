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
	"github.com/ReconfigureIO/fixed"
)

// FIXME To be set dynamically
const INP_LAYER_SIZE uint32 = 4
const HID_LAYER_SIZE uint32 = 3
const OUT_LAYER_SIZE uint32 = 3

//essentially a link connecting neurons
type Synapse struct {
	//weight associated with the synapse
	Weight fixed.Int26_6
	//no of the input/output neuron
	In  int
	Out int
}

//FIXME calculate deltas locally per neuron for BP
//TODO Inputs and Outputs useful for a sparse net
type Neuron struct {
	//activation function
	Activation string
	//no of inputs and outputs per neuron
	Inps [4]int
	Outs [4]int
	//for calculating deltas
	DeltaTemp fixed.Int26_6
	//neuron's output
	OutVal fixed.Int26_6
}

//TODO extend to support any activation type
func ActivationFunction(x fixed.Int26_6) fixed.Int26_6 {

	if x > 0 {
		return x
	} else {
		return 0
	}
}

//inference takes an input image and uses the weights from training
//FIXME add bias
//FIXME pass array of layers
func Inference(weights [4][4]Synapse, input [4][4]fixed.Int26_6, network [4][4]Neuron) [3]fixed.Int26_6 {

	var output [OUT_LAYER_SIZE]fixed.Int26_6

	//calculate out values for the first layer (i = 0)
	for j := uint32(0); j < HID_LAYER_SIZE; j++ {
		for i := uint32(0); i < INP_LAYER_SIZE; i++ {
			network[1][j].OutVal += weights[0][i].Weight * input[0][i]
		}
	}
	//use the weights to calculate the output of neurons in hidden layers
	for j := uint32(0); j < OUT_LAYER_SIZE; j++ {
		for i := uint32(0); i < HID_LAYER_SIZE; i++ {
			network[2][j].OutVal += weights[1][i].Weight * network[1][i].OutVal
		}
	}

	//use the weights to calculate the output of neurons in final layer (i = last)
	for i := uint32(0); i < OUT_LAYER_SIZE; i++ {
		output[i] = network[2][i].OutVal
	}
	return output
}

//trains the network of layers based on the input batches
//compares the output based on the test in the dataset
//TODO add bias and weight distributions as input
//TODO pass a pointer to the network
/*func TrainNetwork(image []byte, test []byte, network [][]Neuron) ([][]Synapse, fixed.Int26_6) {

	var accuracy fixed.Int26_6
	var weights [][]Synapse

	//TODO initialise weights using a random function

	//calculate deltas per neuron
	for i := len(network); i >= 0; i-- {
		for j, _ := range network {

			acc := fixed.Int26_6(0)
			for k, _ := range network {
				acc += weights[i+1][k].Weight * network[i][j].DeltaTemp
			}
			network[i][j].DeltaTemp = acc
		}
	}

	//calculate new weights and update
	for i, layer := range network {
		for j, neuron := range layer {
			weights[i][j].Weight += weights[i][j].Weight * neuron.OutVal
		}
	}

	return weights, accuracy
}*/

//reshapes images based on the resize factors should support:
//padding, flipping, rotation, transpose, etc.
//FIXME resize to be implemented as a struct wrt alignment factors
//FIXME implement it as a separate package
func ReshapeImage(image []byte) []byte {
	return image
}
