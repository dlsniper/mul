package mul2

import (
	"strconv"
	"strings"
)

func prodSeq(seq []string, c chan int64) {
	var prod int64
	prodSet := false

	for _, v := range seq {
		if prodSet == false {
			prod = 1
			prodSet = true
		}

		num, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(err)
		}

		prod = prod * num
	}
	c <- prod
}

func Mul(str string) int64 {
	seq := strings.Split(str, "")
	len := len(seq)
	consecutiveNum := 13
	c := make(chan int64, len/consecutiveNum)

	for i := 0; i+consecutiveNum < len; i++ {
		go prodSeq(seq[i:i+consecutiveNum], c)
	}

	var maxProd int64
	maxProdSet := false

	for j := 0; j < len-consecutiveNum; j++ {
		x := <-c
		if maxProdSet == false {
			maxProd = x
			maxProdSet = true
		} else if x > maxProd {
			maxProd = x
		}
	}

	return maxProd
}
