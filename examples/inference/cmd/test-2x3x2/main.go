package main

import (
	"encoding/binary"
	"fmt"
	"github.com/ReconfigureIO/sdaccel/xcl"
	"os"
	"reflect"
	"testing"
	"github.com/ReconfigureIO/fixed"
	"github.com/ReconfigureIO/fixed/host"
)

//Partition example dataset based on BATCH_SIZE
//NUM_EPOCHS is practical and may vary based on
//the output accuracy achieved from the model
const NUM_EPOCHS int = 100
const BATCH_SIZE int = 500
const LENGTH uint = 50

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
	// Pass the total length of the input
	krnl.SetArg(7, uint32(LENGTH))

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
	actives := [200]fixed.Int26_6{host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0),
		host.I26Float64(0.2),
		host.I26Float64(0.6),
		host.I26Float64(0.16),
		host.I26Float64(0.45),
		host.I26Float64(0.123),
		host.I26Float64(0.335),
		host.I26Float64(0.911),
		host.I26Float64(0.2472),
		host.I26Float64(0.6692),
		host.I26Float64(0.17986),
		host.I26Float64(0.47425),
		host.I26Float64(0.119202),
		host.I26Float64(0.268941),
		host.I26Float64(0.500000),
		host.I26Float64(0.731058),
		host.I26Float64(0.880797),
		host.I26Float64(0.952574),
		host.I26Float64(0.982013),
		host.I26Float64(0.993307),
		host.I26Float64(0.997527),
		host.I26Float64(0.999088),
		host.I26Float64(0.999664),
		host.I26Float64(0.999876),
		host.I26Float64(0.999954),
		host.I26Float64(0.999983),
		host.I26Float64(0.999993),
		host.I26Float64(0.999997),
		host.I26Float64(0.999999),
		host.I26Float64(0.999999),
		host.I26Float64(0.999999),
		host.I26Float64(0.999999),
		host.I26Float64(0.999999),
		host.I26Float64(0.999999),
		host.I26Float64(0.999999),
		host.I26Float64(0.999999),
		host.I26Float64(0.999999),
		host.I26Float64(0.999999),
		host.I26Float64(0.999999),
		host.I26Float64(0.999999),
		host.I26Float64(0.999999),
		host.I26Float64(0.999999),
		host.I26Float64(0.999999),
		host.I26Float64(0.999999),
		host.I26Float64(0.999999),
		host.I26Float64(0.999999),
		host.I26Float64(0.999999),
		host.I26Float64(0.999999),
		host.I26Float64(0.999999),
		host.I26Float64(0.999999),
		host.I26Float64(0.999999),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1),
		host.I26Float64(1)}

	// Input batch size = 4 - famous iris flower dataset
	inp := [4]fixed.Int26_6{host.I26Float64(0.194444444444),
		host.I26Float64(0.583333333333),
		host.I26Float64(0.101694915254),
		host.I26Float64(0.125)}

	// From the training stage of datadan.io network model
	weightH := [12]fixed.Int26_6{host.I26Float64(-9.664649023),
		host.I26Float64(-3.331748963),
		host.I26Float64(0.0479873455),
		host.I26Float64(-9.16865031),
		host.I26Float64(7.526678142),
		host.I26Float64(-11.25986268),
		host.I26Float64(35.375442386),
		host.I26Float64(-8.377651024),
		host.I26Float64(0.585733147),
		host.I26Float64(14.56994879),
		host.I26Float64(-6.97787149),
		host.I26Float64(-3.59583457)}

	biasH := [3]fixed.Int26_6{host.I26Float64(-24.257317924),
		host.I26Float64(2.84156948),
		host.I26Float64(1.23444153)}

	weightO := [9]fixed.Int26_6{host.I26Float64(-6.421059674),
		host.I26Float64(10.430255115),
		host.I26Float64(-10.466201644),
		host.I26Float64(12.384611656),
		host.I26Float64(-5.610451231),
		host.I26Float64(-11.310357612),
		host.I26Float64(-12.155845492),
		host.I26Float64(-4.313406702),
		host.I26Float64(1.379336036)}

	biasO := [3]fixed.Int26_6{host.I26Float64(-4.70394103),
		host.I26Float64(-6.059314294),
		host.I26Float64(6.20161972)}

	// Allocate a buffer on the FPGA to store the return value of our computation
	// The activations is a 200-uint32 set, so we need 4 * 200 bytes to store it
	buffActs := world.Malloc(xcl.ReadOnly, 800)
	defer buffActs.Free()

	// Allocate a buffer on the FPGA to store the return value of our computation
	// The input is a 4-uint32 set, so we need 4 * 4 bytes to store it
	buffIn := world.Malloc(xcl.ReadOnly, 16*LENGTH)
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
	// The output is a 3-int32 set, so we need 4 * 3 bytes to store it
	buffOut := world.Malloc(xcl.WriteOnly, 12)
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
		BenchmarkKernel(
			world, krnl, B, buffActs, buffIn, buffWeightH, buffBiasH, buffWeightO, buffBiasO, buffOut)
	}
	// Benchmark it
	result := testing.Benchmark(f)

	// Print the result
	fmt.Printf("%s\n", result.String())

	// Decode that byte slice into the uint32 we're expecting
	//var ret [3]fixed.Int26_6
	ret := make([]fixed.Int26_6, 3*LENGTH)
	err := binary.Read(buffOut.Reader(), binary.LittleEndian, ret)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	// Compute the expected result
	expected := make([]fixed.Int26_6, 3*LENGTH)

	// Exit with an error if the value is not correct
	if !reflect.DeepEqual(expected, ret) {
		// Print the value we got from the FPGA
		fmt.Printf("Expected %b, got %b (in binary)\n", expected, ret)
		os.Exit(1)
	}

}
