package ReturnRequestResponse_camt_029_001_09

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

var returnRequestResponseWrapper = common.MessageWrapper[MessageModel, any]{
	Config: common.MessageWrapperConfig[MessageModel, any]{
		PathMap:         PathMap,
		DocumentFactory: DocumentFactory,
		DataFactory:     DataFactory,
		RequireFields:   RequireFileds,
		BuildHelper:     func() any { return BuildMessageHelper() },
	},
}

type ReturnRequestResponseWrapper struct{}

func (w *ReturnRequestResponseWrapper) CreateDocument(modelJson []byte) ([]byte, error) {
	return returnRequestResponseWrapper.CreateDocument(modelJson)
}
func (w *ReturnRequestResponseWrapper) ValidateDocument(modelJson []byte) error {
	return returnRequestResponseWrapper.ValidateDocument(modelJson)
}
func (w *ReturnRequestResponseWrapper) CheckRequireField(modelJson []byte) error {
	return returnRequestResponseWrapper.CheckRequireField(modelJson)
}
func (w *ReturnRequestResponseWrapper) GetDataModel(xmlData []byte) (modelJson []byte, err error) {
	return returnRequestResponseWrapper.GetDataModel(xmlData)
}
func (w *ReturnRequestResponseWrapper) GetHelp() ([]byte, error) {
	return returnRequestResponseWrapper.GetHelp()
}
