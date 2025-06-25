package AccountActivityDetailsReport_camt_052_001_08

import (
	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
)

var accountActivityDetailsReportWrapper = common.MessageWrapper[MessageModel, any]{
	Config: common.MessageWrapperConfig[MessageModel, any]{
		PathMap:         PathMap,
		DocumentFactory: DocumentFactory,
		DataFactory:     DataFactory,
		RequireFields:   RequireFileds,
		BuildHelper:     func() any { return BuildMessageHelper() },
	},
}

type AccountActivityDetailsReportWrapper struct{}

func (w *AccountActivityDetailsReportWrapper) CreateDocument(modelJson []byte) ([]byte, error) {
	return accountActivityDetailsReportWrapper.CreateDocument(modelJson)
}
func (w *AccountActivityDetailsReportWrapper) ValidateDocument(modelJson []byte) error {
	return accountActivityDetailsReportWrapper.ValidateDocument(modelJson)
}
func (w *AccountActivityDetailsReportWrapper) CheckRequireField(modelJson []byte) error {
	return accountActivityDetailsReportWrapper.CheckRequireField(modelJson)
}
func (w *AccountActivityDetailsReportWrapper) GetDataModel(xmlData []byte) (modelJson []byte, err error) {
	return accountActivityDetailsReportWrapper.GetDataModel(xmlData)
}
func (w *AccountActivityDetailsReportWrapper) GetHelp() ([]byte, error) {
	return accountActivityDetailsReportWrapper.GetHelp()
}
