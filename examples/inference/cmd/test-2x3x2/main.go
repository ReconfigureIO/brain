package main

import (
	"encoding/binary"
	"fmt"
	"xcl"
	"os"
//	"github.com/reconfigureio/brain/bnn"
//	"github.com/reconfigureio/brain/utils"
	"github.com/reconfigureio/fixed"
)

//Partition example dataset based on BATCH_SIZE
//NUM_EPOCHS is practical and may vary based on
//the output accuracy achieved from the model    
const NUM_EPOCHS int = 100
const BATCH_SIZE int = 500

func main() {
	world := xcl.NewWorld()
	defer world.Release()

	krnl := world.Import("kernel_test").GetKernel("reconfigure_io_sdaccel_builder_stub_0_1")
	defer krnl.Release()

//	var fpath string
//	input = utils.load_data(fpath)

	//load validations 
//	test := bnn.ReadImage("dataset")
//	fmt.Println(test)

	//reshape image 
//	nw_image:= bnn.ReshapeImage(image)
//	fmt.Println(nw_image)


	// Allocate a buffer on the FPGA to store the return value of our computation
	// The output is a uint32, so we need 4 bytes to store it
	buff := world.Malloc(xcl.WriteOnly, 4)
	defer buff.Free()

	// Pass the arguments to the kernel
	a := fixed.I26(0)
	b := fixed.I26(1)

	// Set the first operand
	krnl.SetArg(0, uint32(a))
	// Set the second operand
	krnl.SetArg(1, uint32(b))
	// Set the pointer to the output buffer
	krnl.SetMemoryArg(2, buff)

	// Run the kernel with the supplied arguments
	krnl.Run(1, 1, 1)

	// Decode that byte slice into the uint32 we're expecting
	var ret fixed.Int26_6
	err := binary.Read(buff.Reader(), binary.LittleEndian, &ret)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	expected := a ^ b 

	// Exit with an error if the value is not correct
	if expected != ret {
		// Print the value we got from the FPGA
		fmt.Printf("Expected %d, got %d\n", expected, ret)
		os.Exit(1)
	}
}
