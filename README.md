# exchange

exchange is a Go library for converting exchange, with a focus on exchange . It provides functions to exchange.

## Installation

To use exchange in your Go project, you can import it using the following command:

```shell
go get github.com/panyingyun/exchange
```

## Usage

To use the exchange functions, first import this package:

```go
import "github.com/chengchuu/exchange"
```

Here's an example usage of the `USD2CNY` function:

```go
	dor := 10.0
	cny := exchange.USD2CNY(dor)
	fmt.Println(cny) 
```

In this example, we're converting a Shanghai time string ("08:00") to its equivalent UTC time string ("00:00"). We pass 8 as the `offset` argument since Shanghai's UTC offset is +8.

## Run Test 


```go
function test

go test -v

benchmark test
go test -bench=.
```


## Contributing

If you find a bug or have an idea for a new feature, please open an issue on the GitHub repository. Pull requests are also welcome!

## License

This library is licensed under the  Apache License Version 2.0. See the LICENSE file for more details.