package mul4

import (
	"strconv"
	"strings"
)

func prodSeq(seq []string) int64 {
	prod := int64(1) // AARGH! Go!
	for _, v := range seq {
		if v == "0" {
			return 0
		}

		num, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(err)
		}

		prod = prod * num
	}

	return prod
}

func Mul(str string, consecutiveNum int) int64 {
	seq := strings.Split(str, "")
	len := len(seq)+1

	maxProd := int64(0) // AARGH! Go!
	for i := 0; i < len-consecutiveNum; i++ {
		prod := prodSeq(seq[i : i+consecutiveNum])
		if prod > maxProd {
			maxProd = prod
		}
	}

	return maxProd
}
