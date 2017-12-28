package asynchronous

import (
	"traveAgency"
)

func Run() [10]traveAgency.Info {
	customer := traveAgency.GetCustomerDetails()
	destinations := traveAgency.GetRecommendedDestinations(customer)
	var infos [10]traveAgency.Info

	quotes := [10]chan traveAgency.Quoting{}
	weathers := [10]chan traveAgency.Weather{}

	for i := range quotes {
		quotes[i] = make(chan traveAgency.Quoting)
	}

	for i := range weathers {
		weathers[i] = make(chan traveAgency.Weather)
	}

	for index, dest := range destinations {
		i := index
		d := dest
		go func() {
			quotes[i] <- traveAgency.GetQuote(d)
		}()

		go func() {
			weathers[i] <- traveAgency.GetWeatherForecast(d)
		}()
	}

	for index, dest := range destinations {
		infos[index] = traveAgency.Info{Destinations:dest, Quote:<-quotes[index], Weather:<-weathers[index]}
	}

	return infos
}
