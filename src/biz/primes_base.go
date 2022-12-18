package biz

import (
	"git.garena.com/shopee/loan-service/credit_backend/Goldbach/src/model"
	"math"
	"strconv"
)

func IsPrime(num int64) bool {
	if num == 2 || num == 3 {
		return true
	}
	if num < 2 || num%2 == 0 {
		return false
	}
	/*if num%6 == 1 || num%6 == 5 {
		return false
	}*/
	iLimit := int64(math.Sqrt(float64(num) + 0.1))
	i := int64(2)
	for ; i <= iLimit; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func SumOfTwoPrimes(num int64) []*model.Goldbach {
	if num < 4 || num%2 == 1 {
		panic("num < 4 || num % 2 == 1 in SumOfFistTwoPrime")
	}
	sumPrimesList := make([]*model.Goldbach, 0)
	for i := int64(2); i <= num/2; i++ {
		if IsPrime(i) && IsPrime(num-i) {
			sumPrimes := &model.Goldbach{
				A:   i,
				B:   num - i,
				Num: num,
			}
			sumPrimesList = append(sumPrimesList, sumPrimes)
		}
	}
	fullfillSeq(sumPrimesList)
	return sumPrimesList
}

func fullfillSeq(sumPrimesList []*model.Goldbach) {
	for i, sumPrimes := range sumPrimesList {
		sumPrimes.Seq = int64(i + 1)
		sumPrimes.NumByte = CountOfOneInByte(sumPrimes.Num)
		sumPrimes.AByte = CountOfOneInByte(sumPrimes.A)
		sumPrimes.BByte = CountOfOneInByte(sumPrimes.B)
		sumPrimes.ATri = strconv.FormatInt(int64(sumPrimes.A), 3)
		sumPrimes.BTri = strconv.FormatInt(int64(sumPrimes.B), 3)
		sumPrimes.ASix = strconv.FormatInt(int64(sumPrimes.A), 6)
		sumPrimes.BSix = strconv.FormatInt(int64(sumPrimes.B), 6)
		sumPrimes.AAndB = sumPrimes.A & sumPrimes.B
		sumPrimes.AOrB = sumPrimes.A | sumPrimes.B
		sumPrimes.AXorB = sumPrimes.A ^ sumPrimes.B
		sumPrimes.AAndNotB = sumPrimes.A &^ sumPrimes.B
	}

}
