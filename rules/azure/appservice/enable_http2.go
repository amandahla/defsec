package appservice

import (
	"github.com/aquasecurity/defsec/provider"
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/severity"
	"github.com/aquasecurity/defsec/state"
)

var CheckEnableHttp2 = rules.Register(
	rules.Rule{
		Provider:    provider.AzureProvider,
		Service:     "appservice",
		ShortCode:   "enable-http2",
		Summary:     "Web App uses the latest HTTP version",
		Impact:      "Outdated versions of HTTP has security vulnerabilities",
		Resolution:  "Use the latest version of HTTP",
		Explanation: `Use the latest version of HTTP to ensure you are benefiting from security fixes`,
		Links: []string{ 
		},
		Severity: severity.Low,
	},
	func(s *state.State) (results rules.Results) {
		for _, x := range s.AWS.S3.Buckets {
			if x.Encryption.Enabled.IsFalse() {
				results.Add(
					"",
					x.Encryption.Enabled.Metadata(),
					x.Encryption.Enabled.Value(),
				)
			}
		}
		return
	},
)