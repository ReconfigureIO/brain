# bnn-fpga

This repo contains Go implementation and corresponding APIs for acceleration of Binarized Neural Network (BNN) on FPGAs.

The current bnn package provides a set of functions for neural network processing including network constructors, inference, training and input data manipulation.

##Notes

Implementation wise the following assumptions are made wrt the network (for now):

    * neurons are considered to be fully connected (no sparsity).
    * a bias matrix is to be considered along with the weight matrix.
    * Initial weight (bias) distributions (aka normalization) are randomly defined.
    * an input data set from [the UCI machine learning repo](http://archive.ics.uci.edu/ml/datasets)
    * initial data alignment is to be done by the CPU and get loaded into the FPGA-visible memory space.
    * the inference function is selected as the kernel (FPGA resident action).
    * the design is checked by `reco check` to ensure it is synthesisable by our compiler.

##Install 

First off install the base package from GoLearn which is used for data augmentation etc.
```bash
   go get -t -u -v github.com/sjwhitworth/golearn/base
```

##Getting Started

##Examples
    
##Docs    
    
[TODO] a list of Go features that our compiler dosen't support.


