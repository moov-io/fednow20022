package PaymentStatusRequest_pacs_028_001_03

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

var paymentStatusRequestWrapper = common.MessageWrapper[MessageModel, any]{
	Config: common.MessageWrapperConfig[MessageModel, any]{
		PathMap:         PathMap,
		DocumentFactory: DocumentFactory,
		DataFactory:     DataFactory,
		RequireFields:   RequireFileds,
		BuildHelper:     func() any { return BuildMessageHelper() },
	},
}

type PaymentStatusRequestWrapper struct{}

func (w *PaymentStatusRequestWrapper) CreateDocument(modelJson []byte) ([]byte, error) {
	return paymentStatusRequestWrapper.CreateDocument(modelJson)
}
func (w *PaymentStatusRequestWrapper) ValidateDocument(modelJson []byte) error {
	return paymentStatusRequestWrapper.ValidateDocument(modelJson)
}
func (w *PaymentStatusRequestWrapper) CheckRequireField(modelJson []byte) error {
	return paymentStatusRequestWrapper.CheckRequireField(modelJson)
}
func (w *PaymentStatusRequestWrapper) GetDataModel(xmlData []byte) (modelJson []byte, err error) {
	return paymentStatusRequestWrapper.GetDataModel(xmlData)
}
func (w *PaymentStatusRequestWrapper) GetHelp() ([]byte, error) {
	return paymentStatusRequestWrapper.GetHelp()
}
