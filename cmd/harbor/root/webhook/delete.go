package webhook

import (
	"github.com/goharbor/harbor-cli/pkg/api"
	"github.com/goharbor/harbor-cli/pkg/prompt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strconv"
)

func DeleteWebhookCmd() *cobra.Command {

	var projectName string
	var webhookId string
	var webhookIdInt int64

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "webhook delete",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			var err error

			if projectName != "" && webhookId != "" {
				webhookIdInt, err = strconv.ParseInt(webhookId, 10, 64)
				err = api.DeleteWebhook(projectName, webhookIdInt)
			} else {
				projectName = prompt.GetProjectNameFromUser()
				webhookIdInt = prompt.GetWebhookFromUser(projectName)
				err = api.DeleteWebhook(projectName, webhookIdInt)

			}
			if err != nil {
				log.Errorf("failed to delete webhook: %v", err)
			}

		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&projectName, "project", "", "", "Project Name")
	flags.StringVarP(&webhookId, "webhook", "", "", "Webhook Id")

	return cmd
}
