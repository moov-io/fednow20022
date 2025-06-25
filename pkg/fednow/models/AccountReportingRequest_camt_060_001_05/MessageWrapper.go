package AccountReportingRequest_camt_060_001_05

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

var accountReportingRequestWrapper = common.MessageWrapper[MessageModel, any]{
	Config: common.MessageWrapperConfig[MessageModel, any]{
		PathMap:         PathMap,
		DocumentFactory: DocumentFactory,
		DataFactory:     DataFactory,
		RequireFields:   RequireFileds,
		BuildHelper:     func() any { return BuildMessageHelper() },
	},
}

type AccountReportingRequestWrapper struct{}

func (w *AccountReportingRequestWrapper) CreateDocument(modelJson []byte) ([]byte, error) {
	return accountReportingRequestWrapper.CreateDocument(modelJson)
}
func (w *AccountReportingRequestWrapper) ValidateDocument(modelJson []byte) error {
	return accountReportingRequestWrapper.ValidateDocument(modelJson)
}
func (w *AccountReportingRequestWrapper) CheckRequireField(modelJson []byte) error {
	return accountReportingRequestWrapper.CheckRequireField(modelJson)
}
func (w *AccountReportingRequestWrapper) GetDataModel(xmlData []byte) (modelJson []byte, err error) {
	return accountReportingRequestWrapper.GetDataModel(xmlData)
}
func (w *AccountReportingRequestWrapper) GetHelp() ([]byte, error) {
	return accountReportingRequestWrapper.GetHelp()
}
