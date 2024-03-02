package main

import (
	"bufio"
	"crewFinder/command"
	"crewFinder/db"
	"crewFinder/encryption"
	"crewFinder/httpServ"
	"fmt"
	"os"
	"strings"
	"time"

	"net/http"
)

// Initialize http handler functions, connection to database, pull values from cot.conf, setup the parser, connect to twilio, and begin serving requests
func main() {
	http.HandleFunc("/", httpServ.ReceiveText)
	http.HandleFunc("/status", httpServ.ReceiveTest)
	encryption.InitConf()
	db.DBAdminConnect()
	command.ParserSetup()

	go http.ListenAndServe(":3000", nil)

	phonePrefix := "%2b"
	for {
		defaultPhone := "11234567890"
		var inText string
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("\n\nEnter Text: ")
		if scanner.Scan() {
			inText = scanner.Text()
		}

		// If user input phone number, parse it for use
		inputPhone, request := splitStringByFirstColon(inText)
		if inputPhone != "" {
			// Need to prepend a 1 to non-default phones to pass verification
			defaultPhone = "1" + inputPhone
			inText = request
		}
		defaultPhone = phonePrefix + defaultPhone

		fmt.Printf("Phone number: %s\nMessage: %s\n", defaultPhone, inText)

		response := command.ValidateAndParse(inText, defaultPhone, time.Now().UnixMilli())
		fmt.Printf("\nRESPONSE (%d chars): %s\n", len(response), response)
	}

}

// Used to allow us to specify what phone number we want to message as
func splitStringByFirstColon(input string) (string, string) {
	index := strings.Index(input, ":")
	if index == -1 {
		// Return empty strings if colon is not found
		return "", ""
	}

	phone_number := input[:index]
	request := input[index+1:]

	return phone_number, request
}
