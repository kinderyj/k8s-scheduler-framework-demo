package plugins

import (
	"github.com/kinderyj/k8s-scheduler-framework-demo/pkg/plugins/demoscheduler"
	"github.com/spf13/cobra"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
)

func Register() *cobra.Command {
	return app.NewSchedulerCommand(
		app.WithPlugin(demoscheduler.Name, demoscheduler.New),
	)
}
