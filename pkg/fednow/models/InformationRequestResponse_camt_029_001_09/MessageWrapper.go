package InformationRequestResponse_camt_029_001_09

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

var informationRequestResponseWrapper = common.MessageWrapper[MessageModel, any]{
	Config: common.MessageWrapperConfig[MessageModel, any]{
		PathMap:         PathMap,
		DocumentFactory: DocumentFactory,
		DataFactory:     DataFactory,
		RequireFields:   RequireFileds,
		BuildHelper:     func() any { return BuildMessageHelper() },
	},
}

type InformationRequestResponseWrapper struct{}

func (w *InformationRequestResponseWrapper) CreateDocument(modelJson []byte) ([]byte, error) {
	return informationRequestResponseWrapper.CreateDocument(modelJson)
}
func (w *InformationRequestResponseWrapper) ValidateDocument(modelJson []byte) error {
	return informationRequestResponseWrapper.ValidateDocument(modelJson)
}
func (w *InformationRequestResponseWrapper) CheckRequireField(modelJson []byte) error {
	return informationRequestResponseWrapper.CheckRequireField(modelJson)
}
func (w *InformationRequestResponseWrapper) GetDataModel(xmlData []byte) (modelJson []byte, err error) {
	return informationRequestResponseWrapper.GetDataModel(xmlData)
}
func (w *InformationRequestResponseWrapper) GetHelp() ([]byte, error) {
	return informationRequestResponseWrapper.GetHelp()
}
