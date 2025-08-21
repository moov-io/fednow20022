package Message_FedNowBroadcast_admi_004_001_02

import (
	"encoding/xml"
	"fmt"

	FedNowBroadcast_admi_004_001_02_Model "github.com/moov-io/fednow20022/gen/FedNowBroadcast_admi_004_001_02"
	"github.com/moov-io/fednow20022/gen/Message_FedNowBroadcast_admi_004_001_02"
	"github.com/moov-io/fednow20022/pkg/fednow/models/FedNowBroadcast_admi_004_001_02"
	"github.com/moov-io/fednow20022/pkg/fednow/models/head_001_001_02"
)

var XMLAttributes = []xml.Attr{
	{Name: xml.Name{Local: "xmlns:head"}, Value: "urn:iso:std:iso:20022:tech:xsd:head.001.001.02"},
	{Name: xml.Name{Local: "xmlns:fnb"}, Value: "urn:iso:std:iso:20022:tech:xsd:admi.004.001.02"},
	{Name: xml.Name{Local: "xmlns"}, Value: "urn:fednow"},
}

var (
	HeaderPathMap         = head_001_001_02.PathMap
	HeaderDocumentFactory = head_001_001_02.DocumentFactory
	HeaderDataFactory     = head_001_001_02.DataFactory
	HeaderWrapper         = &head_001_001_02.HeadWrapper{}

	BodyPathMap         = FedNowBroadcast_admi_004_001_02.PathMap
	BodyDocumentFactory = FedNowBroadcast_admi_004_001_02.DocumentFactory
	BodyDataFactory     = FedNowBroadcast_admi_004_001_02.DataFactory
	BodyWrapper         = &FedNowBroadcast_admi_004_001_02.FedNowBroadcastWrapper{}
)

type (
	GeneratedMsgType       = Message_FedNowBroadcast_admi_004_001_02.Message
	GeneratedBodyType      = FedNowBroadcast_admi_004_001_02_Model.Document
	HeaderMessageModelType = *head_001_001_02.MessageModel
	BodyMessageModelType   = *FedNowBroadcast_admi_004_001_02.MessageModel
)

type MessageModel struct {
	AppHdr          head_001_001_02.MessageModel                 `json:"app_hdr,omitempty"`
	FedNowBroadcast FedNowBroadcast_admi_004_001_02.MessageModel `json:"data,omitempty"`
}
type MessageHelper struct {
	AppHdr          head_001_001_02.MessageHelper
	FedNowBroadcast FedNowBroadcast_admi_004_001_02.MessageHelper
}

func BuildHelper() MessageHelper {
	return MessageHelper{
		AppHdr:          head_001_001_02.BuildMessageHelper(),
		FedNowBroadcast: FedNowBroadcast_admi_004_001_02.BuildMessageHelper(),
	}
}

func NewMessageModel(header any, data any) (MessageModel, error) {
	if docheaderStruct, ok := header.(*head_001_001_02.MessageModel); ok {
		if dataStruct, ok := data.(BodyMessageModelType); ok {
			return MessageModel{
				AppHdr:          *docheaderStruct,
				FedNowBroadcast: *dataStruct,
			}, nil
		}
	}
	return MessageModel{}, fmt.Errorf("invalid header or data type")
}
