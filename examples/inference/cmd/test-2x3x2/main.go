package main

import (
	"encoding/binary"
	"fmt"
	"xcl"
	"os"
        "testing"
//	"github.com/reconfigureio/brain/bnn"
//	"github.com/reconfigureio/brain/utils"
	"github.com/reconfigureio/fixed"
)

//Partition example dataset based on BATCH_SIZE
//NUM_EPOCHS is practical and may vary based on
//the output accuracy achieved from the model    
const NUM_EPOCHS int = 100
const BATCH_SIZE int = 500

func BenchmarkKernel(world xcl.World, krnl *xcl.Kernel, B *testing.B, buffIn *xcl.Memory, buffOut *xcl.Memory) {



/*	// Set the first operand
	krnl.SetArg(0, uint32(a))
	// Set the second operand
	krnl.SetArg(1, uint32(b))
*/

	// Set the pointer to the output buffer
	krnl.SetMemoryArg(0, buffIn)
	// Set the pointer to the output buffer
	krnl.SetMemoryArg(1, buffOut)

	// Reset the timer so that we only measure runtime of the kernel
	B.ResetTimer()
	krnl.Run(1, 1, 1)
}

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

	inp := []fixed.Int26_6{0, 1}
//	inpSize := binary.Size(inp)

        // Allocate a buffer on the FPGA to store the return value of our computation
        // The output is a uint32, so we need 4 bytes to store it
        buffIn := world.Malloc(xcl.ReadOnly, 8)
        defer buffIn.Free()

        // Allocate a buffer on the FPGA to store the return value of our computation
        // The output is a uint32, so we need 4 bytes to store it
        buffOut := world.Malloc(xcl.WriteOnly, 4)
        defer buffOut.Free()

/*
	// Set the arguments to the kernel
	a := fixed.I26(0)
	b := fixed.I26(1)
	
*/
	binary.Write(buffIn.Writer(), binary.LittleEndian, inp)
	//numBlocks := uint32(inpSize / 64)

	// Create a function that the benchmarking machinery can call
	f := func(B *testing.B) {
		BenchmarkKernel(world, krnl, B, buffIn, buffOut)
	}
	// Benchmark it
	result := testing.Benchmark(f)

	// Print the result
	fmt.Printf("%s\n", result.String())

	// Decode that byte slice into the uint32 we're expecting
	var ret fixed.Int26_6
	err := binary.Read(buffOut.Reader(), binary.LittleEndian, &ret)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	// Compute the expected result 
	expected := inp[0] ^ inp[1] 

	// Exit with an error if the value is not correct
	if expected != ret {
		// Print the value we got from the FPGA
		fmt.Printf("Expected %d, got %d\n", expected, ret)
		os.Exit(1)
	}

}
