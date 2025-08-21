package Message_PaymentReturn_pacs_004_001_10

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
	"github.com/moov-io/fednow20022/pkg/fednow/models/head_001_001_02"
)

type MessagePaymentReturnWrapper struct{}

func (w *MessagePaymentReturnWrapper) CreateDocument(modelJson []byte) ([]byte, error) {
	var model MessageModel
	if err := json.Unmarshal(modelJson, &model); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	headerDoc, err := common.DocumentWith(model.AppHdr, HeaderPathMap, HeaderDocumentFactory)
	if err != nil {
		return nil, fmt.Errorf("failed to create header document: %w", err)
	}

	hdr, err := head_001_001_02.GetBusinessApplicationHeaderV02(headerDoc)
	if err != nil {
		return nil, fmt.Errorf("failed to get business application header: %w", err)
	}

	dataDoc, err := common.DocumentWith(model.PaymentReturn, BodyPathMap, BodyDocumentFactory)
	if err != nil {
		return nil, fmt.Errorf("failed to create data document: %w", err)
	}

	data, ok := dataDoc.(*GeneratedBodyType)
	if !ok {
		return nil, fmt.Errorf("data document is not of type *AccountActivityDetailsReport_camt_052_001_08.Document")
	}

	doc := GeneratedMsgType{
		Xmlns:         XMLAttributes,
		AppHdr:        hdr,
		PaymentReturn: data,
	}

	xmlData, err := xml.MarshalIndent(doc, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal document to XML: %w", err)
	}
	return xmlData, nil
}

func (w *MessagePaymentReturnWrapper) ValidateDocument(modelJson []byte) error {
	var model MessageModel
	if err := json.Unmarshal(modelJson, &model); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	docXML, err := w.CreateDocument(modelJson)
	if err != nil {
		return fmt.Errorf("failed to create document: %w", err)
	}
	var doc GeneratedMsgType
	if err := xml.Unmarshal(docXML, &doc); err != nil {
		return fmt.Errorf("failed to unmarshal XML: %w", err)
	}
	return doc.Validate()
}

func (w *MessagePaymentReturnWrapper) CheckRequireField(modelJson []byte) error {
	var model MessageModel
	if err := json.Unmarshal(modelJson, &model); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	headerJson, err := json.Marshal(model.AppHdr)
	if err != nil {
		return fmt.Errorf("failed to marshal header JSON: %w", err)
	}
	if err = HeaderWrapper.CheckRequireField(headerJson); err != nil {
		return fmt.Errorf("header validation failed: %w", err)
	}

	dataJson, err := json.Marshal(model.PaymentReturn)
	if err != nil {
		return fmt.Errorf("failed to marshal data JSON: %w", err)
	}
	if err = BodyWrapper.CheckRequireField(dataJson); err != nil {
		return fmt.Errorf("data validation failed: %w", err)
	}
	return nil
}

func (w *MessagePaymentReturnWrapper) GetDataModel(xmlData []byte) (modelJson []byte, err error) {
	var doc GeneratedMsgType
	if err := xml.Unmarshal(xmlData, &doc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal XML: %w", err)
	}

	docheader, err := common.MessageModelWith(doc.AppHdr, HeaderDataFactory, HeaderPathMap)
	if err != nil {
		return nil, err
	}
	bodyModelModel, err := common.MessageModelWith(doc.PaymentReturn, BodyDataFactory, BodyPathMap)
	if err != nil {
		return nil, err
	}

	msgModel, err := NewMessageModel(docheader, bodyModelModel)
	if err != nil {
		return nil, err
	}
	modelJson, err = json.Marshal(msgModel)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal message model: %w", err)
	}
	return modelJson, nil
}

func (w *MessagePaymentReturnWrapper) GetHelp() ([]byte, error) {
	helper := BuildHelper()
	jsonData, err := json.MarshalIndent(helper, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal MessageHelper to JSON: %w", err)
	}
	return jsonData, nil
}
