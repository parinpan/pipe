package main

import (
	"github.com/parinpan/pipe"
	"testing"
)

func Benchmark(b *testing.B) {
	fn := func(prices []int) int {
		var totalPrice = 0

		for _, price := range prices {
			totalPrice += price
		}

		return totalPrice
	}

	b.Run("ExecutedWithoutPipe", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			getFoodTotalPriceText(fn(getFoodPrices(getFoodsWithFilter(10, getRestaurants(), 5))))
		}
	})

	b.Run("ExecutedWithPipe", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pipe.Do(
				pipe.Apply(getRestaurants),
				pipe.Apply(getFoodsWithFilter, 10, pipe.Pass(), 5),
				pipe.Apply(getFoodPrices),
				pipe.Apply(getFoodTotalPriceText, pipe.Pass(fn)))
		}
	})
}
