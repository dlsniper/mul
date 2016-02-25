package mul5

import (
	"strconv"
	"strings"
)

func Mul(str string) int64 {
	seq := strings.Split(str, "")
	len := len(seq)
	consecutiveNum := 13

	maxProd := int64(0) // AARGH! Go!
	for i := 0; i < len-consecutiveNum; i++ {
		prod := int64(1) // AARGH! Go!
		for _, v := range seq[i : i+consecutiveNum] {
			if v == "0" {
				prod = 0
				break
			}

			num, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				panic(err)
			}

			prod = prod * num
		}

		if prod > maxProd {
			maxProd = prod
		}
	}

	return maxProd
}
