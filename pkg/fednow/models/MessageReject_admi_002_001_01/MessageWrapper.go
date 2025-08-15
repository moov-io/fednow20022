package MessageReject_admi_002_001_01

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

var messageRejectWrapper = common.MessageWrapper[MessageModel, any]{
	Config: common.MessageWrapperConfig[MessageModel, any]{
		PathMap:         PathMap,
		DocumentFactory: DocumentFactory,
		DataFactory:     DataFactory,
		RequireFields:   RequireFileds,
		BuildHelper:     func() any { return BuildMessageHelper() },
	},
}

type MessageRejectWrapper struct{}

func (w *MessageRejectWrapper) CreateDocument(modelJson []byte) ([]byte, error) {
	return messageRejectWrapper.CreateDocument(modelJson)
}
func (w *MessageRejectWrapper) ValidateDocument(modelJson []byte) error {
	return messageRejectWrapper.ValidateDocument(modelJson)
}
func (w *MessageRejectWrapper) CheckRequireField(modelJson []byte) error {
	return messageRejectWrapper.CheckRequireField(modelJson)
}
func (w *MessageRejectWrapper) GetDataModel(xmlData []byte) (modelJson []byte, err error) {
	return messageRejectWrapper.GetDataModel(xmlData)
}
func (w *MessageRejectWrapper) GetHelp() ([]byte, error) {
	return messageRejectWrapper.GetHelp()
}
