package synchronous

import "traveAgency"

func Run() [10]traveAgency.Info {
	customer := traveAgency.GetCustomerDetails()
	destinations := traveAgency.GetRecommendedDestinations(customer)
	var infos [10]traveAgency.Info
	for index, dest := range destinations {
		q := traveAgency.GetQuote(dest)
		w := traveAgency.GetWeatherForecast(dest)
		infos[index] = traveAgency.Info{Destinations:dest, Quote:q, Weather:w}
	}
	return infos
}
