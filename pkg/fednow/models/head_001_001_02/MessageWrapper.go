package head_001_001_02

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

var headWrapper = common.MessageWrapper[MessageModel, any]{
	Config: common.MessageWrapperConfig[MessageModel, any]{
		PathMap:         PathMap,
		DocumentFactory: DocumentFactory,
		DataFactory:     DataFactory,
		RequireFields:   RequireFileds,
		BuildHelper:     func() any { return BuildMessageHelper() },
	},
}

type HeadWrapper struct{}

func (w *HeadWrapper) CreateDocument(modelJson []byte) ([]byte, error) {
	return headWrapper.CreateDocument(modelJson)
}
func (w *HeadWrapper) ValidateDocument(modelJson []byte) error {
	return headWrapper.ValidateDocument(modelJson)
}
func (w *HeadWrapper) CheckRequireField(modelJson []byte) error {
	return headWrapper.CheckRequireField(modelJson)
}
func (w *HeadWrapper) GetDataModel(xmlData []byte) (modelJson []byte, err error) {
	return headWrapper.GetDataModel(xmlData)
}
func (w *HeadWrapper) GetHelp() ([]byte, error) {
	return headWrapper.GetHelp()
}
