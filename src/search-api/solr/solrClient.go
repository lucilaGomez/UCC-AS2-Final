package solr

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/stevenferrer/solr-go"
	sdto "search-api/searchDto"
)

type SolrClient struct {
	Client     *solr.JSONClient
	Collection string
}

func (sc *SolrClient) Add(message sdto.QueueMessageDto) { //e.ApiError {
	//var addHotelDto dto.AddDto
	//addHotelDto.Add = dto.DocDto{Doc: hotelDto}
	//data, err := json.Marshal(addHotelDto)

	docSolr := make(map[string]interface{})

	docSolr["id"] = message.Id
	docSolr["name"] = message.Name
	docSolr["city"] = message.City
	docSolr["description"] = message.Description
	docSolr["room_amount"] = message.RoomAmount
	docSolr["street_name"] = message.StreetName
	docSolr["street_number"] = message.StreetNumber
	docSolr["rate"] = message.Rate

	data, err := json.Marshal(docSolr)

	reader := bytes.NewReader(data)
	if err != nil {
		//return e.NewBadRequestApiError("Error getting json")
	}
	_, err = sc.Client.Update(context.TODO(), sc.Collection, solr.JSON, reader)
	//logger.Debug(resp)
	if err != nil {
		//return e.NewBadRequestApiError("Error in solr")
	}

	er := sc.Client.Commit(context.TODO(), sc.Collection)
	if er != nil {
		//logger.Debug("Error committing load")
		//return e.NewInternalServerApiError("Error committing to solr", er)
	}
	//return nil
}
