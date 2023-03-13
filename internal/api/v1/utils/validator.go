package utils

import (
	"fmt"
	"log"

	"github.com/xeipuuv/gojsonschema"
)

func ValidateJSON(schemaRule []byte, data []byte) error {
	schemaLoader := gojsonschema.NewBytesLoader(schemaRule)
	schema, err := gojsonschema.NewSchema(schemaLoader)
	if err != nil {
		log.Printf("schema: %+v, error: %+v", schema, err)
		return fmt.Errorf("failed to load schema: %v", err)
	}

	// Validate the data against the schema
	dataLoader := gojsonschema.NewBytesLoader(data)
	result, err := schema.Validate(dataLoader)
	if err != nil {
		log.Printf("Got error while loading validation data %+v", err.Error())
		return fmt.Errorf("failed to validate data against schema: %v", err)
	}

	// Check if validation succeeded
	if !result.Valid() {
		for _, err := range result.Errors() {
			log.Printf("Got error while validation %+v", err.String())
		}
		return fmt.Errorf("Event provided does not satisfy the eventRule: %v", result.Errors())
	}

	return nil
}
