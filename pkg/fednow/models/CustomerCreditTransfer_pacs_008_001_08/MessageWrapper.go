package CustomerCreditTransfer_pacs_008_001_08

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

var customerCreditTransferWrapper = common.MessageWrapper[MessageModel, any]{
	Config: common.MessageWrapperConfig[MessageModel, any]{
		PathMap:         PathMap,
		DocumentFactory: DocumentFactory,
		DataFactory:     DataFactory,
		RequireFields:   RequireFileds,
		BuildHelper:     func() any { return BuildMessageHelper() },
	},
}

type CustomerCreditTransferWrapper struct{}

func (w *CustomerCreditTransferWrapper) CreateDocument(modelJson []byte) ([]byte, error) {
	return customerCreditTransferWrapper.CreateDocument(modelJson)
}
func (w *CustomerCreditTransferWrapper) ValidateDocument(modelJson []byte) error {
	return customerCreditTransferWrapper.ValidateDocument(modelJson)
}
func (w *CustomerCreditTransferWrapper) CheckRequireField(modelJson []byte) error {
	return customerCreditTransferWrapper.CheckRequireField(modelJson)
}
func (w *CustomerCreditTransferWrapper) GetDataModel(xmlData []byte) (modelJson []byte, err error) {
	return customerCreditTransferWrapper.GetDataModel(xmlData)
}
func (w *CustomerCreditTransferWrapper) GetHelp() ([]byte, error) {
	return customerCreditTransferWrapper.GetHelp()
}
