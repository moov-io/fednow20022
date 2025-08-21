package Message_InformationRequestResponse_camt_029_001_09

import (
	"encoding/xml"
	"fmt"

	InformationRequestResponse_camt_029_001_09_Model "github.com/moov-io/fednow20022/gen/InformationRequestResponse_camt_029_001_09"
	"github.com/moov-io/fednow20022/gen/Message_InformationRequestResponse_camt_029_001_09"
	"github.com/moov-io/fednow20022/pkg/fednow/models/InformationRequestResponse_camt_029_001_09"
	"github.com/moov-io/fednow20022/pkg/fednow/models/head_001_001_02"
)

var XMLAttributes = []xml.Attr{
	{Name: xml.Name{Local: "xmlns:head"}, Value: "urn:iso:std:iso:20022:tech:xsd:head.001.001.01"},
	{Name: xml.Name{Local: "xmlns:irr"}, Value: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.09"},
	{Name: xml.Name{Local: "xmlns"}, Value: "urn:fednow"},
}

var (
	HeaderPathMap         = head_001_001_02.PathMap
	HeaderDocumentFactory = head_001_001_02.DocumentFactory
	HeaderDataFactory     = head_001_001_02.DataFactory
	HeaderWrapper         = &head_001_001_02.HeadWrapper{}

	BodyPathMap         = InformationRequestResponse_camt_029_001_09.PathMap
	BodyDocumentFactory = InformationRequestResponse_camt_029_001_09.DocumentFactory
	BodyDataFactory     = InformationRequestResponse_camt_029_001_09.DataFactory
	BodyWrapper         = &InformationRequestResponse_camt_029_001_09.InformationRequestResponseWrapper{}
)

type (
	GeneratedMsgType       = Message_InformationRequestResponse_camt_029_001_09.Message
	GeneratedBodyType      = InformationRequestResponse_camt_029_001_09_Model.Document
	HeaderMessageModelType = *head_001_001_02.MessageModel
	BodyMessageModelType   = *InformationRequestResponse_camt_029_001_09.MessageModel
)

type MessageModel struct {
	AppHdr                     head_001_001_02.MessageModel                            `json:"app_hdr,omitempty"`
	InformationRequestResponse InformationRequestResponse_camt_029_001_09.MessageModel `json:"data,omitempty"`
}
type MessageHelper struct {
	AppHdr                     head_001_001_02.MessageHelper
	InformationRequestResponse InformationRequestResponse_camt_029_001_09.MessageHelper
}

func BuildHelper() MessageHelper {
	return MessageHelper{
		AppHdr:                     head_001_001_02.BuildMessageHelper(),
		InformationRequestResponse: InformationRequestResponse_camt_029_001_09.BuildMessageHelper(),
	}
}

func NewMessageModel(header any, data any) (MessageModel, error) {
	if docheaderStruct, ok := header.(*head_001_001_02.MessageModel); ok {
		if dataStruct, ok := data.(BodyMessageModelType); ok {
			return MessageModel{
				AppHdr:                     *docheaderStruct,
				InformationRequestResponse: *dataStruct,
			}, nil
		}
	}
	return MessageModel{}, fmt.Errorf("invalid header or data type")
}
