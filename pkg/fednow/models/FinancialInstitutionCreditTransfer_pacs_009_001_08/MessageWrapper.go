package FinancialInstitutionCreditTransfer_pacs_009_001_08

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

type FinancialInstitutionCreditTransferWrapper struct{}

func (w *FinancialInstitutionCreditTransferWrapper) CreateDocument(modelJson []byte) ([]byte, error) {
	return customerCreditTransferWrapper.CreateDocument(modelJson)
}
func (w *FinancialInstitutionCreditTransferWrapper) ValidateDocument(modelJson []byte) error {
	return customerCreditTransferWrapper.ValidateDocument(modelJson)
}
func (w *FinancialInstitutionCreditTransferWrapper) CheckRequireField(modelJson []byte) error {
	return customerCreditTransferWrapper.CheckRequireField(modelJson)
}
func (w *FinancialInstitutionCreditTransferWrapper) GetDataModel(xmlData []byte) (modelJson []byte, err error) {
	return customerCreditTransferWrapper.GetDataModel(xmlData)
}
func (w *FinancialInstitutionCreditTransferWrapper) GetHelp() ([]byte, error) {
	return customerCreditTransferWrapper.GetHelp()
}
