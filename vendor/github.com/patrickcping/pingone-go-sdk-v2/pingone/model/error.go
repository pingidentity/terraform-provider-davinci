package model

import (
	"encoding/json"
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/authorize"
	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/patrickcping/pingone-go-sdk-v2/mfa"
)

// GenericOpenAPIError Provides access to the body, error and model on returned errors.
type GenericOpenAPIError struct {
	body  []byte
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericOpenAPIError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericOpenAPIError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericOpenAPIError) Model() interface{} {
	return e.model
}

func RemarshalGenericOpenAPIErrorObj(errorInf interface{}) (*GenericOpenAPIError, error) {

	var errorJSON []byte
	var err error

	var errorObj *GenericOpenAPIError

	switch t := errorInf.(type) {
	case *management.GenericOpenAPIError:
		errorObj = &GenericOpenAPIError{
			body:  t.Body(),
			error: t.Error(),
		}
		errorJSON, err = json.Marshal(t.Model())
	case *mfa.GenericOpenAPIError:
		errorObj = &GenericOpenAPIError{
			body:  t.Body(),
			error: t.Error(),
		}
		errorJSON, err = json.Marshal(t.Model())
	case *authorize.GenericOpenAPIError:
		errorObj = &GenericOpenAPIError{
			body:  t.Body(),
			error: t.Error(),
		}
		errorJSON, err = json.Marshal(t.Model())
	case GenericOpenAPIError:
		errorObj = &GenericOpenAPIError{
			body:  t.Body(),
			error: t.Error(),
		}
		errorJSON, err = json.Marshal(t.Model())
	default:
		return nil, fmt.Errorf("Cannot Marshal OpenAPI error interface for Remarshal, unknown type - %T", t)
	}

	if err != nil {
		return nil, fmt.Errorf("Cannot Marshal OpenAPI error interface for Remarshal - %s", err)
	}

	if errorJSON != nil {

		var model P1Error

		err = json.Unmarshal(errorJSON, &model)
		if err != nil {
			return nil, fmt.Errorf("Cannot Unmarshal OpenAPI error interface for Remarshal - %s", err)
		}

		errorObj.model = model
	}

	return errorObj, nil
}

func RemarshalErrorObj(errorInf interface{}) (*P1Error, error) {

	var errorJSON []byte
	var err error

	errorJSON, err = json.Marshal(errorInf)
	if err != nil {
		return nil, fmt.Errorf("Cannot Marshal error interface for Remarshal - %s", err)
	}

	var errorObj *P1Error

	err = json.Unmarshal(errorJSON, &errorObj)
	if err != nil {
		return nil, fmt.Errorf("Cannot Marshal error interface for Remarshal - %s", err)
	}

	return errorObj, nil
}
