package abolapi

import (
	"context"
	"encoding/base64"
	"encoding/xml"
	"log"
	"time"
)

// Messages

type AbolRateSoapIn struct {
	Parameters AbolRateElement
}

type AbolRateSoapOut struct {
	Parameters AbolRateResponseElement
}

type AbolShipperSignupSoapIn struct {
	Parameters AbolShipperSignupElement
}

type AbolShipperSignupSoapOut struct {
	Parameters AbolShipperSignupResponseElement
}

type AbolValidateAddressSoapIn struct {
	Parameters AbolValidateAddressElement
}

type AbolValidateAddressSoapOut struct {
	Parameters AbolValidateAddressResponseElement
}

type AbolOrderSoapIn struct {
	Parameters AbolOrderElement
}

type AbolOrderSoapOut struct {
	Parameters AbolOrderResponseElement
}

type AbolShipmentSoapIn struct {
	Parameters AbolShipmentElement
}

func (x AbolSimpleShipmentSoapIn) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	log.Println(start)
	var err error
	err = e.Encode(x.Parameters)
	if err != nil {
		return err
	}

	return nil
}

type AbolShipmentSoapOut struct {
	Parameters AbolShipmentResponseElement
}

type AbolTrackPackageSoapIn struct {
	Parameters AbolTrackPackageElement
}

type AbolTrackPackageSoapOut struct {
	Parameters AbolTrackPackageResponseElement
}

type AbolVoidPackageSoapIn struct {
	Parameters AbolVoidPackageElement
}

type AbolVoidPackageSoapOut struct {
	Parameters AbolVoidPackageResponseElement
}

type AbolCloseOutSoapIn struct {
	Parameters AbolCloseOutElement
}

type AbolCloseOutSoapOut struct {
	Parameters AbolCloseOutResponseElement
}

type AbolClassifySoapIn struct {
	Parameters AbolClassifyElement
}

type AbolClassifySoapOut struct {
	Parameters AbolClassifyResponseElement
}

type AbolDutySoapIn struct {
	Parameters AbolDutyElement
}

type AbolDutySoapOut struct {
	Parameters AbolDutyResponseElement
}

type AbolSimpleShipmentSoapIn struct {
	Parameters AbolSimpleShipmentElement
}

type AbolSimpleShipmentSoapOut struct {
	Parameters AbolSimpleShipmentResponseElement `xml:"http://abolsoft.com/webservices/ AbolSimpleShipmentResponse,omitempty"`
}

type AbolSimpleRateSoapIn struct {
	Parameters AbolSimpleRateElement
}

type AbolSimpleRateSoapOut struct {
	Parameters AbolSimpleRateResponseElement
}

type AbolOrderOMSoapIn struct {
	Parameters AbolOrderOMElement
}

type AbolOrderOMSoapOut struct {
	Parameters AbolOrderOMResponseElement
}

// Elements
type AbolRateElement struct {
	Rate_Request Rate_Request
}

type AbolRateResponseElement struct {
	Rate_Response Rate_Response
}

type AbolShipperSignupElement struct {
	AddShipper_Request AddShipper_Request
}

type AbolShipperSignupResponseElement struct {
	AddShipper_Response AddShipper_Response
}

type AbolValidateAddressElement struct {
	AddVal_Request AddVal_Request
}

type AbolValidateAddressResponseElement struct {
	AddVal_Response AddVal_Response
}

type AbolOrderElement struct {
	AbolImportOrderRequest AbolImportOrderRequest
}

type AbolOrderResponseElement struct {
	AbolImportOrderResponse AbolImportOrderResponse
}

type AbolShipmentElement struct {
	AbolApiShipmentRequest AbolApiShipmentRequest
}

type AbolShipmentResponseElement struct {
	AbolApiShipmentResponse AbolApiShipmentResponse
}

type AbolTrackPackageElement struct {
	AbolApiTrackRequest AbolApiTrackRequest
}

type AbolTrackPackageResponseElement struct {
	AbolApiTrackResponse AbolApiTrackResponse
}

type AbolVoidPackageElement struct {
	AbolApiVoidRequest AbolApiVoidRequest
}

type AbolVoidPackageResponseElement struct {
	AbolApiVoidResponse AbolApiVoidResponse
}

type AbolCloseOutElement struct {
	AbolApiCloseOutRequest AbolApiCloseOutRequest
}

type AbolCloseOutResponseElement struct {
	AbolApiCloseOutResponse AbolApiCloseOutResponse
}

type AbolClassifyElement struct {
	AbolClassify_Request AbolClassify_Request
}

type AbolClassifyResponseElement struct {
	AbolClassify_Response AbolClassify_Response
}

type AbolDutyElement struct {
	AbolDuty_Request AbolDuty_Request
}

type AbolDutyResponseElement struct {
	AbolDuty_Response AbolDuty_Response
}

type AbolSimpleShipmentElement struct {
	XMLName               string                `xml:"http://abolsoft.com/webservices/ AbolSimpleShipment,omitempty"`
	AbolApiSimpleShipment AbolApiSimpleShipment `xml:"AbolSimpleShipment,omitempty"`
}

type AbolSimpleShipmentResponseElement struct {
	AbolApiSimpleShipmentResponse AbolApiSimpleShipmentResponse `xml:"http://abolsoft.com/webservices/ AbolSimpleShipmentResponse,omitempty"`
}

type AbolSimpleRateElement struct {
	AbolSimpleRate AbolSimpleRate
}

type AbolSimpleRateResponseElement struct {
	AbolRateSimpleResponse AbolRateSimpleResponse
}

type AbolOrderOMElement struct {
	AbolImportOrderRequest AbolImportOrderRequest
}

type AbolOrderOMResponseElement struct {
	AbolImportOrderResponse AbolImportOrderResponse
}

// Simple types
type AmountUnit string

const AmountUnit_m3 AmountUnit = "m3"
const AmountUnit_cm3 AmountUnit = "cm3"
const AmountUnit_in3 AmountUnit = "in3"
const AmountUnit_l AmountUnit = "l"
const AmountUnit_pcs AmountUnit = "pcs"
const AmountUnit_doz AmountUnit = "doz"
const AmountUnit_gr AmountUnit = "gr"
const AmountUnit_pairs AmountUnit = "pairs"
const AmountUnit_jwl AmountUnit = "jwl"
const AmountUnit_m2 AmountUnit = "m2"

type RequestTypeEnum string

const RequestTypeEnum_addShipper RequestTypeEnum = "addShipper"
const RequestTypeEnum_addSite RequestTypeEnum = "addSite"

type OrderRequestType string

const OrderRequestType_Add OrderRequestType = "Add"
const OrderRequestType_Edit OrderRequestType = "Edit"
const OrderRequestType_Delete OrderRequestType = "Delete"

type DuplicateOrderType string

const DuplicateOrderType_NONE DuplicateOrderType = "NONE"
const DuplicateOrderType_UNKNOWN DuplicateOrderType = "UNKNOWN"
const DuplicateOrderType_UPDATE DuplicateOrderType = "UPDATE"
const DuplicateOrderType_DELETE DuplicateOrderType = "DELETE"
const DuplicateOrderType_IGNORE DuplicateOrderType = "IGNORE"

type CloseoutRequestType string

const CloseoutRequestType_ACCOUNTSETUP CloseoutRequestType = "ACCOUNTSETUP"
const CloseoutRequestType_TRACKINGNBR CloseoutRequestType = "TRACKINGNBR"
const CloseoutRequestType_CARRIERID CloseoutRequestType = "CARRIERID"
const CloseoutRequestType_MANIFESTGROUPNBR CloseoutRequestType = "MANIFESTGROUPNBR"
const CloseoutRequestType_SERVICEGROUP CloseoutRequestType = "SERVICEGROUP"
const CloseoutRequestType_SERVICEGROUP_ACCOUNTSETUP CloseoutRequestType = "SERVICEGROUP_ACCOUNTSETUP"
const CloseoutRequestType_SERVICECODE CloseoutRequestType = "SERVICECODE"

type ClassifyBy string

const ClassifyBy_ItemCode ClassifyBy = "ItemCode"
const ClassifyBy_HarmonizedCode ClassifyBy = "HarmonizedCode"
const ClassifyBy_HarmonizedCodeAndDesc ClassifyBy = "HarmonizedCodeAndDesc"
const ClassifyBy_Desc ClassifyBy = "Desc"
const ClassifyBy_CodeAndDesc ClassifyBy = "CodeAndDesc"
const ClassifyBy_SubCategoryItemId ClassifyBy = "SubCategoryItemId"

// Complex types
type Rate_Request struct {
	Authentication *Authentication `xml:"Authentication,omitempty"`
	Rate           *Rate           `xml:"Rate,omitempty"`
}

type Authentication struct {
	ActivationKey string `xml:"ActivationKey,omitempty"`
	LoginName     string `xml:"LoginName,omitempty"`
	Password      string `xml:"Password,omitempty"`
	AltID         string `xml:"AltID,omitempty"`
}

type Rate struct {
	RequestDetail *RequestDetail `xml:"RequestDetail,omitempty"`
}

type RequestDetail struct {
	ShipmentList *Shipmentlist `xml:"ShipmentList,omitempty"`
}

type Shipmentlist struct {
	Shipments []Shipment `xml:"Shipment,omitempty"`
}

type Shipment struct {
	BillType                            int                   `xml:"BillType,omitempty"`
	BillAccountNumber                   string                `xml:"BillAccountNumber,omitempty"`
	BillAccountZipCode                  string                `xml:"BillAccountZipCode,omitempty"`
	CreditCardNumber                    string                `xml:"CreditCardNumber,omitempty"`
	CreditCardType                      int                   `xml:"CreditCardType,omitempty"`
	CreditCardExpireDate                *time.Time            `xml:"CreditCardExpireDate,omitempty"`
	CreditCardCode                      string                `xml:"CreditCardCode,omitempty"`
	DriverInstructions                  string                `xml:"DriverInstructions,omitempty"`
	ResidentialFlag                     bool                  `xml:"ResidentialFlag,omitempty"`
	SaturdayDeliveryFlag                bool                  `xml:"SaturdayDeliveryFlag,omitempty"`
	SaturdayPickupFlag                  bool                  `xml:"SaturdayPickupFlag,omitempty"`
	InsidePickupFlag                    bool                  `xml:"InsidePickupFlag,omitempty"`
	InsideDeliveryFlag                  bool                  `xml:"InsideDeliveryFlag,omitempty"`
	CertifiedMailFlag                   bool                  `xml:"CertifiedMailFlag,omitempty"`
	ElectronicReturnReceiptFlag         bool                  `xml:"ElectronicReturnReceiptFlag,omitempty"`
	DeliveryType                        int                   `xml:"DeliveryType,omitempty"`
	DropOffType                         int                   `xml:"DropOffType,omitempty"`
	LabelType                           int                   `xml:"LabelType,omitempty"`
	LabelSize                           int                   `xml:"LabelSize,omitempty"`
	LabelDPI                            int                   `xml:"LabelDPI,omitempty"`
	SortCode1                           string                `xml:"SortCode1,omitempty"`
	SortCode2                           string                `xml:"SortCode2,omitempty"`
	SortCode3                           string                `xml:"SortCode3,omitempty"`
	SortCode4                           string                `xml:"SortCode4,omitempty"`
	SortCode5                           string                `xml:"SortCode5,omitempty"`
	SortCode6                           string                `xml:"SortCode6,omitempty"`
	AuthorizedReturnsNumber             string                `xml:"AuthorizedReturnsNumber,omitempty"`
	ReturnsPackCollectFlag              bool                  `xml:"ReturnsPackCollectFlag,omitempty"`
	NDCPresortDiscountFlag              bool                  `xml:"NDCPresortDiscountFlag,omitempty"`
	ODNCPresortDiscountFlag             bool                  `xml:"ODNCPresortDiscountFlag,omitempty"`
	HolidayDeliveryFlag                 bool                  `xml:"HolidayDeliveryFlag,omitempty"`
	SundayDeliveryFlag                  bool                  `xml:"SundayDeliveryFlag,omitempty"`
	CODReferenceIndicatorType           int                   `xml:"CODReferenceIndicatorType,omitempty"`
	CODRecipientAccountNumber           string                `xml:"CODRecipientAccountNumber,omitempty"`
	CODAddlChargeType                   int                   `xml:"CODAddlChargeType,omitempty"`
	CODAmt                              float64               `xml:"CODAmt,omitempty"`
	CODPayType                          int                   `xml:"CODPayType,omitempty"`
	ServiceCode                         string                `xml:"ServiceCode,omitempty"`
	AccountNumber                       string                `xml:"AccountNumber,omitempty"`
	DimensionUnit                       string                `xml:"DimensionUnit,omitempty"`
	WeightUnit                          string                `xml:"WeightUnit,omitempty"`
	ShipDate                            *time.Time            `xml:"ShipDate,omitempty"`
	TrailerNumber                       string                `xml:"TrailerNumber,omitempty"`
	ManifestGroupNbr                    string                `xml:"ManifestGroupNbr,omitempty"`
	DropZip                             string                `xml:"DropZip,omitempty"`
	ACSFlag                             bool                  `xml:"ACSFlag,omitempty"`
	ReturnLabelRequiredFlag             bool                  `xml:"ReturnLabelRequiredFlag,omitempty"`
	MultiPieceShipmentTrackingNumber    string                `xml:"MultiPieceShipmentTrackingNumber,omitempty"`
	MultiPieceShipmentTotalCount        int                   `xml:"MultiPieceShipmentTotalCount,omitempty"`
	MultiPieceShipmentPieceCounter      int                   `xml:"MultiPieceShipmentPieceCounter,omitempty"`
	InboundLoadId                       string                `xml:"InboundLoadId,omitempty"`
	InboundLoadNumber                   string                `xml:"InboundLoadNumber,omitempty"`
	OutboundLoadId                      string                `xml:"OutboundLoadId,omitempty"`
	OutboundLoadNumber                  string                `xml:"OutboundLoadNumber,omitempty"`
	AppointmentFlag                     bool                  `xml:"AppointmentFlag,omitempty"`
	EveningDeliveryFlag                 bool                  `xml:"EveningDeliveryFlag,omitempty"`
	BookingNumber                       int                   `xml:"BookingNumber,omitempty"`
	ShipperReleaseFlag                  bool                  `xml:"ShipperReleaseFlag,omitempty"`
	ReturnType                          int                   `xml:"ReturnType,omitempty"`
	DeclarationStatementFlag            bool                  `xml:"DeclarationStatementFlag,omitempty"`
	RoutedTransactionIndFlag            bool                  `xml:"RoutedTransactionIndFlag,omitempty"`
	UseAddlCommentsFlag                 bool                  `xml:"UseAddlCommentsFlag,omitempty"`
	ProductCode                         string                `xml:"ProductCode,omitempty"`
	CurrencyCode                        string                `xml:"CurrencyCode,omitempty"`
	ImageRotation                       int                   `xml:"ImageRotation,omitempty"`
	PriorityAlertFlag                   bool                  `xml:"PriorityAlertFlag,omitempty"`
	OverrideShipToAddressValidationFlag bool                  `xml:"OverrideShipToAddressValidationFlag,omitempty"`
	SupplyUsageFlag                     bool                  `xml:"SupplyUsageFlag,omitempty"`
	SignatureReleaseNumber              string                `xml:"SignatureReleaseNumber,omitempty"`
	PackageListEnclosedFlag             bool                  `xml:"PackageListEnclosedFlag,omitempty"`
	RegulatoryControlType               int                   `xml:"RegulatoryControlType,omitempty"`
	IsUPSHundredWeightFlag              bool                  `xml:"IsUPSHundredWeightFlag,omitempty"`
	ShipFrom                            *Address              `xml:"ShipFrom,omitempty"`
	ShipTo                              *Address              `xml:"ShipTo,omitempty"`
	Return                              *Address              `xml:"Return,omitempty"`
	ShipFromHub                         *Address              `xml:"ShipFromHub,omitempty"`
	USPSLocation                        *Address              `xml:"USPSLocation,omitempty"`
	HoldAtLocation                      *Address              `xml:"HoldAtLocation,omitempty"`
	CODRecipient                        *Address              `xml:"CODRecipient,omitempty"`
	DangerousGoods                      *DangerousGoodsDetail `xml:"DangerousGoods,omitempty"`
	OriginallyCapturedPackageId         string                `xml:"OriginallyCapturedPackageId,omitempty"`
	IsSackService                       bool                  `xml:"IsSackService,omitempty"`
	ExportLineItemList                  *ExportLineItems      `xml:"ExportLineItemList,omitempty"`
	PackageList                         *Packages             `xml:"PackageList,omitempty"`
	InternationalShipment               *InternationalDetail  `xml:"InternationalShipment,omitempty"`
	AddlShipmentReference1              string                `xml:"AddlShipmentReference1,omitempty"`
	AddlShipmentReference2              string                `xml:"AddlShipmentReference2,omitempty"`
	AddlShipmentReference3              string                `xml:"AddlShipmentReference3,omitempty"`
	AddlShipmentReference4              string                `xml:"AddlShipmentReference4,omitempty"`
	AddlShipmentReference5              string                `xml:"AddlShipmentReference5,omitempty"`
	AddlShipmentReference6              string                `xml:"AddlShipmentReference6,omitempty"`
	AddlShipmentReference7              string                `xml:"AddlShipmentReference7,omitempty"`
	AddlShipmentReference8              string                `xml:"AddlShipmentReference8,omitempty"`
	DeliveryConfirmationFlag            bool                  `xml:"DeliveryConfirmationFlag,omitempty"`
	SignatureConfirmationFlag           bool                  `xml:"SignatureConfirmationFlag,omitempty"`
	SignatureConfirmationType           int                   `xml:"SignatureConfirmationType,omitempty"`
	BalloonFlag                         bool                  `xml:"BalloonFlag,omitempty"`
	OversizeFlag                        bool                  `xml:"OversizeFlag,omitempty"`
	NonMachinableFlag                   bool                  `xml:"NonMachinableFlag,omitempty"`
	USPSRateIndicator                   string                `xml:"USPSRateIndicator,omitempty"`
	DryIceWeight                        float64               `xml:"DryIceWeight,omitempty"`
	SignatureReleaseFlag                bool                  `xml:"SignatureReleaseFlag,omitempty"`
	InsuranceAmount                     float64               `xml:"InsuranceAmount,omitempty"`
	InsurancePayType                    int                   `xml:"InsurancePayType,omitempty"`
	DeclaredValue                       float64               `xml:"DeclaredValue,omitempty"`
	InsuranceType                       int                   `xml:"InsuranceType,omitempty"`
	EndorsementType                     int                   `xml:"EndorsementType,omitempty"`
	FragileFlag                         bool                  `xml:"FragileFlag,omitempty"`
	USPSSortLevel                       int                   `xml:"USPSSortLevel,omitempty"`
	USPSODEnclosedMailClass             int                   `xml:"USPSODEnclosedMailClass,omitempty"`
	ProcessCategory                     int                   `xml:"ProcessCategory,omitempty"`
	CarrierId                           string                `xml:"CarrierId,omitempty"`
	AdditionalAccountNumber             string                `xml:"AdditionalAccountNumber,omitempty"`
	CarriageValue                       float64               `xml:"CarriageValue,omitempty"`
	ShipmentAccountNbr                  string                `xml:"ShipmentAccountNbr,omitempty"`
}

type Address struct {
	Company         string `xml:"Company,omitempty"`
	Attn            string `xml:"Attn,omitempty"`
	AddressLine1    string `xml:"AddressLine1,omitempty"`
	AddressLine2    string `xml:"AddressLine2,omitempty"`
	AddressLine3    string `xml:"AddressLine3,omitempty"`
	City            string `xml:"City,omitempty"`
	StateCode       string `xml:"StateCode,omitempty"`
	Zip             string `xml:"Zip,omitempty"`
	Zip4            string `xml:"Zip4,omitempty"`
	CountryCode     string `xml:"CountryCode,omitempty"`
	CountryName     string `xml:"CountryName,omitempty"`
	PhoneAreaCode   string `xml:"PhoneAreaCode,omitempty"`
	Phone           string `xml:"Phone,omitempty"`
	PhoneExt        string `xml:"PhoneExt,omitempty"`
	Fax             string `xml:"Fax,omitempty"`
	EMail           string `xml:"EMail,omitempty"`
	Department      string `xml:"Department,omitempty"`
	Reference       string `xml:"Reference,omitempty"`
	ResidentialFlag bool   `xml:"ResidentialFlag,omitempty"`
	IsPOBox         bool   `xml:"IsPOBox,omitempty"`
}

type DangerousGoodsDetail struct {
	DangerousGoodsTechnicalName       string  `xml:"DangerousGoodsTechnicalName,omitempty"`
	DangerousGoodsUnNumber            string  `xml:"DangerousGoodsUnNumber,omitempty"`
	DangerousGoodsNumberOfUnits       int     `xml:"DangerousGoodsNumberOfUnits,omitempty"`
	DangerousGoodsPackingType         string  `xml:"DangerousGoodsPackingType,omitempty"`
	DangerousGoodsPackingInstructions string  `xml:"DangerousGoodsPackingInstructions,omitempty"`
	DangerousGoodsTitleOfSignatory    string  `xml:"DangerousGoodsTitleOfSignatory,omitempty"`
	DangerousGoodsProperShippingName  string  `xml:"DangerousGoodsProperShippingName,omitempty"`
	DangerousGoodsNameOfSignatory     string  `xml:"DangerousGoodsNameOfSignatory,omitempty"`
	DangerousGoodsPercentageNumber    float64 `xml:"DangerousGoodsPercentageNumber,omitempty"`
	DangerousGoodsPlaceOfSignatory    string  `xml:"DangerousGoodsPlaceOfSignatory,omitempty"`
	DangerousGoodsCommodityCount      int     `xml:"DangerousGoodsCommodityCount,omitempty"`
	HazmatCode                        int     `xml:"HazmatCode,omitempty"`
	HazmatDOTShippingName             string  `xml:"HazmatDOTShippingName,omitempty"`
	HazmatClassOrDivision             string  `xml:"HazmatClassOrDivision,omitempty"`
	HazmatLabelType                   string  `xml:"HazmatLabelType,omitempty"`
	HazmatUnitQuantity                float64 `xml:"HazmatUnitQuantity,omitempty"`
	HazmatUnitType                    string  `xml:"HazmatUnitType,omitempty"`
	HazmatContactName                 string  `xml:"HazmatContactName,omitempty"`
	HazmatContactPhone                string  `xml:"HazmatContactPhone,omitempty"`
	HazmatCargoAircraftOnly           bool    `xml:"HazmatCargoAircraftOnly,omitempty"`
	AlcoholPackagingType              int     `xml:"AlcoholPackagingType,omitempty"`
	AlcoholType                       int     `xml:"AlcoholType,omitempty"`
	AlcoholVolumeInLiters             float64 `xml:"AlcoholVolumeInLiters,omitempty"`
	HazmatWeight                      float64 `xml:"HazmatWeight,omitempty"`
	HazmatPackagingInstructions       string  `xml:"HazmatPackagingInstructions,omitempty"`
	HazmatPackingGroup                string  `xml:"HazmatPackingGroup,omitempty"`
	HazmatRegulationSet               string  `xml:"HazmatRegulationSet,omitempty"`
}

type ExportLineItems struct {
	ExportLineItems []ExportLineItem `xml:"ExportLineItem,omitempty"`
}

type ExportLineItem struct {
	Code                     string      `xml:"Code,omitempty"`
	Description              string      `xml:"Description,omitempty"`
	Weight                   float64     `xml:"Weight,omitempty"`
	UnitWeight               string      `xml:"UnitWeight,omitempty"`
	WeightUnit               string      `xml:"WeightUnit,omitempty"`
	Price                    float64     `xml:"Price,omitempty"`
	Quantity                 int         `xml:"Quantity,omitempty"`
	UnitQuantity             *AmountUnit `xml:"UnitQuantity,omitempty"`
	UnitPrice                float64     `xml:"UnitPrice,omitempty"`
	CountryOfManufacture     string      `xml:"CountryOfManufacture,omitempty"`
	SedLineAmt               float64     `xml:"SedLineAmt,omitempty"`
	CoType                   string      `xml:"CoType,omitempty"`
	SedInd                   string      `xml:"SedInd,omitempty"`
	SpecialCommodities       int         `xml:"SpecialCommodities,omitempty"`
	NaftaPreferenceCriteria  int         `xml:"NaftaPreferenceCriteria,omitempty"`
	NaftaCoNetCost           float64     `xml:"NaftaCoNetCost,omitempty"`
	CoMarksAndNumbers        string      `xml:"CoMarksAndNumbers,omitempty"`
	NaftaCoNetcost_BeginDate *time.Time  `xml:"NaftaCoNetcost_BeginDate,omitempty"`
	NaftaCoNetcost_EndDate   *time.Time  `xml:"NaftaCoNetcost_EndDate,omitempty"`
	Harmonizedcode           string      `xml:"Harmonizedcode,omitempty"`
	SalesTaxAmount           float64     `xml:"SalesTaxAmount,omitempty"`
	DutyAmount               float64     `xml:"DutyAmount,omitempty"`
	Currency                 string      `xml:"Currency,omitempty"`
	OutputCurrency           string      `xml:"OutputCurrency,omitempty"`
	Amount                   float64     `xml:"Amount,omitempty"`
}

type Packages struct {
	Packages []Package `xml:"Package,omitempty"`
}

type Package struct {
	ShipToAttn              string  `xml:"ShipToAttn,omitempty"`
	DimWeight               float64 `xml:"DimWeight,omitempty"`
	TrackingNumber          string  `xml:"TrackingNumber,omitempty"`
	Length                  float64 `xml:"Length,omitempty"`
	Height                  float64 `xml:"Height,omitempty"`
	Width                   float64 `xml:"Width,omitempty"`
	PackagingType           int     `xml:"PackagingType,omitempty"`
	ReferenceNumber         string  `xml:"ReferenceNumber,omitempty"`
	Weight                  float64 `xml:"Weight,omitempty"`
	AddlPackageReference1   string  `xml:"AddlPackageReference1,omitempty"`
	AddlPackageReference2   string  `xml:"AddlPackageReference2,omitempty"`
	AddlPackageReference3   string  `xml:"AddlPackageReference3,omitempty"`
	AddlPackageReference4   string  `xml:"AddlPackageReference4,omitempty"`
	AddlPackageReference5   string  `xml:"AddlPackageReference5,omitempty"`
	AddlPackageReference6   string  `xml:"AddlPackageReference6,omitempty"`
	AddlPackageReference7   string  `xml:"AddlPackageReference7,omitempty"`
	AddlPackageReference8   string  `xml:"AddlPackageReference8,omitempty"`
	USPSRateIndicator       string  `xml:"USPSRateIndicator,omitempty"`
	USPSSortLevel           int     `xml:"USPSSortLevel,omitempty"`
	USPSODEnclosedMailClass int     `xml:"USPSODEnclosedMailClass,omitempty"`
	ContentDescription      string  `xml:"ContentDescription,omitempty"`
	AdditionalHandlingFlag  bool    `xml:"AdditionalHandlingFlag,omitempty"`
	NotificationType        int     `xml:"NotificationType,omitempty"`
	NotificationValue       string  `xml:"NotificationValue,omitempty"`
}

type InternationalDetail struct {
	CN22DescriptionType                  int        `xml:"CN22DescriptionType,omitempty"`
	ShipperECCN                          string     `xml:"ShipperECCN,omitempty"`
	ShipperXTN                           string     `xml:"ShipperXTN,omitempty"`
	ShipperFTSR                          string     `xml:"ShipperFTSR,omitempty"`
	Incoterms                            string     `xml:"Incoterms,omitempty"`
	ExportLicense                        string     `xml:"ExportLicense,omitempty"`
	ImportLicense                        string     `xml:"ImportLicense,omitempty"`
	ConsigneeCode                        string     `xml:"ConsigneeCode,omitempty"`
	ConsigneeEIN                         string     `xml:"ConsigneeEIN,omitempty"`
	ConsigneeVAT                         string     `xml:"ConsigneeVAT,omitempty"`
	SignatureName                        string     `xml:"SignatureName,omitempty"`
	SignatureTitle                       string     `xml:"SignatureTitle,omitempty"`
	ExportReasonCode                     int        `xml:"ExportReasonCode,omitempty"`
	ShipperEin                           string     `xml:"ShipperEin,omitempty"`
	SEDFilingOption                      int        `xml:"SEDFilingOption,omitempty"`
	SignaturePhone                       string     `xml:"SignaturePhone,omitempty"`
	SignatureEmailAddress                string     `xml:"SignatureEmailAddress,omitempty"`
	IsDepartmentOfStateShipment          bool       `xml:"IsDepartmentOfStateShipment,omitempty"`
	IsDepartmentOfStateExempt            bool       `xml:"IsDepartmentOfStateExempt,omitempty"`
	DOS22CFRExemptNumber                 string     `xml:"DOS22CFRExemptNumber,omitempty"`
	CommercialInvoiceType                string     `xml:"CommercialInvoiceType,omitempty"`
	CommercialInvoiceNumber              string     `xml:"CommercialInvoiceNumber,omitempty"`
	CommercialInvoiceOtherCharges        float64    `xml:"CommercialInvoiceOtherCharges,omitempty"`
	Use3rdPartyBrokerFlag                bool       `xml:"Use3rdPartyBrokerFlag,omitempty"`
	BrokerShipAlertFlag                  bool       `xml:"BrokerShipAlertFlag,omitempty"`
	BrokerDeliveryNotificationFlag       bool       `xml:"BrokerDeliveryNotificationFlag,omitempty"`
	BrokerExceptionNotificationFlag      bool       `xml:"BrokerExceptionNotificationFlag,omitempty"`
	BrokerEmailFormat                    string     `xml:"BrokerEmailFormat,omitempty"`
	BrokerAccountNumber                  string     `xml:"BrokerAccountNumber,omitempty"`
	AdditionalComments                   string     `xml:"AdditionalComments,omitempty"`
	DutiesTaxAccountNumber               string     `xml:"DutiesTaxAccountNumber,omitempty"`
	AdmissibilityPackagingType           int        `xml:"AdmissibilityPackagingType,omitempty"`
	SenderTINType                        string     `xml:"SenderTINType,omitempty"`
	SenderTINNumber                      string     `xml:"SenderTINNumber,omitempty"`
	RecipientTINNumber                   string     `xml:"RecipientTINNumber,omitempty"`
	AESTransactionNumber                 string     `xml:"AESTransactionNumber,omitempty"`
	PartiesToTrans                       string     `xml:"PartiesToTrans,omitempty"`
	ExportInformationCode                int        `xml:"ExportInformationCode,omitempty"`
	ExportLicenceExpDate                 string     `xml:"ExportLicenceExpDate,omitempty"`
	BrokerageId                          string     `xml:"BrokerageId,omitempty"`
	WorldEaseFlag                        bool       `xml:"WorldEaseFlag,omitempty"`
	WorldEaseCounter                     int        `xml:"WorldEaseCounter,omitempty"`
	GCCNNumber                           string     `xml:"GCCNNumber,omitempty"`
	WorldEaseTrackingNumber              string     `xml:"WorldEaseTrackingNumber,omitempty"`
	WorldEasePortNumber                  string     `xml:"WorldEasePortNumber,omitempty"`
	WorldEasePortZip                     string     `xml:"WorldEasePortZip,omitempty"`
	WorldEasePortName                    string     `xml:"WorldEasePortName,omitempty"`
	WorldEasePortCountry                 string     `xml:"WorldEasePortCountry,omitempty"`
	WorldEasePackageCount                int        `xml:"WorldEasePackageCount,omitempty"`
	FTRExemtionCode                      string     `xml:"FTRExemtionCode,omitempty"`
	ShippingFormContent1                 string     `xml:"ShippingFormContent1,omitempty"`
	ShippingFormContent2                 string     `xml:"ShippingFormContent2,omitempty"`
	ShippingFormContent3                 string     `xml:"ShippingFormContent3,omitempty"`
	ShippingFormContent4                 string     `xml:"ShippingFormContent4,omitempty"`
	ShippingFormContent5                 string     `xml:"ShippingFormContent5,omitempty"`
	ShippingFormContent6                 string     `xml:"ShippingFormContent6,omitempty"`
	FormRequested                        string     `xml:"FormRequested,omitempty"`
	FormGenerationType                   int        `xml:"FormGenerationType,omitempty"`
	NaftaProducer                        *Address   `xml:"NaftaProducer,omitempty"`
	NaftaBlanketPeriod_BeginDate         *time.Time `xml:"NaftaBlanketPeriod_BeginDate,omitempty"`
	NaftaBlanketPeriod_EndDate           *time.Time `xml:"NaftaBlanketPeriod_EndDate,omitempty"`
	NaftaProducerTaxId                   string     `xml:"NaftaProducerTaxId,omitempty"`
	NaftaProducerDeterminationType       int        `xml:"NaftaProducerDeterminationType,omitempty"`
	ExportDate                           *time.Time `xml:"ExportDate,omitempty"`
	ExportingCarrier                     string     `xml:"ExportingCarrier,omitempty"`
	InbondCode                           int        `xml:"InbondCode,omitempty"`
	EntryNumber                          string     `xml:"EntryNumber,omitempty"`
	PointOfOrigin                        string     `xml:"PointOfOrigin,omitempty"`
	ModeOfTransport                      string     `xml:"ModeOfTransport,omitempty"`
	PortOfExport                         string     `xml:"PortOfExport,omitempty"`
	PortOfUnloading                      string     `xml:"PortOfUnloading,omitempty"`
	LoadingPier                          string     `xml:"LoadingPier,omitempty"`
	RoutedExportTransactionIndicator     bool       `xml:"RoutedExportTransactionIndicator,omitempty"`
	ContainerizedIndicator               bool       `xml:"ContainerizedIndicator,omitempty"`
	ExportType                           int        `xml:"ExportType,omitempty"`
	ForwarderTaxId                       string     `xml:"ForwarderTaxId,omitempty"`
	ExceptionCode                        int        `xml:"ExceptionCode,omitempty"`
	SoldToOption                         int        `xml:"SoldToOption,omitempty"`
	TradeDirectProductType               int        `xml:"TradeDirectProductType,omitempty"`
	ImporterAccountNumber                string     `xml:"ImporterAccountNumber,omitempty"`
	CustomFormType                       int        `xml:"CustomFormType,omitempty"`
	ImportControlFlag                    bool       `xml:"ImportControlFlag,omitempty"`
	CommercialInvoiceRemovalFlag         bool       `xml:"CommercialInvoiceRemovalFlag,omitempty"`
	CommercialInvoiceFreightChargeAmount float64    `xml:"CommercialInvoiceFreightChargeAmount,omitempty"`
	IssuingAgencyCode                    string     `xml:"IssuingAgencyCode,omitempty"`
	LicencePlateNumber                   string     `xml:"LicencePlateNumber,omitempty"`
	Importer                             *Address   `xml:"Importer,omitempty"`
	Broker                               *Address   `xml:"Broker,omitempty"`
	UltimateConsignee                    *Address   `xml:"UltimateConsignee,omitempty"`
	UltimateConsigneeAccountNumber       string     `xml:"UltimateConsigneeAccountNumber,omitempty"`
	UltimateConsigneeTaxId               string     `xml:"UltimateConsigneeTaxId,omitempty"`
	Forwarder                            *Address   `xml:"Forwarder,omitempty"`
	IntermediateConsignee                *Address   `xml:"IntermediateConsignee,omitempty"`
	SccAccountNumber                     string     `xml:"SccAccountNumber,omitempty"`
	DutyAccountNumber                    string     `xml:"DutyAccountNumber,omitempty"`
	DutyBillType                         int        `xml:"DutyBillType,omitempty"`
	DutyServiceType                      int        `xml:"DutyServiceType,omitempty"`
	ImporterVATNumber                    string     `xml:"ImporterVATNumber,omitempty"`
	DutyBillPayor                        *Address   `xml:"DutyBillPayor,omitempty"`
	RecipientCustomIdType                int        `xml:"RecipientCustomIdType,omitempty"`
	IncludeDutyCalculation               bool       `xml:"IncludeDutyCalculation,omitempty"`
	BillPayor                            *Address   `xml:"BillPayor,omitempty"`
	CommercialImporterFlag               bool       `xml:"CommercialImporterFlag,omitempty"`
	TotalMonthlyImportedWeight           float64    `xml:"TotalMonthlyImportedWeight,omitempty"`
	TotalMonthlyImportedValue            float64    `xml:"TotalMonthlyImportedValue,omitempty"`
	ExportLicenseName                    string     `xml:"ExportLicenseName,omitempty"`
}

type Rate_Response struct {
	Rate_Results []Rate_Result `xml:"Rate_Result,omitempty"`
}

type Rate_Result struct {
	Error          *error          `xml:"Error,omitempty"`
	RateResultList *RateResultList `xml:"RateResultList,omitempty"`
	Version        string          `xml:"Version,omitempty"`
}

type RateResultList struct {
	RateResults []RateResult `xml:"RateResult,omitempty"`
}

type RateResult struct {
	ReferenceNbr  string            `xml:"ReferenceNbr,omitempty"`
	RateErrorList *ArrayOfRateError `xml:"RateErrorList,omitempty"`
	Services      []Service         `xml:"Service,omitempty"`
}

type ArrayOfRateError struct {
	RateErrors []RateError `xml:"RateError,omitempty"`
}

type RateError struct {
	Code      int    `xml:"Code,omitempty"`
	Message   string `xml:"Message,omitempty"`
	CarrierId string `xml:"CarrierId,omitempty"`
}

type Service struct {
	Carrier           string           `xml:"Carrier,omitempty"`
	ServiceCode       string           `xml:"ServiceCode,omitempty"`
	ServiceDesc       string           `xml:"ServiceDesc,omitempty"`
	DeliveryDate      *time.Time       `xml:"DeliveryDate,omitempty"`
	DeliveryTime      string           `xml:"DeliveryTime,omitempty"`
	GuarnateedService bool             `xml:"GuarnateedService,omitempty"`
	ShipmentCharges   *ShipmentCharges `xml:"ShipmentCharges,omitempty"`
}

type ShipmentCharges struct {
	Rate     float64     `xml:"Rate,omitempty"`
	BaseRate float64     `xml:"BaseRate,omitempty"`
	Fees     *FeeDetails `xml:"Fees,omitempty"`
}

type FeeDetails struct {
	Fee *ArrayOfFeeDetail `xml:"Fee,omitempty"`
}

type ArrayOfFeeDetail struct {
	FeeDetails []FeeDetail `xml:"FeeDetail,omitempty"`
}

type FeeDetail struct {
	Name   string `xml:"Name,attr,ommitempty"`
	Charge string `xml:"Charge,attr,ommitempty"`
}

type AddShipper_Request struct {
	Request *Request `xml:"Request,omitempty"`
}

type Request struct {
	Authentication *Authentication  `xml:"Authentication,omitempty"`
	RequestType    *RequestTypeEnum `xml:"RequestType,omitempty"`
	EULAAccept     bool             `xml:"EULAAccept,omitempty"`
	PlanDesc       string           `xml:"PlanDesc,omitempty"`
	Address        *Address         `xml:"Address,omitempty"`
	User           *User            `xml:"User,omitempty"`
}

type User struct {
	FirstName    string `xml:"FirstName,omitempty"`
	LastName     string `xml:"LastName,omitempty"`
	UserPhone    string `xml:"UserPhone,omitempty"`
	UserFax      string `xml:"UserFax,omitempty"`
	UserEmail    string `xml:"UserEmail,omitempty"`
	UserName     string `xml:"UserName,omitempty"`
	UserPassword string `xml:"UserPassword,omitempty"`
}

type AddShipper_Response struct {
	ShipperResult *ShipperResult `xml:"ShipperResult,omitempty"`
}

type ShipperResult struct {
	Error       *error       `xml:"Error,omitempty"`
	ShipperInfo *ShipperInfo `xml:"ShipperInfo,omitempty"`
	UserInfo    *UserInfo    `xml:"UserInfo,omitempty"`
	Version     string       `xml:"Version,omitempty"`
}

type ShipperInfo struct {
	ShipperId int `xml:"ShipperId,omitempty"`
}

type UserInfo struct {
	UserId       int    `xml:"UserId,omitempty"`
	UserName     string `xml:"UserName,omitempty"`
	UserPassword string `xml:"UserPassword,omitempty"`
}

type AddVal_Request struct {
	Authentication  *Authentication `xml:"Authentication,omitempty"`
	Provider        string          `xml:"Provider,omitempty"`
	DeniedPartyFlag bool            `xml:"DeniedPartyFlag,omitempty"`
	Address         *Address        `xml:"Address,omitempty"`
}

type AddVal_Response struct {
	AddVal_Result *AddVal_Result `xml:"AddVal_Result,omitempty"`
}

type AddVal_Result struct {
	Error              *error       `xml:"Error,omitempty"`
	ConfirmationNumber string       `xml:"ConfirmationNumber,omitempty"`
	Version            string       `xml:"Version,omitempty"`
	AddressList        *AddressList `xml:"AddressList,omitempty"`
}

type AddressList struct {
	Addresss           []Address                  `xml:"Address,omitempty"`
	AddressSuggestList *ArrayOfAddressSuggestList `xml:"AddressSuggestList,omitempty"`
}

type ArrayOfAddressSuggestList struct {
	AddressSuggestions []AddressSuggestList `xml:"AddressSuggestion,omitempty"`
}

type AddressSuggestList struct {
	Company           string             `xml:"Company,omitempty"`
	Attn              string             `xml:"Attn,omitempty"`
	AddressLine1      string             `xml:"AddressLine1,omitempty"`
	AddressLine2      string             `xml:"AddressLine2,omitempty"`
	AddressLine3      string             `xml:"AddressLine3,omitempty"`
	City              string             `xml:"City,omitempty"`
	StateCode         string             `xml:"StateCode,omitempty"`
	Zip               string             `xml:"Zip,omitempty"`
	Zip4              string             `xml:"Zip4,omitempty"`
	CountryCode       string             `xml:"CountryCode,omitempty"`
	CountryName       string             `xml:"CountryName,omitempty"`
	PhoneAreaCode     string             `xml:"PhoneAreaCode,omitempty"`
	Phone             string             `xml:"Phone,omitempty"`
	PhoneExt          string             `xml:"PhoneExt,omitempty"`
	Fax               string             `xml:"Fax,omitempty"`
	EMail             string             `xml:"EMail,omitempty"`
	Department        string             `xml:"Department,omitempty"`
	Reference         string             `xml:"Reference,omitempty"`
	ResidentialFlag   bool               `xml:"ResidentialFlag,omitempty"`
	StreetNumberRange *StreetNumberRange `xml:"StreetNumberRange,omitempty"`
	ZipRange          *ZipRange          `xml:"ZipRange,omitempty"`
}

type StreetNumberRange struct {
	LowNumber  string `xml:"LowNumber,omitempty"`
	HighNumber string `xml:"HighNumber,omitempty"`
}

type ZipRange struct {
	LowNumber  string `xml:"LowNumber,omitempty"`
	HighNumber string `xml:"HighNumber,omitempty"`
}

type AbolImportOrderRequest struct {
	Request *Order_Request `xml:"Request,omitempty"`
}

type Order_Request struct {
	Authentication *Authentication   `xml:"Authentication,omitempty"`
	RequestType    *OrderRequestType `xml:"RequestType,omitempty"`
	OrderList      *OrderList        `xml:"OrderList,omitempty"`
	MapName        string            `xml:"MapName,omitempty"`
}

type OrderList struct {
	Orders []OrderV3 `xml:"Order,omitempty"`
}

type OrderV3 struct {
	OrderID                             int                        `xml:"OrderID,omitempty"`
	OrderNumber                         string                     `xml:"OrderNumber,omitempty"`
	OrderNumberExternal                 string                     `xml:"OrderNumberExternal,omitempty"`
	OrderStatus                         int                        `xml:"OrderStatus,omitempty"`
	OrderDate                           *time.Time                 `xml:"OrderDate,omitempty"`
	OrderAmount                         float64                    `xml:"OrderAmount,omitempty"`
	TaxAmount                           float64                    `xml:"TaxAmount,omitempty"`
	DutyAmount                          float64                    `xml:"DutyAmount,omitempty"`
	ShippingAmount                      float64                    `xml:"ShippingAmount,omitempty"`
	Weight                              float64                    `xml:"Weight,omitempty"`
	ShippingCarrier                     string                     `xml:"ShippingCarrier,omitempty"`
	ShippingMethod                      string                     `xml:"ShippingMethod,omitempty"`
	Servicecode                         string                     `xml:"Servicecode,omitempty"`
	BatchNumber                         int                        `xml:"BatchNumber,omitempty"`
	AmountPaid                          float64                    `xml:"AmountPaid,omitempty"`
	PaymentMethod                       int                        `xml:"PaymentMethod,omitempty"`
	CheckoutStatus                      int                        `xml:"CheckoutStatus,omitempty"`
	PaymentStatus                       int                        `xml:"PaymentStatus,omitempty"`
	BuyerEmail                          string                     `xml:"BuyerEmail,omitempty"`
	BuyerUserID                         string                     `xml:"BuyerUserID,omitempty"`
	PurchaseOrder                       string                     `xml:"PurchaseOrder,omitempty"`
	OrderTerms                          string                     `xml:"OrderTerms,omitempty"`
	SalesAgent                          string                     `xml:"SalesAgent,omitempty"`
	LineItemList                        *LineItemList              `xml:"LineItemList,omitempty"`
	BillType                            int                        `xml:"BillType,omitempty"`
	BillAccountNumber                   string                     `xml:"BillAccountNumber,omitempty"`
	BillAccountZipCode                  string                     `xml:"BillAccountZipCode,omitempty"`
	CreditCardNumber                    string                     `xml:"CreditCardNumber,omitempty"`
	CreditCardType                      int                        `xml:"CreditCardType,omitempty"`
	CreditCardExpireDate                *time.Time                 `xml:"CreditCardExpireDate,omitempty"`
	CreditCardCode                      string                     `xml:"CreditCardCode,omitempty"`
	DriverInstructions                  string                     `xml:"DriverInstructions,omitempty"`
	ResidentialFlag                     bool                       `xml:"ResidentialFlag,omitempty"`
	SaturdayDeliveryFlag                bool                       `xml:"SaturdayDeliveryFlag,omitempty"`
	SaturdayPickupFlag                  bool                       `xml:"SaturdayPickupFlag,omitempty"`
	InsidePickupFlag                    bool                       `xml:"InsidePickupFlag,omitempty"`
	InsideDeliveryFlag                  bool                       `xml:"InsideDeliveryFlag,omitempty"`
	CertifiedMailFlag                   bool                       `xml:"CertifiedMailFlag,omitempty"`
	ElectronicReturnReceiptFlag         bool                       `xml:"ElectronicReturnReceiptFlag,omitempty"`
	DeliveryType                        int                        `xml:"DeliveryType,omitempty"`
	DropOffType                         int                        `xml:"DropOffType,omitempty"`
	SortCode1                           string                     `xml:"SortCode1,omitempty"`
	SortCode2                           string                     `xml:"SortCode2,omitempty"`
	SortCode3                           string                     `xml:"SortCode3,omitempty"`
	SortCode4                           string                     `xml:"SortCode4,omitempty"`
	SortCode5                           string                     `xml:"SortCode5,omitempty"`
	SortCode6                           string                     `xml:"SortCode6,omitempty"`
	AuthorizedReturnsNumber             string                     `xml:"AuthorizedReturnsNumber,omitempty"`
	ReturnsPackCollectFlag              bool                       `xml:"ReturnsPackCollectFlag,omitempty"`
	NDCPresortDiscountFlag              bool                       `xml:"NDCPresortDiscountFlag,omitempty"`
	ODNCPresortDiscountFlag             bool                       `xml:"ODNCPresortDiscountFlag,omitempty"`
	HolidayDeliveryFlag                 bool                       `xml:"HolidayDeliveryFlag,omitempty"`
	SundayDeliveryFlag                  bool                       `xml:"SundayDeliveryFlag,omitempty"`
	CODReferenceIndicatorType           int                        `xml:"CODReferenceIndicatorType,omitempty"`
	CODRecipientAccountNumber           string                     `xml:"CODRecipientAccountNumber,omitempty"`
	CODAddlChargeType                   int                        `xml:"CODAddlChargeType,omitempty"`
	CODAmt                              float64                    `xml:"CODAmt,omitempty"`
	CODPayType                          int                        `xml:"CODPayType,omitempty"`
	AccountNumber                       string                     `xml:"AccountNumber,omitempty"`
	DimensionUnit                       string                     `xml:"DimensionUnit,omitempty"`
	WeightUnit                          string                     `xml:"WeightUnit,omitempty"`
	ShipDate                            *time.Time                 `xml:"ShipDate,omitempty"`
	DropZip                             string                     `xml:"DropZip,omitempty"`
	ACSFlag                             bool                       `xml:"ACSFlag,omitempty"`
	ReturnLabelRequiredFlag             bool                       `xml:"ReturnLabelRequiredFlag,omitempty"`
	AppointmentFlag                     bool                       `xml:"AppointmentFlag,omitempty"`
	EveningDeliveryFlag                 bool                       `xml:"EveningDeliveryFlag,omitempty"`
	BookingNumber                       string                     `xml:"BookingNumber,omitempty"`
	ShipperReleaseFlag                  bool                       `xml:"ShipperReleaseFlag,omitempty"`
	AdditionalHandlingFlag              bool                       `xml:"AdditionalHandlingFlag,omitempty"`
	ReturnType                          int                        `xml:"ReturnType,omitempty"`
	DeclarationStatementFlag            bool                       `xml:"DeclarationStatementFlag,omitempty"`
	ProductCode                         string                     `xml:"ProductCode,omitempty"`
	CurrencyCode                        string                     `xml:"CurrencyCode,omitempty"`
	PriorityAlertFlag                   bool                       `xml:"PriorityAlertFlag,omitempty"`
	OverrideShipToAddressValidationFlag bool                       `xml:"OverrideShipToAddressValidationFlag,omitempty"`
	SupplyUsageFlag                     bool                       `xml:"SupplyUsageFlag,omitempty"`
	SignatureReleaseNumber              string                     `xml:"SignatureReleaseNumber,omitempty"`
	PackageListEnclosedFlag             bool                       `xml:"PackageListEnclosedFlag,omitempty"`
	RegulatoryControlType               int                        `xml:"RegulatoryControlType,omitempty"`
	IsUPSHundredWeightFlag              bool                       `xml:"IsUPSHundredWeightFlag,omitempty"`
	ShipFrom                            *Address                   `xml:"ShipFrom,omitempty"`
	ShipTo                              *Address                   `xml:"ShipTo,omitempty"`
	Return                              *Address                   `xml:"Return,omitempty"`
	ShipFromHub                         *Address                   `xml:"ShipFromHub,omitempty"`
	UspsLocation                        *Address                   `xml:"UspsLocation,omitempty"`
	HoldAtLocation                      *Address                   `xml:"HoldAtLocation,omitempty"`
	CODRecipient                        *Address                   `xml:"CODRecipient,omitempty"`
	BillPayor                           *Address                   `xml:"BillPayor,omitempty"`
	DangerousGoods                      *OrderDangerousGoodsDetail `xml:"DangerousGoods,omitempty"`
	OriginallyCapturedPackageId         string                     `xml:"OriginallyCapturedPackageId,omitempty"`
	PackageList                         *PackagesList              `xml:"PackageList,omitempty"`
	InternationalDetail                 *OrderInternationalDetail  `xml:"InternationalDetail,omitempty"`
	ProcessStatus                       int                        `xml:"ProcessStatus,omitempty"`
	ProcessMessage                      string                     `xml:"ProcessMessage,omitempty"`
	BatchImportNumber                   string                     `xml:"BatchImportNumber,omitempty"`
	BatchExportNumber                   string                     `xml:"BatchExportNumber,omitempty"`
	InsuranceAmount                     float64                    `xml:"InsuranceAmount,omitempty"`
	TransactionID                       string                     `xml:"TransactionID,omitempty"`
	LookupValue                         string                     `xml:"LookupValue,omitempty"`
	DuplicateType                       *DuplicateOrderType        `xml:"DuplicateType,omitempty"`
	FreightShipment                     *OrderFreightShipment      `xml:"FreightShipment,omitempty"`
	HeatedServiceFlag                   bool                       `xml:"HeatedServiceFlag,omitempty"`
	HoldForPickupFlag                   bool                       `xml:"HoldForPickupFlag,omitempty"`
	HydraulicLiftFlag                   bool                       `xml:"HydraulicLiftFlag,omitempty"`
	TradeShowPickupDeliveryFlag         bool                       `xml:"TradeShowPickupDeliveryFlag,omitempty"`
	LumperUnloadingFlag                 bool                       `xml:"LumperUnloadingFlag,omitempty"`
	FrozenServiceFlag                   bool                       `xml:"FrozenServiceFlag,omitempty"`
	ChequeReturnVASAmount               float64                    `xml:"ChequeReturnVASAmount,omitempty"`
	KeepFreshFlag                       bool                       `xml:"KeepFreshFlag,omitempty"`
	Oversize2Flag                       bool                       `xml:"Oversize2Flag,omitempty"`
	Misc1                               string                     `xml:"Misc1,omitempty"`
	Misc2                               string                     `xml:"Misc2,omitempty"`
	Misc3                               string                     `xml:"Misc3,omitempty"`
	Misc4                               float64                    `xml:"Misc4,omitempty"`
	Misc5                               float64                    `xml:"Misc5,omitempty"`
	Misc6                               float64                    `xml:"Misc6,omitempty"`
	Misc7                               int                        `xml:"Misc7,omitempty"`
	Misc8                               int                        `xml:"Misc8,omitempty"`
	Misc9                               bool                       `xml:"Misc9,omitempty"`
	Misc10                              *time.Time                 `xml:"Misc10,omitempty"`
}

type LineItemList struct {
	LineItems []OrderItemV3 `xml:"LineItem,omitempty"`
}

type OrderItemV3 struct {
	DisplayName              string     `xml:"DisplayName,omitempty"`
	Description              string     `xml:"Description,omitempty"`
	VendorName               string     `xml:"VendorName,omitempty"`
	QuantityOrdered          int        `xml:"QuantityOrdered,omitempty"`
	QuantityShipped          int        `xml:"QuantityShipped,omitempty"`
	TaxAmount                float64    `xml:"TaxAmount,omitempty"`
	DutyAmount               float64    `xml:"DutyAmount,omitempty"`
	SKU                      string     `xml:"SKU,omitempty"`
	HarmonizedCode           string     `xml:"HarmonizedCode,omitempty"`
	Length                   float64    `xml:"Length,omitempty"`
	Width                    float64    `xml:"Width,omitempty"`
	Height                   float64    `xml:"Height,omitempty"`
	Weight                   float64    `xml:"Weight,omitempty"`
	WeightUnit               string     `xml:"WeightUnit,omitempty"`
	DimensionUnit            string     `xml:"DimensionUnit,omitempty"`
	UnitWeight               string     `xml:"UnitWeight,omitempty"`
	UnitPrice                float64    `xml:"UnitPrice,omitempty"`
	UnitQuantity             string     `xml:"UnitQuantity,omitempty"`
	Price                    float64    `xml:"Price,omitempty"`
	ProductId                string     `xml:"ProductId,omitempty"`
	ItemId                   string     `xml:"ItemId,omitempty"`
	CountryOfManufacture     string     `xml:"CountryOfManufacture,omitempty"`
	SedLineAmt               float64    `xml:"SedLineAmt,omitempty"`
	COType                   string     `xml:"COType,omitempty"`
	COMarksAndNumbers        string     `xml:"COMarksAndNumbers,omitempty"`
	SEDInd                   string     `xml:"SEDInd,omitempty"`
	SpecialCommodities       int        `xml:"SpecialCommodities,omitempty"`
	NaftaPreferenceCriteria  string     `xml:"NaftaPreferenceCriteria,omitempty"`
	NaftaCoNetCost           string     `xml:"NaftaCoNetCost,omitempty"`
	NaftaCoNetcost_BeginDate *time.Time `xml:"NaftaCoNetcost_BeginDate,omitempty"`
	NaftaCoNetcost_EndDate   *time.Time `xml:"NaftaCoNetcost_EndDate,omitempty"`
	Currency                 string     `xml:"Currency,omitempty"`
	OutputCurrency           string     `xml:"OutputCurrency,omitempty"`
	Amount                   float64    `xml:"Amount,omitempty"`
	PictureURL               string     `xml:"PictureURL,omitempty"`
	ThumbnailURL             string     `xml:"ThumbnailURL,omitempty"`
}

type OrderDangerousGoodsDetail struct {
	DangerousGoodsTechnicalName       string  `xml:"DangerousGoodsTechnicalName,omitempty"`
	DangerousGoodsUnNumber            string  `xml:"DangerousGoodsUnNumber,omitempty"`
	DangerousGoodsNumberOfUnits       int     `xml:"DangerousGoodsNumberOfUnits,omitempty"`
	DangerousGoodsPackingType         string  `xml:"DangerousGoodsPackingType,omitempty"`
	DangerousGoodsPackingInstructions string  `xml:"DangerousGoodsPackingInstructions,omitempty"`
	DangerousGoodsTitleOfSignatory    string  `xml:"DangerousGoodsTitleOfSignatory,omitempty"`
	DangerousGoodsProperShippingName  string  `xml:"DangerousGoodsProperShippingName,omitempty"`
	DangerousGoodsNameOfSignatory     string  `xml:"DangerousGoodsNameOfSignatory,omitempty"`
	DangerousGoodsPercentageNumber    float64 `xml:"DangerousGoodsPercentageNumber,omitempty"`
	DangerousGoodsPlaceOfSignatory    string  `xml:"DangerousGoodsPlaceOfSignatory,omitempty"`
	DangerousGoodsCommodityCount      int     `xml:"DangerousGoodsCommodityCount,omitempty"`
	HazmatCode                        int     `xml:"HazmatCode,omitempty"`
	HazmatDOTShippingName             string  `xml:"HazmatDOTShippingName,omitempty"`
	HazmatClassOrDivision             string  `xml:"HazmatClassOrDivision,omitempty"`
	HazmatLabelType                   string  `xml:"HazmatLabelType,omitempty"`
	HazmatUnitQuantity                float64 `xml:"HazmatUnitQuantity,omitempty"`
	HazmatUnitType                    string  `xml:"HazmatUnitType,omitempty"`
	HazmatContactName                 string  `xml:"HazmatContactName,omitempty"`
	HazmatContactPhone                string  `xml:"HazmatContactPhone,omitempty"`
	HazmatCargoAircraftOnly           bool    `xml:"HazmatCargoAircraftOnly,omitempty"`
	AlcoholPackagingType              int     `xml:"AlcoholPackagingType,omitempty"`
	AlcoholType                       int     `xml:"AlcoholType,omitempty"`
	AlcoholVolumeInLiters             float64 `xml:"AlcoholVolumeInLiters,omitempty"`
	HazmatWeight                      float64 `xml:"HazmatWeight,omitempty"`
	HazmatPackagingInstructions       string  `xml:"HazmatPackagingInstructions,omitempty"`
	HazmatPackingGroup                string  `xml:"HazmatPackingGroup,omitempty"`
	HazmatRegulationSet               int     `xml:"HazmatRegulationSet,omitempty"`
	HazmatFlag                        bool    `xml:"HazmatFlag,omitempty"`
}

type PackagesList struct {
	Packages []OrderPackage `xml:"Package,omitempty"`
}

type OrderPackage struct {
	DimWeight                 float64 `xml:"DimWeight,omitempty"`
	TrackingNumber            string  `xml:"TrackingNumber,omitempty"`
	NotificationType          int     `xml:"NotificationType,omitempty"`
	NotificationValue         string  `xml:"NotificationValue,omitempty"`
	Length                    float64 `xml:"Length,omitempty"`
	Height                    float64 `xml:"Height,omitempty"`
	Width                     float64 `xml:"Width,omitempty"`
	PackagingType             int     `xml:"PackagingType,omitempty"`
	DeclaredValue             float64 `xml:"DeclaredValue,omitempty"`
	ContentDescription        string  `xml:"ContentDescription,omitempty"`
	ReferenceNumber           string  `xml:"ReferenceNumber,omitempty"`
	Weight                    float64 `xml:"Weight,omitempty"`
	WeightUnit                string  `xml:"WeightUnit,omitempty"`
	AddlPackageReference1     string  `xml:"AddlPackageReference1,omitempty"`
	AddlPackageReference2     string  `xml:"AddlPackageReference2,omitempty"`
	AddlPackageReference3     string  `xml:"AddlPackageReference3,omitempty"`
	AddlPackageReference4     string  `xml:"AddlPackageReference4,omitempty"`
	AddlPackageReference5     string  `xml:"AddlPackageReference5,omitempty"`
	AddlPackageReference6     string  `xml:"AddlPackageReference6,omitempty"`
	AddlPackageReference7     string  `xml:"AddlPackageReference7,omitempty"`
	AddlPackageReference8     string  `xml:"AddlPackageReference8,omitempty"`
	DeliveryConfirmationFlag  bool    `xml:"DeliveryConfirmationFlag,omitempty"`
	SignatureConfirmationFlag bool    `xml:"SignatureConfirmationFlag,omitempty"`
	SignatureConfirmationType int     `xml:"SignatureConfirmationType,omitempty"`
	SignatureReleaseFlag      bool    `xml:"SignatureReleaseFlag,omitempty"`
	BalloonFlag               bool    `xml:"BalloonFlag,omitempty"`
	OversizeFlag              bool    `xml:"OversizeFlag,omitempty"`
	NonMachinableFlag         bool    `xml:"NonMachinableFlag,omitempty"`
	USPSRateIndicator         string  `xml:"USPSRateIndicator,omitempty"`
	DryIceWeight              float64 `xml:"DryIceWeight,omitempty"`
	InsuranceAmount           float64 `xml:"InsuranceAmount,omitempty"`
	InsuranceType             int     `xml:"InsuranceType,omitempty"`
	InsurancePayType          int     `xml:"InsurancePayType,omitempty"`
	EndorsementType           int     `xml:"EndorsementType,omitempty"`
	FragileFlag               bool    `xml:"FragileFlag,omitempty"`
	AdditionalHandlingFlag    bool    `xml:"AdditionalHandlingFlag,omitempty"`
	USPSSortLevel             int     `xml:"USPSSortLevel,omitempty"`
	USPSODEnclosedMailClass   int     `xml:"USPSODEnclosedMailClass,omitempty"`
	ProcessCategory           int     `xml:"ProcessCategory,omitempty"`
	ShipToAttn                string  `xml:"ShipToAttn,omitempty"`
}

type OrderInternationalDetail struct {
	CN22DescriptionType                  int        `xml:"CN22DescriptionType,omitempty"`
	ShipperECCN                          string     `xml:"ShipperECCN,omitempty"`
	ShipperXTN                           string     `xml:"ShipperXTN,omitempty"`
	ShipperFTSR                          string     `xml:"ShipperFTSR,omitempty"`
	Incoterms                            string     `xml:"Incoterms,omitempty"`
	ExportLicense                        string     `xml:"ExportLicense,omitempty"`
	ImportLicense                        string     `xml:"ImportLicense,omitempty"`
	ConsigneeEIN                         string     `xml:"ConsigneeEIN,omitempty"`
	ConsigneeVAT                         string     `xml:"ConsigneeVAT,omitempty"`
	SignatureName                        string     `xml:"SignatureName,omitempty"`
	SignatureTitle                       string     `xml:"SignatureTitle,omitempty"`
	ExportReasonCode                     int        `xml:"ExportReasonCode,omitempty"`
	ShipperEIN                           string     `xml:"ShipperEIN,omitempty"`
	SEDFilingOption                      int        `xml:"SEDFilingOption,omitempty"`
	SignaturePhone                       string     `xml:"SignaturePhone,omitempty"`
	SignatureEmailAddress                string     `xml:"SignatureEmailAddress,omitempty"`
	IsDepartmentOfStateShipment          bool       `xml:"IsDepartmentOfStateShipment,omitempty"`
	IsDepartmentOfStateExempt            bool       `xml:"IsDepartmentOfStateExempt,omitempty"`
	CommercialInvoiceNumber              string     `xml:"CommercialInvoiceNumber,omitempty"`
	CommercialInvoiceOtherCharges        float64    `xml:"CommercialInvoiceOtherCharges,omitempty"`
	Use3rdPartyBrokerFlag                bool       `xml:"Use3rdPartyBrokerFlag,omitempty"`
	BrokerShipAlertFlag                  bool       `xml:"BrokerShipAlertFlag,omitempty"`
	BrokerDeliveryNotificationFlag       bool       `xml:"BrokerDeliveryNotificationFlag,omitempty"`
	BrokerExceptionNotificationFlag      bool       `xml:"BrokerExceptionNotificationFlag,omitempty"`
	BrokerEmailFormat                    string     `xml:"BrokerEmailFormat,omitempty"`
	BrokerAccountNumber                  string     `xml:"BrokerAccountNumber,omitempty"`
	AdditionalComments                   string     `xml:"AdditionalComments,omitempty"`
	AdmissibilityPackagingType           int        `xml:"AdmissibilityPackagingType,omitempty"`
	SenderTINType                        string     `xml:"SenderTINType,omitempty"`
	SenderTINNumber                      string     `xml:"SenderTINNumber,omitempty"`
	RecipientTINNumber                   string     `xml:"RecipientTINNumber,omitempty"`
	AESTransactionNumber                 string     `xml:"AESTransactionNumber,omitempty"`
	PartiesToTrans                       string     `xml:"PartiesToTrans,omitempty"`
	ExportInformationCode                int        `xml:"ExportInformationCode,omitempty"`
	ExportLicenseExpDate                 *time.Time `xml:"ExportLicenseExpDate,omitempty"`
	BrokerageId                          string     `xml:"BrokerageId,omitempty"`
	WorldEaseFlag                        bool       `xml:"WorldEaseFlag,omitempty"`
	WorldEaseCounter                     int        `xml:"WorldEaseCounter,omitempty"`
	GCCNNumber                           string     `xml:"GCCNNumber,omitempty"`
	WorldEaseTrackingNumber              string     `xml:"WorldEaseTrackingNumber,omitempty"`
	WorldEasePortNumber                  string     `xml:"WorldEasePortNumber,omitempty"`
	WorldEasePortZip                     string     `xml:"WorldEasePortZip,omitempty"`
	WorldEasePortName                    string     `xml:"WorldEasePortName,omitempty"`
	WorldEasePortCountry                 string     `xml:"WorldEasePortCountry,omitempty"`
	WorldEasePackageCount                int        `xml:"WorldEasePackageCount,omitempty"`
	FTRExemtionCode                      string     `xml:"FTRExemtionCode,omitempty"`
	ShippingFormContent1                 string     `xml:"ShippingFormContent1,omitempty"`
	ShippingFormContent2                 string     `xml:"ShippingFormContent2,omitempty"`
	ShippingFormContent3                 string     `xml:"ShippingFormContent3,omitempty"`
	ShippingFormContent4                 string     `xml:"ShippingFormContent4,omitempty"`
	ShippingFormContent5                 string     `xml:"ShippingFormContent5,omitempty"`
	ShippingFormContent6                 string     `xml:"ShippingFormContent6,omitempty"`
	FormRequested                        string     `xml:"FormRequested,omitempty"`
	FormGenerationType                   int        `xml:"FormGenerationType,omitempty"`
	NaftaProducer                        *Address   `xml:"NaftaProducer,omitempty"`
	NaftaBlanketPeriod_BeginDate         *time.Time `xml:"NaftaBlanketPeriod_BeginDate,omitempty"`
	NaftaBlanketPeriod_EndDate           *time.Time `xml:"NaftaBlanketPeriod_EndDate,omitempty"`
	NaftaProducerTaxId                   string     `xml:"NaftaProducerTaxId,omitempty"`
	NaftaProducerDeterminationType       int        `xml:"NaftaProducerDeterminationType,omitempty"`
	ExportDate                           *time.Time `xml:"ExportDate,omitempty"`
	ExportingCarrier                     string     `xml:"ExportingCarrier,omitempty"`
	InbondCode                           int        `xml:"InbondCode,omitempty"`
	EntryNumber                          string     `xml:"EntryNumber,omitempty"`
	PointOfOrigin                        string     `xml:"PointOfOrigin,omitempty"`
	ModeOfTransport                      string     `xml:"ModeOfTransport,omitempty"`
	PortOfExport                         string     `xml:"PortOfExport,omitempty"`
	PortOfUnloading                      string     `xml:"PortOfUnloading,omitempty"`
	LoadingPier                          string     `xml:"LoadingPier,omitempty"`
	RoutedExportTransactionIndicator     bool       `xml:"RoutedExportTransactionIndicator,omitempty"`
	ContainerizedIndicator               bool       `xml:"ContainerizedIndicator,omitempty"`
	ExportType                           int        `xml:"ExportType,omitempty"`
	ForwarderTaxId                       string     `xml:"ForwarderTaxId,omitempty"`
	ExceptionCode                        int        `xml:"ExceptionCode,omitempty"`
	SoldToOption                         int        `xml:"SoldToOption,omitempty"`
	TradeDirectProductType               int        `xml:"TradeDirectProductType,omitempty"`
	ImporterAccountNumber                string     `xml:"ImporterAccountNumber,omitempty"`
	ImportControlFlag                    bool       `xml:"ImportControlFlag,omitempty"`
	CommercialInvoiceRemovalFlag         bool       `xml:"CommercialInvoiceRemovalFlag,omitempty"`
	CommercialInvoiceFreightChargeAmount float64    `xml:"CommercialInvoiceFreightChargeAmount,omitempty"`
	IssuingAgencyCode                    string     `xml:"IssuingAgencyCode,omitempty"`
	LicencePlateNumber                   string     `xml:"LicencePlateNumber,omitempty"`
	Importer                             *Address   `xml:"Importer,omitempty"`
	Broker                               *Address   `xml:"Broker,omitempty"`
	UltimateConsignee                    *Address   `xml:"UltimateConsignee,omitempty"`
	UltimateConsigneeAccountNumber       string     `xml:"UltimateConsigneeAccountNumber,omitempty"`
	UltimateConsigneeTaxId               string     `xml:"UltimateConsigneeTaxId,omitempty"`
	Forwarder                            *Address   `xml:"Forwarder,omitempty"`
	IntermediateConsignee                *Address   `xml:"IntermediateConsignee,omitempty"`
	DutyAccountNumber                    string     `xml:"DutyAccountNumber,omitempty"`
	DutyBillType                         int        `xml:"DutyBillType,omitempty"`
	DutyServiceType                      int        `xml:"DutyServiceType,omitempty"`
	ImporterVATNumber                    string     `xml:"ImporterVATNumber,omitempty"`
	DutyBillPayor                        *Address   `xml:"DutyBillPayor,omitempty"`
	RecipientCustomIdType                int        `xml:"RecipientCustomIdType,omitempty"`
}

type OrderFreightShipment struct {
	DeliveryInstructions    string               `xml:"DeliveryInstructions,omitempty"`
	PickupInstructions      string               `xml:"PickupInstructions,omitempty"`
	BillOfLadingNumber      string               `xml:"BillOfLadingNumber,omitempty"`
	TotalHandlingUnitsCount int                  `xml:"TotalHandlingUnitsCount,omitempty"`
	CollectTermsType        int                  `xml:"CollectTermsType,omitempty"`
	DeclaredValueUnits      string               `xml:"DeclaredValueUnits,omitempty"`
	CoverageDetail          *CoverageDetail      `xml:"CoverageDetail,omitempty"`
	BillingContactAddress   *Address             `xml:"BillingContactAddress,omitempty"`
	FreightBrokerAddress    *Address             `xml:"FreightBrokerAddress,omitempty"`
	HandlingUnits           *ArrayOfHandlingUnit `xml:"HandlingUnits,omitempty"`
	ShippingUnits           *ArrayOfShippingUnit `xml:"ShippingUnits,omitempty"`
}

type CoverageDetail struct {
	CoverageType   int     `xml:"CoverageType,omitempty"`
	CoverageAmount float64 `xml:"CoverageAmount,omitempty"`
}

type ArrayOfHandlingUnit struct {
	HandlingUnits []HandlingUnit `xml:"HandlingUnit,omitempty"`
}

type HandlingUnit struct {
	Type            int     `xml:"Type,omitempty"`
	Weight          float64 `xml:"Weight,omitempty"`
	WeightUnit      string  `xml:"WeightUnit,omitempty"`
	Length          float64 `xml:"Length,omitempty"`
	Width           float64 `xml:"Width,omitempty"`
	Height          float64 `xml:"Height,omitempty"`
	DimUnit         string  `xml:"DimUnit,omitempty"`
	DeclaredValue   float64 `xml:"DeclaredValue,omitempty"`
	InsuranceAmount float64 `xml:"InsuranceAmount,omitempty"`
}

type ArrayOfShippingUnit struct {
	ShippingUnits []ShippingUnit `xml:"ShippingUnit,omitempty"`
}

type ShippingUnit struct {
	ShippingUnitType                      int     `xml:"ShippingUnitType,omitempty"`
	ShippingUnitQuantity                  string  `xml:"ShippingUnitQuantity,omitempty"`
	ShippingUnitQuantityUnit              int     `xml:"ShippingUnitQuantityUnit,omitempty"`
	UnitQuantityPerShippingUnit           int     `xml:"UnitQuantityPerShippingUnit,omitempty"`
	HandlingUnitsPerShippingUnit          int     `xml:"HandlingUnitsPerShippingUnit,omitempty"`
	Description                           string  `xml:"Description,omitempty"`
	CommercialInvoiceCommodityDescription string  `xml:"CommercialInvoiceCommodityDescription,omitempty"`
	Weight                                float64 `xml:"Weight,omitempty"`
	WeightUnit                            string  `xml:"WeightUnit,omitempty"`
	NmfcClassCode                         int     `xml:"NmfcClassCode,omitempty"`
	NmfcClassCodeSpecified                bool    `xml:"NmfcClassCodeSpecified,omitempty"`
	NmfcItemCode                          string  `xml:"NmfcItemCode,omitempty"`
	CountryOfManufacture                  string  `xml:"CountryOfManufacture,omitempty"`
	HarmonizedCode                        string  `xml:"HarmonizedCode,omitempty"`
	Length                                float64 `xml:"Length,omitempty"`
	Width                                 float64 `xml:"Width,omitempty"`
	Height                                float64 `xml:"Height,omitempty"`
	DimUnit                               string  `xml:"DimUnit,omitempty"`
	Comment                               string  `xml:"Comment,omitempty"`
	DeclaredValue                         float64 `xml:"DeclaredValue,omitempty"`
	InsuranceAmount                       float64 `xml:"InsuranceAmount,omitempty"`
	PurchaseOrderNumber                   string  `xml:"PurchaseOrderNumber,omitempty"`
	CommodityWeight                       float64 `xml:"CommodityWeight,omitempty"`
	CommodityPieces                       string  `xml:"CommodityPieces,omitempty"`
}

type AbolImportOrderResponse struct {
	AbolImportOrderResult *AbolImportOrderResult `xml:"AbolImportOrderResult,omitempty"`
}

type AbolImportOrderResult struct {
	ImportOrderResult *ImportOrderResult `xml:"ImportOrderResult,omitempty"`
}

type ImportOrderResult struct {
	Errors  *Errors `xml:"Errors,omitempty"`
	Status  *Status `xml:"Status,omitempty"`
	Order   *order  `xml:"Order,omitempty"`
	Version string  `xml:"Version,omitempty"`
}

type Errors struct {
	Error *Error `xml:"Error,omitempty"`
}

type Error struct {
	ErrorType *ErrorType `xml:"ErrorType,omitempty"`
}

type ErrorType struct {
	Code    int    `xml:"Code,omitempty"`
	Message string `xml:"Message,omitempty"`
}

type Status struct {
	Code    int    `xml:"Code,omitempty"`
	Message string `xml:"Message,omitempty"`
}

type order struct {
	OrderResponses []OrderResponse `xml:"OrderResponse,omitempty"`
}

type OrderResponse struct {
	OrderNbr string  `xml:"OrderNbr,omitempty"`
	OrderID  int     `xml:"OrderID,omitempty"`
	BatchNbr int     `xml:"BatchNbr,omitempty"`
	Status   *Status `xml:"Status,omitempty"`
}

type AbolApiShipmentRequest struct {
	Request *AbolApiShipment_Request `xml:"Request,omitempty"`
}

type AbolApiShipment_Request struct {
	Authentication *Authentication `xml:"Authentication,omitempty"`
	Shipment       *Shipment       `xml:"Shipment,omitempty"`
}

type AbolApiShipmentResponse struct {
	ShipResult *ShipResult `xml:"ShipResult,omitempty"`
}

type ShipResult struct {
	Error    *ShipError         `xml:"Error,omitempty"`
	Shipment *Shipment_Response `xml:"Shipment,omitempty"`
	Version  string             `xml:"Version,omitempty"`
}

type ShipError struct {
	Code    int    `xml:"Code,omitempty"`
	Message string `xml:"Message,omitempty"`
}

type Shipment_Response struct {
	DASFlag             bool             `xml:"DASFlag,omitempty"`
	RemoteAreaFlag      bool             `xml:"RemoteAreaFlag,omitempty"`
	HandlingFlag        bool             `xml:"HandlingFlag,omitempty"`
	PieceCountFlag      bool             `xml:"PieceCountFlag,omitempty"`
	FuelSurchargeFlag   bool             `xml:"FuelSurchargeFlag,omitempty"`
	ManifestGroupNbr    string           `xml:"ManifestGroupNbr,omitempty"`
	SortCode1           string           `xml:"SortCode1,omitempty"`
	SortCode2           string           `xml:"SortCode2,omitempty"`
	SortCode3           string           `xml:"SortCode3,omitempty"`
	SortCode4           string           `xml:"SortCode4,omitempty"`
	SortCode5           string           `xml:"SortCode5,omitempty"`
	SortCode6           string           `xml:"SortCode6,omitempty"`
	CODFee              float64          `xml:"CODFee,omitempty"`
	BillWeight          float64          `xml:"BillWeight,omitempty"`
	DimWeight           float64          `xml:"DimWeight,omitempty"`
	Charge              float64          `xml:"Charge,omitempty"`
	DASFee              float64          `xml:"DASFee,omitempty"`
	DCFee               float64          `xml:"DCFee,omitempty"`
	ExcessDimFee1       float64          `xml:"ExcessDimFee1,omitempty"`
	ExcessDimFee2       float64          `xml:"ExcessDimFee2,omitempty"`
	ExcessDimFee3       float64          `xml:"ExcessDimFee3,omitempty"`
	ExcessDimFee4       float64          `xml:"ExcessDimFee4,omitempty"`
	ExcessWeightFee     float64          `xml:"ExcessWeightFee,omitempty"`
	FuelCharge          float64          `xml:"FuelCharge,omitempty"`
	DangerousGoodsFee   float64          `xml:"DangerousGoodsFee,omitempty"`
	Oversize1Fee        float64          `xml:"Oversize1Fee,omitempty"`
	Oversize2Fee        float64          `xml:"Oversize2Fee,omitempty"`
	Oversize3Fee        float64          `xml:"Oversize3Fee,omitempty"`
	ResidentialFee      float64          `xml:"ResidentialFee,omitempty"`
	SaturdayPickupFee   float64          `xml:"SaturdayPickupFee,omitempty"`
	DeliveryDutyPaidFee float64          `xml:"DeliveryDutyPaidFee,omitempty"`
	NeutralDeliveryFee  float64          `xml:"NeutralDeliveryFee,omitempty"`
	RemoteAreaFee       float64          `xml:"RemoteAreaFee,omitempty"`
	SplitDutyFee        float64          `xml:"SplitDutyFee,omitempty"`
	TenThirtyAMFee      float64          `xml:"TenThirtyAMFee,omitempty"`
	HoldAtLocationFee   float64          `xml:"HoldAtLocationFee,omitempty"`
	PieceCntFee         float64          `xml:"PieceCntFee,omitempty"`
	ServiceCode         string           `xml:"ServiceCode,omitempty"`
	InsuranceCharge     float64          `xml:"InsuranceCharge,omitempty"`
	ServiceId           int              `xml:"ServiceId,omitempty"`
	DocumentList        *ArrayOfDocument `xml:"DocumentList,omitempty"`
	Pieces              *ArrayOfPiece    `xml:"Pieces,omitempty"`
}

type ArrayOfDocument struct {
	Documents []Document `xml:"Document,omitempty"`
}

type Document struct {
	Type    string `xml:"Type,omitempty"`
	URL     string `xml:"URL,omitempty"`
	Content string `xml:"Content,omitempty"`
}

type ArrayOfPiece struct {
	Pieces []Piece `xml:"Piece,omitempty"`
}

type Piece struct {
	TrackingNbr           string           `xml:"TrackingNbr,omitempty"`
	ShipmentCharges       *ShipmentCharges `xml:"ShipmentCharges,omitempty"`
	LabelList             *LabelList       `xml:"LabelList,omitempty"`
	ReferenceNumber       string           `xml:"ReferenceNumber,omitempty"`
	AddlPackageReference1 string           `xml:"AddlPackageReference1,omitempty"`
	AddlPackageReference2 string           `xml:"AddlPackageReference2,omitempty"`
	AddlPackageReference3 string           `xml:"AddlPackageReference3,omitempty"`
	AddlPackageReference4 string           `xml:"AddlPackageReference4,omitempty"`
	AddlPackageReference5 string           `xml:"AddlPackageReference5,omitempty"`
	AddlPackageReference6 string           `xml:"AddlPackageReference6,omitempty"`
	AddlPackageReference7 string           `xml:"AddlPackageReference7,omitempty"`
	AddlPackageReference8 string           `xml:"AddlPackageReference8,omitempty"`
}

type LabelImage string

func (li *LabelImage) Decode() ([]byte, error) {
	d := make([]byte, base64.StdEncoding.DecodedLen(len(*li)))
	_, err := base64.StdEncoding.Decode(d, []byte(*li))

	if err != nil {
		return nil, err
	}
	return d, nil
}

type LabelList struct {
	Label1                   LabelImage `xml:"Label1,omitempty"`
	Label2                   LabelImage `xml:"Label2,omitempty"`
	Label3                   LabelImage `xml:"Label3,omitempty"`
	Label4                   LabelImage `xml:"Label4,omitempty"`
	Label5                   LabelImage `xml:"Label5,omitempty"`
	IntlShippingFormContent1 string     `xml:"IntlShippingFormContent1,omitempty"`
	IntlShippingFormContent2 string     `xml:"IntlShippingFormContent2,omitempty"`
	IntlShippingFormContent3 string     `xml:"IntlShippingFormContent3,omitempty"`
	IntlShippingFormContent4 string     `xml:"IntlShippingFormContent4,omitempty"`
	IntlShippingFormContent5 string     `xml:"IntlShippingFormContent5,omitempty"`
	IntlShippingFormContent6 string     `xml:"IntlShippingFormContent6,omitempty"`
}

type AbolApiTrackRequest struct {
	Authentication                   *Authentication `xml:"Authentication,omitempty"`
	AcceptCarriersTermsAndConditions bool            `xml:"AcceptCarriersTermsAndConditions,omitempty"`
	Track                            *Track          `xml:"Track,omitempty"`
}

type Track struct {
	Packages *ArrayOfRequestPackages `xml:"Packages,omitempty"`
}

type ArrayOfRequestPackages struct {
	Packages []RequestPackages `xml:"Package,omitempty"`
}

type RequestPackages struct {
	PackageId   string `xml:"PackageId,omitempty"`
	ServiceCode string `xml:"ServiceCode,omitempty"`
}

type AbolApiTrackResponse struct {
	TrackResult *TrackResult `xml:"TrackResult,omitempty"`
}

type TrackResult struct {
	Error    *TrackError             `xml:"Error,omitempty"`
	Packages *ArrayOfResponsePackage `xml:"Packages,omitempty"`
	Version  string                  `xml:"Version,omitempty"`
}

type TrackError struct {
	Code        string `xml:"Code,omitempty"`
	Description string `xml:"Description,omitempty"`
}

type ArrayOfResponsePackage struct {
	Packages []ResponsePackage `xml:"Package,omitempty"`
}

type ResponsePackage struct {
	Status             string  `xml:"Status,omitempty"`
	PackageId          string  `xml:"PackageId,omitempty"`
	TrackingReference1 string  `xml:"TrackingReference1,omitempty"`
	TrackingReference2 string  `xml:"TrackingReference2,omitempty"`
	Events             *Events `xml:"Events,omitempty"`
}

type Events struct {
	Events []Event `xml:"Event,omitempty"`
}

type Event struct {
	Code          string     `xml:"Code,omitempty"`
	Description   string     `xml:"Description,omitempty"`
	EventDateTime *time.Time `xml:"EventDateTime,omitempty"`
	Location      *Location  `xml:"Location,omitempty"`
}

type Location struct {
	City        string `xml:"City,omitempty"`
	State       string `xml:"State,omitempty"`
	CountryCode string `xml:"CountryCode,omitempty"`
	Country     string `xml:"Country,omitempty"`
}

type AbolApiVoidRequest struct {
	Authentication *AbolApiVoid_Authentication `xml:"Authentication,omitempty"`
	Void           *AbolApiVoid_TrackingNo     `xml:"Void,omitempty"`
}

type AbolApiVoid_Authentication struct {
	ActivationKey string `xml:"ActivationKey,omitempty"`
	LoginName     string `xml:"LoginName,omitempty"`
	Password      string `xml:"Password,omitempty"`
	AltID         string `xml:"AltID,omitempty"`
}

type AbolApiVoid_TrackingNo struct {
	TrackingNumber string `xml:"TrackingNumber,omitempty"`
}

type AbolApiVoidResponse struct {
	AbolVoidResult *AbolVoidResult `xml:"AbolVoidResult,omitempty"`
}

type AbolVoidResult struct {
	Error   *ResponseError `xml:"Error,omitempty"`
	Version string         `xml:"Version,omitempty"`
}

type ResponseError struct {
	Code    int    `xml:"Code,omitempty"`
	Message string `xml:"Message,omitempty"`
}

type AbolApiCloseOutRequest struct {
	Request *AbolApiCloseOut_Request `xml:"Request,omitempty"`
}

type AbolApiCloseOut_Request struct {
	Authentication *Authentication `xml:"Authentication,omitempty"`
	Closeout       *Closeout       `xml:"Closeout,omitempty"`
}

type Closeout struct {
	RequestType          *CloseoutRequestType  `xml:"RequestType,omitempty"`
	RequestDetail        string                `xml:"RequestDetail,omitempty"`
	CloseOutTrackingNbrs *CloseOutTrackingNbrs `xml:"CloseOutTrackingNbrs,omitempty"`
}

type CloseOutTrackingNbrs struct {
	TrackingNbrs []string `xml:"TrackingNbr,omitempty"`
}

type AbolApiCloseOutResponse struct {
	AbolCloseoutResult *AbolCloseoutResult `xml:"AbolCloseoutResult,omitempty"`
}

type AbolCloseoutResult struct {
	Error              *CloseOutError    `xml:"Error,omitempty"`
	ConfirmationNumber string            `xml:"ConfirmationNumber,omitempty"`
	Version            string            `xml:"Version,omitempty"`
	ManifestDetails    []CloseOutDetails `xml:"ManifestDetails,omitempty"`
}

type CloseOutError struct {
	Code    int    `xml:"Code,omitempty"`
	Message string `xml:"Message,omitempty"`
}

type CloseOutDetails struct {
	Error        *CloseOutError        `xml:"Error,omitempty"`
	ManifestNbr  string                `xml:"ManifestNbr,omitempty"`
	ManifestType string                `xml:"ManifestType,omitempty"`
	TotalPieces  int                   `xml:"TotalPieces,omitempty"`
	TrackingNbrs *CloseOutTrackingNbrs `xml:"TrackingNbrs,omitempty"`
	DocumentList *DocumentList         `xml:"DocumentList,omitempty"`
}

type DocumentList struct {
	Documents []Document `xml:"Document,omitempty"`
}

type AbolClassify_Request struct {
	Authentication *Authentication `xml:"Authentication,omitempty"`
	Description    string          `xml:"Description,omitempty"`
}

type AbolClassify_Response struct {
	AbolClassifyResult *AbolClassifyResult `xml:"AbolClassifyResult,omitempty"`
}

type AbolClassifyResult struct {
	Error                *error                `xml:"Error,omitempty"`
	CategoriesResultList *CategoriesResultList `xml:"CategoriesResultList,omitempty"`
}

type CategoriesResultList struct {
	Categories *ArrayOfCategory `xml:"Categories,omitempty"`
}

type ArrayOfCategory struct {
	Categorys []Category `xml:"Category,omitempty"`
}

type Category struct {
	Id            string              `xml:"Id,attr,ommitempty"`
	Name          string              `xml:"Name,attr,ommitempty"`
	SubCategories *ArrayOfSubCategory `xml:"SubCategories,omitempty"`
}

type ArrayOfSubCategory struct {
	SubCategorys []SubCategory `xml:"SubCategory,omitempty"`
}

type SubCategory struct {
	Id    string                  `xml:"Id,attr,ommitempty"`
	Name  string                  `xml:"Name,attr,ommitempty"`
	Items *ArrayOfSubCategoryItem `xml:"Items,omitempty"`
}

type ArrayOfSubCategoryItem struct {
	SubCategoryItems []SubCategoryItem `xml:"SubCategoryItem,omitempty"`
}

type SubCategoryItem struct {
	Id   string `xml:"Id,attr,ommitempty"`
	Name string `xml:"Name,omitempty"`
}

type AbolDuty_Request struct {
	Authentication *Authentication    `xml:"Authentication,omitempty"`
	RequestDetail  *DutyRequestDetail `xml:"RequestDetail,omitempty"`
}

type DutyRequestDetail struct {
	ShipFromCountryCode        string                      `xml:"ShipFromCountryCode,omitempty"`
	ShipToCountryCode          string                      `xml:"ShipToCountryCode,omitempty"`
	ShipToState                string                      `xml:"ShipToState,omitempty"`
	CommercialImporterFlag     bool                        `xml:"CommercialImporterFlag,omitempty"`
	TotalMonthlyImportedWeight float64                     `xml:"TotalMonthlyImportedWeight,omitempty"`
	TotalMonthlyImportedValue  float64                     `xml:"TotalMonthlyImportedValue,omitempty"`
	ClassificationType         *ClassifyBy                 `xml:"ClassificationType,omitempty"`
	ItemSubCategoryItemIds     *ItemSubCategoryItemIds     `xml:"ItemSubCategoryItemIds,omitempty"`
	ItemHarmonizedCodes        *ItemHarmonizedCodes        `xml:"ItemHarmonizedCodes,omitempty"`
	ItemDescriptions           *ItemDescriptions           `xml:"ItemDescriptions,omitempty"`
	ItemCodes                  *ItemCodes                  `xml:"ItemCodes,omitempty"`
	ItemUnitPrices             *ItemUnitPrices             `xml:"ItemUnitPrices,omitempty"`
	ItemWeights                *ItemWeights                `xml:"ItemWeights,omitempty"`
	ItemQuantities             *ItemQuantities             `xml:"ItemQuantities,omitempty"`
	ItemUnitQuantities         *ItemUnitQuantities         `xml:"ItemUnitQuantities,omitempty"`
	ItemAmounts                *ItemAmounts                `xml:"ItemAmounts,omitempty"`
	ItemCountriesOfManufacture *ItemCountriesOfManufacture `xml:"ItemCountriesOfManufacture,omitempty"`
	ReferenceNumbers           *ReferenceNumbers           `xml:"ReferenceNumbers,omitempty"`
	TotalPostage               float64                     `xml:"TotalPostage,omitempty"`
	InsuranceAmount            float64                     `xml:"InsuranceAmount,omitempty"`
	CurrencyCode               string                      `xml:"CurrencyCode,omitempty"`
	Weight                     float64                     `xml:"Weight,omitempty"`
	DetailedResponseFlag       bool                        `xml:"DetailedResponseFlag,omitempty"`
	HSCodeFlag                 bool                        `xml:"HSCodeFlag,omitempty"`
}

type ItemSubCategoryItemIds struct {
	SubCategoryItemIds []int `xml:"SubCategoryItemId,omitempty"`
}

type ItemHarmonizedCodes struct {
	HarmonizedCodes []string `xml:"HarmonizedCode,omitempty"`
}

type ItemDescriptions struct {
	Descriptions []string `xml:"Description,omitempty"`
}

type ItemCodes struct {
	Codes []string `xml:"Code,omitempty"`
}

type ItemUnitPrices struct {
	UnitPrices []float64 `xml:"UnitPrice,omitempty"`
}

type ItemWeights struct {
	Weights []float64 `xml:"Weight,omitempty"`
}

type ItemQuantities struct {
	Quantitys []float64 `xml:"Quantity,omitempty"`
}

type ItemUnitQuantities struct {
	UnitQuantitys []AmountUnit `xml:"UnitQuantity,omitempty"`
}

type ItemAmounts struct {
	ItemAmounts []string `xml:"ItemAmount,omitempty"`
}

type ItemCountriesOfManufacture struct {
	CountryOfManufactures []string `xml:"CountryOfManufacture,omitempty"`
}

type ReferenceNumbers struct {
	ReferenceNumbers []string `xml:"ReferenceNumber,omitempty"`
}

type AbolDuty_Response struct {
	DutyTaxResult *DutyTaxResult `xml:"DutyTaxResult,omitempty"`
}

type DutyTaxResult struct {
	Id           string             `xml:"Id,attr,ommitempty"`
	Error        *error             `xml:"Error,omitempty"`
	Items        *ArrayOfItemResult `xml:"Items,omitempty"`
	TotalCharges *TotalCharges      `xml:"TotalCharges,omitempty"`
}

type ArrayOfItemResult struct {
	ItemResults []ItemResult `xml:"ItemResult,omitempty"`
}

type ItemResult struct {
	Id                    string           `xml:"Id,attr,ommitempty"`
	Hscode                string           `xml:"Hscode,attr,ommitempty"`
	Reference             string           `xml:"Reference,attr,ommitempty"`
	Match                 string           `xml:"Match,attr,ommitempty"`
	Matchcode             string           `xml:"Matchcode,attr,ommitempty"`
	Dutyspread            string           `xml:"Dutyspread,attr,ommitempty"`
	Customsvalue          *CustomsValue    `xml:"Customsvalue,omitempty"`
	Duty                  *Duty            `xml:"Duty,omitempty"`
	Salestax              *SalesTax        `xml:"Salestax,omitempty"`
	Additionalimporttaxes *AdditionalTaxes `xml:"Additionalimporttaxes,omitempty"`
	Total                 *Total           `xml:"Total,omitempty"`
}

type CustomsValue struct {
	Name   string  `xml:"Name,attr,ommitempty"`
	Amount *Amount `xml:"Amount,omitempty"`
}

type Amount struct {
	Currency string  `xml:"Currency,attr,ommitempty"`
	Value    float64 `xml:"Value,omitempty"`
}

type Duty struct {
	Amount *Amount `xml:"Amount,omitempty"`
}

type SalesTax struct {
	Name   string  `xml:"Name,attr,ommitempty"`
	Amount *Amount `xml:"Amount,omitempty"`
}

type AdditionalTaxes struct {
	Taxes *ArrayOfTax `xml:"Taxes,omitempty"`
}

type ArrayOfTax struct {
	Taxs []Tax `xml:"Tax,omitempty"`
}

type Tax struct {
	Name   string  `xml:"Name,attr,ommitempty"`
	Amount *Amount `xml:"Amount,omitempty"`
}

type Total struct {
	Amount *Amount `xml:"Amount,omitempty"`
}

type TotalCharges struct {
	Customsvalue          *CustomsValue    `xml:"Customsvalue,omitempty"`
	Duty                  *Duty            `xml:"Duty,omitempty"`
	Salestax              *SalesTax        `xml:"Salestax,omitempty"`
	Additionalimporttaxes *AdditionalTaxes `xml:"Additionalimporttaxes,omitempty"`
	Total                 *Total           `xml:"Total,omitempty"`
}

type AbolApiSimpleShipment struct {
	AbolSimpleShip *AbolApiSimpleShipmentRequest `xml:"AbolSimpleShip,omitempty"`
}

type AbolApiSimpleShipmentRequest struct {
	Request *AbolApiSimpleShipment_Request `xml:"Request,omitempty"`
}

type AbolApiSimpleShipment_Request struct {
	Authentication *Authentication `xml:"Authentication,omitempty"`
	SimpleShipment *SimpleShipment `xml:"SimpleShipment,omitempty"`
}

type SimpleShipment struct {
	CarrierId                           string                       `xml:"CarrierId,omitempty"`
	BillType                            int                          `xml:"BillType,omitempty"`
	BillAccountNumber                   string                       `xml:"BillAccountNumber,omitempty"`
	BillAccountZipCode                  string                       `xml:"BillAccountZipCode,omitempty"`
	DriverInstructions                  string                       `xml:"DriverInstructions,omitempty"`
	ResidentialFlag                     bool                         `xml:"ResidentialFlag,omitempty"`
	SaturdayDeliveryFlag                bool                         `xml:"SaturdayDeliveryFlag,omitempty"`
	SaturdayPickupFlag                  bool                         `xml:"SaturdayPickupFlag,omitempty"`
	LabelType                           int                          `xml:"LabelType,omitempty"`
	LabelSize                           int                          `xml:"LabelSize,omitempty"`
	HolidayDeliveryFlag                 bool                         `xml:"HolidayDeliveryFlag,omitempty"`
	SundayDeliveryFlag                  bool                         `xml:"SundayDeliveryFlag,omitempty"`
	DimensionUnit                       string                       `xml:"DimensionUnit,omitempty"`
	WeightUnit                          string                       `xml:"WeightUnit,omitempty"`
	ShipDate                            *time.Time                   `xml:"ShipDate,omitempty"`
	ServiceCode                         string                       `xml:"ServiceCode,omitempty"`
	DeclaredValue                       float64                      `xml:"DeclaredValue,omitempty"`
	CurrencyCode                        string                       `xml:"CurrencyCode,omitempty"`
	DimWeight                           float64                      `xml:"DimWeight,omitempty"`
	TrackingNumber                      string                       `xml:"TrackingNumber,omitempty"`
	Length                              float64                      `xml:"Length,omitempty"`
	Height                              float64                      `xml:"Height,omitempty"`
	Width                               float64                      `xml:"Width,omitempty"`
	PackagingType                       int                          `xml:"PackagingType,omitempty"`
	ReferenceNumber                     string                       `xml:"ReferenceNumber,omitempty"`
	Weight                              float64                      `xml:"Weight,omitempty"`
	AddlPackageReference1               string                       `xml:"AddlPackageReference1,omitempty"`
	AddlPackageReference2               string                       `xml:"AddlPackageReference2,omitempty"`
	AddlPackageReference3               string                       `xml:"AddlPackageReference3,omitempty"`
	AddlPackageReference4               string                       `xml:"AddlPackageReference4,omitempty"`
	AddlPackageReference5               string                       `xml:"AddlPackageReference5,omitempty"`
	AddlPackageReference6               string                       `xml:"AddlPackageReference6,omitempty"`
	AddlPackageReference7               string                       `xml:"AddlPackageReference7,omitempty"`
	AddlPackageReference8               string                       `xml:"AddlPackageReference8,omitempty"`
	ContentDescription                  string                       `xml:"ContentDescription,omitempty"`
	NotificationType                    int                          `xml:"NotificationType,omitempty"`
	NotificationValue                   string                       `xml:"NotificationValue,omitempty"`
	ImageRotation                       int                          `xml:"ImageRotation,omitempty"`
	EveningDeliveryFlag                 bool                         `xml:"EveningDeliveryFlag,omitempty"`
	ReturnType                          int                          `xml:"ReturnType,omitempty"`
	ReturnServiceCode                   string                       `xml:"ReturnServiceCode,omitempty"`
	DeliveryConfirmationFlag            bool                         `xml:"DeliveryConfirmationFlag,omitempty"`
	SignatureConfirmationFlag           bool                         `xml:"SignatureConfirmationFlag,omitempty"`
	SignatureConfirmationType           int                          `xml:"SignatureConfirmationType,omitempty"`
	InsuranceAmount                     float64                      `xml:"InsuranceAmount,omitempty"`
	InsurancePayType                    int                          `xml:"InsurancePayType,omitempty"`
	InsuranceType                       int                          `xml:"InsuranceType,omitempty"`
	EndorsementType                     int                          `xml:"EndorsementType,omitempty"`
	OverrideShipToAddressValidationFlag bool                         `xml:"OverrideShipToAddressValidationFlag,omitempty"`
	AdditionalAccountNumber             string                       `xml:"AdditionalAccountNumber,omitempty"`
	CarriageValue                       float64                      `xml:"CarriageValue,omitempty"`
	ShipFrom                            *Address                     `xml:"ShipFrom,omitempty"`
	ShipTo                              *Address                     `xml:"ShipTo,omitempty"`
	Return                              *Address                     `xml:"Return,omitempty"`
	SimpleExportLineItemList            *SimpleExportLineItems       `xml:"SimpleExportLineItemList,omitempty"`
	SimpleInternationalShipment         *SimpleInternationalShipment `xml:"SimpleInternationalShipment,omitempty"`
}

type SimpleExportLineItems struct {
	SimpleExportLineItems []SimpleExportLineItem `xml:"SimpleExportLineItem,omitempty"`
}

type SimpleExportLineItem struct {
	Code                 string      `xml:"Code,omitempty"`
	Description          string      `xml:"Description,omitempty"`
	Weight               float64     `xml:"Weight,omitempty"`
	UnitWeight           string      `xml:"UnitWeight,omitempty"`
	Price                float64     `xml:"Price,omitempty"`
	Quantity             int         `xml:"Quantity,omitempty"`
	UnitQuantity         *AmountUnit `xml:"UnitQuantity,omitempty"`
	UnitPrice            float64     `xml:"UnitPrice,omitempty"`
	CountryOfManufacture string      `xml:"CountryOfManufacture,omitempty"`
	Harmonizedcode       string      `xml:"Harmonizedcode,omitempty"`
	SalesTaxAmount       float64     `xml:"SalesTaxAmount,omitempty"`
	DutyAmount           float64     `xml:"DutyAmount,omitempty"`
	Currency             string      `xml:"Currency,omitempty"`
	OutputCurrency       string      `xml:"OutputCurrency,omitempty"`
	Amount               float64     `xml:"Amount,omitempty"`
}

type SimpleInternationalShipment struct {
	DeclaredValue              float64 `xml:"DeclaredValue,omitempty"`
	CN22DescriptionType        int     `xml:"CN22DescriptionType,omitempty"`
	Incoterms                  string  `xml:"Incoterms,omitempty"`
	ExportReasonCode           int     `xml:"ExportReasonCode,omitempty"`
	ShipperEin                 string  `xml:"ShipperEin,omitempty"`
	SEDFilingOption            int     `xml:"SEDFilingOption,omitempty"`
	AESTransactionNumber       string  `xml:"AESTransactionNumber,omitempty"`
	FTRExemtionCode            string  `xml:"FTRExemtionCode,omitempty"`
	FormRequested              string  `xml:"FormRequested,omitempty"`
	FormGenerationType         int     `xml:"FormGenerationType,omitempty"`
	ExportType                 int     `xml:"ExportType,omitempty"`
	ExceptionCode              int     `xml:"ExceptionCode,omitempty"`
	CustomFormType             int     `xml:"CustomFormType,omitempty"`
	DutyAccountNumber          string  `xml:"DutyAccountNumber,omitempty"`
	DutyBillType               int     `xml:"DutyBillType,omitempty"`
	DutyServiceType            int     `xml:"DutyServiceType,omitempty"`
	IncludeDutyCalculation     bool    `xml:"IncludeDutyCalculation,omitempty"`
	CommercialImporterFlag     bool    `xml:"CommercialImporterFlag,omitempty"`
	TotalMonthlyImportedWeight float64 `xml:"TotalMonthlyImportedWeight,omitempty"`
	TotalMonthlyImportedValue  float64 `xml:"TotalMonthlyImportedValue,omitempty"`
}

type AbolApiSimpleShipmentResponse struct {
	SimpleShipmentResult *SimpleShipmentResult `xml:"http://abolsoft.com/webservices/ SimpleShipmentResult,omitempty"`
}

type SimpleShipmentResult struct {
	Error          *ShipError              `xml:"Error,omitempty"`
	SimpleShipment *SimpleShipmentResponse `xml:"SimpleShipment,omitempty"`
	Version        string                  `xml:"Version,omitempty"`
}

type SimpleShipmentResponse struct {
	DASFlag                 bool                   `xml:"DASFlag,omitempty"`
	RemoteAreaFlag          bool                   `xml:"RemoteAreaFlag,omitempty"`
	HandlingFlag            bool                   `xml:"HandlingFlag,omitempty"`
	FuelSurchargeFlag       bool                   `xml:"FuelSurchargeFlag,omitempty"`
	SortCode1               string                 `xml:"SortCode1,omitempty"`
	SortCode2               string                 `xml:"SortCode2,omitempty"`
	SortCode3               string                 `xml:"SortCode3,omitempty"`
	SortCode4               string                 `xml:"SortCode4,omitempty"`
	SortCode5               string                 `xml:"SortCode5,omitempty"`
	SortCode6               string                 `xml:"SortCode6,omitempty"`
	CODFee                  float64                `xml:"CODFee,omitempty"`
	BillWeight              float64                `xml:"BillWeight,omitempty"`
	DimWeight               float64                `xml:"DimWeight,omitempty"`
	Charge                  float64                `xml:"Charge,omitempty"`
	DCFee                   float64                `xml:"DCFee,omitempty"`
	FuelCharge              float64                `xml:"FuelCharge,omitempty"`
	ResidentialFee          float64                `xml:"ResidentialFee,omitempty"`
	SaturdayPickupFee       float64                `xml:"SaturdayPickupFee,omitempty"`
	DeliveryDutyPaidFee     float64                `xml:"DeliveryDutyPaidFee,omitempty"`
	NeutralDeliveryFee      float64                `xml:"NeutralDeliveryFee,omitempty"`
	RemoteAreaFee           float64                `xml:"RemoteAreaFee,omitempty"`
	SplitDutyFee            float64                `xml:"SplitDutyFee,omitempty"`
	TenThirtyAMFee          float64                `xml:"TenThirtyAMFee,omitempty"`
	HoldAtLocationFee       float64                `xml:"HoldAtLocationFee,omitempty"`
	ServiceCode             string                 `xml:"ServiceCode,omitempty"`
	InsuranceCharge         float64                `xml:"InsuranceCharge,omitempty"`
	ServiceId               int                    `xml:"ServiceId,omitempty"`
	DocumentList            *ArrayOfSimpleDocument `xml:"DocumentList,omitempty"`
	TrackingNbr             string                 `xml:"TrackingNbr,omitempty"`
	Fees                    *Simple_FeeDetails     `xml:"Fees,omitempty"`
	LabelList               *SimpleLabelList       `xml:"LabelList,omitempty"`
	AlternateTrackingNumber string                 `xml:"AlternateTrackingNumber,omitempty"`
}

type ArrayOfSimpleDocument struct {
	SimpleDocuments []SimpleDocument `xml:"SimpleDocument,omitempty"`
}

type SimpleDocument struct {
	Type    string `xml:"Type,omitempty"`
	URL     string `xml:"URL,omitempty"`
	Content string `xml:"Content,omitempty"`
}

type Simple_FeeDetails struct {
	Fee *SimpleFeeDetails `xml:"Fee,omitempty"`
}

type SimpleFeeDetails struct {
	FeeDetails []SimpleFeeDetail `xml:"FeeDetail,omitempty"`
}

type SimpleFeeDetail struct {
	Name   string `xml:"Name,attr,ommitempty"`
	Charge string `xml:"Charge,attr,ommitempty"`
}

type SimpleLabelList struct {
	Label1 LabelImage `xml:"Label1,omitempty"`
	Label2 LabelImage `xml:"Label2,omitempty"`
	Label3 LabelImage `xml:"Label3,omitempty"`
	Label4 LabelImage `xml:"Label4,omitempty"`
	Label5 LabelImage `xml:"Label5,omitempty"`
}

type AbolSimpleRate struct {
	Request *AbolSimpleRateRequest `xml:"Request,omitempty"`
}

type AbolSimpleRateRequest struct {
	Authentication *Authentication `xml:"Authentication,omitempty"`
	SimpleRate     *SimpleShipment `xml:"SimpleRate,omitempty"`
}

type AbolRateSimpleResponse struct {
	SimpleRateResults []Simple_RateResult `xml:"SimpleRateResult,omitempty"`
}

type Simple_RateResult struct {
	Error                *error                `xml:"Error,omitempty"`
	SimpleRateResultList *SimpleRateResultList `xml:"SimpleRateResultList,omitempty"`
	Version              string                `xml:"Version,omitempty"`
}

type SimpleRateResultList struct {
	SimpleRateResults []SimpleRateResult `xml:"SimpleRateResult,omitempty"`
}

type SimpleRateResult struct {
	ReferenceNbr string            `xml:"ReferenceNbr,omitempty"`
	Error        *ArrayOfRateError `xml:"Error,omitempty"`
	Services     []SimpleService   `xml:"Service,omitempty"`
}

type SimpleService struct {
	Carrier           string      `xml:"Carrier,omitempty"`
	ServiceCode       string      `xml:"ServiceCode,omitempty"`
	ServiceDesc       string      `xml:"ServiceDesc,omitempty"`
	Rate              float64     `xml:"Rate,omitempty"`
	DeliveryDate      *time.Time  `xml:"DeliveryDate,omitempty"`
	DeliveryTime      string      `xml:"DeliveryTime,omitempty"`
	GuaranteedService bool        `xml:"GuaranteedService,omitempty"`
	Fees              *ArrayOfFee `xml:"Fees,omitempty"`
}

type ArrayOfFee struct {
	Fees []Fee `xml:"Fee,omitempty"`
}

type Fee struct {
	Name   string `xml:"Name,attr,ommitempty"`
	Charge string `xml:"Charge,attr,ommitempty"`
}

// Ports & bindings

type AbolapiSoap struct {
	address string
	client  *SOAPClient
}

func NewAbolapiSoap() *AbolapiSoap {
	addr := "http://testapi.iabol.com/api/abolapi.asmx"
	return &AbolapiSoap{
		address: addr,
		client:  newSOAPClient(addr, false, nil),
	}
}

// AbolRate - Perform rating with abolapi
func (x *AbolapiSoap) AbolRate(ctx context.Context, request AbolRateSoapIn) (*AbolRateSoapOut, error) {
	response := new(AbolRateSoapOut)
	action := "http://abolsoft.com/webservices/AbolRate"
	err := x.client.Call(ctx, action, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// AbolShipperSignup - Shipper Signup
func (x *AbolapiSoap) AbolShipperSignup(ctx context.Context, request AbolShipperSignupSoapIn) (*AbolShipperSignupSoapOut, error) {
	response := new(AbolShipperSignupSoapOut)
	action := "http://abolsoft.com/webservices/AbolShipperSignup"
	err := x.client.Call(ctx, action, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// AbolValidateAddress - Perform address validation with abolapi
func (x *AbolapiSoap) AbolValidateAddress(ctx context.Context, request AbolValidateAddressSoapIn) (*AbolValidateAddressSoapOut, error) {
	response := new(AbolValidateAddressSoapOut)
	action := "http://abolsoft.com/webservices/AbolValidateAddress"
	err := x.client.Call(ctx, action, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// AbolOrder - Perform order import with abolapi
func (x *AbolapiSoap) AbolOrder(ctx context.Context, request AbolOrderSoapIn) (*AbolOrderSoapOut, error) {
	response := new(AbolOrderSoapOut)
	action := "http://abolsoft.com/webservices/AbolOrder"
	err := x.client.Call(ctx, action, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// AbolShipment - Perform shipment with abolapi
func (x *AbolapiSoap) AbolShipment(ctx context.Context, request AbolShipmentSoapIn) (*AbolShipmentSoapOut, error) {
	response := new(AbolShipmentSoapOut)
	action := "http://abolsoft.com/webservices/AbolShipment"
	err := x.client.Call(ctx, action, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// AbolTrackPackage - Track package with abolapi
func (x *AbolapiSoap) AbolTrackPackage(ctx context.Context, request AbolTrackPackageSoapIn) (*AbolTrackPackageSoapOut, error) {
	response := new(AbolTrackPackageSoapOut)
	action := "http://abolsoft.com/webservices/AbolTrackPackage"
	err := x.client.Call(ctx, action, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// AbolVoidPackage - Void package with abolapi
func (x *AbolapiSoap) AbolVoidPackage(ctx context.Context, request AbolVoidPackageSoapIn) (*AbolVoidPackageSoapOut, error) {
	response := new(AbolVoidPackageSoapOut)
	action := "http://abolsoft.com/webservices/AbolVoidPackage"
	err := x.client.Call(ctx, action, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// AbolCloseOut - Perform CloseOut with abolapi
func (x *AbolapiSoap) AbolCloseOut(ctx context.Context, request AbolCloseOutSoapIn) (*AbolCloseOutSoapOut, error) {
	response := new(AbolCloseOutSoapOut)
	action := "http://abolsoft.com/webservices/AbolCloseOut"
	err := x.client.Call(ctx, action, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// AbolClassify - Get duty categories
func (x *AbolapiSoap) AbolClassify(ctx context.Context, request AbolClassifySoapIn) (*AbolClassifySoapOut, error) {
	response := new(AbolClassifySoapOut)
	action := "http://abolsoft.com/webservices/AbolClassify"
	err := x.client.Call(ctx, action, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// AbolDuty - Calculate duty
func (x *AbolapiSoap) AbolDuty(ctx context.Context, request AbolDutySoapIn) (*AbolDutySoapOut, error) {
	response := new(AbolDutySoapOut)
	action := "http://abolsoft.com/webservices/AbolDuty"
	err := x.client.Call(ctx, action, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// AbolSimpleShipment - Simple Abol Shipment
func (x *AbolapiSoap) AbolSimpleShipment(ctx context.Context, request AbolSimpleShipmentSoapIn) (*AbolSimpleShipmentSoapOut, error) {
	response := new(AbolSimpleShipmentSoapOut)
	action := "http://abolsoft.com/webservices/AbolSimpleShipment"
	err := x.client.Call(ctx, action, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// AbolSimpleRate - Simple Abol Rate
func (x *AbolapiSoap) AbolSimpleRate(ctx context.Context, request AbolSimpleRateSoapIn) (*AbolSimpleRateSoapOut, error) {
	response := new(AbolSimpleRateSoapOut)
	action := "http://abolsoft.com/webservices/AbolSimpleRate"
	err := x.client.Call(ctx, action, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// AbolOrderOM - Perform order import with abolapi
func (x *AbolapiSoap) AbolOrderOM(ctx context.Context, request AbolOrderOMSoapIn) (*AbolOrderOMSoapOut, error) {
	response := new(AbolOrderOMSoapOut)
	action := "http://abolsoft.com/webservices/AbolOrderOM"
	err := x.client.Call(ctx, action, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type AbolapiSoap12 struct {
	address string
	client  *SOAPClient
}

func NewAbolapiSoap12() *AbolapiSoap12 {
	addr := "http://testapi.iabol.com/api/abolapi.asmx"
	return &AbolapiSoap12{
		address: addr,
		client:  newSOAPClient(addr, false, nil),
	}
}

// AbolRate - Perform rating with abolapi
func (x *AbolapiSoap12) AbolRate(ctx context.Context, request AbolRateSoapIn) (*AbolRateSoapOut, error) {
	response := new(AbolRateSoapOut)
	action := "http://abolsoft.com/webservices/AbolRate"
	err := x.client.Call(ctx, action, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// AbolShipperSignup - Shipper Signup
func (x *AbolapiSoap12) AbolShipperSignup(ctx context.Context, request AbolShipperSignupSoapIn) (*AbolShipperSignupSoapOut, error) {
	response := new(AbolShipperSignupSoapOut)
	action := "http://abolsoft.com/webservices/AbolShipperSignup"
	err := x.client.Call(ctx, action, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// AbolValidateAddress - Perform address validation with abolapi
func (x *AbolapiSoap12) AbolValidateAddress(ctx context.Context, request AbolValidateAddressSoapIn) (*AbolValidateAddressSoapOut, error) {
	response := new(AbolValidateAddressSoapOut)
	action := "http://abolsoft.com/webservices/AbolValidateAddress"
	err := x.client.Call(ctx, action, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// AbolOrder - Perform order import with abolapi
func (x *AbolapiSoap12) AbolOrder(ctx context.Context, request AbolOrderSoapIn) (*AbolOrderSoapOut, error) {
	response := new(AbolOrderSoapOut)
	action := "http://abolsoft.com/webservices/AbolOrder"
	err := x.client.Call(ctx, action, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// AbolShipment - Perform shipment with abolapi
func (x *AbolapiSoap12) AbolShipment(ctx context.Context, request AbolShipmentSoapIn) (*AbolShipmentSoapOut, error) {
	response := new(AbolShipmentSoapOut)
	action := "http://abolsoft.com/webservices/AbolShipment"
	err := x.client.Call(ctx, action, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// AbolTrackPackage - Track package with abolapi
func (x *AbolapiSoap12) AbolTrackPackage(ctx context.Context, request AbolTrackPackageSoapIn) (*AbolTrackPackageSoapOut, error) {
	response := new(AbolTrackPackageSoapOut)
	action := "http://abolsoft.com/webservices/AbolTrackPackage"
	err := x.client.Call(ctx, action, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// AbolVoidPackage - Void package with abolapi
func (x *AbolapiSoap12) AbolVoidPackage(ctx context.Context, request AbolVoidPackageSoapIn) (*AbolVoidPackageSoapOut, error) {
	response := new(AbolVoidPackageSoapOut)
	action := "http://abolsoft.com/webservices/AbolVoidPackage"
	err := x.client.Call(ctx, action, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// AbolCloseOut - Perform CloseOut with abolapi
func (x *AbolapiSoap12) AbolCloseOut(ctx context.Context, request AbolCloseOutSoapIn) (*AbolCloseOutSoapOut, error) {
	response := new(AbolCloseOutSoapOut)
	action := "http://abolsoft.com/webservices/AbolCloseOut"
	err := x.client.Call(ctx, action, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// AbolClassify - Get duty categories
func (x *AbolapiSoap12) AbolClassify(ctx context.Context, request AbolClassifySoapIn) (*AbolClassifySoapOut, error) {
	response := new(AbolClassifySoapOut)
	action := "http://abolsoft.com/webservices/AbolClassify"
	err := x.client.Call(ctx, action, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// AbolDuty - Calculate duty
func (x *AbolapiSoap12) AbolDuty(ctx context.Context, request AbolDutySoapIn) (*AbolDutySoapOut, error) {
	response := new(AbolDutySoapOut)
	action := "http://abolsoft.com/webservices/AbolDuty"
	err := x.client.Call(ctx, action, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// AbolSimpleShipment - Simple Abol Shipment
func (x *AbolapiSoap12) AbolSimpleShipment(ctx context.Context, request *AbolSimpleShipmentSoapIn) (*AbolSimpleShipmentResponseElement, error) {
	// arse
	response := new(AbolSimpleShipmentResponseElement)
	action := "http://abolsoft.com/webservices/AbolSimpleShipment"
	err := x.client.Call(ctx, action, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// AbolSimpleRate - Simple Abol Rate
func (x *AbolapiSoap12) AbolSimpleRate(ctx context.Context, request AbolSimpleRateSoapIn) (*AbolSimpleRateSoapOut, error) {
	response := new(AbolSimpleRateSoapOut)
	action := "http://abolsoft.com/webservices/AbolSimpleRate"
	err := x.client.Call(ctx, action, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// AbolOrderOM - Perform order import with abolapi
func (x *AbolapiSoap12) AbolOrderOM(ctx context.Context, request AbolOrderOMSoapIn) (*AbolOrderOMSoapOut, error) {
	response := new(AbolOrderOMSoapOut)
	action := "http://abolsoft.com/webservices/AbolOrderOM"
	err := x.client.Call(ctx, action, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type AbolapiHttpGet struct {
	address string
	client  *SOAPClient
}

func NewAbolapiHttpGet() *AbolapiHttpGet {
	addr := "http://testapi.iabol.com/api/abolapi.asmx"
	return &AbolapiHttpGet{
		address: addr,
		client:  newSOAPClient(addr, false, nil),
	}
}

type AbolapiHttpPost struct {
	address string
	client  *SOAPClient
}

func NewAbolapiHttpPost() *AbolapiHttpPost {
	addr := "http://testapi.iabol.com/api/abolapi.asmx"
	return &AbolapiHttpPost{
		address: addr,
		client:  newSOAPClient(addr, false, nil),
	}
}
