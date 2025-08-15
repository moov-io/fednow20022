package FedNowBroadcast_admi_004_001_02

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

var fedNowBroadcastWrapper = common.MessageWrapper[MessageModel, any]{
	Config: common.MessageWrapperConfig[MessageModel, any]{
		PathMap:         PathMap,
		DocumentFactory: DocumentFactory,
		DataFactory:     DataFactory,
		RequireFields:   RequireFileds,
		BuildHelper:     func() any { return BuildMessageHelper() },
	},
}

type FedNowBroadcastWrapper struct{}

func (w *FedNowBroadcastWrapper) CreateDocument(modelJson []byte) ([]byte, error) {
	return fedNowBroadcastWrapper.CreateDocument(modelJson)
}
func (w *FedNowBroadcastWrapper) ValidateDocument(modelJson []byte) error {
	return fedNowBroadcastWrapper.ValidateDocument(modelJson)
}
func (w *FedNowBroadcastWrapper) CheckRequireField(modelJson []byte) error {
	return fedNowBroadcastWrapper.CheckRequireField(modelJson)
}
func (w *FedNowBroadcastWrapper) GetDataModel(xmlData []byte) (modelJson []byte, err error) {
	return fedNowBroadcastWrapper.GetDataModel(xmlData)
}
func (w *FedNowBroadcastWrapper) GetHelp() ([]byte, error) {
	return fedNowBroadcastWrapper.GetHelp()
}
