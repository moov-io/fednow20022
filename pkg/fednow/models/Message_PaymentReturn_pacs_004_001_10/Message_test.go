package Message_PaymentReturn_pacs_004_001_10

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func MergeXML(
	headerTitle string,
	headerXML []byte,
	dataTitle string,
	dataXML []byte) error {
	headerModelJson, err := HeaderWrapper.GetDataModel(headerXML)
	if err != nil {
		return fmt.Errorf("failed to get header model: %w", err)
	}
	dataModelJson, err := BodyWrapper.GetDataModel(dataXML)
	if err != nil {
		return fmt.Errorf("failed to get data model: %w", err)
	}
	var headerModel HeaderMessageModelType
	if err := json.Unmarshal(headerModelJson, &headerModel); err != nil {
		return fmt.Errorf("failed to unmarshal header model: %w", err)
	}
	var dataModel BodyMessageModelType
	if err := json.Unmarshal(dataModelJson, &dataModel); err != nil {
		return fmt.Errorf("failed to unmarshal data model: %w", err)
	}
	msgModel := &MessageModel{
		AppHdr:        *headerModel,
		PaymentReturn: *dataModel,
	}
	msgJson, err := json.MarshalIndent(msgModel, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal message model: %w", err)
	}
	err = common.WriteJSONTo("modelSample", headerTitle+"_"+dataTitle+".json", msgJson)
	return err
}

func TestCreateJson(t *testing.T) {
	var moduleName = "PaymentReturn"
	var headerSamplePath = "../head_001_001_02/swiftSample"
	var bodySamplePath = "../PaymentReturn_pacs_004_001_10/swiftSample"
	type sampleMap struct {
		title string
		path  string
	}
	headerSamplePaths, err := common.GetSubFilePaths(headerSamplePath)
	require.NoError(t, err, "failed to get header sample paths")
	dataSamplePaths, err := common.GetSubFilePaths(bodySamplePath)
	require.NoError(t, err, "failed to get data sample paths")
	var headerSamples []sampleMap
	for _, path := range headerSamplePaths {
		filename := filepath.Base(path)
		title := strings.TrimSuffix(filename, filepath.Ext(filename))
		if strings.Contains(title, moduleName) {
			headerSamples = append(headerSamples, sampleMap{title, path})
		}
	}

	var dataSamples []sampleMap
	for _, path := range dataSamplePaths {
		filename := filepath.Base(path)
		title := strings.TrimSuffix(filename, filepath.Ext(filename))
		dataSamples = append(dataSamples, sampleMap{title, path})
	}
	count := 0
	maxTests := 9

	for _, headerSample := range headerSamples {
		for _, dataSample := range dataSamples {
			if count >= maxTests {
				break
			}
			t.Run(headerSample.title+"_"+dataSample.title, func(t *testing.T) {
				headerXML, err := common.ReadFile(headerSample.path)
				require.NoError(t, err, "failed to read header file %s", headerSample.path)
				dataXML, err := common.ReadFile(dataSample.path)
				require.NoError(t, err, "failed to read data file %s", dataSample.path)
				err = MergeXML(headerSample.title, headerXML, dataSample.title, dataXML)
				require.NoError(t, err, "failed to merge XML files")
			})
			count++
		}
		if count >= maxTests {
			break
		}
	}
}

func TestCreateDocument(t *testing.T) {
	wrapper := &MessagePaymentReturnWrapper{}
	/*Valid documents*/
	modelSamplePaths, err := common.GetSubFilePaths("modelSample")
	if err != nil {
		t.Fatalf("failed to get model sample paths: %v", err)
	}
	for _, path := range modelSamplePaths {
		t.Run("Testing file: "+path, func(t *testing.T) { // Include path in the test title
			filename := filepath.Base(path)                                        // "1.xml"
			nameWithoutExt := strings.TrimSuffix(filename, filepath.Ext(filename)) // "1"
			data, err := common.ReadFile(path)
			require.NoError(t, err, "failed to read file %s", path)
			xmlData, err := wrapper.CreateDocument(data)
			require.NoError(t, err, "failed to create document from model %s", path)
			err = common.WriteXMLTo("xmlSample", nameWithoutExt+".xml", xmlData)
			require.NoError(t, err, "failed to write XML to generate %s", path)
		})
	}
	tests := []struct {
		name        string
		modelJson   []byte
		expectError bool
		errorMsg    string
	}{
		{
			name:        "invalid JSON returns error",
			modelJson:   []byte(`{"invalid": json}`),
			expectError: true,
			errorMsg:    "failed to unmarshal JSON",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   []byte(``),
			expectError: true,
			errorMsg:    "failed to unmarshal JSON",
		},
		{
			name:        "nil JSON returns error",
			modelJson:   nil,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := wrapper.CreateDocument(tt.modelJson)

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorMsg != "" {
					assert.Contains(t, err.Error(), tt.errorMsg)
				}
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)

				// Verify it's valid XML
				var xmlDoc interface{}
				err = xml.Unmarshal(result, &xmlDoc)
				assert.NoError(t, err, "Generated XML should be valid")
			}
		})
	}
}
func TestValidateDocument(t *testing.T) {
	wrapper := &MessagePaymentReturnWrapper{}
	/*Valid documents*/
	modelSamplePaths, err := common.GetSubFilePaths("modelSample")
	require.NoError(t, err, "failed to get sample paths")
	for _, path := range modelSamplePaths {
		t.Run("Testing file: "+path, func(t *testing.T) { // Include path in the test title
			data, err := common.ReadFile(path)
			require.NoError(t, err, "failed to read file %s", path)
			err = wrapper.ValidateDocument(data)
			require.NoError(t, err, "failed to validate document %s", path)
		})
	}
	tests := []struct {
		name        string
		modelJson   []byte
		expectError bool
		errorMsg    string
	}{
		{
			name:        "invalid JSON returns error",
			modelJson:   []byte(`{"invalid": json}`),
			expectError: true,
			errorMsg:    "failed to unmarshal JSON",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   []byte(``),
			expectError: true,
			errorMsg:    "failed to unmarshal JSON",
		},
		{
			name:        "nil JSON returns error",
			modelJson:   nil,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := wrapper.ValidateDocument(tt.modelJson)

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorMsg != "" {
					assert.Contains(t, err.Error(), tt.errorMsg)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCheckRequireField(t *testing.T) {
	wrapper := &MessagePaymentReturnWrapper{}
	/*Valid documents*/
	modelSamplePaths, err := common.GetSubFilePaths("modelSample")
	require.NoError(t, err, "failed to get sample paths")
	for _, path := range modelSamplePaths {
		t.Run("Testing file: "+path, func(t *testing.T) { // Include path in the test title
			data, err := common.ReadFile(path)
			require.NoError(t, err, "failed to read file %s", path)
			err = wrapper.CheckRequireField(data)
			require.NoError(t, err, "failed to check required fields in document %s", path)
		})
	}
	tests := []struct {
		name        string
		modelJson   []byte
		expectError bool
		errorMsg    string
	}{
		{
			name:        "invalid JSON returns error",
			modelJson:   []byte(`{"invalid": json}`),
			expectError: true,
			errorMsg:    "failed to unmarshal JSON",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   []byte(``),
			expectError: true,
			errorMsg:    "failed to unmarshal JSON",
		},
		{
			name:        "nil JSON returns error",
			modelJson:   nil,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := wrapper.CheckRequireField(tt.modelJson)

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorMsg != "" {
					assert.Contains(t, err.Error(), tt.errorMsg)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
func TestGetDataModel(t *testing.T) {
	wrapper := &MessagePaymentReturnWrapper{}
	/*Valid documents*/
	xmlSamplePaths, err := common.GetSubFilePaths("xmlSample")
	require.NoError(t, err, "failed to get sample paths")
	for _, path := range xmlSamplePaths {
		t.Run("Testing file: "+path, func(t *testing.T) { // Include path in the test title
			data, err := common.ReadFile(path)
			require.NoError(t, err, "failed to read file %s", path)
			modelJson, err := wrapper.GetDataModel(data)
			require.NoError(t, err, "failed to check required fields in document %s", path)
			require.NotNil(t, modelJson, "modelJson should not be nil for file %s", path)
		})
	}
	tests := []struct {
		name        string
		xmlData     []byte
		expectError bool
		errorMsg    string
	}{
		{
			name:        "invalid XML returns error",
			xmlData:     []byte(`<invalid>xml`),
			expectError: true,
			errorMsg:    "failed to unmarshal XML",
		},
		{
			name:        "empty XML returns error",
			xmlData:     []byte(``),
			expectError: true,
			errorMsg:    "failed to unmarshal XML",
		},
		{
			name:        "nil XML returns error",
			xmlData:     nil,
			expectError: true,
			errorMsg:    "failed to unmarshal XML",
		},
		{
			name:        "malformed XML returns error",
			xmlData:     []byte(`<?xml version="1.0"?><Document>missing namespace</Document>`),
			expectError: true,
			errorMsg:    "failed to unmarshal XML",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := wrapper.GetDataModel(tt.xmlData)

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorMsg != "" {
					assert.Contains(t, err.Error(), tt.errorMsg)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
func TestGetHelp(t *testing.T) {
	wrapper := &MessagePaymentReturnWrapper{}

	result, err := wrapper.GetHelp()

	assert.NoError(t, err)
	assert.NotEmpty(t, result)

	// Verify it's valid JSON
	var jsonData interface{}
	err = json.Unmarshal(result, &jsonData)
	assert.NoError(t, err, "Help result should be valid JSON")
}
func TestMain(m *testing.M) {
	// Run tests
	exitCode := m.Run()

	// After tests, check code coverage if running with coverage
	if coverProfile := os.Getenv("COVERAGE_CHECK"); coverProfile != "" {
		cmd := exec.Command("go", "tool", "cover", "-func=coverage.out")
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to check coverage: %v\n", err)
			os.Exit(2)
		}
		lines := strings.Split(string(output), "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "total:") {
				var percent float64
				_, err := fmt.Sscanf(line, "total: (statements) %f%%", &percent)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Failed to parse coverage percent: %v\n", err)
					os.Exit(2)
				}
				if percent < 80.0 {
					fmt.Fprintf(os.Stderr, "Code coverage is below 80%%: %.1f%%\n", percent)
					os.Exit(3)
				}
			}
		}
	}
	os.Exit(exitCode)
}
