package AccountBalanceReport_camt_052_001_08

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

var accountBalanceReportReportWrapper = common.MessageWrapper[MessageModel, any]{
	Config: common.MessageWrapperConfig[MessageModel, any]{
		PathMap:         PathMap,
		DocumentFactory: DocumentFactory,
		DataFactory:     DataFactory,
		RequireFields:   RequireFileds,
		BuildHelper:     func() any { return BuildMessageHelper() },
	},
}

type AccountBalanceReportReportWrapper struct{}

func (w *AccountBalanceReportReportWrapper) CreateDocument(modelJson []byte) ([]byte, error) {
	return accountBalanceReportReportWrapper.CreateDocument(modelJson)
}
func (w *AccountBalanceReportReportWrapper) ValidateDocument(modelJson []byte) error {
	return accountBalanceReportReportWrapper.ValidateDocument(modelJson)
}
func (w *AccountBalanceReportReportWrapper) CheckRequireField(modelJson []byte) error {
	return accountBalanceReportReportWrapper.CheckRequireField(modelJson)
}
func (w *AccountBalanceReportReportWrapper) GetDataModel(xmlData []byte) (modelJson []byte, err error) {
	return accountBalanceReportReportWrapper.GetDataModel(xmlData)
}
func (w *AccountBalanceReportReportWrapper) GetHelp() ([]byte, error) {
	return accountBalanceReportReportWrapper.GetHelp()
}
