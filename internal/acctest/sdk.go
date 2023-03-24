package acctest

import (
	"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

type SchemaAttributeFloat64 struct {
	AttributeName string
	ExpectedValue float64
	ActualValue   float64
}
type SchemaAttributeBoolean struct {
	AttributeName string
	ExpectedValue bool
	ActualValue   bool
}
type SchemaAttributeString struct {
	AttributeName string
	ExpectedValue string
	ActualValue   string
}
type SchemaAttributeMap struct {
	AttributeName string
	ExpectedValue map[string]interface{}
	ActualValue   map[string]interface{}
}

func (s *SchemaAttributeFloat64) Compare() error {
	if s.ExpectedValue != s.ActualValue {
		return fmt.Errorf("Expected %s to be %d, got %d", s.AttributeName, int(s.ExpectedValue), int(s.ActualValue))
	}
	return nil
}

func (s *SchemaAttributeBoolean) Compare() error {
	if s.ExpectedValue != s.ActualValue {
		return fmt.Errorf("Expected %s to be %t, got %t", s.AttributeName, bool(s.ExpectedValue), bool(s.ActualValue))
	}
	return nil
}

func (s *SchemaAttributeString) Compare() error {
	if s.ExpectedValue != s.ActualValue {
		return fmt.Errorf("Expected %s to be %s, got %s", s.AttributeName, s.ExpectedValue, s.ActualValue)
	}
	return nil
}

func (s *SchemaAttributeMap) Compare() error {
	if len(s.ExpectedValue) != len(s.ActualValue) {
		return fmt.Errorf("Expected %s to be %d, got %d", s.AttributeName, len(s.ExpectedValue), len(s.ActualValue))
	}
	return nil
}

// Can pull top-level attributes from state and return as string
func GetAttributeFromState(s *terraform.State, resourceFullName, attributeName string) (string, error) {
	rs, ok := s.RootModule().Resources[resourceFullName]
	if !ok {
		return "", fmt.Errorf("Not found: %s", resourceFullName)
	}
	if rs.Primary.Attributes[attributeName] != "" {
		return rs.Primary.Attributes[attributeName], nil
	}
	return "", fmt.Errorf("Attribute %s not found in state", attributeName)
}

type CompareFunc func(interface{}) error

func ComposeCompare(fs ...error) error {
	var result *multierror.Error
	for i, f := range fs {
		if err := f; err != nil {
			result = multierror.Append(result, fmt.Errorf("Attribute Check %d/%d error: %s", i+1, len(fs), err))
		}
	}
	return result.ErrorOrNil()
}
