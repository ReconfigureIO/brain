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

package utils

import (

	"github.com/sjwhitworth/golearn/base"
	"fmt"
	"math"
)

func load_data(path string) (instance *base.DenseInstances){

   rawData, err := base.ParseCSVToInstances("datasets/mnist_test.csv", false)
   if err != nil {
      panic(err)
   }
   return rawData
}

// Continuous sigmoid implementation
func sigmoid(x float64) float64 {  
        return 1.0 / (1.0 + math.Exp(-x))
}

// Discretise sigmoid based on 'z' granularity
// 'z' = 1 returns a 200 entry table of sigmoid
/*func discrete_sigmoid(z int32) {

    i := float64(0)
    index := 0	
    for i = -100.0 ; i < 100.0 ; i= i + (1/z){

       tmp := sigmoid(i)
       index = index + 1	 	
       if tmp == 1 {
   
              fmt.Printf("fixed.I26F(1 , 0 << 0), \n")		        
       }else {fmt.Printf("fixed.I26F(0 , %d << 0), \n", int64(tmp*1000000))
       }
    }
    
}*/

