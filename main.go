package main

import (
	"github.com/rossgray/custom-k8s-scheduler/plugin"
	"k8s.io/klog"
	scheduler "k8s.io/kubernetes/cmd/kube-scheduler/app"
)

func main() {
	command := scheduler.NewSchedulerCommand(
		scheduler.WithPlugin(plugin.Name, plugin.New),
	)
	if err := command.Execute(); err != nil {
		klog.Fatal(err)
	}
}
