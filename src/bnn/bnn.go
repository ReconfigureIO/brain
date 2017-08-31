package bnn

type neuron struct {
    //activation function
    act string
    //no of inputs and outputs per neuron
    inps, outs int
}

func NetworkLayer(size int, act string) []neuron{

  layer := make([]neuron, size)

  //init the array
  for i, _:= range layer {

    layer[i].act = act
    layer[i].inps = 0
    layer[i].outs = 0
  }

  return layer
}

