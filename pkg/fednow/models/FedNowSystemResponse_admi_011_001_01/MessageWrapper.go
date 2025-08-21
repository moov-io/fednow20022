package FedNowSystemResponse_admi_011_001_01

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

var fedNowSystemResponseWrapper = common.MessageWrapper[MessageModel, any]{
	Config: common.MessageWrapperConfig[MessageModel, any]{
		PathMap:         PathMap,
		DocumentFactory: DocumentFactory,
		DataFactory:     DataFactory,
		RequireFields:   RequireFileds,
		BuildHelper:     func() any { return BuildMessageHelper() },
	},
}

type FedNowSystemResponseWrapper struct{}

func (w *FedNowSystemResponseWrapper) CreateDocument(modelJson []byte) ([]byte, error) {
	return fedNowSystemResponseWrapper.CreateDocument(modelJson)
}
func (w *FedNowSystemResponseWrapper) ValidateDocument(modelJson []byte) error {
	return fedNowSystemResponseWrapper.ValidateDocument(modelJson)
}
func (w *FedNowSystemResponseWrapper) CheckRequireField(modelJson []byte) error {
	return fedNowSystemResponseWrapper.CheckRequireField(modelJson)
}
func (w *FedNowSystemResponseWrapper) GetDataModel(xmlData []byte) (modelJson []byte, err error) {
	return fedNowSystemResponseWrapper.GetDataModel(xmlData)
}
func (w *FedNowSystemResponseWrapper) GetHelp() ([]byte, error) {
	return fedNowSystemResponseWrapper.GetHelp()
}
