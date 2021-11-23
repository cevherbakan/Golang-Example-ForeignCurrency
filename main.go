package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.tcmb.gov.tr/kurlar/today.xml")

	if err != nil {
		log.Fatal("HATA: %s", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	var d date
	xml.Unmarshal(data,&d)
	fmt.Println(d.Currencies)
}
type currency struct{
	XMLName xml.Name `xml:"Currency" `
	Isim string `xml:"Isim" `
	ForexBuying string `xml:"ForexBuying" `
	ForexSelling string `xml:"ForexSelling" `
}

type date struct {
	XMLName xml.Name `xml:"Tarih_Date" `
	Currencies []currency `xml:"Currency" `
}

func (c currency) String() string {
	return fmt.Sprintf("\t Isim:%s Alis:%s Satis:%s",c.Isim,c.ForexBuying,c.ForexSelling)
}
