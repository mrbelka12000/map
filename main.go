package main

import (
	"encoding/xml"
	"github.com/graincomg/graincom_logistik/models"
	"github.com/graincomg/graincom_logistik/prettifier"
	"github.com/graincomg/graincom_logistik/visualiser"
	"io/ioutil"
	"log"
)

func main() {
	body, err := ioutil.ReadFile("../../Downloads/map (7)")
	if err != nil {
		log.Fatal(err)
	}
	var mp models.Oms
	err = xml.Unmarshal(body, &mp)
	if err != nil {
		log.Fatal(err)
	}
	prettifier.Сonvenience(mp)

	visualiser.Visualiser(mp.Node, mp.Way, nil)
}
