<p align="center"> 
	<img src="https://user-images.githubusercontent.com/14908455/107840608-caf6a700-6de6-11eb-813d-24dbc1e63981.png" width="150"/>
</p>

Pipe is a clone of Clojure's Threading Macros concept. It offers simplicity to construct a series of function executions in a syntactic sugar way.

## Installation
```
go get -u github.com/parinpan/pipe
```

## Usage
See the full example here: https://github.com/parinpan/pipe/blob/master/example/pipe_implementation.go
```golang
result := pipe.Do(
  pipe.Apply(funcA),
  pipe.Apply(funcB),
  pipe.Apply(funcC),
  pipe.Apply(funcD),
)

fmt.Printf("%+v\n", result)

// more advanced usage

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
```
