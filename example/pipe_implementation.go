package main

import (
	"fmt"
	"github.com/parinpan/pipe"
	"time"
)

type Food struct {
	Name   string
	Price  int
	Rating int
}

type Restaurant struct {
	Name     string
	OpenAt   time.Time
	CloseAt  time.Time
	Location string
	Foods    []Food
}

var (
	RestaurantList = []Restaurant{
		{
			Name:     "McDonald",
			OpenAt:   time.Date(2021, 1, 1, 9, 0, 0, 0, time.Local),
			CloseAt:  time.Date(2021, 1, 1, 21, 0, 0, 0, time.Local),
			Location: "United States",
			Foods: []Food{
				{
					Name:   "McFlurry",
					Price:  5,
					Rating: 3,
				},
				{
					Name:   "BigMac",
					Price:  10,
					Rating: 5,
				},
			},
		},
		{
			Name:     "Taco Bell",
			OpenAt:   time.Date(2021, 1, 1, 10, 0, 0, 0, time.Local),
			CloseAt:  time.Date(2021, 1, 1, 22, 0, 0, 0, time.Local),
			Location: "Indonesia",
			Foods: []Food{
				{
					Name:   "Nachos",
					Price:  5,
					Rating: 3,
				},
				{
					Name:   "Burrito",
					Price:  10,
					Rating: 5,
				},
				{
					Name:  "Mexican Pizza",
					Price: 20,
				},
			},
		},
	}
)

func getRestaurants() []Restaurant {
	return RestaurantList
}

func getFoodsWithFilter(price int, restaurants []Restaurant, rating int) []Food {
	var foods []Food

	for _, restaurant := range restaurants {
		for _, food := range restaurant.Foods {
			if food.Price == price && food.Rating == rating {
				foods = append(foods, food)
			}
		}
	}

	return foods
}

func getFoodWithHighestPrice(foods []Food) Food {
	var selectedFood Food
	var maximumPrice = 0

	for _, food := range foods {
		if food.Price > maximumPrice {
			selectedFood = food
			maximumPrice = food.Price
		}
	}

	return selectedFood
}

func getFoodPrices(foods []Food) []int {
	var prices []int

	for _, food := range foods {
		prices = append(prices, food.Price)
	}

	return prices
}

func getFoodTotalPriceText(totalPrice int) string {
	return fmt.Sprintf("The food total price is: $%d", totalPrice)
}

func main() {
	result := pipe.Do(
		pipe.Apply(getRestaurants),
		pipe.Apply(getFoodsWithFilter, 10, pipe.Pass(), 5),
		pipe.Apply(getFoodPrices),
		pipe.Apply(getFoodTotalPriceText, pipe.Pass(func(prices []int) int {
			var totalPrice = 0

			for _, price := range prices {
				totalPrice += price
			}

			return totalPrice
		})))

	fmt.Printf("%+v\n", result)
}
