package AccountDebitCreditNotification_camt_054_001_08

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

var accountDebitCreditNotificationWrapper = common.MessageWrapper[MessageModel, any]{
	Config: common.MessageWrapperConfig[MessageModel, any]{
		PathMap:         PathMap,
		DocumentFactory: DocumentFactory,
		DataFactory:     DataFactory,
		RequireFields:   RequireFileds,
		BuildHelper:     func() any { return BuildMessageHelper() },
	},
}

type AccountDebitCreditNotificationWrapper struct{}

func (w *AccountDebitCreditNotificationWrapper) CreateDocument(modelJson []byte) ([]byte, error) {
	return accountDebitCreditNotificationWrapper.CreateDocument(modelJson)
}
func (w *AccountDebitCreditNotificationWrapper) ValidateDocument(modelJson []byte) error {
	return accountDebitCreditNotificationWrapper.ValidateDocument(modelJson)
}
func (w *AccountDebitCreditNotificationWrapper) CheckRequireField(modelJson []byte) error {
	return accountDebitCreditNotificationWrapper.CheckRequireField(modelJson)
}
func (w *AccountDebitCreditNotificationWrapper) GetDataModel(xmlData []byte) (modelJson []byte, err error) {
	return accountDebitCreditNotificationWrapper.GetDataModel(xmlData)
}
func (w *AccountDebitCreditNotificationWrapper) GetHelp() ([]byte, error) {
	return accountDebitCreditNotificationWrapper.GetHelp()
}
