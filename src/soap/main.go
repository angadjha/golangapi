package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Sample SOAP request payload
	soapRequestPayload := `
	<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
		<soap:Body>
			<YourSOAPRequestPayloadHere>
				<!-- Add your SOAP request payload here -->
			</YourSOAPRequestPayloadHere>
		</soap:Body>
	</soap:Envelope>`

	// Set the URL for the SOAP service
	url := "https://example.com/your-soap-service-endpoint"

	// Create a SOAP request
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(soapRequestPayload))
	if err != nil {
		fmt.Println("Error creating SOAP request:", err)
		return
	}

	// Set SOAPAction header and Content-Type
	req.Header.Set("SOAPAction", "YourSOAPActionHere")
	req.Header.Set("Content-Type", "text/xml;charset=UTF-8")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending SOAP request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print the SOAP response
	fmt.Println("SOAP Response:")
	fmt.Println(string(body))
}
