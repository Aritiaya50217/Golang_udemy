package prime

import (
	"fmt"
	"math"
)

var Notprimes [100000] bool // ไม่ใช่ primes ทั้งหมด

func init(){
	// ต้องการจะหาค่า primes
	fmt.Println("initialization in Prime package")
	sqrtLen := int(math.Ceil(math.Sqrt(float64(len(Notprimes)))))
	for i := 2; i <  sqrtLen; i++ {
		if Notprimes[i] {
			continue
		}
		Notprimes[i] = false
		for j := i * 2 ; j < len(Notprimes); i++ {
			Notprimes[j] = true
		}
	}
	fmt.Println("initialized")
}

func IsPrime(x int64) bool {
	return !Notprimes[x]
}