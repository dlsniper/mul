package mul6

import (
	"strconv"
	"strings"
	"runtime"
	"sync"
)

func prodSeq(seq []string, c chan int64, wg *sync.WaitGroup) {
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
	wg.Done()
	c <- prod
}

func Mul(str string, consecutiveNum int) int64 {
	seq := strings.Split(str, "")
	len := len(seq)+1
	c := make(chan int64, len-consecutiveNum)
	numCPU := runtime.NumCPU()

	for i := 0; i < len - consecutiveNum; i++ {
		wg := &sync.WaitGroup{}
		for j:=0; j<numCPU && i < len - consecutiveNum; j++ {
			wg.Add(1)
			go prodSeq(seq[i:i+consecutiveNum], c, wg)
			i++
		}
		wg.Wait()
	}
	close(c)

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
