package cardpay

import "encoding/xml"

/**
    Represents configuration of connection with Cardpay API
 */
type Configuration struct {
	UrlPayment      string
	UrlChangeStatus string
	UrlRestApi      string

	PmLogin         string
	PmPassword      string

	WalletId        int
	SecretWord      string
}

/**
    Represents order
 */
type Order struct {
	XMLName        xml.Name  `xml:"order"`
	WalletId       int       `xml:"wallet_id,attr"`
	Number         string    `xml:"number,attr"`
	Description    string    `xml:"description,attr"`
	Currency       string    `xml:"currency,attr"`
	Amount         float64   `xml:"amount,attr"`
	Email          string    `xml:"email,attr"`
	CustomerId     string    `xml:"customer_id,attr,omitempty"`
	IsTwoPhase     bool      `xml:"is_two_phase,attr,omitempty"`
	RecurringBegin bool      `xml:"recurring_begin,attr,omitempty"`
	RecurringId    string    `xml:"recurring_id,attr,omitempty"`
	Note           string    `xml:"note,attr,omitempty"`
	ReturnUrl      string    `xml:"return_url,attr,omitempty"`
	SuccessUrl     string    `xml:"success_url,attr,omitempty"`
	DeclineUrl     string    `xml:"decline_url,attr,omitempty"`
	CancelUrl      string    `xml:"cancel_url,attr,omitempty"`

	Card           []Card     // defining as array allows tag to be optional
	Billing        []Billing  // defining as array allows tag to be optional
	Shipping       []Shipping // defining as array allows tag to be optional
	Items          []Item     `xml:"items>."` // defining as array allows tag to be optional
}

/**
    Represents payment card
 */
type Card struct {
	XMLName xml.Name `xml:"card"`
	Num     int      `xml:"num,attr"`
	Holder  string   `xml:"holder,attr"`
	Cvv     int      `xml:"cvv,attr"`
	Expires string   `xml:"expires,attr"`
}

/**
    Represents billing address
 */
type Billing struct {
	XMLName   xml.Name `xml:"billing"`
	Country   string   `xml:"country,attr"`
	State     string   `xml:"state,attr,omitempty"`
	Zip       string   `xml:"zip,attr"`
	City      string   `xml:"city,attr"`
	Street    string   `xml:"street,attr"`
	Phone     string   `xml:"phone,attr"`
}

/**
    Represents shipping address
 */
type Shipping struct {
	XMLName   xml.Name `xml:"shipping"`
	Country   string   `xml:"country,attr"`
	State     string   `xml:"state,attr,omitempty"`
	Zip       string   `xml:"zip,attr,omitempty"`
	City      string   `xml:"city,attr,omitempty"`
	Street    string   `xml:"street,attr,omitempty"`
	Phone     string   `xml:"phone,attr,omitempty"`
}

/**
    Represents order item
 */
type Item struct {
	XMLName     xml.Name `xml:"item"`
	Name        string   `xml:"name,attr"`
	Description string   `xml:"description,attr,omitempty"`
	Count       int      `xml:"count,attr,omitempty"`
	Price       float64  `xml:"price,attr,omitempty"`
}

/**
    Represents parameters used to change order status
 */
type OrderStatusParameters struct {
	ClientLogin    string
	ClientPassword string
	Id             string
	StatusTo       string
	Amount         float64
	Reason         string
}

/**
    Represents response of changing order status
 */
type OrderStatusResponse struct {
	XMLName    xml.Name        `xml:"response"`
	IsExecuted string          `xml:"is_executed,attr"`
	Details    string          `xml:"details,attr,omitempty"`
	Order      OrderStatusOrder
}

/**
    Represents order that comes with request to change order status
 */
type OrderStatusOrder struct {
	XMLName         xml.Name `xml:"order"`
	Id              string   `xml:"id,attr"`
	StatusTo        string   `xml:"status_to,attr"`
	Currency        string   `xml:"currency,attr,omitempty"`
	RefundAmount    string   `xml:"refund_amount,attr,omitempty"`
	RemainingAmount string   `xml:"remaining_amount,attr,omitempty"`
}

/**
    Represents response from REST API
 */
type RestApiPaymentResponse struct {
	HasMore bool                  `json:"has_more"`
	Data    []RestApiPaymentsData `json:"data"`
}

/**
   Represents data that comes with response from REST API
 */
type RestApiPaymentsData struct {
	Id              string  `json:"id"`
	Number          string  `json:"number"`
	State           string  `json:"state"`
	Date            int64   `json:"date"`
	CustomerId      string  `json:"customerId"`
	DeclineReason   string  `json:"declineReason"`
	DeclineCode     string  `json:"declineCode"`
	AuthCode        string  `json:"authCode"`
	Is3d            bool    `json:"is3d"`
	Currency        string  `json:"currency"`
	Amount          float64 `json:"amount"`
	RefundedAmount  float64 `json:"refundedAmount"`
	Note            string  `json:"note"`
	Email           string  `json:"email"`
	Rrn             string  `json:"rrn"`
	OriginalOrderId string  `json:"originalOrderId"`
}
