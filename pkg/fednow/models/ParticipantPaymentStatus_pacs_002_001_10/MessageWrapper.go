package ParticipantPaymentStatus_pacs_002_001_10

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

var participantPaymentStatusWrapper = common.MessageWrapper[MessageModel, any]{
	Config: common.MessageWrapperConfig[MessageModel, any]{
		PathMap:         PathMap,
		DocumentFactory: DocumentFactory,
		DataFactory:     DataFactory,
		RequireFields:   RequireFileds,
		BuildHelper:     func() any { return BuildMessageHelper() },
	},
}

type ParticipantPaymentStatusWrapper struct{}

func (w *ParticipantPaymentStatusWrapper) CreateDocument(modelJson []byte) ([]byte, error) {
	return participantPaymentStatusWrapper.CreateDocument(modelJson)
}
func (w *ParticipantPaymentStatusWrapper) ValidateDocument(modelJson []byte) error {
	return participantPaymentStatusWrapper.ValidateDocument(modelJson)
}
func (w *ParticipantPaymentStatusWrapper) CheckRequireField(modelJson []byte) error {
	return participantPaymentStatusWrapper.CheckRequireField(modelJson)
}
func (w *ParticipantPaymentStatusWrapper) GetDataModel(xmlData []byte) (modelJson []byte, err error) {
	return participantPaymentStatusWrapper.GetDataModel(xmlData)
}
func (w *ParticipantPaymentStatusWrapper) GetHelp() ([]byte, error) {
	return participantPaymentStatusWrapper.GetHelp()
}
