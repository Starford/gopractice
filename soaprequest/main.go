package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Data struct {
	XMLName xml.Name
	Body    Body
}

type Body struct {
	XMLName                     xml.Name
	CelsiusToFahrenheitResponse completeResponse `xml:"CelsiusToFahrenheitResponse"`
}

type completeResponse struct {
	XMLName                   xml.Name `xml:"CelsiusToFahrenheitResponse"`
	CelsiusToFahrenheitResult string   `xml:"CelsiusToFahrenheitResult"`
	Errorstring               string   `xml:"error"`
}

// type Body struct {
// 	XMLName     xml.Name
// 	GetResponse completeResponse `xml:"activationPack_completeResponse"`
// }

// type completeResponse struct {
// 	XMLName xml.Name `xml:"activationPack_completeResponse"`
// 	Id      string   `xml:"Id,attr"`
// 	MyVar   string   `xml:"activationPack_completeResult"`
// }

func main() {

	url := "https://www.w3schools.com/xml/tempconvert.asmx?wsdl="

	payload := strings.NewReader("<?xml version=\"1.0\" encoding=\"utf-8\"?>\n<soap12:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap12=\"http://www.w3.org/2003/05/soap-envelope\">\n  <soap12:Body>\n    <CelsiusToFahrenheit xmlns=\"https://www.w3schools.com/xml/\">\n      <Celsius>1000</Celsius>\n    </CelsiusToFahrenheit>\n  </soap12:Body>\n</soap12:Envelope>")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "620f29af-ccfe-433b-8e91-25e1bc42acd4")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	var data Data
	xml.Unmarshal([]byte(body), &data)
	jsonData, _ := json.Marshal(data)
	fmt.Println(string(jsonData))

}
