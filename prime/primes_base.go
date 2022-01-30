package prime

import "math"

func IsPrime(num int64) bool {
	if num == 2 || num == 3 {
		return true
	}
	if num < 2 || num%2 == 0 {
		return false
	}
	iLimit := int64(math.Sqrt(float64(num) + 0.1))
	i := int64(2)
	for ; i <= iLimit; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

type Goldbach struct {
	Id  int64
	Seq int
	A   int64
	B   int64
	Num int64
}

func SumOfTwoPrimes(num int64) []*Goldbach {
	if num < 4 || num%2 == 1 {
		panic("num < 4 || num % 2 == 1 in SumOfFistTwoPrime")
	}
	sumPrimesList := make([]*Goldbach, 0)
	for i := int64(2); i <= num/2; i++ {
		if IsPrime(i) && IsPrime(num-i) {
			sumPrimes := &Goldbach{
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

func fullfillSeq(sumPrimesList []*Goldbach) {
	for i, sumPrimes := range sumPrimesList {
		sumPrimes.Seq = i + 1
	}

}
