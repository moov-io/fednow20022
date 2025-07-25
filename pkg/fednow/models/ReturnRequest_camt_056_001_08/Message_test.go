package ReturnRequest_camt_056_001_08

import (
	"path/filepath"
	"testing"

	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
	"github.com/stretchr/testify/require"
)

func TestMappingTable(t *testing.T) {
	swiftSamplePaths, err := common.GetSubFilePaths("swiftSample")
	require.NoError(t, err, "failed to get sample paths")
	for _, path := range swiftSamplePaths {
		t.Run("Testing file: "+path, func(t *testing.T) { // Include path in the test title
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
			/*Create Document from Model*/
			newDocument, err := common.DocumentWith(messageModel, PathMap, DocumentFactory)
			require.NoError(t, err, "failed to create document from message model %s", path)
			require.NotNil(t, newDocument, "new document should not be nil for %s", path)

			/*Compare Document and newDocument*/
			xml_document, err := common.XMLFrom(document)
			require.NoError(t, err, "failed to convert original document to XML for %s", path)

			xml_newDocument, err := common.XMLFrom(newDocument)
			require.NoError(t, err, "failed to convert new document to XML for %s", path)
			/*Create New Generated XML*/
			fileName := filepath.Base(path)
			err = common.WriteXMLToGenerate("Result"+fileName+".xml", xml_newDocument)
			require.NoError(t, err, "failed to write new XML to file for %s", path)
			/*Compare XMLs*/
			isEqual := common.CompareXMLs(xml_document, xml_newDocument)
			require.True(t, isEqual, "XML documents should be equal for %s", path)
		})
	}
}
func TestRequireFeildCounts(t *testing.T) {
	document := DocumentFactory()
	xml_document, err := common.XMLFrom(document)
	require.NoError(t, err, "failed to convert document to XML")
	err = common.WriteXMLToGenerate("Empty"+".xml", xml_document)
	require.NoError(t, err, "failed to write empty XML to file")
	nonEmptyElementCount, err := common.CountLeafElements(xml_document)
	require.NoError(t, err, "failed to count non-empty elements in XML document")
	require.Equal(t, len(RequireFileds), nonEmptyElementCount, "number of non-empty elements should match required fields count")
}
