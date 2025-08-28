package common

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type MessageWrapperConfig[Model any, Doc any] struct {
	PathMap         map[string]any
	DocumentFactory DocumentFactory
	DataFactory     DataFactory
	RequireFields   []string
	BuildHelper     func() any
}

type MessageWrapper[Model any, Doc any] struct {
	Config MessageWrapperConfig[Model, Doc]
}

func (w *MessageWrapper[Model, Doc]) CreateDocument(modelJson []byte) ([]byte, error) {
	var model Model
	if err := json.Unmarshal(modelJson, &model); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	doc, err := DocumentWith(model, w.Config.PathMap, w.Config.DocumentFactory)
	if err != nil {
		return nil, fmt.Errorf("failed to create document: %w", err)
	}
	xmlData, err := xml.MarshalIndent(doc, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal document to XML: %w", err)
	}
	return xmlData, nil
}

func (w *MessageWrapper[Model, Doc]) ValidateDocument(modelJson []byte) error {
	var model Model
	if err := json.Unmarshal(modelJson, &model); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	// doc, err := DocumentWith(model, w.Config.PathMap, w.Config.DocumentFactory)
	// if err != nil {
	// 	return fmt.Errorf("failed to create document: %w", err)
	// }
	// if err := doc.Validate(); err != nil {
	// 	return fmt.Errorf("validation failed: %w", err)
	// }
	return nil
}

func (w *MessageWrapper[Model, Doc]) CheckRequireField(modelJson []byte) error {
	var model Model
	if err := json.Unmarshal(modelJson, &model); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return CheckRequireFields(model, w.Config.RequireFields)
}

func (w *MessageWrapper[Model, Doc]) GetDataModel(xmlData []byte) (modelJson []byte, err error) {
	doc, _, err := DocumentFrom(xmlData, w.Config.DocumentFactory)
	if err != nil {
		return nil, err
	}

	model, err := MessageModelWith(doc, w.Config.DataFactory, w.Config.PathMap)
	if err != nil {
		return nil, err
	}

	modelJson, err = json.Marshal(model)
	if err != nil {
		return nil, err
	}

	return modelJson, nil
}

func (w *MessageWrapper[Model, Doc]) GetHelp() ([]byte, error) {
	helper := w.Config.BuildHelper()
	jsonData, err := json.MarshalIndent(helper, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal MessageHelper to JSON: %w", err)
	}
	return jsonData, nil
}
