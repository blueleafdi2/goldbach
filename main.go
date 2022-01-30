package main

import (
	"fmt"
	"git.garena.com/shopee/loan-service/credit_backend/Goldbach/prime"
	"strconv"
)

var maxPrime int = 100

func main() {

	//
	/*for i:=2; i<maxPrime; i++{
		if prime.IsPrime(i) {
			fmt.Printf("%d\t, %v\t, %v\t\n", i, strconv.FormatInt(int64(i), 2), strconv.FormatInt(int64(i), 3))
		}
	}*/

	for i := 4; i < 1000000; i += 2 {
		for j := 2; j <= i/2; j++ {
			if prime.IsPrime(j) && prime.IsPrime(i-j) {
				fmt.Printf("%8d, %10v, %10v, %10v\n", i, strconv.FormatInt(int64(i), 2), strconv.FormatInt(int64(j), 2), strconv.FormatInt(int64(i-j), 2))
				break
			}
		}
	}

}
