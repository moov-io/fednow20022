package Message_AccountBalanceReport_camt_052_001_08

import (
	"encoding/xml"
	"fmt"

	AccountBalanceReport_camt_052_001_08_Model "github.com/moov-io/fednow20022/gen/AccountBalanceReport_camt_052_001_08"
	"github.com/moov-io/fednow20022/gen/Message_AccountBalanceReport_camt_052_001_08"
	"github.com/moov-io/fednow20022/pkg/fednow/models/AccountBalanceReport_camt_052_001_08"
	"github.com/moov-io/fednow20022/pkg/fednow/models/head_001_001_02"
)

var XMLAttributes = []xml.Attr{
	{Name: xml.Name{Local: "xmlns:head"}, Value: "urn:iso:std:iso:20022:tech:xsd:head.001.001.02"},
	{Name: xml.Name{Local: "xmlns:aadr"}, Value: "urn:iso:std:iso:20022:tech:xsd:camt.052.001.08"},
	{Name: xml.Name{Local: "xmlns"}, Value: "urn:fednow"},
}

var (
	HeaderPathMap         = head_001_001_02.PathMap
	HeaderDocumentFactory = head_001_001_02.DocumentFactory
	HeaderDataFactory     = head_001_001_02.DataFactory
	HeaderWrapper         = &head_001_001_02.HeadWrapper{}

	BodyPathMap         = AccountBalanceReport_camt_052_001_08.PathMap
	BodyDocumentFactory = AccountBalanceReport_camt_052_001_08.DocumentFactory
	BodyDataFactory     = AccountBalanceReport_camt_052_001_08.DataFactory
	BodyWrapper         = &AccountBalanceReport_camt_052_001_08.AccountBalanceReportReportWrapper{}
)

type (
	GeneratedMsgType       = Message_AccountBalanceReport_camt_052_001_08.Message
	GeneratedBodyType      = AccountBalanceReport_camt_052_001_08_Model.Document
	HeaderMessageModelType = *head_001_001_02.MessageModel
	BodyMessageModelType   = *AccountBalanceReport_camt_052_001_08.MessageModel
)

type MessageModel struct {
	AppHdr               head_001_001_02.MessageModel                      `json:"app_hdr,omitempty"`
	AccountBalanceReport AccountBalanceReport_camt_052_001_08.MessageModel `json:"data,omitempty"`
}
type MessageHelper struct {
	AppHdr               head_001_001_02.MessageHelper
	AccountBalanceReport AccountBalanceReport_camt_052_001_08.MessageHelper
}

func BuildHelper() MessageHelper {
	return MessageHelper{
		AppHdr:               head_001_001_02.BuildMessageHelper(),
		AccountBalanceReport: AccountBalanceReport_camt_052_001_08.BuildMessageHelper(),
	}
}

func NewMessageModel(header any, data any) (MessageModel, error) {
	if docheaderStruct, ok := header.(*head_001_001_02.MessageModel); ok {
		if dataStruct, ok := data.(BodyMessageModelType); ok {
			return MessageModel{
				AppHdr:               *docheaderStruct,
				AccountBalanceReport: *dataStruct,
			}, nil
		}
	}
	return MessageModel{}, fmt.Errorf("invalid header or data type")
}
