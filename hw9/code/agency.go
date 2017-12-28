package traveAgency

import "time"

type Weather struct{}

type Destinations struct{}

type Quoting struct{}

type Customers struct{}

type Info struct {
	Destinations    Destinations
	Quote   Quoting
	Weather Weather
}

func GetCustomerDetails() Customers {
	time.Sleep(150 * time.Millisecond)
	return Customers{}
}

func GetRecommendedDestinations(customers Customers) [10]Destinations {
	time.Sleep(250 * time.Millisecond)
	return [10]Destinations{}
}

func GetQuote(destination Destinations) Quoting {
	time.Sleep(170 * time.Millisecond)
	return Quoting{}
}

func GetWeatherForecast(destinations Destinations) Weather {
	time.Sleep(330 * time.Millisecond)
	return Weather{}
}

