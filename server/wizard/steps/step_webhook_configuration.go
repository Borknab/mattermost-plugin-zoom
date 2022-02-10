package steps

import (
	"fmt"
	"net/url"

	"github.com/mattermost/mattermost-plugin-api/experimental/flow/steps"

	"github.com/mattermost/mattermost-plugin-zoom/server/config"
)

const (
	stepNameWebhookConfiguration = "webhook_configuration"

	stepTitleWebhookConfiguration = "Configure webhook in Zoom"

	stepDescriptionWebhookConfiguration = `1. Click on the **Feature** category in the left sidebar.
2. Enable **Event Subscriptions**.
3. Click **Add New Event Subscription** and give it a name \(e.g. "Mattermost events"\).
4. Enter in **Event notification endpoint URL**: %s
5. For the **Event notification receiver** field, select "All users in the account"

%s

We'll select the webhook events in the next step.
`
)

func WebhookConfigurationStep(pluginURL string, getConfiguration config.GetConfigurationFunc) steps.Step {
	secret := getConfiguration().WebhookSecret
	secret = url.QueryEscape(secret)

	eventConfigImage := imagePathToMarkdown(pluginURL, "Event Configuration", "event_configuration.png")

	webhookURL := fmt.Sprintf("`%s/webhook?secret=%s`", pluginURL, secret)
	description := fmt.Sprintf(stepDescriptionWebhookConfiguration, webhookURL, eventConfigImage)

	return steps.NewCustomStepBuilder(stepNameWebhookConfiguration, stepTitleWebhookConfiguration, description).
		WithButton(steps.Button{
			Name:  "Continue",
			Style: steps.Default,
		}).
		Build()
}
