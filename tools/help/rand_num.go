package help

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	rand2 "math/rand"
	"time"
)

func RandNum(length int) string {
	return randNum("", length)
}

func randNum(str string, length int) string {
	i, _ := rand.Int(rand.Reader, big.NewInt(10))

	str = fmt.Sprintf("%s%d", str, i)

	if len(str) < length {
		str = randNum(str, length)
	}

	return str
}

func RandInt(length int) (n int64) {

	for i := 1; i <= length; i++ {
		x, _ := rand.Int(rand.Reader, big.NewInt(10))

		j := x.Int64()
		if i == length && j == 0 {
			j = 1
		}
		n += j * int64(math.Pow10(i-1))
	}

	return n
}

func RandLt(max int) int {
	rand2.Seed(time.Now().UnixNano())
	return rand2.Intn(max)
}

func RandStrForNow() string {
	return fmt.Sprintf("%s%s", time.Now().Format("20060102"), RandNum(4))
}
