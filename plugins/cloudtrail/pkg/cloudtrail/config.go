/*
Copyright (C) 2022 The Falco Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cloudtrail

import (
	"bytes"
	"reflect"
	"text/template"

	awsSessionExtras "github.com/falcosecurity/plugins/shared/go/aws/session"
)

// Struct for plugin init config.
// Do not use these tags directly; use PluginConfigFields() instead.
type PluginConfig struct {
	// either an enum item or an arbitrary string.
	// https://github.com/json-schema-org/json-schema-spec/issues/665
	Profile  string `json:"profile" jsonschema:"title=Credentials profile,description=Specifies the AWS credentials profile to use for authentication (Default: empty / default profile),default={{.Profiles}}"`
	Region              string `json:"region" jsonschema:"title=AWS region,description=Specifies the AWS region to connect to (Default: empty / default profile region),default="`
	S3DownloadConcurrency  int `json:"s3DownloadConcurrency" jsonschema:"title=S3 download concurrency,description=Controls the number of background goroutines used to download S3 files (Default: 1),default=1"`
	SQSDelete             bool `json:"sqsDelete" jsonschema:"title=Delete SQS messages,description=If true then the plugin will delete SQS messages from the queue immediately after receiving them (Default: true),default=true"`
	UseAsync              bool `json:"useAsync" jsonschema:"title=Use async extraction,description=If true then async extraction optimization is enabled (Default: true),default=true"`
	UseS3SNS              bool `json:"useS3SNS" jsonschema:"title=Use S3 SNS,description=If true then the plugin will expect SNS messages to originate from S3 instead of directly from Cloudtrail (Default: false),default=false"`
}

// Reset sets the configuration to its default values
func (p *PluginConfig) Reset() {
	p.Profile = ""
	p.Region = ""
	p.SQSDelete = true
	p.S3DownloadConcurrency = 1
	p.UseAsync = true
	p.UseS3SNS = false
}

// PluginConfigFields returns an array of PluginConfig fields with JSON Schema
// tags built from the current user environment.
func PluginConfigFields() []reflect.StructField {
	pcType := reflect.TypeOf(PluginConfig{})
	tmplFields := make([]reflect.StructField, pcType.NumField())
	tmplValues := awsSessionExtras.GetLocalSchemaValues()
	for i := 0; i < pcType.NumField(); i++ {
		field := pcType.Field(i)
		if tmpl, err := template.New("").Parse(string(field.Tag)); err == nil {
			var tmplTag bytes.Buffer
			tmpl.Execute(&tmplTag, tmplValues)
			field.Tag = reflect.StructTag(tmplTag.String())
		}
		tmplFields[i] = field
	}
	return tmplFields
}