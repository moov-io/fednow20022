package Message_ParticipantBroadcast_admi_004_001_02

import (
	"encoding/xml"
	"fmt"

	"github.com/moov-io/fednow20022/gen/Message_ParticipantBroadcast_admi_004_001_02"
	ParticipantBroadcast_admi_004_001_02_Model "github.com/moov-io/fednow20022/gen/ParticipantBroadcast_admi_004_001_02"
	"github.com/moov-io/fednow20022/pkg/fednow/models/ParticipantBroadcast_admi_004_001_02"
	"github.com/moov-io/fednow20022/pkg/fednow/models/head_001_001_02"
)

var XMLAttributes = []xml.Attr{
	{Name: xml.Name{Local: "xmlns:head"}, Value: "urn:iso:std:iso:20022:tech:xsd:head.001.001.01"},
	{Name: xml.Name{Local: "xmlns:pb"}, Value: "urn:iso:std:iso:20022:tech:xsd:admi.004.001.02"},
	{Name: xml.Name{Local: "xmlns"}, Value: "urn:fednow"},
}

var (
	HeaderPathMap         = head_001_001_02.PathMap
	HeaderDocumentFactory = head_001_001_02.DocumentFactory
	HeaderDataFactory     = head_001_001_02.DataFactory
	HeaderWrapper         = &head_001_001_02.HeadWrapper{}

	BodyPathMap         = ParticipantBroadcast_admi_004_001_02.PathMap
	BodyDocumentFactory = ParticipantBroadcast_admi_004_001_02.DocumentFactory
	BodyDataFactory     = ParticipantBroadcast_admi_004_001_02.DataFactory
	BodyWrapper         = &ParticipantBroadcast_admi_004_001_02.ParticipantBroadcastWrapper{}
)

type (
	GeneratedMsgType       = Message_ParticipantBroadcast_admi_004_001_02.Message
	GeneratedBodyType      = ParticipantBroadcast_admi_004_001_02_Model.Document
	HeaderMessageModelType = *head_001_001_02.MessageModel
	BodyMessageModelType   = *ParticipantBroadcast_admi_004_001_02.MessageModel
)

type MessageModel struct {
	AppHdr               head_001_001_02.MessageModel                      `json:"app_hdr,omitempty"`
	ParticipantBroadcast ParticipantBroadcast_admi_004_001_02.MessageModel `json:"data,omitempty"`
}
type MessageHelper struct {
	AppHdr               head_001_001_02.MessageHelper
	ParticipantBroadcast ParticipantBroadcast_admi_004_001_02.MessageHelper
}

func BuildHelper() MessageHelper {
	return MessageHelper{
		AppHdr:               head_001_001_02.BuildMessageHelper(),
		ParticipantBroadcast: ParticipantBroadcast_admi_004_001_02.BuildMessageHelper(),
	}
}

func NewMessageModel(header any, data any) (MessageModel, error) {
	if docheaderStruct, ok := header.(*head_001_001_02.MessageModel); ok {
		if dataStruct, ok := data.(BodyMessageModelType); ok {
			return MessageModel{
				AppHdr:               *docheaderStruct,
				ParticipantBroadcast: *dataStruct,
			}, nil
		}
	}
	return MessageModel{}, fmt.Errorf("invalid header or data type")
}
