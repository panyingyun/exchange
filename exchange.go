package exchange

func EUR2USD(amount float64) float64 {
	return amount * 1.15
}

func USD2EUR(amount float64) float64 {
	return amount * 0.87
}

func USD2CNY(amount float64) float64 {
	return amount * 7.10
}

func CNY2USD(amount float64) float64 {
	return amount * 0.14
}
