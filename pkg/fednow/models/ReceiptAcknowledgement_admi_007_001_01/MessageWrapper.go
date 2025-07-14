package ReceiptAcknowledgement_admi_007_001_01

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

var receiptAcknowledgementWrapper = common.MessageWrapper[MessageModel, any]{
	Config: common.MessageWrapperConfig[MessageModel, any]{
		PathMap:         PathMap,
		DocumentFactory: DocumentFactory,
		DataFactory:     DataFactory,
		RequireFields:   RequireFileds,
		BuildHelper:     func() any { return BuildMessageHelper() },
	},
}

type ReceiptAcknowledgementWrapper struct{}

func (w *ReceiptAcknowledgementWrapper) CreateDocument(modelJson []byte) ([]byte, error) {
	return receiptAcknowledgementWrapper.CreateDocument(modelJson)
}
func (w *ReceiptAcknowledgementWrapper) ValidateDocument(modelJson []byte) error {
	return receiptAcknowledgementWrapper.ValidateDocument(modelJson)
}
func (w *ReceiptAcknowledgementWrapper) CheckRequireField(modelJson []byte) error {
	return receiptAcknowledgementWrapper.CheckRequireField(modelJson)
}
func (w *ReceiptAcknowledgementWrapper) GetDataModel(xmlData []byte) (modelJson []byte, err error) {
	return receiptAcknowledgementWrapper.GetDataModel(xmlData)
}
func (w *ReceiptAcknowledgementWrapper) GetHelp() ([]byte, error) {
	return receiptAcknowledgementWrapper.GetHelp()
}
