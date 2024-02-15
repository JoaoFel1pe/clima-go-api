package helpers

import (
	"api/src/v1/config"
	"log"
	"net/url"
	"path"
)

func BuildQueryCurrentCondition(lat string, lon string) string {
	base_url := config.AZURE_BASE_URL
	api_key := config.AZURE_API_KEY

	params := url.Values{}
	params.Add("api-version", "1.0")
	params.Add("query", lat+","+lon)
	params.Add("subscription-key", api_key)

	u, err := url.Parse(base_url)
	if err != nil {
		log.Fatal(err)
	}
	u.Path = path.Join(u.Path, "/currentConditions/json")
	u.RawQuery = params.Encode()

	query := u.String()

	decodedURL, err := url.QueryUnescape(query)

	if err != nil {
		log.Fatal(err)
	}

	return decodedURL

}

func BuildQueryForecast(lat string, lon string, days string) string {
	base_url := config.AZURE_BASE_URL
	api_key := config.AZURE_API_KEY

	params := url.Values{}

	params.Add("api-version", "1.0")
	params.Add("query", lat+","+lon)
	params.Add("duration", days)
	params.Add("subscription-key", api_key)

	u, err := url.Parse(base_url)
	if err != nil {
		log.Fatal(err)
	}

	u.Path = path.Join(u.Path, "/forecast/daily/json")
	u.RawQuery = params.Encode()

	query := u.String()

	decodedURL, err := url.QueryUnescape(query)

	if err != nil {
		log.Fatal(err)
	}

	return decodedURL
}
