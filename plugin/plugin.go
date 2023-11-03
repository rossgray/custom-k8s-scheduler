package plugin

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

const (
	Name = "MyCustomPlugin"
	// MatchLabel is the label used to match pods to nodes when scheduling
	// (this is just a dummy example since this functionality doesn't require a custom scheduler;
	// you can use nodeSelector for instance)
	MatchLabel = "nodeGroup"
)

type MyCustomPlugin struct {
	handle framework.Handle
}

var _ framework.FilterPlugin = &MyCustomPlugin{}

func (p *MyCustomPlugin) Name() string {
	return Name
}

func (p *MyCustomPlugin) Filter(
	ctx context.Context, state *framework.CycleState, pod *corev1.Pod, nodeInfo *framework.NodeInfo,
) *framework.Status {

	node := nodeInfo.Node()
	klog.Infof("Applying Filter to pod '%s' and node '%s'", pod.Name, node.Name)

	// first check if pod has required label
	nodeGroupPod, labelFound := pod.GetLabels()[MatchLabel]
	if !labelFound {
		// if pod doesn't have label, assume it can be scheduled anywhere
		klog.Infof("Pod doesn't have required '%s' label; can be scheduled anywhere", MatchLabel)
		return framework.NewStatus(framework.Success, "")
	}

	nodeGroupNode, labelFound := node.GetLabels()[MatchLabel]
	if !labelFound {
		// if we have label on the pod but not on the node, we can't schedule pod here
		klog.Infof("Node doesn't have required '%s' label -> pod cannot be scheduled here", MatchLabel)
		return framework.NewStatus(framework.Unschedulable, "node does not have required label")
	}

	if nodeGroupNode != nodeGroupPod {
		// if nodeGroup of pod and node don't match, we can't schedule pod here
		klog.Infof("Node '%s' label value (%s) does not match pod label value (%s)", MatchLabel, nodeGroupNode, nodeGroupPod)
		return framework.NewStatus(framework.Unschedulable, "nodeGroup of pod and node don't match")
	}

	// if we reach here, pod label matches node label, so we can schedule pod here
	klog.Infof("Node '%s' label value (%s) matches pod label value (%s) -> can be scheduled here", MatchLabel, nodeGroupNode, nodeGroupPod)
	return framework.NewStatus(framework.Success, "")
}

// New initializes a new plugin and returns it.
// (Note that pluginArgs is unused here but could be used to handle input configuration for our plugin)
func New(pluginArgs runtime.Object, handle framework.Handle) (framework.Plugin, error) {
	return &MyCustomPlugin{
		handle: handle,
	}, nil
}
