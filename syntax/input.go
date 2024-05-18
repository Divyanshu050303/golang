package syntax

import (
	"fmt"
)

func Input(){
fmt.Println("Enter your name")
name:=""

fmt.Scan(&name);

fmt.Println(name)

}