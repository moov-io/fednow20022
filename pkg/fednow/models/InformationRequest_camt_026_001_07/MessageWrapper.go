package InformationRequest_camt_026_001_07

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

var informationRequestWrapper = common.MessageWrapper[MessageModel, any]{
	Config: common.MessageWrapperConfig[MessageModel, any]{
		PathMap:         PathMap,
		DocumentFactory: DocumentFactory,
		DataFactory:     DataFactory,
		RequireFields:   RequireFileds,
		BuildHelper:     func() any { return BuildMessageHelper() },
	},
}

type InformationRequestWrapper struct{}

func (w *InformationRequestWrapper) CreateDocument(modelJson []byte) ([]byte, error) {
	return informationRequestWrapper.CreateDocument(modelJson)
}
func (w *InformationRequestWrapper) ValidateDocument(modelJson []byte) error {
	return informationRequestWrapper.ValidateDocument(modelJson)
}
func (w *InformationRequestWrapper) CheckRequireField(modelJson []byte) error {
	return informationRequestWrapper.CheckRequireField(modelJson)
}
func (w *InformationRequestWrapper) GetDataModel(xmlData []byte) (modelJson []byte, err error) {
	return informationRequestWrapper.GetDataModel(xmlData)
}
func (w *InformationRequestWrapper) GetHelp() ([]byte, error) {
	return informationRequestWrapper.GetHelp()
}
