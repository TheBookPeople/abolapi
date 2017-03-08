package abolapi

import (
	"context"
	"log"
	"os"
	"testing"
	"time"
)

var (
	activationKey = getEnv("ABOLAPI_TEST_ACT_KEY")
	login         = getEnv("ABOLAPI_TEST_LOGIN")
	password      = getEnv("ABOLAPI_TEST_PASSWORD")
)

func TestVoidPackage(t *testing.T) {
	//	VoidPackage(ctx context.Context, trackingNo string) (*AbolVoidResult, error) {
	//c := newClient()
}
func TestCloseOut(t *testing.T) {
	//(ctx context.Context, trackingNo string) (*AbolCloseoutResult, error) {
	//	c := newClient()
}
func TestHandleSoapErr(t *testing.T) {
	//(action string, err error) error {
	//c := newClient()
}

func TestSimpleShipment(t *testing.T) {
	//ctx context.Context, ss *SimpleShipment) (*SimpleShipmentResponse, error) {
	c := newClient()

	simpleShipment := &SimpleShipment{
		/* ServiceCode options:
		Spring Packet Light 1356
		Spring Packet 1351
		Spring Packet Plus 1352
		Spring Parcel 1353				// What happened to 1355?
		Spring Express 1354 */
		ServiceCode:        "1351",                             // TODO
		BillType:           100,                                // 100 = Not specified // TODO
		DriverInstructions: "delivery instructions for driver", // TODO
		LabelType:          6,                                  // PNG
		LabelSize:          0,                                  // 4x6
		DimensionUnit:      "CM",
		WeightUnit:         "KG",
		ShipDate:           &time.Time{},
		DeclaredValue:      1.99,  // TODO
		CurrencyCode:       "GBP", // TODO
		Length:             25,
		Height:             60,
		Width:              7,
		PackagingType:      105,
		ReferenceNumber:    "ref1234",
		Weight:             2,
		ContentDescription: "Books",

		//						RegulatoryControlType: 100,   // 100 = Not specified // TODO
		ShipFrom: &Address{
			Company:      "The Book People Ltd",
			AddressLine1: "Parc Menai",
			City:         "Bangor",
			CountryCode:  "GB",
			Zip:          "LL57 4FB",
		},
		ShipTo: &Address{
			Company:         "Ricoh Hong Kong Ltd",
			AddressLine1:    "21/F. One Kowloon",
			AddressLine2:    "1 Wang Yuen Street",
			City:            "Kowloon Bay",
			CountryCode:     "HK",
			Phone:           "(852) 2893-0022", //TODO
			EMail:           "a@b.com",
			ResidentialFlag: true,
			IsPOBox:         false,
		},
		SimpleExportLineItemList: &SimpleExportLineItems{
			SimpleExportLineItems: []SimpleExportLineItem{
				SimpleExportLineItem{
					Price:                1.99,
					CountryOfManufacture: "GB",
					Description:          "Superted",
					Code:                 "STED",
					Quantity:             1,
				},
			},
		},
	}
	resp, err := c.SimpleShipment(context.Background(), simpleShipment)
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		name     string
		expected interface{}
		actual   interface{}
	}{
		{"trackingNbr", 22, len(resp.TrackingNbr)},
		//	{"name", "expected", "actual"},
	}

	for _, test := range tests {
		if test.expected != test.actual {
			t.Errorf("Expected %s to be %q but was %q", test.name, test.expected, test.actual)
		}
	}

}

func TestAuthentication(t *testing.T) {
	c := newClient()
	a := c.authentication()
	tests := []struct{ name, expected, actual string }{
		{"activationKey", activationKey, a.ActivationKey},
		{"login", login, a.LoginName},
		{"password", password, a.Password},
		// AltID?
	}
	for _, test := range tests {
		if test.expected != test.actual {
			t.Errorf("Expected %s to be %q but was %q", test.name, test.expected, test.actual)
		}
	}
}

func TestNewClient(t *testing.T) {
	client := newClient()
	if client.soapClient == nil {
		t.Error("Soap client was not initialized")
	}
}

func newClient() *Client {
	return NewClient(activationKey, login, password)
}

func getEnv(env string) string {
	val := os.Getenv(env)
	if val == "" {
		log.Panicf("Environmental variable %q was not set.", env)
	}
	return val
}
