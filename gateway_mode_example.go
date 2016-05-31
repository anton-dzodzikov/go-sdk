package main

import (
    . "github.com/cardpay/go-sdk/cardpay"
)

/**
    This is a short example of integration with Cardpay
    processing API in Gateway Mode
 */
func main() {
    // set connection parameters (see full list in Configuration struct definition)
    var config Configuration = Configuration{
        UrlPayment: "https://sandbox.cardpay.com/MI/cardpayment.html",
        WalletId: 1234,
        SecretWord: "sdCsFTwSasd",
    }

    // set parameters of order (see full list in Order struct definition)
    var order Order = Order {
        Number: "111",
        Description: "Some product",
        Currency: "USD",
        Amount: 123.40,
        Email: "mr.dzodzikov@gmail.com",
        // to begin recurrent payments (optional)
        RecurringBegin: true,
        Card: []Card { Card {
            Num: 4000000000000077,
            Holder: "JOHN SMITH",
            Cvv: 843,
            Expires: "10/2018",
        }},
        Billing: []Billing { Billing {
            Country: "RU",
            State: "Primorsky Kray",
            Zip: "690089",
            City: "Vladivostok",
            Street: "Kovalchuka 9d",
            Phone: "89242444126",
        }},
        // Shipping is optional in Gateway Mode
        Shipping: []Shipping { Shipping {
            Country: "RU",
            State: "Primorsky Kray",
            Zip: "690089",
            City: "Vladivostok",
            Street: "Kovalchuka 9d",
            Phone: "89242444126",
        }},
        // Items are optional in Gateway Mode
        Items: []Item { Item {
            Name: "Book",
        }},
    }

    // call Cardpay API
    var result string = MakePayment(config, order)

    // use resulting xml-string to proceed payment
}