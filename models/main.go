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
//	"github.com/reconfigureio/brain/bnn"
//	"github.com/reconfigureio/brain/utils"
	"github.com/reconfigureio/fixed"
	"fmt"
	"math"

)

const INP_LAYER_SIZE int = 2
const HID_LAYER_SIZE int = 16
const OUT_LAYER_SIZE int = 1

//TODO Functions to be generalized and moved to bnn package
//essentially a link connecting neurons 
type Synapse struct {
    //weight associated with the synapse
    Weight      fixed.Int52_12
    //no of the input/output neuron
    In, Out     int
}
//FIXME calculate deltas locally per neuron for BP
//TODO Inputs and Outputs useful for a sparse net
type Neuron struct {
    //activation function
    Activation  string
    //no of inputs and outputs per neuron
    Inps, Outs  []int
    //for calculating deltas
    DeltaTemp   fixed.Int52_12
    //neuron's output
    OutVal      fixed.Int52_12
    //outVal= activation (outVal from previous layer * in_weights)  
    In_wights   []fixed.Int52_12
}

//constructs a layer of neurons with arbitrary 'size' and 'activation' functions
func NetworkLayer(size int, act string) []Neuron{

  layer := make([]Neuron, size)

  //init the array
  for i, _:= range layer {

    layer[i].Activation = act
  }
  return layer
}

//TODO extend to support any activation type
func Activations(act string, x fixed.Int52_12) fixed.Int52_12{

   switch act{
     case "relu":
	return fixed.I52(int64(math.Max(0,float64(x))))
     case "sig":
        return fixed.I52(int64(float64(x)/(1 + math.Abs(float64(x)))))
//1.0 / (1.0 + math.Exp(-x))
     default:
	return fixed.Int52_12(0)
   }
}

//TODO extend to support other activation' for BackPropagation
func Activations_(act string, x fixed.Int52_12) fixed.Int52_12{

   switch act{
     case "relu":
	if x > 0 {return 1} else {return 0} 
     case "sig":
        return x * (1.0 - x)
     default:
	return fixed.Int52_12(0)
   }
}


//inference takes an input image and uses the weights from training  
//FIXME add bias
func Inference(ptr *[][]Neuron, tdata []fixed.Int52_12, weights [][]fixed.Int52_12) fixed.Int52_12{

   //Initialize the first layer
   for i, _ := range (*ptr)[0] {
 
	(*ptr)[0][i].OutVal = tdata[i]
   }

   //Calculate outvals for the hidden layer
   for i, _ := range (*ptr)[1] {
 
	inp0 := (*ptr)[0][0].OutVal * weights[0][i]
	inp1 := (*ptr)[0][1].OutVal * weights[0][i]
        (*ptr)[1][i].OutVal = Activations((*ptr)[1][0].Activation, inp0 + inp1)
   }

   //Calculate outval for the output layer
   sum := fixed.Int52_12(0)
   for i, _ := range (*ptr)[1] {
      sum += (*ptr)[1][i].OutVal
   }
   (*ptr)[2][0].OutVal = Activations((*ptr)[2][0].Activation, sum * weights[1][0])

   output :=  (*ptr)[2][0].OutVal

   return output
}

func main(
	// The first set of arguments will be the ports for interacting with host 
	//output fixed.Int26_6,
	// The second set of arguments will be the ports for interacting with memory
//	memReadAddr chan<- axiprotocol.Addr,
//	memReadData <-chan axiprotocol.ReadData,

//	memWriteAddr chan<- axiprotocol.Addr,
//	memWriteData chan<- axiprotocol.WriteData,
//	memWriteResp <-chan axiprotocol.WriteResp
){

	//cast rawdate to input vars
	training_data := [][]fixed.Int52_12{
    		 {0, 0},
    		 {0, 1},
		 {1, 0},
		 {1, 1}}
	target_data := []fixed.Int52_12{
    		 0,
    		 1,
		 1,
		 0}
	test_data := [][]fixed.Int52_12{
    		 {0, 1},
    		 {1, 1},
		 {1, 0},
		 {0, 0}}
	acc_data := []fixed.Int52_12{
    		 1,
    		 0,
		 1,
		 0}

	//weights exported from xornet on KERAS (epoch size = 500 - sgd)
/*	weights := [][]fixed.Int52_12{
 		[]fixed.Int52_12{-0.35589939,
       		  0.13612342,
       		 -0.27676189,
       		 -0.06193029,
       		 -0.37450755,
       		  0.48630142,
       		  0.40621114,
       		  0.11644399,
       		 -0.33843306,
       		  0.34775987,
       		 -0.14313582,
       		 -0.04034447,
       		  0.54061526,
       		 -0.42877936,
       		  0.54952145,
       		  0.19469711},[]fixed.Int52_12{-0.08784658}}*/

	//weights exported from xornet on KERAS (epoch size = 5000 - adam)
	weights := [][]fixed.Int52_12{
		[]fixed.Int52_12{fixed.I52F(0, 1726144),
		 fixed.I52F(-2, 10709),
       		 fixed.I52F(0, 43040475),
	         fixed.I52F(-0, 36798), //?
       		 fixed.I52F(-2 ,14761877),
       		 fixed.I52F(1 ,65221334),
       		 fixed.I52F(-0, 47918937),
       		 fixed.I52F(-2 ,28618431),
      		 fixed.I52F(-1, 64216483),
       		 fixed.I52F(1, 45400071),
       		 fixed.I52F(0, 8930543), //?
       		 fixed.I52F(-1,85224831),
       		 fixed.I52F(1,3171016),
     		 fixed.I52F(-1,74173605),
       		 fixed.I52F(-0,37978798),
       		 fixed.I52F(-2,9490085)},[]fixed.Int52_12{fixed.I52F(0, 46938747)}} //09490085??

	//build a network with 3 layers of input, hidden, and output
	layer_in := NetworkLayer(INP_LAYER_SIZE,"na")
	layer_hidden := NetworkLayer(HID_LAYER_SIZE,"relu")
	layer_out := NetworkLayer(OUT_LAYER_SIZE,"sig")

	//Ignore training phase for now
	var _ = training_data
	var _ = target_data
	//.Print(training_data, target_data, test_data)

	network := [][]Neuron{layer_in, layer_hidden, layer_out}

	//train the network 
	//FIXME add initial weight and bias distribution
	//weights, acc := bnn.TrainNetwork(training_data, target_data, network)
	loss := fixed.Int52_12(0)
	for i, _ := range test_data {
		//Prediction/Inference based on the input test dataset 
	        ret := Inference(&network, test_data[i], weights)
		loss += fixed.Int52_12(math.Abs(float64(ret  - acc_data[i]))) 
	}

        //Output the Accuracy value to standard out.
	acc := 100 * (1-loss/4)/1
	fmt.Printf("Accuracy is : %F% \n\n", float64(acc))
}
