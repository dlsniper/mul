package mul1

import (
	"strconv"
	"strings"
)

func prodSeq(seq []string, c chan float64) {
	var prod float64
	prodSet := false

	for _, v := range seq {
		if prodSet == false {
			prod = 1
			prodSet = true
		}

		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		prod = prod * float64(num)
	}
	c <- prod
}

func Mul(str string) float64 {
	seq := strings.Split(str, "")
	len := len(seq)
	consecutiveNum := 13
	c := make(chan float64, len/consecutiveNum)

	for i := 0; i+consecutiveNum < len; i++ {
		go prodSeq(seq[i:i+consecutiveNum], c)
	}

	var maxProd float64
	maxProdSet := false

	for j := 0; j+consecutiveNum < len; j++ {
		select {
		case x := <-c:
			if maxProdSet == false {
				maxProd = x
				maxProdSet = true
			} else if x > maxProd {
				maxProd = x
			}

		}
	}

	return maxProd
}
