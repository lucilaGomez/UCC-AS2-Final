package service

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"hotel-api/dto"
	"io/ioutil"
	"net/http"
)

type amadeusService struct{}

type amadeusServiceInterface interface {
	getHotelByCity(city string) dto.AmadeusDto
}

var AmadeusService amadeusServiceInterface

func init() {
	AmadeusService = &amadeusService{}
}

func (s *amadeusService) getHotelByCity(city string) dto.AmadeusDto {
	var url = "https://test.api.amadeus.com/v1/reference-data/locations/hotels/by-city?cityCode="
	url = url + city

	req, err := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", "Bearer quioRqGRy0V1fXN8pdfTh2wj63OU")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("error en get Amadeus", err)
	}
	log.Info("request a Amadeus: ", req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Error al leer body de respuesta de Amadeus:", err)
	}

	var amadeusDto dto.AmadeusDto
	jsonResp := json.Unmarshal(body, &amadeusDto)

	log.Info("response Amadeus: ", jsonResp)

	return amadeusDto

}
