[![Build Status](https://travis-ci.org/ReconfigureIO/brain.svg?branch=master)](https://travis-ci.org/ReconfigureIO/brain)

# bnn-fpga

This repo contains Go implementation and corresponding APIs for acceleration of Binarized Neural Network (BNN) on FPGAs.

The current bnn package provides a set of functions for neural network processing including network constructors, inference, training and input data manipulation.

## Notes

Implementation wise the following assumptions are made wrt the network (for now):

    * neurons are considered to be fully connected (no sparsity).
    * a bias matrix is to be considered along with the weight matrix.
    * Initial weight (bias) distributions (aka normalization) are randomly defined.
    * an input data set from [the UCI machine learning repo](http://archive.ics.uci.edu/ml/datasets)
    * initial data alignment is to be done by the CPU and get loaded into the FPGA-visible memory space.
    * the inference function is selected as the kernel (FPGA resident action).
    * the design is checked by `reco check` to ensure it is synthesisable by our compiler.

## Install 

First off install the base package from GoLearn which is used for data augmentation etc.
```bash
   export GOPATH=`dir`
  
   uzip ./src/datasets/*.zip && rm ./src/datasets/*.zip 
   
   go get -t -u -v github.com/sjwhitworth/golearn/base

   go get -t -u -v github.com/reconfigureio/brain/bnn

   (Optional: utils package provides functions for loading datasets)
   go get -t -u -v github.com/reconfigureio/brain/utils
```

## Getting Started

## Examples

/examples dir contains two examples (inference and training) of implementing a neural network using Reconfigure.io's provided `brain` library. Each example is consist of `cmd` and `kernel` parts. It is usually a good practice to think in advance about partitioning your computation and data. In the current examples we have decided to locate data (incl. weight and bias), from training phase, on the host side (`/cmd/main.go`) so that we could apply pre-inference data augmentation, such as re formatting the entires from float to fixed.   

for more information about implementing, simulating, and running kernels please refer to our documentation:
http://docs.reconfigure.io/


## Docs    
    
[TODO] a list of Go features that our compiler dosen't support.
http://docs.reconfigure.io/go_support.html

## References

Daniel Whitenack, "Machine Learning With Go", Packt Publishing, September 2017
Data Set: Iris flower data set: https://en.wikipedia.org/wiki/Iris_flower_data_set



