package common

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
)

type dummyModel struct {
	Field string `json:"field"`
}

type dummyDoc struct{}

func (d *dummyDoc) Validate() error {
	if reflect.ValueOf(d).IsZero() {
		return errors.New("invalid doc")
	}
	return nil
}

func dummyDocumentFactory() ISODocument {
	return &dummyDoc{}
}

func dummyDataFactory() any {
	return &dummyModel{}
}

func TestCreateDocument(t *testing.T) {
	w := MessageWrapper[dummyModel, *dummyDoc]{
		Config: MessageWrapperConfig[dummyModel, *dummyDoc]{
			PathMap:         map[string]any{},
			DocumentFactory: dummyDocumentFactory,
		},
	}
	model := dummyModel{Field: "value"}
	b, _ := json.Marshal(model)
	_, err := w.CreateDocument(b)
	if err != nil {
		t.Errorf("CreateDocument failed: %v", err)
	}
	// Negative: invalid JSON
	_, err = w.CreateDocument([]byte(`{"bad":}`))
	if err == nil {
		t.Error("CreateDocument should fail on invalid JSON")
	}
}

func TestValidateDocument(t *testing.T) {
	w := MessageWrapper[dummyModel, *dummyDoc]{
		Config: MessageWrapperConfig[dummyModel, *dummyDoc]{
			PathMap:         map[string]any{},
			DocumentFactory: dummyDocumentFactory,
		},
	}
	model := dummyModel{Field: "value"}
	b, _ := json.Marshal(model)
	err := w.ValidateDocument(b)
	if err != nil {
		t.Errorf("ValidateDocument failed: %v", err)
	}
	// Negative: invalid JSON
	err = w.ValidateDocument([]byte(`{"bad":}`))
	if err == nil {
		t.Error("ValidateDocument should fail on invalid JSON")
	}
}

func TestCheckRequireField(t *testing.T) {
	w := MessageWrapper[dummyModel, *dummyDoc]{
		Config: MessageWrapperConfig[dummyModel, *dummyDoc]{
			RequireFields: []string{"Field"},
		},
	}
	model := dummyModel{Field: "value"}
	b, _ := json.Marshal(model)
	err := w.CheckRequireField(b)
	if err != nil {
		t.Errorf("CheckRequireField failed: %v", err)
	}
	// Negative: missing required field
	model = dummyModel{Field: ""}
	b, _ = json.Marshal(model)
	err = w.CheckRequireField(b)
	if err == nil {
		t.Error("CheckRequireField should fail when required field is missing or empty")
	}
}

func TestGetDataModel(t *testing.T) {
	w := MessageWrapper[dummyModel, *dummyDoc]{
		Config: MessageWrapperConfig[dummyModel, *dummyDoc]{
			DocumentFactory: dummyDocumentFactory,
			DataFactory:     dummyDataFactory,
			PathMap:         map[string]any{},
		},
	}
	// Valid XML with xmlns
	xmlData := []byte(`<dummyDoc xmlns="urn:test"></dummyDoc>`)
	_, err := w.GetDataModel(xmlData)
	if err != nil {
		t.Errorf("GetDataModel failed: %v", err)
	}
	// Negative: missing xmlns
	xmlData = []byte(`<dummyDoc></dummyDoc>`)
	_, err = w.GetDataModel(xmlData)
	if err == nil {
		t.Error("GetDataModel should fail without xmlns")
	}
	// Negative: invalid XML
	xmlData = []byte(`<dummyDoc><bad></dummyDoc>`)
	_, err = w.GetDataModel(xmlData)
	if err == nil {
		t.Error("GetDataModel should fail on invalid XML")
	}
}

func TestGetHelp(t *testing.T) {
	w := MessageWrapper[dummyModel, *dummyDoc]{
		Config: MessageWrapperConfig[dummyModel, *dummyDoc]{
			BuildHelper: func() any { return map[string]string{"help": "info"} },
		},
	}
	_, err := w.GetHelp()
	if err != nil {
		t.Errorf("GetHelp failed: %v", err)
	}
	// Negative: BuildHelper returns non-marshalable value
	w.Config.BuildHelper = func() any { return make(chan int) }
	_, err = w.GetHelp()
	if err == nil {
		t.Error("GetHelp should fail if BuildHelper returns non-marshalable value")
	}
}
