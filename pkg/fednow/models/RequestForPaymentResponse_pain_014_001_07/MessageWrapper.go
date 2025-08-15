package RequestForPaymentResponse_pain_014_001_07

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

var requestForPaymentResponseWrapper = common.MessageWrapper[MessageModel, any]{
	Config: common.MessageWrapperConfig[MessageModel, any]{
		PathMap:         PathMap,
		DocumentFactory: DocumentFactory,
		DataFactory:     DataFactory,
		RequireFields:   RequireFileds,
		BuildHelper:     func() any { return BuildMessageHelper() },
	},
}

type RequestForPaymentResponseWrapper struct{}

func (w *RequestForPaymentResponseWrapper) CreateDocument(modelJson []byte) ([]byte, error) {
	return requestForPaymentResponseWrapper.CreateDocument(modelJson)
}
func (w *RequestForPaymentResponseWrapper) ValidateDocument(modelJson []byte) error {
	return requestForPaymentResponseWrapper.ValidateDocument(modelJson)
}
func (w *RequestForPaymentResponseWrapper) CheckRequireField(modelJson []byte) error {
	return requestForPaymentResponseWrapper.CheckRequireField(modelJson)
}
func (w *RequestForPaymentResponseWrapper) GetDataModel(xmlData []byte) (modelJson []byte, err error) {
	return requestForPaymentResponseWrapper.GetDataModel(xmlData)
}
func (w *RequestForPaymentResponseWrapper) GetHelp() ([]byte, error) {
	return requestForPaymentResponseWrapper.GetHelp()
}
