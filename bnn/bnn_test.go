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

package bnn

import (
	"github.com/ReconfigureIO/fixed"
	"github.com/sjwhitworth/golearn/base"
	"testing"
)

func load_data(path string) (instance *base.DenseInstances) {

	rawData, err := base.ParseCSVToInstances(path, false)
	if err != nil {
		panic(err)
	}
	return rawData
}

func TestInference(t *testing.T) {

	// Prepare data matrices for inferenece
	_ = load_data("../datasets/mnist_train.csv")

	weights := [4][4]Synapse{}
	input := [4][4]fixed.Int26_6{}
	network := [4][4]Neuron{}

	ret := Inference(weights, input, network)
	expected := [3]fixed.Int26_6{}

	if ret != expected {
		t.Errorf("Expected %d got %d", expected, ret)
	}

}
