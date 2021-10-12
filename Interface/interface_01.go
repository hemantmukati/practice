package main
import "fmt"

type person interface {
	specification()
}

type hemant struct {
	surname string
	age int
	height float64
}
 func(h hemant) specification (){
	 fmt.Println(h)
 }
 func what(x person ) (y interface{}) {
	 return y
 }

func main (){
	h:= hemant {
		surname : "mukati" ,
		age : 25 ,
		height: 5.3 ,
	}
	var p person = h
	 h.specification()
     fmt.Println(what(p))

}
