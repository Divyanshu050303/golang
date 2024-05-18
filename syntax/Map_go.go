package syntax

import (
	"fmt"
)

func Gomap(){
//  to make the empty map in golang
emptyMap := make(map[string]int)
fmt.Println(emptyMap)

//  to make the valued map in golang

valueMap:=map[string]int{
"one":1,
"two":2,
"three":3,
"four":4,
}
fmt.Println(valueMap)

//  map iteration 

for i, v:=range valueMap{
fmt.Println(i, v)
}

//  accessing the particular value in the map
fmt.Println(valueMap["one"])

// add the another value to the map
valueMap["five"]=5
valueMap["six"]=6
valueMap["seven"]=7
fmt.Println(valueMap)

//  delete the particular value from the map

delete(valueMap, "one");
fmt.Println(valueMap)


}