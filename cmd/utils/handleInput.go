/*
Copyright © 2024 Epyklab contact@epyklab.com
*/
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

func SetEntryTypeByTag(tag string, author string) {

	switch tag := tag; {

	case tag == "tagIPAddress":
		fmt.Print("Name: ")
		name := getUserInput()

		fmt.Print("Description: ")
		description := getUserInput()

		fmt.Print("IP: ")
		ip := getUserInput()

		//BUG: author value is not being accessed here.
		obj := entries.IPAddress{Name: name, Author: viper.GetString("team.self"), Description: description, IP: ip, Tags: tag}
		marshalIn(obj)

	case tag == "tagDomainName":
		fmt.Print("Name: ")
		name := getUserInput()

		fmt.Print("Description: ")
		description := getUserInput()

		fmt.Print("Domain Name: ")
		domainname := getUserInput()

		obj := entries.DomainName{Author: viper.GetString("team.self"), Name: name, Description: description, Domain: domainname}
		marshalIn(obj)

	case tag == "tagUsername":
		fmt.Print("Name: ")
		name := getUserInput()

		fmt.Print("Description: ")
		description := getUserInput()

		fmt.Print("Username: ")
		username := getUserInput()

		obj := entries.Username{Author: viper.GetString("team.self"), Name: name, Description: description, Username: username}
		marshalIn(obj)

	case tag == "tagPassword":
		fmt.Print("Name: ")
		name := getUserInput()

		fmt.Print("Description: ")
		description := getUserInput()

		fmt.Print("Password: ")
		password := getUserInput()

		obj := entries.Password{Author: viper.GetString("team.self"), Name: name, Description: description, Password: password}
		marshalIn(obj)

	case tag == "tagSTSToken":
		fmt.Print("Name: ")
		name := getUserInput()

		fmt.Print("Description: ")
		description := getUserInput()

		fmt.Print("STS Token: ")
		ststoken := getUserInput()

		obj := entries.STSToken{Author: viper.GetString("team.self"), Name: name, Description: description, Token: ststoken}
		marshalIn(obj)

	case tag == "tagAPIKey":
		fmt.Print("Name: ")
		name := getUserInput()

		fmt.Print("Description: ")
		description := getUserInput()

		fmt.Print("API Key: ")
		apikey := getUserInput()

		obj := entries.APIKey{Author: viper.GetString("team.self"), Name: name, Description: description, Key: apikey}
		marshalIn(obj)

	case tag == "tagCertificate":
		fmt.Print("Name: ")
		name := getUserInput()

		fmt.Print("Description: ")
		description := getUserInput()

		fmt.Print("Certificate: ")
		certificate := getUserInput()

		obj := entries.Certificate{Author: viper.GetString("team.self"), Name: name, Description: description, Certificate: certificate}
		marshalIn(obj)

	case tag == "tagVulnerability":
		fmt.Print("Name: ")
		name := getUserInput()

		fmt.Print("Description: ")
		description := getUserInput()

		fmt.Print("Vulnerability: ")
		vuln := getUserInput()

		obj := entries.Vulnerability{Author: viper.GetString("team.self"), Name: name, Description: description, Vuln: vuln}
		marshalIn(obj)

	case tag == "tagExploit":
		fmt.Print("Name: ")
		name := getUserInput()

		fmt.Print("Description: ")
		description := getUserInput()

		fmt.Print("Exploit: ")
		exploit := getUserInput()

		obj := entries.Exploit{Author: viper.GetString("team.self"), Name: name, Description: description, Exploit: exploit}
		marshalIn(obj)

	case tag == "tagPayload":
		fmt.Print("Name: ")
		name := getUserInput()

		fmt.Print("Description: ")
		description := getUserInput()

		fmt.Print("Payload: ")
		payload := getUserInput()

		obj := entries.Payload{Author: viper.GetString("team.self"), Name: name, Description: description, Payload: payload}
		marshalIn(obj)

	case tag == "tagLog":
		fmt.Print("Name: ")
		name := getUserInput()

		fmt.Print("Description: ")
		description := getUserInput()

		fmt.Print("Log: ")
		log := getUserInput()

		obj := entries.Log{Author: viper.GetString("team.self"), Name: name, Description: description, Log: log}
		marshalIn(obj)

	case tag == "tagReport":
		fmt.Print("Name: ")
		name := getUserInput()

		fmt.Print("Description: ")
		description := getUserInput()

		fmt.Print("Report: ")
		report := getUserInput()

		obj := entries.Report{Author: viper.GetString("team.self"), Name: name, Description: description, Report: report}
		marshalIn(obj)

	case tag == "tagConfiguration":
		fmt.Print("Name: ")
		name := getUserInput()

		fmt.Print("Description: ")
		description := getUserInput()

		fmt.Print("Configuration: ")
		configuration := getUserInput()

		obj := entries.Configuration{Author: viper.GetString("team.self"), Name: name, Description: description, Config: configuration}
		marshalIn(obj)

	case tag == "tagBackup":
		fmt.Print("Name: ")
		name := getUserInput()

		fmt.Print("Description: ")
		description := getUserInput()

		fmt.Print("Backup: ")
		backup := getUserInput()

		obj := entries.Backup{Author: viper.GetString("team.self"), Name: name, Description: description, Backup: backup}
		marshalIn(obj)

	case tag == "tagMalwareSample":
		fmt.Print("Name: ")
		name := getUserInput()

		fmt.Print("Description: ")
		description := getUserInput()

		fmt.Print("Malware Sample: ")
		malsample := getUserInput()

		obj := entries.MalwareSample{Author: viper.GetString("team.self"), Name: name, Description: description, Sample: malsample}
		marshalIn(obj)

	case tag == "":
		fmt.Println("No tag provided")

	}

}

func getUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
