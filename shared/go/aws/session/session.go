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

package session

import (
	"os"
	"path"
	"strings"

	awsSession "github.com/aws/aws-sdk-go/aws/session"
	"gopkg.in/ini.v1"
)

type SchemaValues struct {
	Profiles string
}

// CreateSession returns a session for AWS API
func CreateSession(region, profile string) *awsSession.Session {
	if region != "" {
		os.Setenv("AWS_REGION", region)
	}
	if profile != "" {
		os.Setenv("AWS_PROFILE", profile)
	}
	return awsSession.Must(awsSession.NewSessionWithOptions(
		awsSession.Options{
			SharedConfigState: awsSession.SharedConfigEnable,
		},
	))
}

// getProfiles returns an array of AWS profile names gathered from the following sources:
//   - The AWS_PROFILE environment variable
//   - Sections in ~/.aws/credentials
//   - "profile ..." sections in ~/.aws/config
// The "default" profile is represented by an empty string. Entries are unique.
func getProfiles() []string {
	profiles := []string{""}
	unique := make(map[string]bool)
	homeDir, _ := os.UserHomeDir()

	if envProfile, ok := os.LookupEnv("AWS_PROFILE"); ok {
		profiles = append(profiles, envProfile)
		unique[envProfile] = true
	}

	credFile := os.Getenv("AWS_SHARED_CREDENTIALS_FILE");
	if credFile == "" {
		credFile = path.Join(homeDir, ".aws", "credentials")
	}
	if credIni, err := ini.Load(credFile); err == nil {
		for _, section := range credIni.SectionStrings() {
			if strings.ToLower(section) != "default" && !unique[section] {
				profiles = append(profiles, section)
				unique[section] = true
			}
		}
	}

	confFile := os.Getenv("AWS_CONFIG_FILE");
	if confFile == "" {
		confFile = path.Join(homeDir, ".aws", "config")
	}
	if confIni, err := ini.Load(confFile); err == nil {
		for _, section := range confIni.SectionStrings() {
			if strings.HasPrefix(section, "profile ") {
				section = strings.TrimPrefix(section, "profile ")
				if strings.ToLower(section) != "default" && !unique[section] {
					profiles = append(profiles, section)
					unique[section] = true
				}
			}
		}
	}
	return profiles
}

func GetLocalSchemaValues() SchemaValues {
	profiles := getProfiles()
	profileEnums := ""
	for i := range profiles {
		profileEnums += ",enum=" + profiles[i]
	}
	localSchemaValues := SchemaValues{
		profileEnums,
	}
	return localSchemaValues
}