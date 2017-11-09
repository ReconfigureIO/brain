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
//	addrWH uintptr,
	addrBH uintptr,
//	addrWO uintptr,
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

	// From the training stage of datadan.io network model
	weightH := [12]fixed.Int26_6{fixed.I26F(-9 , 664649023 << 0),
		fixed.I26F(-3 , 331748963 << 0),
		fixed.I26F(0 , 479873455 >> 1),
		fixed.I26F(-9 , 16865031 << 0),
		fixed.I26F(7 , 526678142 << 0),
		fixed.I26F(-11 , 25986268 << 0),
		fixed.I26F(35 , 375442386 << 0),
		fixed.I26F(-8 , 377651024 << 0),
		fixed.I26F(0 , 585733147 << 0),
		fixed.I26F(14 , 56994879 << 0),
		fixed.I26F(-6 , 97787149 << 0),
		fixed.I26F(-3 , 59583457 << 0)}

	weightO := [9]fixed.Int26_6{fixed.I26F(-6 , 421059674 << 0),
		fixed.I26F(10 , 430255115 << 0),
		fixed.I26F(-10 , 466201644 << 0),
		fixed.I26F(12 , 384611656 << 0),
		fixed.I26F(-5 , 610451231 << 0),
		fixed.I26F(-11 , 310357612 << 0),
		fixed.I26F(-12 , 155845492 << 0),
		fixed.I26F(-4 , 313406702 << 0),
		fixed.I26F(1 , 379336036 << 0)}


	// Since we're not reading anything from memory, disable those reads
	//go axiprotocol.ReadDisable(memReadAddr, memReadData)


	//Read in the first input batch
	for i := 0; i < INP_LAYER_SIZE ; i++{
	 
	 layer_in[i] = fixed.Int26_6(aximemory.ReadUInt32(memReadAddr, memReadData, false, addrIn + uintptr(4*i)))
	}

	//Calculate outvals for the hidden layer
	for i := 0; i < HID_LAYER_SIZE ; i++{
	 

		p0 := layer_in[0] * weightH[0 + i] //i + HID_LAYER_SIZE * i
		p1 := layer_in[1] * weightH[3 + i] 
		p2 := layer_in[2] * weightH[6 + i] 
		p3 := layer_in[3] * weightH[9 + i]

		// Add corresponding Bias
		bias := fixed.Int26_6(aximemory.ReadUInt32(memReadAddr, memReadData, false, addrBH + uintptr(4*i)))

		// Calculate biased sum of products per neuron in hidden layer
		out := p0 + p1 + p2 + p3 + bias			
		
		// Apply Sigmoid function 
		layer_hidden[i] = fixed.Int26_6(aximemory.ReadUInt32(memReadAddr, memReadData, false, addrAct + uintptr(4*out)))
 	}

	//Calculate outval for the output layer
	for i := 0; i < OUT_LAYER_SIZE ; i++{
	 
		p0 := layer_hidden[0] * weightO[0 + i] //i + OUT_LAYER_SIZE * i
		p1 := layer_hidden[1] * weightO[3 + i]
		p2 := layer_hidden[2] * weightO[6 + i]

		// Add corresponding Bias
		bias := fixed.Int26_6(aximemory.ReadUInt32(memReadAddr, memReadData, false, addrBO + uintptr(4*i)))

		// Calculate biased sum of products per neuron in hidden layer
		out := p0 + p1 + p2 + bias			
		
		// Apply Sigmoid function 
		layer_out[i] = fixed.Int26_6(aximemory.ReadUInt32(memReadAddr, memReadData, false, addrAct + uintptr(4*out)))
	}

	// Write it back to the pointer the host requests
	aximemory.WriteUInt32(
		memWriteAddr, memWriteData, memWriteResp, false, addrOut, layer_out)
}
