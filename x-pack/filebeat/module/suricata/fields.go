// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package suricata

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "suricata", asset.ModuleFieldsPri, AssetSuricata); err != nil {
		panic(err)
	}
}

// AssetSuricata returns asset data.
// This is the base64 encoded gzipped contents of module/suricata.
func AssetSuricata() string {
	return "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"
}
