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

func (c *Client) authentication() *Authentication {
	return &Authentication{
		ActivationKey: c.activationKey,
		LoginName:     c.login,
		Password:      c.password,
	}
}

func (c *Client) VoidPackage(ctx context.Context, trackingNo string) (*AbolVoidResult, error) {
	in := AbolVoidPackageSoapIn{
		Parameters: AbolVoidPackageElement{
			AbolApiVoidRequest: AbolApiVoidRequest{
				Authentication: &AbolApiVoid_Authentication{
					ActivationKey: c.activationKey,
					LoginName:     c.login,
					Password:      c.password,
				},
				Void: &AbolApiVoid_TrackingNo{
					TrackingNumber: trackingNo,
				},
			},
		},
	}

	resp, err := c.soapClient.AbolVoidPackage(ctx, &in)
	if err = handleSoapErr("VoidPackage", err); err != nil {
		return nil, err
	}
	vResp := resp.AbolApiVoidResponse.AbolVoidResult
	if vResp.Error != nil && vResp.Error.Code != 0 {
		return nil, fmt.Errorf("VoidPackage Error %d: %q", vResp.Error.Code, strings.TrimSuffix(vResp.Error.Message, "\n"))
	}

	return vResp, nil
}

func (c *Client) CloseOut(ctx context.Context, trackingNo string) (*AbolCloseoutResult, error) {

	in := AbolCloseOutSoapIn{
		Parameters: AbolCloseOutElement{
			AbolApiCloseOutRequest: AbolApiCloseOutRequest{
				Request: &AbolApiCloseOut_Request{
					Authentication: c.authentication(),
					Closeout: &Closeout{
						RequestType: CloseoutRequestType_TRACKINGNBR,
						CloseOutTrackingNbrs: &CloseOutTrackingNbrs{
							TrackingNbrs: []string{trackingNo},
						},
					},
				},
			},
		},
	}

	resp, err := c.soapClient.AbolCloseOut(ctx, &in)
	if err != nil {
		if sf, ok := err.(SOAPFault); ok {
			return nil, fmt.Errorf("CloseOut failed with SOAPFault: %d - %s\n%v", sf.Code, sf.Detail, err)
		}
		return nil, err
	}
	coResp := resp.AbolApiCloseOutResponse.AbolCloseoutResult
	if coResp.Error != nil && coResp.Error.Code != 0 {
		return nil, fmt.Errorf("CloseOut Error %d: %q", coResp.Error.Code, strings.TrimSuffix(coResp.Error.Message, "\n"))
	}

	return coResp, nil
}

func handleSoapErr(action string, err error) error {
	if err != nil {
		if sf, ok := err.(SOAPFault); ok {
			return fmt.Errorf("%s failed with SOAPFault: %d - %s\n%v", action, sf.Code, sf.Detail, err)
		}
	}
	return err
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
		return nil, fmt.Errorf("SimpleShipment Error %d: %q", ssResp.Error.Code, strings.TrimSuffix(ssResp.Error.Message, "\n"))
	}

	return ssResp.SimpleShipment, nil
}

func NewClient(activationKey, login, password string) *Client {
	soapClient := NewAbolapiSoap12()
	return &Client{activationKey, login, password, soapClient}
}
