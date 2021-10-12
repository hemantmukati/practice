package main
import "fmt"

type person1 interface {
	specification()
}

type hemant1 struct {
	surname string
	age int
	height float64
}
func(h hemant1) specification (){
	fmt.Println(h)
}
func what1(x person1 ) (y interface{}) {
	return y
}

func main (){
	h:= hemant1 {
		surname : "mukati" ,
		age : 25 ,
		height: 5.3 ,
	}
	var p person1 = h
	h.specification()
	fmt.Println(what1(p))

}
