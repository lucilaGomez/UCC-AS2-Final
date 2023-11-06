package solr

import (
	"github.com/rtt/Go-Solr"
	log "github.com/sirupsen/logrus"
)

func InitSolr() {
	_, err := solr.Init("http://localhost/solr/hotels", 8983, "hotels")
	if err != nil {
		log.Info("Failed to connect to Solr")
		log.Fatal(err)
	} else {
		log.Info("Connected to Solr successfully")
	}
}
