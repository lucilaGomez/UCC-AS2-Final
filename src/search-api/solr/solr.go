package solr

import (
	"github.com/rtt/Go-Solr"
	log "github.com/sirupsen/logrus"
)

var SolrClient *solr.Connection

func InitSolr() {

	var err error

	SolrClient, err = solr.Init("http://localhost:8983/solr/hotels", 8983, "hotels")
	if err != nil {
		log.Info("Failed to connect to Solr")
		log.Fatal(err)
	} else {
		log.Info("Connected to Solr successfully")
	}
}
