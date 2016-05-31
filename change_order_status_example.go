package main

import (
	. "github.com/cardpay/go-sdk/cardpay"
	"fmt"
)

/**
    This example describes how to change status of
    order via Cardpay API
 */
func main() {
	// set connection parameters (see full list in Configuration struct definition)
	var config Configuration = Configuration {
		UrlChangeStatus: "https://sandbox.cardpay.com/MI/service/order-change-status",
		PmLogin: "logintopm",
		PmPassword: "qwerty123",
	}

	// set parameters of new order status
	var parameters OrderStatusParameters = OrderStatusParameters {
		Id: "70091",
		StatusTo: "refund",
		Amount: 100.0,
		Reason: "Cancel subscription",
	}

	// call Cardpay API to change status of order
	var result OrderStatusResponse = ChangeOrderStatus(config, parameters)

	// you may use resulting OrderStatusResponse object
	// to determine whether changing was executed or not
	// and other details
	fmt.Println(result.IsExecuted + ":" + result.Details)
}