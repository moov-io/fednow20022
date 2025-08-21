package FedNowPaymentStatus_pacs_002_001_10

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

var fedNowPaymentStatusWrapper = common.MessageWrapper[MessageModel, any]{
	Config: common.MessageWrapperConfig[MessageModel, any]{
		PathMap:         PathMap,
		DocumentFactory: DocumentFactory,
		DataFactory:     DataFactory,
		RequireFields:   RequireFileds,
		BuildHelper:     func() any { return BuildMessageHelper() },
	},
}

type FedNowPaymentStatusWrapper struct{}

func (w *FedNowPaymentStatusWrapper) CreateDocument(modelJson []byte) ([]byte, error) {
	return fedNowPaymentStatusWrapper.CreateDocument(modelJson)
}
func (w *FedNowPaymentStatusWrapper) ValidateDocument(modelJson []byte) error {
	return fedNowPaymentStatusWrapper.ValidateDocument(modelJson)
}
func (w *FedNowPaymentStatusWrapper) CheckRequireField(modelJson []byte) error {
	return fedNowPaymentStatusWrapper.CheckRequireField(modelJson)
}
func (w *FedNowPaymentStatusWrapper) GetDataModel(xmlData []byte) (modelJson []byte, err error) {
	return fedNowPaymentStatusWrapper.GetDataModel(xmlData)
}
func (w *FedNowPaymentStatusWrapper) GetHelp() ([]byte, error) {
	return fedNowPaymentStatusWrapper.GetHelp()
}
