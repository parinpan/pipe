package main

import (
	"fachr.in/pipe"
	"fmt"
	"time"
)

type Food struct {
	Name  string
	Price int64
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
					Name:  "McFlurry",
					Price: 5,
				},
				{
					Name:  "BigMac",
					Price: 10,
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
					Name:  "Nachos",
					Price: 5,
				},
				{
					Name:  "Burrito",
					Price: 10,
				},
				{
					Name:  "Mexican Pizza",
					Price: 20,
				},
			},
		},
	}
)

func getMcDonaldFromRestaurants(restaurants []Restaurant) Restaurant {
	for _, restaurant := range restaurants {
		if restaurant.Name == "McDonald" {
			return restaurant
		}
	}

	return Restaurant{}
}

func getFoodsFromRestaurant(restaurant Restaurant) []Food {
	return restaurant.Foods
}

func getFoodWithHighestPrice(foods []Food) Food {
	var selectedFood Food
	var maximumPrice = int64(0)

	for _, food := range foods {
		if food.Price > maximumPrice {
			selectedFood = food
			maximumPrice = food.Price
		}
	}

	return selectedFood
}

func isFoodHasExactName(food Food, foodName string) bool {
	return food.Name == foodName
}

func main() {
	result := pipe.Do(
		pipe.Apply(getMcDonaldFromRestaurants, RestaurantList),
		pipe.Apply(getFoodsFromRestaurant),
		pipe.Apply(getFoodWithHighestPrice),
		pipe.Apply(isFoodHasExactName, pipe.Pass(), "BigMac"))

	fmt.Printf("%+v\n", result)
}
