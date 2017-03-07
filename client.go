package abolapi

import (
	"context"
	"fmt"
	"strings"
)

type Client struct {
	activationKey string
	login         string
	password      string
	soapClient    *AbolapiSoap12
}

func (c *Client) SimpleShipment(ctx context.Context, ss *SimpleShipment) (*SimpleShipmentResponse, error) {
	in := AbolSimpleShipmentSoapIn{
		Parameters: AbolSimpleShipmentElement{
			AbolApiSimpleShipment: AbolApiSimpleShipment{
				AbolSimpleShip: &AbolApiSimpleShipmentRequest{
					Request: &AbolApiSimpleShipment_Request{
						Authentication: &Authentication{
							ActivationKey: c.activationKey,
							LoginName:     c.login,
							Password:      c.password,
						},
						SimpleShipment: ss,
					},
				},
			},
		},
	}

	resp, err := c.soapClient.AbolSimpleShipment(ctx, &in)
	if err != nil {
		if sf, ok := err.(SOAPFault); ok {
			return nil, fmt.Errorf("SimpleShipment failed with SOAPFault: %d - %s\n%v", sf.Code, sf.Detail, err)
		}
		return nil, err
	}

	ssResp := resp.AbolApiSimpleShipmentResponse.SimpleShipmentResult
	if ssResp.Error != nil && ssResp.Error.Code != 0 {
		return nil, fmt.Errorf("SimpleShipmentError %d: %q", ssResp.Error.Code, strings.TrimSuffix(ssResp.Error.Message, "\n"))
	}

	return ssResp.SimpleShipment, nil
}

func NewClient(activationKey, login, password string) *Client {
	soapClient := NewAbolapiSoap12()
	return &Client{activationKey, login, password, soapClient}
}
