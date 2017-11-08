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

/*
// TODO Use Xilinx's full precision LogiCore IP
// for math.Exp. For now a lookup table is imported instead.
func sigmoid_act(x fixed.Int26_6) fixed.Int26_6{

  //generated table by util/disretise_sig(1)
  prime := [16]fixed.Int26_6{

/*fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 
		fixed.I26F(0 , 0 << 0), 

		fixed.I26F(0 , 2 << 0), 
		fixed.I26F(0 , 6 << 0), 
		fixed.I26F(0 , 16 << 0), 
		fixed.I26F(0 , 45 << 0), 
		fixed.I26F(0 , 123 << 0), 
		fixed.I26F(0 , 335 << 0), 
		fixed.I26F(0 , 911 << 0), 
		fixed.I26F(0 , 2472 << 0), 
		fixed.I26F(0 , 6692 << 0), 
		fixed.I26F(0 , 17986 << 0), 
		fixed.I26F(0 , 47425 << 0), 
		fixed.I26F(0 , 119202 << 0), 
		fixed.I26F(0 , 268941 << 0), 
		fixed.I26F(0 , 500000 << 0), 
		fixed.I26F(0 , 731058 << 0), 
		fixed.I26F(0 , 880797 << 0)}
/* 
		fixed.I26F(0 , 952574 << 0), 
		fixed.I26F(0 , 982013 << 0), 
		fixed.I26F(0 , 993307 << 0), 
		fixed.I26F(0 , 997527 << 0), 
		fixed.I26F(0 , 999088 << 0)*/

/*, 
		fixed.I26F(0 , 999664 << 0), 
		fixed.I26F(0 , 999876 << 0), 
		fixed.I26F(0 , 999954 << 0), 
		fixed.I26F(0 , 999983 << 0), 
		fixed.I26F(0 , 999993 << 0), 
		fixed.I26F(0 , 999997 << 0), 
		fixed.I26F(0 , 999999 << 0), 
		fixed.I26F(0 , 999999 << 0), 
		fixed.I26F(0 , 999999 << 0), 
		fixed.I26F(0 , 999999 << 0), 
		fixed.I26F(0 , 999999 << 0), 
		fixed.I26F(0 , 999999 << 0), 
		fixed.I26F(0 , 999999 << 0), 
		fixed.I26F(0 , 999999 << 0), 
		fixed.I26F(0 , 999999 << 0), 
		fixed.I26F(0 , 999999 << 0), 
		fixed.I26F(0 , 999999 << 0), 
		fixed.I26F(0 , 999999 << 0), 
		fixed.I26F(0 , 999999 << 0), 
		fixed.I26F(0 , 999999 << 0), 
		fixed.I26F(0 , 999999 << 0), 
		fixed.I26F(0 , 999999 << 0), 
		fixed.I26F(0 , 999999 << 0), 
		fixed.I26F(0 , 999999 << 0), 
		fixed.I26F(0 , 999999 << 0), 
		fixed.I26F(0 , 999999 << 0), 
		fixed.I26F(0 , 999999 << 0), 
		fixed.I26F(0 , 999999 << 0), 
		fixed.I26F(0 , 999999 << 0)
} 


		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0), 
		fixed.I26F(1 , 0 << 0)

        // Convert x (26_6) to index
	// add by (200*z)/2 index  
	return prime[uint8(x) + 100]

	if (uint8(x) + 100) < 87 {

		return 0

	} else if (uint8(x) + 100) < 137 {

		return prime[uint8(x) + 100 - 87]

	} else {

		return 1	
	}

}*/

func Top(
	addrIn uintptr,
	addrAct uintptr,
	addrOut uintptr,
	// The first set of arguments will be the ports for interacting with host 
	//output fixed.Int26_6,
	// The second set of arguments will be the ports for interacting with memory
	memReadAddr chan<- axiprotocol.Addr,
	memReadData <-chan axiprotocol.ReadData,

	memWriteAddr chan<- axiprotocol.Addr,
	memWriteData chan<- axiprotocol.WriteData,
	memWriteResp <-chan axiprotocol.WriteResp){


	//cast rawdate to input vars
/*	training_data := [][]fixed.Int26_6{
    		 []fixed.Int26_6{0, 0},
    		 []fixed.Int26_6{0, 1},
		 []fixed.Int26_6{1, 0},
		 []fixed.Int26_6{1, 1}}

	target_data := []fixed.Int26_6{
    		 0,
    		 1,
		 1,
		 0}
*/
/*	test_data := [4]fixed.Int26_6{
    		 fixed.Int26_6{0, 1},
    		 fixed.Int26_6{1, 1},
		 fixed.Int26_6{1, 0},
		 fixed.Int26_6{0, 0}}

	acc_data := []fixed.Int26_6{
    		 1,
    		 0,
		 1,
		 0}
*/
	//weights exported from xornet on KERAS (epoch size = 500 - sgd)
/*	weights := [][]fixed.Int26_6{
 		[]fixed.Int26_6{-0.35589939,
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
       		  0.19469711},[]fixed.Int26_6{-0.08784658}}*/

	//weights exported from xornet on KERAS (epoch size = 5000 - adam)
	weights_h := [16]fixed.Int26_6{
		 fixed.I26F(0, 1726144),
		 fixed.I26F(-2, 10709),
       		 fixed.I26F(0, 43040475),
	         fixed.I26F(-0, 36798), //?
       		 fixed.I26F(-2 ,14761877),
       		 fixed.I26F(1 ,65221334),
       		 fixed.I26F(-0, 47918937),
       		 fixed.I26F(-2 ,28618431),
      		 fixed.I26F(-1, 64216483),
       		 fixed.I26F(1, 45400071),
       		 fixed.I26F(0, 8930543), //?
       		 fixed.I26F(-1,85224831),
       		 fixed.I26F(1,3171016),
     		 fixed.I26F(-1,74173605),
       		 fixed.I26F(-0,37978798),
       		 fixed.I26F(-2,9490085)}

	weights_o := fixed.I26F(0, 46938747) //09490085??

	acts := [200]fixed.Int26_6{0}

	ch := make(chan uint32, 2)

	//build a network with 3 layers of input, hidden, and output
//	layer_in := NetworkLayer(INP_LAYER_SIZE,"relu")
//	layer_hidden := NetworkLayer(HID_LAYER_SIZE,"relu")
//	layer_out := NetworkLayer(OUT_LAYER_SIZE,"sig")

	var layer_in [INP_LAYER_SIZE]fixed.Int26_6  //"relu"
	var layer_hidden [HID_LAYER_SIZE]fixed.Int26_6 //"relu")
	var layer_out [OUT_LAYER_SIZE]fixed.Int26_6 //"sig"



	// Since we're not reading anything from memory, disable those reads
//	go axiprotocol.ReadDisable(memReadAddr, memReadData)

	// Write it back to the pointer the host requests
	
	//Initialize the first layer
//	layer_in[0] = fixed.Int26_6(aximemory.ReadUInt32(memReadAddr, memReadData, false, addrIn))
//	layer_in[1] = fixed.Int26_6(aximemory.ReadUInt32(memReadAddr, memReadData, false, addrIn+4))

        aximemory.ReadBurstUInt32(memReadAddr, memReadData, false, addrAct, 2, ch )

	for i := 0; i < 2 ; i++{
	 
	 acts[i] = fixed.Int26_6(<-ch)
	}


	for i := 0; i < INP_LAYER_SIZE ; i++{
	 
	 layer_in[i] = fixed.Int26_6(aximemory.ReadUInt32(memReadAddr, memReadData, false, addrIn + uintptr(4*i)))
	}

	//Calculate outvals for the hidden layer
	for i := 0; i < HID_LAYER_SIZE ; i++{
	 
		inp0 := layer_in[0] * weights_h[i]
		inp1 := layer_in[1] * weights_h[i]
		out := inp0 + inp1 
 		//FIXME add activations - relu
		layer_hidden[i] = out  
 	}
	//Calculate outval for the output layer
	sum := fixed.Int26_6(0)
	for i := 0; i < OUT_LAYER_SIZE ; i++{
	 
		sum += layer_hidden[i]
	}
	//FIXME add activations - sig
	 
	layer_out[0] = acts[uint8(sum * weights_o)]

	output := layer_out[0]

	// Write it back to the pointer the host requests
	aximemory.WriteUInt32(
		memWriteAddr, memWriteData, memWriteResp, false, addrOut, uint32(output))
}
