package Message_InformationRequest_camt_026_001_07

import (
	"encoding/xml"
	"fmt"

	InformationRequest_camt_026_001_07_Model "github.com/moov-io/fednow20022/gen/InformationRequest_camt_026_001_07"
	"github.com/moov-io/fednow20022/gen/Message_InformationRequest_camt_026_001_07"
	"github.com/moov-io/fednow20022/pkg/fednow/models/InformationRequest_camt_026_001_07"
	"github.com/moov-io/fednow20022/pkg/fednow/models/head_001_001_02"
)

var XMLAttributes = []xml.Attr{
	{Name: xml.Name{Local: "xmlns:head"}, Value: "urn:iso:std:iso:20022:tech:xsd:head.001.001.01"},
	{Name: xml.Name{Local: "xmlns:ir"}, Value: "urn:iso:std:iso:20022:tech:xsd:camt.026.001.07"},
	{Name: xml.Name{Local: "xmlns"}, Value: "urn:fednow"},
}

var (
	HeaderPathMap         = head_001_001_02.PathMap
	HeaderDocumentFactory = head_001_001_02.DocumentFactory
	HeaderDataFactory     = head_001_001_02.DataFactory
	HeaderWrapper         = &head_001_001_02.HeadWrapper{}

	BodyPathMap         = InformationRequest_camt_026_001_07.PathMap
	BodyDocumentFactory = InformationRequest_camt_026_001_07.DocumentFactory
	BodyDataFactory     = InformationRequest_camt_026_001_07.DataFactory
	BodyWrapper         = &InformationRequest_camt_026_001_07.InformationRequestWrapper{}
)

type (
	GeneratedMsgType       = Message_InformationRequest_camt_026_001_07.Message
	GeneratedBodyType      = InformationRequest_camt_026_001_07_Model.Document
	HeaderMessageModelType = *head_001_001_02.MessageModel
	BodyMessageModelType   = *InformationRequest_camt_026_001_07.MessageModel
)

type MessageModel struct {
	AppHdr             head_001_001_02.MessageModel                    `json:"app_hdr,omitempty"`
	InformationRequest InformationRequest_camt_026_001_07.MessageModel `json:"data,omitempty"`
}
type MessageHelper struct {
	AppHdr             head_001_001_02.MessageHelper
	InformationRequest InformationRequest_camt_026_001_07.MessageHelper
}

func BuildHelper() MessageHelper {
	return MessageHelper{
		AppHdr:             head_001_001_02.BuildMessageHelper(),
		InformationRequest: InformationRequest_camt_026_001_07.BuildMessageHelper(),
	}
}

func NewMessageModel(header any, data any) (MessageModel, error) {
	if docheaderStruct, ok := header.(*head_001_001_02.MessageModel); ok {
		if dataStruct, ok := data.(BodyMessageModelType); ok {
			return MessageModel{
				AppHdr:             *docheaderStruct,
				InformationRequest: *dataStruct,
			}, nil
		}
	}
	return MessageModel{}, fmt.Errorf("invalid header or data type")
}
