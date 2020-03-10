package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
	"regexp"
)

// CallOutput struct
type CallOutput struct {
	VendorDetails struct {
		Oui            string `json:"oui"`
		IsPrivate      bool   `json:"isPrivate"`
		CompanyName    string `json:"companyName"`
		CompanyAddress string `json:"companyAddress"`
		CountryCode    string `json:"countryCode"`
	} `json:"vendorDetails"`
	BlockDetails struct {
		BlockFound          bool   `json:"blockFound"`
		BorderLeft          string `json:"borderLeft"`
		BorderRight         string `json:"borderRight"`
		BlockSize           int    `json:"blockSize"`
		AssignmentBlockSize string `json:"assignmentBlockSize"`
		DateCreated         string `json:"dateCreated"`
		DateUpdated         string `json:"dateUpdated"`
	} `json:"blockDetails"`
	MacAddressDetails struct {
		SearchTerm         string   `json:"searchTerm"`
		IsValid            bool     `json:"isValid"`
		VirtualMachine     string   `json:"virtualMachine"`
		Applications       []string `json:"applications"`
		TransmissionType   string   `json:"transmissionType"`
		AdministrationType string   `json:"administrationType"`
		WiresharkNotes     string   `json:"wiresharkNotes"`
		Comment            string   `json:"comment"`
	} `json:"macAddressDetails"`
}

// Output struct
type Output struct {
	MACAddress  string `json:"MACAddress"`
	CompanyName string `json:"CompanyName"`
}

func main() {

	address := flag.String("address", "", "MAC Address to analise (Required)")
	apikey := flag.String("apikey", "", "API key to connect to macaddress.io (Required)")
	outputType := flag.String("output", "text", "Type of output. {text|json}")
	help := flag.Bool("help", false, "Show this help menu.")
	flag.Parse()
	
	if flag.NFlag() == 0 || *help || *address == "" || *apikey == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Regex check of the MAC address
	match, _ := regexp.MatchString("^([0-9a-fA-F][0-9a-fA-F]:){5}([0-9a-fA-F][0-9a-fA-F])$", *address)

	if !match {
		fmt.Println(*address, "is not a valid MAC address")
		os.Exit(1)
	}

	client := resty.New()

	resp, _ := client.R().
		// SetHeader("X-Authentication-Token", *apikey).
		// SetQueryString("&output=json&search=" + *address).
		SetQueryString("apiKey=" + *apikey + "&output=json&search=" + *address).
		Get("https://api.macaddress.io/v1")

	// Convert response body to Todo struct
	var callOutputStruct CallOutput
	json.Unmarshal(resp.Body(), &callOutputStruct)

	output := Output{
		MACAddress:  *address,
		CompanyName: callOutputStruct.VendorDetails.CompanyName,
	}

	switch *outputType {
	case "json":
		outputM, _ := json.Marshal(output)
		fmt.Println(string(outputM))
	default:
		fmt.Println("MAC Address:", output.MACAddress, "\nCompany Name:", output.CompanyName)
	}
}
