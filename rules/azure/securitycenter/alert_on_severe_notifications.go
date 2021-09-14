package securitycenter

import (
	"github.com/aquasecurity/defsec/provider"
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/severity"
	"github.com/aquasecurity/defsec/state"
)

var CheckAlertOnSevereNotifications = rules.Register(
	rules.Rule{
		Provider:    provider.AzureProvider,
		Service:     "security-center",
		ShortCode:   "alert-on-severe-notifications",
		Summary:     "Send notification emails for high severity alerts",
		Impact:      "The ability to react to high severity notifications could be delayed",
		Resolution:  " Set alert notifications to be on",
		Explanation: `It is recommended that at least one valid contact is configured for the security center. 
Microsoft will notify the security contact directly in the event of a security incident using email and require alerting to be turned on.`,
		Links: []string{ 
			"https://azure.microsoft.com/en-us/services/security-center/",
		},
		Severity: severity.Medium,
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