package main

import (
	"github.com/parinpan/pipe"
	"testing"
)

func Benchmark(b *testing.B) {
	b.Run("ExecutedWithoutPipe", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isFoodHasExactName(
				getFoodWithHighestPrice(
					getFoodsFromRestaurant(
						getMcDonaldFromRestaurants(RestaurantList))),
				"BigMac")
		}
	})

	b.Run("ExecutedWithPipe", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pipe.Do(
				pipe.Apply(getMcDonaldFromRestaurants, RestaurantList),
				pipe.Apply(getFoodsFromRestaurant),
				pipe.Apply(getFoodWithHighestPrice),
				pipe.Apply(isFoodHasExactName, pipe.Pass(), "BigMac"))
		}
	})
}
