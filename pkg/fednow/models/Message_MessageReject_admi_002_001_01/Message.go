package Message_MessageReject_admi_002_001_01

import (
	"encoding/xml"
	"fmt"

	MessageReject_admi_002_001_01_Model "github.com/moov-io/fednow20022/gen/MessageReject_admi_002_001_01"
	"github.com/moov-io/fednow20022/gen/Message_MessageReject_admi_002_001_01"
	"github.com/moov-io/fednow20022/pkg/fednow/models/MessageReject_admi_002_001_01"
	"github.com/moov-io/fednow20022/pkg/fednow/models/head_001_001_02"
)

var XMLAttributes = []xml.Attr{
	{Name: xml.Name{Local: "xmlns:head"}, Value: "urn:iso:std:iso:20022:tech:xsd:head.001.001.01"},
	{Name: xml.Name{Local: "xmlns:mr"}, Value: "urn:iso:std:iso:20022:tech:xsd:admi.002.001.01"},
	{Name: xml.Name{Local: "xmlns"}, Value: "urn:fednow"},
}

var (
	HeaderPathMap         = head_001_001_02.PathMap
	HeaderDocumentFactory = head_001_001_02.DocumentFactory
	HeaderDataFactory     = head_001_001_02.DataFactory
	HeaderWrapper         = &head_001_001_02.HeadWrapper{}

	BodyPathMap         = MessageReject_admi_002_001_01.PathMap
	BodyDocumentFactory = MessageReject_admi_002_001_01.DocumentFactory
	BodyDataFactory     = MessageReject_admi_002_001_01.DataFactory
	BodyWrapper         = &MessageReject_admi_002_001_01.MessageRejectWrapper{}
)

type (
	GeneratedMsgType       = Message_MessageReject_admi_002_001_01.Message
	GeneratedBodyType      = MessageReject_admi_002_001_01_Model.Document
	HeaderMessageModelType = *head_001_001_02.MessageModel
	BodyMessageModelType   = *MessageReject_admi_002_001_01.MessageModel
)

type MessageModel struct {
	AppHdr        head_001_001_02.MessageModel               `json:"app_hdr,omitempty"`
	MessageReject MessageReject_admi_002_001_01.MessageModel `json:"data,omitempty"`
}
type MessageHelper struct {
	AppHdr        head_001_001_02.MessageHelper
	MessageReject MessageReject_admi_002_001_01.MessageHelper
}

func BuildHelper() MessageHelper {
	return MessageHelper{
		AppHdr:        head_001_001_02.BuildMessageHelper(),
		MessageReject: MessageReject_admi_002_001_01.BuildMessageHelper(),
	}
}

func NewMessageModel(header any, data any) (MessageModel, error) {
	if docheaderStruct, ok := header.(*head_001_001_02.MessageModel); ok {
		if dataStruct, ok := data.(BodyMessageModelType); ok {
			return MessageModel{
				AppHdr:        *docheaderStruct,
				MessageReject: *dataStruct,
			}, nil
		}
	}
	return MessageModel{}, fmt.Errorf("invalid header or data type")
}
