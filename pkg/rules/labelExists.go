package rules

import (
	"bytes"

	"github.com/thedevsaddam/gojsonq/v2"
)

func LabelExists(json []byte) int {
	paths := gojsonq.New().Reader(bytes.NewReader(json)).Find("metadata.labels")

	if paths == nil {
		return 1
	}

	labels := paths.(map[string]any)

	for k := range labels {
		if k == "project" {
			return 0
		}
	}

	return 1
}
