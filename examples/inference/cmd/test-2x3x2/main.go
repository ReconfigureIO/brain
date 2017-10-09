package main

import (
	"fmt"
	"xcl"
	"github.com/reconfigureio/brain/bnn"
	"github.com/reconfigureio/brain/utils"
)

func main() {
	world := xcl.NewWorld()
	defer world.Release()

	krnl := world.Import("kernel_test").GetKernel("reconfigure_io_sdaccel_builder_stub_0_1")
	defer krnl.Release()

	var fpath string
	input = utils.load_data(fpath)

	//load validations 
	test := bnn.ReadImage("dataset")
	fmt.Println(test)

	//reshape image 
	nw_image:= bnn.ReshapeImage(image)
	fmt.Println(nw_image)

	krnl.Run(1, 1, 1)
	fmt.Println("job's done!")
}
