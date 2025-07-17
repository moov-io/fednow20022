package Message_AccountActivityDetailsReport_camt_052_001_08

// import (
// 	"encoding/xml"
// 	"fmt"

// 	AccountActivityDetailsReport_camt_052_001_08_Model "github.com/moov-io/fednow20022/gen/AccountActivityDetailsReport_camt_052_001_08"
// 	"github.com/moov-io/fednow20022/gen/Message_AccountActivityDetailsReport_camt_052_001_08"
// 	"github.com/moov-io/fednow20022/pkg/fednow/models/AccountActivityDetailsReport_camt_052_001_08"
// 	"github.com/moov-io/fednow20022/pkg/fednow/models/head_001_001_02"
// )

// var XMLAttributes = []xml.Attr{
// 	{Name: xml.Name{Local: "xmlns:head"}, Value: "urn:iso:std:iso:20022:tech:xsd:head.001.001.02"},
// 	{Name: xml.Name{Local: "xmlns:aadr"}, Value: "urn:iso:std:iso:20022:tech:xsd:camt.052.001.08"},
// 	{Name: xml.Name{Local: "xmlns"}, Value: "urn:fednow"},
// }

// var (
// 	HeaderPathMap         = head_001_001_02.PathMap
// 	HeaderDocumentFactory = head_001_001_02.DocumentFactory
// 	HeaderDataFactory     = head_001_001_02.DataFactory
// 	HeaderWrapper         = &head_001_001_02.HeadWrapper{}

// 	BodyPathMap         = AccountActivityDetailsReport_camt_052_001_08.PathMap
// 	BodyDocumentFactory = AccountActivityDetailsReport_camt_052_001_08.DocumentFactory
// 	BodyDataFactory     = AccountActivityDetailsReport_camt_052_001_08.DataFactory
// 	BodyWrapper         = &AccountActivityDetailsReport_camt_052_001_08.AccountActivityDetailsReportWrapper{}
// )

// type (
// 	GeneratedMsgType       = Message_AccountActivityDetailsReport_camt_052_001_08.Message
// 	GeneratedBodyType      = AccountActivityDetailsReport_camt_052_001_08_Model.Document
// 	HeaderMessageModelType = *head_001_001_02.MessageModel
// 	BodyMessageModelType   = *AccountActivityDetailsReport_camt_052_001_08.MessageModel
// )

// type MessageModel struct {
// 	AppHdr                       head_001_001_02.MessageModel                              `json:"app_hdr,omitempty"`
// 	AccountActivityDetailsReport AccountActivityDetailsReport_camt_052_001_08.MessageModel `json:"data,omitempty"`
// }
// type MessageHelper struct {
// 	AppHdr                       head_001_001_02.MessageHelper
// 	AccountActivityDetailsReport AccountActivityDetailsReport_camt_052_001_08.MessageHelper
// }

// func BuildHelper() MessageHelper {
// 	return MessageHelper{
// 		AppHdr:                       head_001_001_02.BuildMessageHelper(),
// 		AccountActivityDetailsReport: AccountActivityDetailsReport_camt_052_001_08.BuildMessageHelper(),
// 	}
// }

// func NewMessageModel(header any, data any) (MessageModel, error) {
// 	if docheaderStruct, ok := header.(*head_001_001_02.MessageModel); ok {
// 		if dataStruct, ok := data.(BodyMessageModelType); ok {
// 			return MessageModel{
// 				AppHdr:                       *docheaderStruct,
// 				AccountActivityDetailsReport: *dataStruct,
// 			}, nil
// 		}
// 	}
// 	return MessageModel{}, fmt.Errorf("invalid header or data type")
// }
