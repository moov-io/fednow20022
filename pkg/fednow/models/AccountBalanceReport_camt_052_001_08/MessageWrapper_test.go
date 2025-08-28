package AccountBalanceReport_camt_052_001_08

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"testing"

	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMakeValidDataJson(t *testing.T) {
	swiftSamplePaths, err := common.GetSubFilePaths("swiftSample")
	require.NoError(t, err, "failed to get sample paths")
	path := swiftSamplePaths[0] // Use the first sample path for testing
	data, err := common.ReadFile(path)
	require.NoError(t, err, "failed to read file %s", path)
	/*Create Document from Swift Sample File*/
	document, xlns, err := common.DocumentFrom(data, DocumentFactory)
	require.NoError(t, err, "failed to parse document from %s", path)
	require.NotNil(t, document, "document should not be nil for %s", path)
	require.Equal(t, xlns, XLNS, "unexpected XML namespace for %s", path)
	/*Create Model from Document*/
	messageModel, err := common.MessageModelWith(document, DataFactory, PathMap)
	require.NoError(t, err, "failed to create message model from document %s", path)
	require.NotNil(t, messageModel, "message model should not be nil for %s", path)

	jsonData, err := common.JSONFrom(messageModel)
	require.NoError(t, err, "failed to convert message model to JSON for %s", path)
	require.NotEmpty(t, jsonData, "JSON data should not be empty for %s", path)

	err = common.WriteJSONToGenerate("ValidData.json", jsonData)
	require.NoError(t, err, "failed to write valid data JSON to file")
	require.NotEmpty(t, data, "data should not be empty")
}
func TestCreateDocument(t *testing.T) {
	wrapper := &AccountBalanceReportReportWrapper{}
	validData, err := common.ReadFile("./modelSample/ValidData.json")
	require.NoError(t, err, "failed to read valid data JSON")

	tests := []struct {
		name        string
		modelJson   []byte
		expectError bool
		errorMsg    string
	}{
		{
			name:        "valid model creates document successfully",
			modelJson:   validData,
			expectError: false,
			errorMsg:    "",
		},
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
	wrapper := &AccountBalanceReportReportWrapper{}
	validData, err := common.ReadFile("./modelSample/ValidData.json")
	require.NoError(t, err, "failed to read valid data JSON")

	tests := []struct {
		name        string
		modelJson   []byte
		expectError bool
		errorMsg    string
	}{
		{
			name:        "valid model validates successfully",
			modelJson:   validData,
			expectError: false,
			errorMsg:    "",
		},
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
	wrapper := &AccountBalanceReportReportWrapper{}
	validData, err := common.ReadFile("./modelSample/ValidData.json")
	require.NoError(t, err, "failed to read valid data JSON")

	tests := []struct {
		name        string
		modelJson   []byte
		expectError bool
		errorMsg    string
	}{
		{
			name:        "valid model has all required fields",
			modelJson:   validData,
			expectError: false,
			errorMsg:    "",
		},
		{
			name: "model with missing required field fails validation",
			modelJson: []byte(
				`{
					"created_date_time": "2023-09-02T19:31:13-04:00",
					"pagenation": {
							"page_number": "1",
							"last_page_indicator": true
					}
				}`),
			expectError: true,
			errorMsg:    "required field",
		},
		{
			name: "model with missing required field fails validation",
			modelJson: []byte(
				`{
					"message_id": "12345",
					"pagenation": {
							"page_number": "1",
							"last_page_indicator": true
					}
				}`),
			expectError: true,
			errorMsg:    "required field",
		},
		{
			name: "model with missing required field fails validation",
			modelJson: []byte(
				`{
					"message_id": "12345",
					"created_date_time": "2023-09-02T19:31:13-04:00",
					"pagenation": {
							"last_page_indicator": true
					}
				}`),
			expectError: true,
			errorMsg:    "required field",
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
				/*Check document validate*/
				err = wrapper.ValidateDocument(tt.modelJson)
				assert.NoError(t, err, "Document validation should pass for valid model")
			}
		})
	}
}
func TestGetDataModel(t *testing.T) {
	wrapper := &AccountBalanceReportReportWrapper{}
	swiftSamplePaths, err := common.GetSubFilePaths("swiftSample")
	require.NoError(t, err, "failed to get sample paths")
	path := swiftSamplePaths[0] // Use the first sample path for testing
	validData, err := common.ReadFile(path)
	require.NoError(t, err, "failed to read file %s", path)
	tests := []struct {
		name        string
		xmlData     []byte
		expectError bool
		errorMsg    string
	}{
		{
			name:        "valid XML converts successfully",
			xmlData:     validData,
			expectError: false,
		},
		{
			name:        "invalid XML returns error",
			xmlData:     []byte(`<invalid>xml`),
			expectError: true,
			errorMsg:    "XML decode failed for document",
		},
		{
			name:        "empty XML returns error",
			xmlData:     []byte(``),
			expectError: true,
			errorMsg:    "XML decode failed for document",
		},
		{
			name:        "nil XML returns error",
			xmlData:     nil,
			expectError: true,
			errorMsg:    "XML decode failed for document",
		},
		{
			name:        "malformed XML returns error",
			xmlData:     []byte(`<?xml version="1.0"?><Document>missing namespace</Document>`),
			expectError: true,
			errorMsg:    "XML document missing xmlns attribute",
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

func TestCustomerCreditTransferWrapper_GetHelp(t *testing.T) {
	wrapper := &AccountBalanceReportReportWrapper{}

	result, err := wrapper.GetHelp()

	assert.NoError(t, err)
	assert.NotEmpty(t, result)

	// Verify it's valid JSON
	var jsonData interface{}
	err = json.Unmarshal(result, &jsonData)
	assert.NoError(t, err, "Help result should be valid JSON")

	// Verify it contains expected fields
	assert.Contains(t, string(result), "MessageId")
	assert.Contains(t, string(result), "CreatedDateTime")
}
func Test_MessageModel_fields_exist_in_MessageHelper(t *testing.T) {
	modelType := reflect.TypeOf(MessageModel{})
	helperType := reflect.TypeOf(BuildMessageHelper())

	helperFields := map[string]bool{}
	for i := 0; i < helperType.NumField(); i++ {
		helperFields[helperType.Field(i).Name] = true
	}

	missing := []string{}
	for i := 0; i < modelType.NumField(); i++ {
		fieldName := modelType.Field(i).Name
		if !helperFields[fieldName] {
			missing = append(missing, fieldName)
		}
	}

	if len(missing) > 0 {
		t.Errorf("Fields in MessageModel missing in MessageHelper: %v", missing)
	}
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
