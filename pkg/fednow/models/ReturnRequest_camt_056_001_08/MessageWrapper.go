package ReturnRequest_camt_056_001_08

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

var returnRequestWrapper = common.MessageWrapper[MessageModel, any]{
	Config: common.MessageWrapperConfig[MessageModel, any]{
		PathMap:         PathMap,
		DocumentFactory: DocumentFactory,
		DataFactory:     DataFactory,
		RequireFields:   RequireFileds,
		BuildHelper:     func() any { return BuildMessageHelper() },
	},
}

type ReturnRequestWrapper struct{}

func (w *ReturnRequestWrapper) CreateDocument(modelJson []byte) ([]byte, error) {
	return returnRequestWrapper.CreateDocument(modelJson)
}
func (w *ReturnRequestWrapper) ValidateDocument(modelJson []byte) error {
	return returnRequestWrapper.ValidateDocument(modelJson)
}
func (w *ReturnRequestWrapper) CheckRequireField(modelJson []byte) error {
	return returnRequestWrapper.CheckRequireField(modelJson)
}
func (w *ReturnRequestWrapper) GetDataModel(xmlData []byte) (modelJson []byte, err error) {
	return returnRequestWrapper.GetDataModel(xmlData)
}
func (w *ReturnRequestWrapper) GetHelp() ([]byte, error) {
	return returnRequestWrapper.GetHelp()
}
