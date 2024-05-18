package syntax

import ("fmt")

func Loops(){
//  for loop in golang

for i:=0;i<5;i++{
fmt.Println(i)
}
//  while loop in golang

j:=0

for j<5{
fmt.Println("divyanshu singh");
j++
}
//  infinite loop in golang
// for {
// fmt.Println("infinite loop in golang")
// }

//  range loop in golang
numbers:=[]string{"One","Two","Three","Four","Five","Six"}
for i, n:=range numbers{

fmt.Println(i+1,n)

}




}