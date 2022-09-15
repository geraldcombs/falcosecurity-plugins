module github.com/falcosecurity/plugins/plugins/cloudtrail

go 1.15

require github.com/aws/aws-sdk-go v1.44.51

require (
	github.com/alecthomas/jsonschema v0.0.0-20220216202328-9eeeec9d044b
	github.com/aws/aws-lambda-go v1.34.1
	github.com/aws/aws-sdk-go-v2/config v1.15.7
	github.com/aws/aws-sdk-go-v2/service/sqs v1.18.5
	github.com/falcosecurity/plugin-sdk-go v0.5.0
	github.com/falcosecurity/plugins/shared/go/aws/session v0.0.0-00010101000000-000000000000
	github.com/valyala/fastjson v1.6.3
)

replace github.com/falcosecurity/plugins/shared/go/aws/session => ../../shared/go/aws/session
