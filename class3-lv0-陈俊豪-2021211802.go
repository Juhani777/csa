//@program:     class3-lv0
//@author:      chenjunhao
//@create:      2021-11-20
package main



import (
	"fmt"
)

var (
	myres = make(map[int] int ,20)
	ch =make(chan int,1)
)


func factorial(n int ){
	var res =1
	for i:=1;i<=n;i++{
		res *=i
	}
	ch<-res
	myres[n] = res
	<-ch
}

func main(){
	for i:=1;i<=20;i++{
		go factorial(i)
	}
	for i,v:= range myres{
		fmt.Printf("myres[%d] = %d\n",i,v)
	}
}