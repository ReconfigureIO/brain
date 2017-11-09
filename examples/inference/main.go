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

	// Import the entire framework (including bundled verilog)
	_ "sdaccel"

	// Use the new AXI protocol package
	aximemory "axi/memory"
	axiprotocol "axi/protocol"

//	"github.com/reconfigureio/add"
)

const INP_LAYER_SIZE int = 4
const HID_LAYER_SIZE int = 3
const OUT_LAYER_SIZE int = 3

func Top(
	addrAct uintptr,
	addrIn uintptr,
	addrWH uintptr,
	addrBH uintptr,
	addrWO uintptr,
	addrBO uintptr,
	addrOut uintptr,

	// The first set of arguments will be the ports for interacting with host 
	//output fixed.Int26_6,
	// The second set of arguments will be the ports for interacting with memory
	memReadAddr chan<- axiprotocol.Addr,
	memReadData <-chan axiprotocol.ReadData,

	memWriteAddr chan<- axiprotocol.Addr,
	memWriteData chan<- axiprotocol.WriteData,
	memWriteResp <-chan axiprotocol.WriteResp){

	//build a network with 3 layers of input, hidden, and output
//	layer_in := NetworkLayer(INP_LAYER_SIZE,"relu")
//	layer_hidden := NetworkLayer(HID_LAYER_SIZE,"relu")
//	layer_out := NetworkLayer(OUT_LAYER_SIZE,"sig")

	var layer_in [INP_LAYER_SIZE]fixed.Int26_6  
	var layer_hidden [HID_LAYER_SIZE]fixed.Int26_6 
	var layer_out [OUT_LAYER_SIZE]fixed.Int26_6 


	// Since we're not reading anything from memory, disable those reads
	//go axiprotocol.ReadDisable(memReadAddr, memReadData)


	//Read in the first input batch
	for i := 0; i < INP_LAYER_SIZE ; i++{
	 
	 layer_in[i] = fixed.Int26_6(aximemory.ReadUInt32(memReadAddr, memReadData, false, addrIn + uintptr(4*i)))
	}

	//Calculate outvals for the hidden layer
	for i := 0; i < HID_LAYER_SIZE ; i++{
	 

		p0 := layer_in[0] * weights_h[i]
		p1 := layer_in[1] * weights_h[i]
		p2 := layer_in[2] * weights_h[i]
		p3 := layer_in[3] * weights_h[i]

		// Add corresponding Bias
		bias := fixed.Int26_6(aximemory.ReadUInt32(memReadAddr, memReadData, false, addrBH + uintptr(4*i)))

		// Calculate biased sum of products per neuron in hidden layer
		out := p0 + p1 + p2 + p3 + bias			
		
		// Apply Sigmoid function 
		layer_hidden[i] = fixed.Int26_6(aximemory.ReadUInt32(memReadAddr, memReadData, false, addrAct + uintptr(4*out)))
 	}

	//Calculate outval for the output layer
	for i := 0; i < OUT_LAYER_SIZE ; i++{
	 
		p0 := layer_hidden[0] * weights_o[i]
		p1 := layer_hidden[1] * weights_o[i]
		p2 := layer_hidden[2] * weights_o[i]

		// Add corresponding Bias
		bias := fixed.Int26_6(aximemory.ReadUInt32(memReadAddr, memReadData, false, addrBO + uintptr(4*i)))

		// Calculate biased sum of products per neuron in hidden layer
		out := p0 + p1 + p2 + bias			
		
		// Apply Sigmoid function 
		layer_out[i] = fixed.Int26_6(aximemory.ReadUInt32(memReadAddr, memReadData, false, addrAct + uintptr(4*out)))
	}

	// Write it back to the pointer the host requests
	aximemory.WriteUInt32(
		memWriteAddr, memWriteData, memWriteResp, false, addrOut, uint32({layer_out[2],layer_out[1],layer_out[0]}))
}
