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

func BenchmarkKernel(world xcl.World, 
		krnl *xcl.Kernel,
		B *testing.B, 
		buffActs *xcl.Memory, 
		buffIn *xcl.Memory, 
		buffWeightH *xcl.Memory,
		buffBiasH *xcl.Memory,
		buffWeightO *xcl.Memory,
		buffBiasO *xcl.Memory,
		buffOut *xcl.Memory) {


	// Set the pointer to the output buffer
	krnl.SetMemoryArg(0, buffActs)
	// Set the pointer to the output buffer
	krnl.SetMemoryArg(1, buffIn)
	// Set the pointer to the output buffer
	krnl.SetMemoryArg(2, buffWeightH)
	// Set the pointer to the output buffer
	krnl.SetMemoryArg(3, buffBiasH)
	// Set the pointer to the output buffer
	krnl.SetMemoryArg(4, buffWeightO)
	// Set the pointer to the output buffer
	krnl.SetMemoryArg(5, buffBiasO)
	// Set the pointer to the output buffer
	krnl.SetMemoryArg(6, buffOut)


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


	// Generated table by util/disretise_sig(1)
  	actives := [200]fixed.Int26_6{fixed.I26F(0 , 0 << 0), 
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
		fixed.I26F(0 , 880797 << 0),
		fixed.I26F(0 , 952574 << 0), 
		fixed.I26F(0 , 982013 << 0), 
		fixed.I26F(0 , 993307 << 0), 
		fixed.I26F(0 , 997527 << 0), 
		fixed.I26F(0 , 999088 << 0), 
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
		fixed.I26F(0 , 999999 << 0),
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
		fixed.I26F(1 , 0 << 0)}

	// Input batch size = 4 - famous iris flower dataset
	inp := [4]fixed.Int26_6{fixed.I26F(0 , 583333333333 << 0),
		fixed.I26F(0 , 291666666667 << 0),
		fixed.I26F(0 , 728813559322 << 0),
		fixed.I26F(0 , 75 << 0)}

	// From the training stage of datadan.io network model
	weightH := [12]fixed.Int26_6{fixed.I26F(-9 , 664649023077196 << 0),
		fixed.I26F(-3 , 331748963332474 << 0),
		fixed.I26F(0 , 4798734558954657 >> 1),
		fixed.I26F(-9 , 16865031545236 << 0),
		fixed.I26F(7 , 526678142973989 << 0),
		fixed.I26F(-11 , 259862681673592 << 0),
		fixed.I26F(35 , 375442386538865 << 0),
		fixed.I26F(-8 , 377651024226779 << 0),
		fixed.I26F(0 , 5857331473302626 << 0),
		fixed.I26F(14 , 569948798276545 << 0),
		fixed.I26F(-6 , 977871499870339 << 0),
		fixed.I26F(-3 , 595834572863486 << 0)}

	biasH := [3]fixed.Int26_6{fixed.I26F(-24 , 257317924965545 << 0),
		fixed.I26F(2 , 841569482220003 << 0),
		fixed.I26F(1 , 2344415303891234 << 0)}

	weightO := [9]fixed.Int26_6{fixed.I26F(-6 , 421059674610277 << 0),
		fixed.I26F(10 , 430255115994242 << 0),
		fixed.I26F(-10 , 466201644889994 << 0),
		fixed.I26F(12 , 38461165673367 << 0),
		fixed.I26F(-5 , 610451231603859 << 0),
		fixed.I26F(-11 , 310357612664564 << 0),
		fixed.I26F(-12 , 155845492268348 << 0),
		fixed.I26F(-4 , 313406702633787 << 0),
		fixed.I26F(1 , 3793360362443159 << 0)}

	biasO := [3]fixed.Int26_6{fixed.I26F(-4 , 703941034009741 << 0),
		fixed.I26F(-6 , 59314294559707 >> 1),
		fixed.I26F(6 , 201619726339744 << 0)}


        // Allocate a buffer on the FPGA to store the return value of our computation
        // The activations is a 200-uint32 set, so we need 4 * 200 bytes to store it
        buffActs := world.Malloc(xcl.ReadOnly, 800)
        defer buffActs.Free()

        // Allocate a buffer on the FPGA to store the return value of our computation
        // The input is a 4-uint32 set, so we need 4 * 4 bytes to store it
        buffIn := world.Malloc(xcl.ReadOnly, 16)
        defer buffIn.Free()

        // Allocate a buffer on the FPGA to store the return value of our computation
        // The hidden-weights is a 12-uint32 set, so we need 4 * 12 bytes to store it
        buffWeightH := world.Malloc(xcl.ReadOnly, 48)
        defer buffActs.Free()

        // Allocate a buffer on the FPGA to store the return value of our computation
        // The hidden biases is a 3-uint32 set, so we need 4 * 3 bytes to store it
        buffBiasH := world.Malloc(xcl.ReadOnly, 12)
        defer buffActs.Free()

        // Allocate a buffer on the FPGA to store the return value of our computation
        // The out weights is a 9-uint32 set, so we need 4 * 9 bytes to store it
        buffWeightO := world.Malloc(xcl.ReadOnly, 36)
        defer buffActs.Free()

        // Allocate a buffer on the FPGA to store the return value of our computation
        // The out biases is a 3-uint32 set, so we need 4 * 3 bytes to store it
        buffBiasO := world.Malloc(xcl.ReadOnly, 12)
        defer buffActs.Free()

        // Allocate a buffer on the FPGA to store the return value of our computation
        // The output is a uint32, so we need 4 bytes to store it
        buffOut := world.Malloc(xcl.WriteOnly, 4)
        defer buffOut.Free()


	// Write into the allocated buffers of inps, acts, weights and bias

	binary.Write(buffActs.Writer(), binary.LittleEndian, actives)

	binary.Write(buffIn.Writer(), binary.LittleEndian, inp)

	binary.Write(buffWeightH.Writer(), binary.LittleEndian, weightH)

	binary.Write(buffBiasH.Writer(), binary.LittleEndian, biasH)

	binary.Write(buffWeightO.Writer(), binary.LittleEndian, weightO)

	binary.Write(buffBiasO.Writer(), binary.LittleEndian, biasO)


	// Create a function that the benchmarking machinery can call
	f := func(B *testing.B) {
		BenchmarkKernel(world, krnl, B, buffActs, buffIn, buffWeightH, buffBiasH, buffWeightO, buffBiasO, buffOut)
	}
	// Benchmark it
	result := testing.Benchmark(f)

	// Print the result
	fmt.Printf("%s\n", result.String())

	// Decode that byte slice into the uint32 we're expecting
	var ret []fixed.Int26_6
	err := binary.Read(buffOut.Reader(), binary.LittleEndian, &ret)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	// Compute the expected result 
	expected := [3]fixed.Int26_6{0,1,0}

	// Exit with an error if the value is not correct
	if expected[1] != ret[1] {
		// Print the value we got from the FPGA
		fmt.Printf("Expected %d, got %d\n", expected[1], ret[1])
		os.Exit(1)
	}

}
