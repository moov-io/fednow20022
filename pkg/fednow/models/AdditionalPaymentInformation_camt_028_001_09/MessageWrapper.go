package AdditionalPaymentInformation_camt_028_001_09

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

var additionalPaymentInformationWrapper = common.MessageWrapper[MessageModel, any]{
	Config: common.MessageWrapperConfig[MessageModel, any]{
		PathMap:         PathMap,
		DocumentFactory: DocumentFactory,
		DataFactory:     DataFactory,
		RequireFields:   RequireFileds,
		BuildHelper:     func() any { return BuildMessageHelper() },
	},
}

type AdditionalPaymentInformationWrapper struct{}

func (w *AdditionalPaymentInformationWrapper) CreateDocument(modelJson []byte) ([]byte, error) {
	return additionalPaymentInformationWrapper.CreateDocument(modelJson)
}
func (w *AdditionalPaymentInformationWrapper) ValidateDocument(modelJson []byte) error {
	return additionalPaymentInformationWrapper.ValidateDocument(modelJson)
}
func (w *AdditionalPaymentInformationWrapper) CheckRequireField(modelJson []byte) error {
	return additionalPaymentInformationWrapper.CheckRequireField(modelJson)
}
func (w *AdditionalPaymentInformationWrapper) GetDataModel(xmlData []byte) (modelJson []byte, err error) {
	return additionalPaymentInformationWrapper.GetDataModel(xmlData)
}
func (w *AdditionalPaymentInformationWrapper) GetHelp() ([]byte, error) {
	return additionalPaymentInformationWrapper.GetHelp()
}
