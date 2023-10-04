package utils

import (
	"fmt"
	"regexp"
	"strings"
)

type ImportComponent struct {
	Label     string
	Regexp    *regexp.Regexp
	PrimaryID bool
}

// Parse Import ID format
func ParseImportID(id string, components ...ImportComponent) (map[string]string, error) {

	keys := make([]string, len(components))
	regexpList := make([]string, len(components))

	i := 0
	for _, v := range components {
		keys[i] = v.Label
		regexpList[i] = v.Regexp.String()
		i++
	}

	compiledRegexpString := fmt.Sprintf("^%s$", strings.Join(regexpList, `\/`))

	m, err := regexp.MatchString(compiledRegexpString, id)
	if err != nil {
		return nil, fmt.Errorf("Cannot verify import ID regex: %s", err)
	}

	if !m {
		return nil, fmt.Errorf("Invalid import ID specified (\"%s\").  The ID should be in the format \"%s\" and must match regex: %s", id, strings.Join(keys, "/"), compiledRegexpString)
	}

	attributeValues := strings.SplitN(id, "/", len(components))

	if len(attributeValues) != len(components) {
		return nil, fmt.Errorf("Invalid import ID specified (\"%s\").  The ID should be in the format \"%s\".", id, strings.Join(keys, "/"))
	}

	attributes := make(map[string]string)

	i = 0
	for _, v := range components {
		attributes[v.Label] = attributeValues[i]
		i++
	}

	return attributes, nil
}
