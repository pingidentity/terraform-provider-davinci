package acctest

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/samir-gandhi/davinci-client-go/davinci"
)

// takes path to json file, adjusts flow name attribute to include unique resource name, and returns json string
func ReadFlowJsonFile(path string) (string, error) {
	_, currentPath, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("failed to get current path")
	}
	flowByte, err := os.ReadFile(filepath.Join(filepath.Dir(currentPath), filepath.Clean(path)))
	if err != nil {
		return "", err
	}

	var flowMap map[string]interface{}
	err = json.Unmarshal(flowByte, &flowMap)
	if err != nil {
		return "", err
	}

	flowMap["name"] = flowMap["name"].(string)

	flowByte, err = json.Marshal(flowMap)
	if err != nil {
		return "", err
	}

	// Form up to object
	var flow davinci.Flow
	err = json.Unmarshal(flowByte, &flow)
	if err != nil {
		return "", err
	}

	// Back to string
	mainFlowJsonBytes, err := json.Marshal(flow)
	if err != nil {
		return "", err
	}

	return string(mainFlowJsonBytes), nil
}

// takes path to json file, adjusts flow name attribute to include unique resource name, and returns json string
func ReadFlowJsonFileRaw(path string) (string, error) {
	_, currentPath, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("failed to get current path")
	}
	flowByte, err := os.ReadFile(filepath.Join(filepath.Dir(currentPath), filepath.Clean(path)))
	if err != nil {
		return "", err
	}

	return string(flowByte), nil
}
