package cardpay

import (
	"encoding/xml"
	"encoding/base64"
	"crypto/sha512"
	"encoding/hex"
	"net/http"
	"net/url"
	"io/ioutil"
	"crypto/sha256"
	"encoding/json"
	"strconv"
)

/**
    Get salted SHA512 hash
 */
func GetSha512FromString(initial string, salt string) string {
	hash := sha512.Sum512([]byte(initial + salt))

	return hex.EncodeToString(hash[:64])
}

/**
    Get simple (without salt) SHA256 hash
 */
func GetSha256FromString(initial string) string {
	hash := sha256.Sum256([]byte(initial))

	return  hex.EncodeToString(hash[:32])
}

/**
    Convert domain object to XML string
 */
func ConvertStructToXmlString(data interface{}) string {
	result, _ := xml.MarshalIndent(data, "", " ")

	return string(result)
}

/**
    Encode string to Base64-encoded string
 */
func EncodeStringToBase64(initial string) string {
	return base64.StdEncoding.EncodeToString([]byte(initial))
}

/**
    Make payment call to API
 */
func MakePayment(config Configuration, order Order) string {
	order.WalletId = config.WalletId

	orderXml := ConvertStructToXmlString(order)

	resp, _ := http.PostForm(config.UrlPayment,
		url.Values{
			"orderXML": {EncodeStringToBase64(orderXml)},
			"sha512": {GetSha512FromString(orderXml, config.SecretWord)}})

	defer resp.Body.Close()

	contents, _ := ioutil.ReadAll(resp.Body)

	return string(contents);
}

/**
    Change status of order
 */
func ChangeOrderStatus(config Configuration, parameters OrderStatusParameters) OrderStatusResponse {
	resp, _ := http.PostForm(config.UrlChangeStatus, map[string][]string {
		"client_login": { config.PmLogin },
		"client_password": { GetSha256FromString(config.PmPassword) },
		"id": { parameters.Id },
		"status_to": { parameters.StatusTo },
		"amount": { strconv.FormatFloat(parameters.Amount, 'f', 6, 64) },
		"reason": { parameters.Reason },
	})

	defer resp.Body.Close()

	contents, _ := ioutil.ReadAll(resp.Body)

	result := OrderStatusResponse{}

	xml.Unmarshal(contents, &result)

	return result;
}

/**
    Get list of payments
 */
func GetPayments(config Configuration) RestApiPaymentResponse {
    return CallRestApi(config, "/payments")
}

/**
    Get list of refunds
 */
func GetRefunds(config Configuration) RestApiPaymentResponse {
    return CallRestApi(config, "/refunds")
}

/**
    Make call to REST API
 */
func CallRestApi(config Configuration, url string) RestApiPaymentResponse {
	client := http.Client{}

	request, _ := http.NewRequest("GET", config.UrlRestApi + url, nil)
	request.Header.Set("Authorization",
		"Basic " + EncodeStringToBase64(config.PmLogin + ":" + config.PmPassword))

	response, _ := client.Do(request)

	defer response.Body.Close()

	contents, _ := ioutil.ReadAll(response.Body)

	result := RestApiPaymentResponse{}

	json.Unmarshal(contents, &result)

	return result;
}