package cardpay

import (
	"testing"
	"strings"
)

func TestGetSha512FromString(t *testing.T) {
	type TestData struct {
		GivenData string
		GivenSalt string
		Expected  string
	}

	var testData = []TestData {
		{ "Test", "Salt",
			"e5b3d99adbaee54ff8375a606cf0ab4a5052d03f6cd2129c33002fd799962f80d44418ea9b584f5027922ec71acf3f2e8c704bdf3e300bbc6369d827b4e519ce" },
		{ " ", " ",
			"16b7aa7f7e549ba129c776bb91ce1e692da103271242d44a9bc145cf338450c90132496ead2530f527b1bd7f50544f37e7d27a2d2bbb58099890aa320f40aca9" },
		{ "Give me my hash", "Take salt",
			"05ab467e0767071e9745dca127ac683d35fd3f77b8ef219e7130c02e39f39b58c083df861504f14354037f1b4569b732f7d4aa35a491eca4650f3c1b00b1a88a" },
	}

	for _, data := range testData {
		testResult := GetSha512FromString(data.GivenData, data.GivenSalt)

		if testResult != data.Expected {
			t.Error(
				"For", data.GivenData + " and " + data.GivenSalt,
				"expected", data.Expected,
				"got", testResult,
			)
		}
	}
}

func TestGetSha256FromString(t *testing.T) {
	type TestData struct {
		GivenData string
		Expected  string
	}

	var testData = []TestData {
		{ "Test", "532eaabd9574880dbf76b9b8cc00832c20a6ec113d682299550d7a6e0f345e25" },
		{ " ", "36a9e7f1c95b82ffb99743e0c5c4ce95d83c9a430aac59f84ef3cbfab6145068" },
		{ "Some long string", "3d3a717ba0928cb1720d9e0db602fa254a43921a5d0b6df25a2171bbb0ccb119" },
	}

	for _, data := range testData {
		testResult := GetSha256FromString(data.GivenData)

		if testResult != data.Expected {
			t.Error(
				"For", data.GivenData,
				"expected", data.Expected,
				"got", testResult,
			)
		}
	}
}

func TestConvertStructToXmlString(t *testing.T) {
	type TestData struct {
		GivenData interface{}
		Expected  string
	}

	var testData = []TestData {
		{ Billing { Country: "UK" }, `<billing country="UK" zip="" city="" street="" phone=""></billing>` },
		{ Shipping { Country: "RU" }, `<shipping country="RU"></shipping>` },
	}

	for _, data := range testData {
		testResult := ConvertStructToXmlString(data.GivenData)

		if testResult != data.Expected {
			t.Error(
				"For", data.GivenData,
				"expected", data.Expected,
				"got", testResult,
			)
		}
	}
}

func TestEncodeStringToBase64(t *testing.T) {
	type TestData struct {
		GivenData string
		Expected  string
	}

	var testData = []TestData {
		{ "<order id=\"123\"></order>", "PG9yZGVyIGlkPSIxMjMiPjwvb3JkZXI+" },
		{ " ", "IA==" },
		{ "Test", "VGVzdA==" },
	}

	for _, data := range testData {
		testResult := EncodeStringToBase64(data.GivenData)

		if testResult != data.Expected {
			t.Error(
				"For", data.GivenData,
				"expected", data.Expected,
				"got", testResult,
			)
		}
	}
}

func TestMakePayment(t *testing.T) {
	type TestData struct {
		GivenOrder  Order
		GivenConfig Configuration
		Expected    string
	}

	var testData = []TestData {
		{
			Order {
			    WalletId: 99999999,
			    Number: "111",
				Description: "Some product",
				Currency: "USD",
				Amount: 123.00,
				Email: "john.smith@example.com",
		    }, Configuration {
			    UrlPayment: "https://sandbox.cardpay.com/MI/cardpayment.html",
			    SecretWord: "sdRFsD",
		    }, "Unknown shop" }}

	for _, data := range testData {
		testResult := MakePayment(data.GivenConfig, data.GivenOrder)

		if !strings.Contains(testResult, data.Expected) {
			t.Error(
				"For", data.GivenOrder,
				"expected", data.Expected,
				"got", testResult,
			)
		}
	}
}

func TestChangeOrderStatus(t *testing.T) {
	type TestData struct {
		GivenConfig     Configuration
		GivenParameters OrderStatusParameters
		Expected OrderStatusResponse
	}

	var testData = []TestData {
		{
            Configuration {
			    PmLogin: "abcd",
				PmPassword: "qwerty",
				UrlChangeStatus: "https://sandbox.cardpay.com/MI/service/order-change-status",
		    }, OrderStatusParameters {
			    Id: "123456",
			    StatusTo: "refund",
			    Amount: 100.0,
			    Reason: "I want it back!",
			}, OrderStatusResponse {
			    Details: "Login Failed",
		    },
		}}

	for _, data := range testData {
		testResult := ChangeOrderStatus(data.GivenConfig, data.GivenParameters)

		if testResult.Details != data.Expected.Details {
			t.Error(
				"For", data.GivenParameters,
				"expected", data.Expected.Details,
				"got", testResult.Details,
			)
		}
	}
}

func TestGetPayments(t *testing.T) {
	type TestData struct {
		GivenConfig Configuration
		Expected    bool
	}

	var testData = []TestData {
		{
			Configuration {
				UrlRestApi: "https://sandbox.cardpay.com/MI/api/v2",
				PmLogin: "login",
				PmPassword: "pass",
			}, false,
		}}

	for _, data := range testData {
		testResult := GetPayments(data.GivenConfig)

		if testResult.HasMore {
			t.Error(
				"For", data.GivenConfig,
				"expected", data.Expected,
				"got", testResult.HasMore,
			)
		}
	}
}

func TestGetRefunds(t *testing.T) {
	type TestData struct {
		GivenConfig Configuration
		Expected    bool
	}

	var testData = []TestData {
		{
			Configuration {
				UrlRestApi: "https://sandbox.cardpay.com/MI/api/v2",
				PmLogin: "login",
				PmPassword: "pass",
			}, false,
		}}

	for _, data := range testData {
		testResult := GetRefunds(data.GivenConfig)

		if testResult.HasMore {
			t.Error(
				"For", data.GivenConfig,
				"expected", data.Expected,
				"got", testResult.HasMore,
			)
		}
	}
}
