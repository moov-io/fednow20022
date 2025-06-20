package AccountActivityTotalsReport_camt_052_001_08

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
)

type AccountActivityTotalsReportWrapper struct{}

func (w *AccountActivityTotalsReportWrapper) CreateDocument(modelJson []byte) ([]byte, error) {
	var model MessageModel
	err := json.Unmarshal(modelJson, &model)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON to MessageModel: %w", err)
	}
	doc, err := common.DocumentWith(model, PathMap, DocumentFactory)
	if err != nil {
		return nil, fmt.Errorf("failed to create document: %w", err)
	}
	xmlData, err := xml.MarshalIndent(doc, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal document to XML: %w", err)
	}

	return xmlData, nil
}
func (w *AccountActivityTotalsReportWrapper) ValidateDocument(modelJson []byte) error {
	var model MessageModel
	err := json.Unmarshal(modelJson, &model)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON to MessageModel: %w", err)
	}
	doc, err := common.DocumentWith(model, PathMap, DocumentFactory)
	if err != nil {
		return fmt.Errorf("failed to create document: %w", err)
	}
	if err := doc.Validate(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	return nil
}
func (w *AccountActivityTotalsReportWrapper) CheckRequireField(modelJson []byte) error {
	var model MessageModel
	err := json.Unmarshal(modelJson, &model)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON to MessageModel: %w", err)
	}
	return common.CheckRequireFields(model, RequireFileds)
}
func (w *AccountActivityTotalsReportWrapper) GetDataModel(xmlData []byte) (modelJson []byte, err error) {
	doc, _, err := common.DocumentFrom(xmlData, DocumentFactory)
	if err != nil {
		return nil, err
	}

	model, err := common.MessageModelWith(doc, DataFactory, PathMap)
	if err != nil {
		return nil, err
	}

	modelJson, err = json.Marshal(model)
	if err != nil {
		return nil, err
	}

	return modelJson, nil
}
func (w *AccountActivityTotalsReportWrapper) GetHelp() ([]byte, error) {
	// Build the MessageHelper structure
	helper := BuildMessageHelper()

	// Marshal the structure into a JSON string
	jsonData, err := json.MarshalIndent(helper, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal MessageHelper to JSON: %w", err)
	}

	return jsonData, nil
}
