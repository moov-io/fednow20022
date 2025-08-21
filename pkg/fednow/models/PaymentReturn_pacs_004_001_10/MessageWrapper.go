package PaymentReturn_pacs_004_001_10

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

var paymentReturnWrapper = common.MessageWrapper[MessageModel, any]{
	Config: common.MessageWrapperConfig[MessageModel, any]{
		PathMap:         PathMap,
		DocumentFactory: DocumentFactory,
		DataFactory:     DataFactory,
		RequireFields:   RequireFileds,
		BuildHelper:     func() any { return BuildMessageHelper() },
	},
}

type PaymentReturnWrapper struct{}

func (w *PaymentReturnWrapper) CreateDocument(modelJson []byte) ([]byte, error) {
	return paymentReturnWrapper.CreateDocument(modelJson)
}
func (w *PaymentReturnWrapper) ValidateDocument(modelJson []byte) error {
	return paymentReturnWrapper.ValidateDocument(modelJson)
}
func (w *PaymentReturnWrapper) CheckRequireField(modelJson []byte) error {
	return paymentReturnWrapper.CheckRequireField(modelJson)
}
func (w *PaymentReturnWrapper) GetDataModel(xmlData []byte) (modelJson []byte, err error) {
	return paymentReturnWrapper.GetDataModel(xmlData)
}
func (w *PaymentReturnWrapper) GetHelp() ([]byte, error) {
	return paymentReturnWrapper.GetHelp()
}
