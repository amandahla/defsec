package keyvault

import (
	"github.com/aquasecurity/defsec/provider"
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/severity"
	"github.com/aquasecurity/defsec/state"
)

var CheckContentTypeForSecret = rules.Register(
	rules.Rule{
		Provider:    provider.AzureProvider,
		Service:     "keyvault",
		ShortCode:   "content-type-for-secret",
		Summary:     "Key vault Secret should have a content type set",
		Impact:      "The secret's type is unclear without a content type",
		Resolution:  "Provide content type for secrets to aid interpretation on retrieval",
		Explanation: `Content Type is an optional Key Vault Secret behavior and is not enabled by default.

Clients may specify the content type of a secret to assist in interpreting the secret data when it's retrieved. The maximum length of this field is 255 characters. There are no pre-defined values. The suggested usage is as a hint for interpreting the secret data.`,
		Links: []string{ 
			"https://docs.microsoft.com/en-us/azure/key-vault/secrets/about-secrets",
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