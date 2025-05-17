package plugins

import (
	"github.com/rwynn/monstache/v5/monstachemap"
	"strings"
)

// Map transforms the document before it's indexed in Elasticsearch
func Map(input *monstachemap.MapperPluginInput) (output *monstachemap.MapperPluginOutput, err error) {
	doc := input.Document
	ns := input.Namespace

	// Check if this is from our target collection
	if ns == "translation.record" {
		// Create a new document with our desired schema
		newDoc := make(map[string]interface{})

		// Set default value for scene
		newDoc["scene"] = "tape"

		// Map the fields we want to keep according to the schema
		fieldsToMap := []string{"tapeId", "deviceId", "userId", "timestamp"}
		for _, field := range fieldsToMap {
			if val, ok := doc[field]; ok {
				newDoc[field] = val
			}
		}

		// Extract srcText if it exists
		var srcText, targetText string

		if val, ok := doc["srcText"]; ok {
			if str, ok := val.(string); ok {
				srcText = str
				// Map to sourceText field in Elasticsearch
				newDoc["sourceText"] = srcText
			}
		}

		// Extract targetText if it exists
		if val, ok := doc["targetText"]; ok {
			if str, ok := val.(string); ok {
				targetText = str
				// Map to translatedText field in Elasticsearch
				newDoc["translatedText"] = targetText
			}
		}

		// Combine srcText and targetText into translationsText
		translationsText := strings.TrimSpace(srcText + " " + targetText)
		if translationsText != "" {
			newDoc["translationsText"] = translationsText
		}

		// Keep the _id field to maintain document identity
		if id, ok := doc["_id"]; ok {
			newDoc["_id"] = id
		}

		output = &monstachemap.MapperPluginOutput{
			Document: newDoc,
			Index:    "translations",
		}
		return
	}

	// For non-target collections, pass through unchanged
	output = &monstachemap.MapperPluginOutput{Document: doc}
	return
}

// Filter determines if a document should be indexed
func Filter(input *monstachemap.MapperPluginInput) (keep bool, err error) {
	// Include all documents
	return true, nil
}

// Pipeline defines an aggregation pipeline for change streams
func Pipeline(ns string, changeStream bool) (stages []interface{}, err error) {
	// No custom pipeline
	return nil, nil
}

// Process allows pre-processing of documents

// Version tells Monstache the API level supported by this plugin
func Version() int {
	return 2
}

// FilterDocument returns which documents to include from a certain namespace
func FilterDocument(doc map[string]interface{}, namespace string) (bool, error) {
	return true, nil
}

// ProcessDocument allows for custom processing of documents before mapping
func ProcessDocument(doc map[string]interface{}, namespace string) error {
	return nil
}

// MapNamespace allows for changing the target namespace
func MapNamespace(namespace string) (string, error) {
	if namespace == "translation.record" {
		return "translations", nil
	}
	return namespace, nil
}
