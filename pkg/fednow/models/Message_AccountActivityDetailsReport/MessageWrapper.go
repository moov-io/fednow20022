package Message_AccountActivityDetailsReport

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/moov-io/fednow20022/gen/AccountActivityDetailsReport_camt_052_001_08"
	"github.com/moov-io/fednow20022/gen/Message_AccountActivityDetailsReport_camt_052_001_08"
	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
	"github.com/moov-io/fednow20022/pkg/fednow/models/head_001_001_02"
)

type MessageAccountActivityDetailsReportWrapper struct{}

func (w *MessageAccountActivityDetailsReportWrapper) CreateDocument(modelJson []byte) ([]byte, error) {
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

	dataDoc, err := common.DocumentWith(model.AccountActivityDetailsReport, DataPathMap, DataDocumentFactory)
	if err != nil {
		return nil, fmt.Errorf("failed to create data document: %w", err)
	}
	if data, ok := dataDoc.(*AccountActivityDetailsReport_camt_052_001_08.Document); ok {
		doc := Message_AccountActivityDetailsReport_camt_052_001_08.Message{
			Xmlns:                        XMLAttributes,
			AppHdr:                       hdr,
			AccountActivityDetailsReport: data,
		}
		xmlData, err := xml.MarshalIndent(doc, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("failed to marshal document to XML: %w", err)
		}
		return xmlData, nil
	} else {
		// dataDoc is not of the expected type
		return nil, fmt.Errorf("data document is not of type *AccountActivityDetailsReport_camt_052_001_08.Document")
	}
}
func (w *MessageAccountActivityDetailsReportWrapper) ValidateDocument(modelJson []byte) error {
	var model MessageModel
	if err := json.Unmarshal(modelJson, &model); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	docXML, err := w.CreateDocument(modelJson)
	if err != nil {
		return fmt.Errorf("failed to create document: %w", err)
	}
	var doc Message_AccountActivityDetailsReport_camt_052_001_08.Message
	if err := xml.Unmarshal(docXML, &doc); err != nil {
		return fmt.Errorf("failed to unmarshal XML: %w", err)
	}
	return doc.Validate()
}
func (w *MessageAccountActivityDetailsReportWrapper) CheckRequireField(modelJson []byte) error {
	var model MessageModel
	if err := json.Unmarshal(modelJson, &model); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	headerJson, err := json.Marshal(model.AppHdr)
	if err != nil {
		return fmt.Errorf("failed to marshal header JSON: %w", err)
	}
	err = HeaderWrapper.CheckRequireField(headerJson)
	if err != nil {
		return fmt.Errorf("header validation failed: %w", err)
	}
	dataJson, err := json.Marshal(model.AccountActivityDetailsReport)
	if err != nil {
		return fmt.Errorf("failed to marshal data JSON: %w", err)
	}
	err = DataWrapper.CheckRequireField(dataJson)
	if err != nil {
		return fmt.Errorf("data validation failed: %w", err)
	}
	return err
}
func (w *MessageAccountActivityDetailsReportWrapper) GetDataModel(xmlData []byte) (modelJson []byte, err error) {
	var doc Message_AccountActivityDetailsReport_camt_052_001_08.Message
	if err := xml.Unmarshal(xmlData, &doc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal XML: %w", err)
	}
	header := doc.AppHdr
	dataModel := doc.AccountActivityDetailsReport

	docheader, err := common.MessageModelWith(header, HeaderDataFactory, HeaderPathMap)
	if err != nil {
		return nil, err
	}
	dataModelModel, err := common.MessageModelWith(dataModel, DataDataFactory, DataPathMap)
	if err != nil {
		return nil, err
	}
	msgModel, err := NewMessageModel(docheader, dataModelModel)
	if err != nil {
		return nil, err
	}
	modelJson, err = json.Marshal(msgModel)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal message model: %w", err)
	}
	return modelJson, nil

}
func (w *MessageAccountActivityDetailsReportWrapper) GetHelp()  ([]byte, error) {
	helper := BuildHelper()
	jsonData, err := json.MarshalIndent(helper, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal MessageHelper to JSON: %w", err)
	}
	return jsonData, nil
}