package ParticipantBroadcast_admi_004_001_02

import "github.com/moov-io/fednow20022/pkg/fednow/models/common"

var participantBroadcastWrapper = common.MessageWrapper[MessageModel, any]{
	Config: common.MessageWrapperConfig[MessageModel, any]{
		PathMap:         PathMap,
		DocumentFactory: DocumentFactory,
		DataFactory:     DataFactory,
		RequireFields:   RequireFileds,
		BuildHelper:     func() any { return BuildMessageHelper() },
	},
}

type ParticipantBroadcastWrapper struct{}

func (w *ParticipantBroadcastWrapper) CreateDocument(modelJson []byte) ([]byte, error) {
	return participantBroadcastWrapper.CreateDocument(modelJson)
}
func (w *ParticipantBroadcastWrapper) ValidateDocument(modelJson []byte) error {
	return participantBroadcastWrapper.ValidateDocument(modelJson)
}
func (w *ParticipantBroadcastWrapper) CheckRequireField(modelJson []byte) error {
	return participantBroadcastWrapper.CheckRequireField(modelJson)
}
func (w *ParticipantBroadcastWrapper) GetDataModel(xmlData []byte) (modelJson []byte, err error) {
	return participantBroadcastWrapper.GetDataModel(xmlData)
}
func (w *ParticipantBroadcastWrapper) GetHelp() ([]byte, error) {
	return participantBroadcastWrapper.GetHelp()
}
