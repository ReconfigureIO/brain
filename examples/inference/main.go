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
	 
		inp0 := layer_in[0] * weights_h[i]
		inp1 := layer_in[1] * weights_h[i]
		out := inp0 + inp1 
 		
		layer_hidden[i] = fixed.Int26_6(aximemory.ReadUInt32(memReadAddr, memReadData, false, addrAct + uintptr(4*out)))
 	}
	//Calculate outval for the output layer
	sum := fixed.Int26_6(0)
	for i := 0; i < OUT_LAYER_SIZE ; i++{
	 
		sum += layer_hidden[i]
	}
	
	layer_out[0] = fixed.Int26_6(aximemory.ReadUInt32(memReadAddr, memReadData, false, addrAct + uintptr(4 * uint8(sum * weights_o))))

	output := layer_out[0]

	// Write it back to the pointer the host requests
	aximemory.WriteUInt32(
		memWriteAddr, memWriteData, memWriteResp, false, addrOut, uint32(output))
}
