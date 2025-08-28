package RequestForPayment_pain_013_001_07

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

var requestForPaymentWrapper = common.MessageWrapper[MessageModel, any]{
	Config: common.MessageWrapperConfig[MessageModel, any]{
		PathMap:         PathMap,
		DocumentFactory: DocumentFactory,
		DataFactory:     DataFactory,
		RequireFields:   RequireFileds,
		BuildHelper:     func() any { return BuildMessageHelper() },
	},
}

type RequestForPaymentWrapper struct{}

func (w *RequestForPaymentWrapper) CreateDocument(modelJson []byte) ([]byte, error) {
	return requestForPaymentWrapper.CreateDocument(modelJson)
}
func (w *RequestForPaymentWrapper) ValidateDocument(modelJson []byte) error {
	return requestForPaymentWrapper.ValidateDocument(modelJson)
}
func (w *RequestForPaymentWrapper) CheckRequireField(modelJson []byte) error {
	return requestForPaymentWrapper.CheckRequireField(modelJson)
}
func (w *RequestForPaymentWrapper) GetDataModel(xmlData []byte) (modelJson []byte, err error) {
	return requestForPaymentWrapper.GetDataModel(xmlData)
}
func (w *RequestForPaymentWrapper) GetHelp() ([]byte, error) {
	return requestForPaymentWrapper.GetHelp()
}
