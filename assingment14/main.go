package main

import (
	"errors"
	"fmt"
)

func divide(num, dnum float64) float64 {
	if dnum == 0 {
		panic("Denomenator = 0 can't compute")
	}
	return num/dnum
}
func safeDivide(num,dnum float64) (float64, error) {
	defer recoveryFunc()
	if dnum == 0 {
		err := errors.New("Denomenator is zero!")
		return -1,err
	}
	result := divide(num, dnum)
	return result, nil
}

func recoveryFunc (){
	recoveryMsg := recover()
	if recoveryMsg != nil{
		fmt.Println("recovery msg:",recoveryMsg)
		return
	}
}

func main() {
	result,err := safeDivide(3,0)
	if err != nil{
		fmt.Println("Instead of panic u will get error")
		fmt.Println("Erro: ",err)
		return
	}
	fmt.Println("safe divide result: ",result)
	fmt.Println("To see panic we are calling diivde below")
	fmt.Println("safe divide result: ",divide(3,0))

	result,err =safeDivide()
}
