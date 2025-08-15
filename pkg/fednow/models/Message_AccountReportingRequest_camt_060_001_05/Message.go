package Message_AccountReportingRequest_camt_060_001_05

import (
	"encoding/xml"
	"fmt"

	AccountReportingRequest_camt_060_001_05_Model "github.com/moov-io/fednow20022/gen/AccountReportingRequest_camt_060_001_05"
	"github.com/moov-io/fednow20022/gen/Message_AccountReportingRequest_camt_060_001_05"
	"github.com/moov-io/fednow20022/pkg/fednow/models/AccountReportingRequest_camt_060_001_05"
	"github.com/moov-io/fednow20022/pkg/fednow/models/head_001_001_02"
)

var XMLAttributes = []xml.Attr{
	{Name: xml.Name{Local: "xmlns:head"}, Value: "urn:iso:std:iso:20022:tech:xsd:head.001.001.01"},
	{Name: xml.Name{Local: "xmlns:arr"}, Value: "urn:iso:std:iso:20022:tech:xsd:camt.060.001.05"},
	{Name: xml.Name{Local: "xmlns"}, Value: "urn:fednow"},
}

var (
	HeaderPathMap         = head_001_001_02.PathMap
	HeaderDocumentFactory = head_001_001_02.DocumentFactory
	HeaderDataFactory     = head_001_001_02.DataFactory
	HeaderWrapper         = &head_001_001_02.HeadWrapper{}

	BodyPathMap         = AccountReportingRequest_camt_060_001_05.PathMap
	BodyDocumentFactory = AccountReportingRequest_camt_060_001_05.DocumentFactory
	BodyDataFactory     = AccountReportingRequest_camt_060_001_05.DataFactory
	BodyWrapper         = &AccountReportingRequest_camt_060_001_05.AccountReportingRequestWrapper{}
)

type (
	GeneratedMsgType       = Message_AccountReportingRequest_camt_060_001_05.Message
	GeneratedBodyType      = AccountReportingRequest_camt_060_001_05_Model.Document
	HeaderMessageModelType = *head_001_001_02.MessageModel
	BodyMessageModelType   = *AccountReportingRequest_camt_060_001_05.MessageModel
)

type MessageModel struct {
	AppHdr                  head_001_001_02.MessageModel                         `json:"app_hdr,omitempty"`
	AccountReportingRequest AccountReportingRequest_camt_060_001_05.MessageModel `json:"data,omitempty"`
}
type MessageHelper struct {
	AppHdr                  head_001_001_02.MessageHelper
	AccountReportingRequest AccountReportingRequest_camt_060_001_05.MessageHelper
}

func BuildHelper() MessageHelper {
	return MessageHelper{
		AppHdr:                  head_001_001_02.BuildMessageHelper(),
		AccountReportingRequest: AccountReportingRequest_camt_060_001_05.BuildMessageHelper(),
	}
}

func NewMessageModel(header any, data any) (MessageModel, error) {
	if docheaderStruct, ok := header.(*head_001_001_02.MessageModel); ok {
		if dataStruct, ok := data.(BodyMessageModelType); ok {
			return MessageModel{
				AppHdr:                  *docheaderStruct,
				AccountReportingRequest: *dataStruct,
			}, nil
		}
	}
	return MessageModel{}, fmt.Errorf("invalid header or data type")
}
