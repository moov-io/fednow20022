package RequestForPaymentCancellationRequestResponse_camt_029_001_09

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

var requestForPaymentCancellationRequestResponseWrapper = common.MessageWrapper[MessageModel, any]{
	Config: common.MessageWrapperConfig[MessageModel, any]{
		PathMap:         PathMap,
		DocumentFactory: DocumentFactory,
		DataFactory:     DataFactory,
		RequireFields:   RequireFileds,
		BuildHelper:     func() any { return BuildMessageHelper() },
	},
}

type RequestForPaymentCancellationRequestResponseWrapper struct{}

func (w *RequestForPaymentCancellationRequestResponseWrapper) CreateDocument(modelJson []byte) ([]byte, error) {
	return requestForPaymentCancellationRequestResponseWrapper.CreateDocument(modelJson)
}
func (w *RequestForPaymentCancellationRequestResponseWrapper) ValidateDocument(modelJson []byte) error {
	return requestForPaymentCancellationRequestResponseWrapper.ValidateDocument(modelJson)
}
func (w *RequestForPaymentCancellationRequestResponseWrapper) CheckRequireField(modelJson []byte) error {
	return requestForPaymentCancellationRequestResponseWrapper.CheckRequireField(modelJson)
}
func (w *RequestForPaymentCancellationRequestResponseWrapper) GetDataModel(xmlData []byte) (modelJson []byte, err error) {
	return requestForPaymentCancellationRequestResponseWrapper.GetDataModel(xmlData)
}
func (w *RequestForPaymentCancellationRequestResponseWrapper) GetHelp() ([]byte, error) {
	return requestForPaymentCancellationRequestResponseWrapper.GetHelp()
}
