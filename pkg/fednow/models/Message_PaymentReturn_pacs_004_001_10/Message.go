package Message_PaymentReturn_pacs_004_001_10

import (
	"encoding/xml"
	"fmt"

	"github.com/moov-io/fednow20022/gen/Message_PaymentReturn_pacs_004_001_10"
	PaymentReturn_pacs_004_001_10_Model "github.com/moov-io/fednow20022/gen/PaymentReturn_pacs_004_001_10"
	"github.com/moov-io/fednow20022/pkg/fednow/models/PaymentReturn_pacs_004_001_10"
	"github.com/moov-io/fednow20022/pkg/fednow/models/head_001_001_02"
)

var XMLAttributes = []xml.Attr{
	{Name: xml.Name{Local: "xmlns:head"}, Value: "urn:iso:std:iso:20022:tech:xsd:head.001.001.01"},
	{Name: xml.Name{Local: "xmlns:pr"}, Value: "urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10"},
	{Name: xml.Name{Local: "xmlns"}, Value: "urn:fednow"},
}

var (
	HeaderPathMap         = head_001_001_02.PathMap
	HeaderDocumentFactory = head_001_001_02.DocumentFactory
	HeaderDataFactory     = head_001_001_02.DataFactory
	HeaderWrapper         = &head_001_001_02.HeadWrapper{}

	BodyPathMap         = PaymentReturn_pacs_004_001_10.PathMap
	BodyDocumentFactory = PaymentReturn_pacs_004_001_10.DocumentFactory
	BodyDataFactory     = PaymentReturn_pacs_004_001_10.DataFactory
	BodyWrapper         = &PaymentReturn_pacs_004_001_10.PaymentReturnWrapper{}
)

type (
	GeneratedMsgType       = Message_PaymentReturn_pacs_004_001_10.Message
	GeneratedBodyType      = PaymentReturn_pacs_004_001_10_Model.Document
	HeaderMessageModelType = *head_001_001_02.MessageModel
	BodyMessageModelType   = *PaymentReturn_pacs_004_001_10.MessageModel
)

type MessageModel struct {
	AppHdr        head_001_001_02.MessageModel               `json:"app_hdr,omitempty"`
	PaymentReturn PaymentReturn_pacs_004_001_10.MessageModel `json:"data,omitempty"`
}
type MessageHelper struct {
	AppHdr        head_001_001_02.MessageHelper
	PaymentReturn PaymentReturn_pacs_004_001_10.MessageHelper
}

func BuildHelper() MessageHelper {
	return MessageHelper{
		AppHdr:        head_001_001_02.BuildMessageHelper(),
		PaymentReturn: PaymentReturn_pacs_004_001_10.BuildMessageHelper(),
	}
}

func NewMessageModel(header any, data any) (MessageModel, error) {
	if docheaderStruct, ok := header.(*head_001_001_02.MessageModel); ok {
		if dataStruct, ok := data.(BodyMessageModelType); ok {
			return MessageModel{
				AppHdr:        *docheaderStruct,
				PaymentReturn: *dataStruct,
			}, nil
		}
	}
	return MessageModel{}, fmt.Errorf("invalid header or data type")
}
