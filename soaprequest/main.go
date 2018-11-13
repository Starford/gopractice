package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// type Data struct {
// 	XMLName xml.Name
// 	Body    Body
// }

// type Body struct {
// 	XMLName                     xml.Name
// 	CelsiusToFahrenheitResponse completeResponse `xml:"CelsiusToFahrenheitResponse"`
// }

// type completeResponse struct {
// 	XMLName                   xml.Name `xml:"CelsiusToFahrenheitResponse"`
// 	CelsiusToFahrenheitResult string   `xml:"CelsiusToFahrenheitResult"`
// 	Errorstring               string   `xml:"error"`
// }

type Data struct {
	Fahrenheit string `xml:"CelsiusToFahrenheitResult" json:"fahrenheit"`
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

	celcius := "1059"

	url := "https://www.w3schools.com/xml/tempconvert.asmx?wsdl="

	payload := strings.NewReader("<?xml version=\"1.0\" encoding=\"utf-8\"?>\n<soap12:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap12=\"http://www.w3.org/2003/05/soap-envelope\">\n  <soap12:Body>\n    <CelsiusToFahrenheit xmlns=\"https://www.w3schools.com/xml/\">\n      <Celsius>" + celcius + "</Celsius>\n    </CelsiusToFahrenheit>\n  </soap12:Body>\n</soap12:Envelope>")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "620f29af-ccfe-433b-8e91-25e1bc42acd4")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Printf("res.StatusCode: %d\n", res.StatusCode)
		panic("error ")
	}

	defer res.Body.Close()

	fmt.Printf("res.StatusCode: %d\n", res.StatusCode)

	respbody, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(res)
	// fmt.Println(string(respbody))
	body := string(respbody)

	// // var body = `<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><CelsiusToFahrenheitResponse xmlns="https://www.w3schools.com/xml/"><CelsiusToFahrenheitResult></CelsiusToFahrenheitResult></CelsiusToFahrenheitResponse>`

	body = strings.Replace(body, "<?xml version=\"1.0\" encoding=\"utf-8\"?>", "", -1)
	body = strings.Replace(body, "<soap:Envelope xmlns:soap=\"http://www.w3.org/2003/05/soap-envelope\" xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\">", "", -1)
	body = strings.Replace(body, string("<?xml version=\"1.0\" encoding=\"UTF-8\"?>"), "", -1)
	body = strings.Replace(body, "<soap:Envelope xmlns:soap=\"http://www.w3.org/2003/05/soap-envelope\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\">", "", -1)

	body = strings.Replace(body, string("<soap:Body>"), "", -1)
	body = strings.Replace(body, string("</soap:Body>"), "", -1)
	body = strings.Replace(body, string("</soap:Envelope>"), "", -1)
	body = strings.TrimSpace(body)
	// fmt.Println(body)

	var data Data
	xml.Unmarshal([]byte(body), &data)
	jsonData, _ := json.Marshal(data)
	fmt.Println(string(jsonData))

}
