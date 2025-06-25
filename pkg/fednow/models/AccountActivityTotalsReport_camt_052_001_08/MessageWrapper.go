package AccountActivityTotalsReport_camt_052_001_08

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

var accountActivityTotalsReportWrapper = common.MessageWrapper[MessageModel, any]{
	Config: common.MessageWrapperConfig[MessageModel, any]{
		PathMap:         PathMap,
		DocumentFactory: DocumentFactory,
		DataFactory:     DataFactory,
		RequireFields:   RequireFileds,
		BuildHelper:     func() any { return BuildMessageHelper() },
	},
}

type AccountActivityTotalsReportWrapper struct{}

func (w *AccountActivityTotalsReportWrapper) CreateDocument(modelJson []byte) ([]byte, error) {
	return accountActivityTotalsReportWrapper.CreateDocument(modelJson)
}
func (w *AccountActivityTotalsReportWrapper) ValidateDocument(modelJson []byte) error {
	return accountActivityTotalsReportWrapper.ValidateDocument(modelJson)
}
func (w *AccountActivityTotalsReportWrapper) CheckRequireField(modelJson []byte) error {
	return accountActivityTotalsReportWrapper.CheckRequireField(modelJson)
}
func (w *AccountActivityTotalsReportWrapper) GetDataModel(xmlData []byte) (modelJson []byte, err error) {
	return accountActivityTotalsReportWrapper.GetDataModel(xmlData)
}
func (w *AccountActivityTotalsReportWrapper) GetHelp() ([]byte, error) {
	return accountActivityTotalsReportWrapper.GetHelp()
}
