package main

import (
	"math/rand"
	"time"
	"fmt"
)

type randItem struct {
	start float64
	end   float64
}

func GetRandomItem(items map[interface{}]float64) interface{} {
	nums := make(map[interface{}]randItem)

	var s float64
	for k, f := range items {
		if f <= 0 {
			continue
		}
		nums[k] = randItem{
			start: s,
			end:   s + f,
		}
		s = nums[k].end
	}
	rnd := rand.Float64() * s
	for k, f := range nums {
		if f.start <= rnd && f.end > rnd {
			return k
		}
	}
	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	type info struct {
		Start int
		End   int
	}

	mItems := map[interface{}]float64{
		info{
			Start: 0,
			End:   15,
		}: 0.5,
		info{
			Start: 16,
			End:   31,
		}: 0.2,
		info{
			Start: 32,
			End:   64,
		}: 0.2,
		info{
			Start: 65,
			End:   128,
		}: 0.1,
	}

	item := GetRandomItem(mItems).(info)
	fmt.Println("===", item)
}
