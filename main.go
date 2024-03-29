package main

import (
	"git.garena.com/shopee/loan-service/credit_backend/Goldbach/src/biz"
	"git.garena.com/shopee/loan-service/credit_backend/Goldbach/src/dao"
)

func main() {
	defer func() {
		dao.Close()
	}()

	for num := int64(4); num <= 100; num += 2 {
		goldbachs := biz.SumOfTwoPrimes(num)
		dao.Insert(goldbachs)
	}

	//
	/*for i:=2; i<maxPrime; i++{
		if prime.IsPrime(i) {
			fmt.Printf("%d\t, %v\t, %v\t\n", i, strconv.FormatInt(int64(i), 2), strconv.FormatInt(int64(i), 3))
		}
	}*/

	/*	for i := 4; i < 1000000; i += 2 {
		for j := 2; j <= i/2; j++ {
			if prime.IsPrime(j) && prime.IsPrime(i-j) {
				fmt.Printf("%8d, %10v, %10v, %10v\n", i, strconv.FormatInt(int64(i), 2), strconv.FormatInt(int64(j), 2), strconv.FormatInt(int64(i-j), 2))
				break
			}
		}
	}*/

}
