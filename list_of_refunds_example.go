package main

import (
	. "github.com/cardpay/go-sdk/cardpay"
	"fmt"
)

/**
    This example describes how to get list of refunds
    via Cardpay REST API
 */
func main() {
	// set connection parameters (see full list in Configuration struct definition)
	var config Configuration = Configuration {
		UrlRestApi: "https://sandbox.cardpay.com/MI/api/v2",
		PmLogin: "logintopm",
		PmPassword: "qwerty123",
	}

	// call Cardpay API to get list of refunds
	var result RestApiPaymentResponse = GetRefunds(config)

	// you may use resulting RestApiPaymentResponse object
	// to know the detailed information for each payment in
	// the list
	for _, data := range result {
		fmt.Println(data.Number)
		fmt.Println(data.CustomerId)
		fmt.Println(data.Amount)
	}
}