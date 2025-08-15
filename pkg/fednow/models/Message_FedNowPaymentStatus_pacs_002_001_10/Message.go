package Message_FedNowPaymentStatus_pacs_002_001_10

import (
	"encoding/xml"
	"fmt"

	FedNowPaymentStatus_pacs_002_001_10_Model "github.com/moov-io/fednow20022/gen/FedNowPaymentStatus_pacs_002_001_10"
	"github.com/moov-io/fednow20022/gen/Message_FedNowPaymentStatus_pacs_002_001_10"
	"github.com/moov-io/fednow20022/pkg/fednow/models/FedNowPaymentStatus_pacs_002_001_10"
	"github.com/moov-io/fednow20022/pkg/fednow/models/head_001_001_02"
)

var XMLAttributes = []xml.Attr{
	{Name: xml.Name{Local: "xmlns:head"}, Value: "urn:iso:std:iso:20022:tech:xsd:head.001.001.01"},
	{Name: xml.Name{Local: "xmlns:fnps"}, Value: "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10"},
	{Name: xml.Name{Local: "xmlns"}, Value: "urn:fednow"},
}

var (
	HeaderPathMap         = head_001_001_02.PathMap
	HeaderDocumentFactory = head_001_001_02.DocumentFactory
	HeaderDataFactory     = head_001_001_02.DataFactory
	HeaderWrapper         = &head_001_001_02.HeadWrapper{}

	BodyPathMap         = FedNowPaymentStatus_pacs_002_001_10.PathMap
	BodyDocumentFactory = FedNowPaymentStatus_pacs_002_001_10.DocumentFactory
	BodyDataFactory     = FedNowPaymentStatus_pacs_002_001_10.DataFactory
	BodyWrapper         = &FedNowPaymentStatus_pacs_002_001_10.FedNowPaymentStatusWrapper{}
)

type (
	GeneratedMsgType       = Message_FedNowPaymentStatus_pacs_002_001_10.Message
	GeneratedBodyType      = FedNowPaymentStatus_pacs_002_001_10_Model.Document
	HeaderMessageModelType = *head_001_001_02.MessageModel
	BodyMessageModelType   = *FedNowPaymentStatus_pacs_002_001_10.MessageModel
)

type MessageModel struct {
	AppHdr              head_001_001_02.MessageModel                     `json:"app_hdr,omitempty"`
	FedNowPaymentStatus FedNowPaymentStatus_pacs_002_001_10.MessageModel `json:"data,omitempty"`
}
type MessageHelper struct {
	AppHdr              head_001_001_02.MessageHelper
	FedNowPaymentStatus FedNowPaymentStatus_pacs_002_001_10.MessageHelper
}

func BuildHelper() MessageHelper {
	return MessageHelper{
		AppHdr:              head_001_001_02.BuildMessageHelper(),
		FedNowPaymentStatus: FedNowPaymentStatus_pacs_002_001_10.BuildMessageHelper(),
	}
}

func NewMessageModel(header any, data any) (MessageModel, error) {
	if docheaderStruct, ok := header.(*head_001_001_02.MessageModel); ok {
		if dataStruct, ok := data.(BodyMessageModelType); ok {
			return MessageModel{
				AppHdr:              *docheaderStruct,
				FedNowPaymentStatus: *dataStruct,
			}, nil
		}
	}
	return MessageModel{}, fmt.Errorf("invalid header or data type")
}
