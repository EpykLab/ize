package userio

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	entries "github.com/Epyklab/ize/cmd/entries"
	"github.com/spf13/viper"
)

type Marshalable interface{}

func ConvertToJSON(input Marshalable) ([]byte, error) {
	return json.Marshal(input)
}

func marshalIn(obj Marshalable) {

	// Convert IPAddress to JSON
	jsonData, err := ConvertToJSON(obj)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
	}
	entries.MakeMongoEntry(jsonData)
}

func SetEntryTypeByTag(tag string) {

	switch tag := tag; {

	case tag == "tagIPAddress":
		fmt.Print("Description: ")
		description := getUserInput()

		fmt.Print("IP: ")
		ip := getUserInput()

		//BUG: author value is not being accessed here.
		obj := entries.IPAddress{Author: viper.GetString("self"), Description: description, IP: ip, Tags: tag}
		marshalIn(obj)

	case tag == "tagDomainName":
		fallthrough

	case tag == "tagUsername":
		fallthrough

	case tag == "tagPassword":
		fallthrough

	case tag == "tagSTSToken":
		fallthrough

	case tag == "tagAPIKey":
		fallthrough

	case tag == "tagCertificate":
		fallthrough

	case tag == "tagVulnerability":
		fallthrough

	case tag == "tagExploit":
		fallthrough

	case tag == "tagPayload":
		fallthrough

	case tag == "tagLog":
		fallthrough

	case tag == "tagReport":
		fallthrough

	case tag == "tagConfiguration":
		fallthrough

	case tag == "tagBackup":
		fallthrough

	case tag == "tagMalwareSample":
		fallthrough

	case tag == "":
		fmt.Println("No tag provided")

	}

}

func getUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
