package demoscheduler

import (
	"github.com/kinderyj/k8s-scheduler-framework-demo/pkg/plugins/filter"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
)

const (
	Name = "demoscheduler"
)

type Args struct {
	KubeConfig string `json:"kubeconfig,omitempty"`
	Master     string `json:"master,omitempty"`
}

var _ framework.FilterPlugin = &DemoScheduler{}
var _ framework.PreBindPlugin = &DemoScheduler{}

type DemoScheduler struct {
	args   *Args
	handle framework.FrameworkHandle
}

func New(configuration *runtime.Unknown, handle framework.FrameworkHandle) (framework.Plugin, error) {
	args := &Args{}
	if err := framework.DecodeInto(configuration, args); err != nil {
		return nil, err
	}
	klog.V(3).Infof("The Decoded args: %+v", args)
	return &DemoScheduler{
		args:   args,
		handle: handle,
	}, nil
}

func (s *DemoScheduler) Name() string {
	return Name
}

func (s *DemoScheduler) Filter(pc *framework.PluginContext, pod *v1.Pod, nodename string) *framework.Status {
	klog.V(4).Infof("Try to filter pod: %v, node: %v", pod.Name, nodename)
	nodeInfoMap := s.handle.NodeInfoSnapshot().NodeInfoMap
	if node, ok := nodeInfoMap[nodename]; ok {
		klog.Infof("The node info is %v", node)
		if ok, msg := filter.FilterWithGPU(node); ok {
			return framework.NewStatus(framework.Success, "")
		} else {
			return framework.NewStatus(framework.Unschedulable, "Node:"+nodename+msg)
		}
	}
	return framework.NewStatus(framework.Unschedulable, "Node: "+nodename+" Not Found.")
}

func (s *DemoScheduler) PreBind(pc *framework.PluginContext, pod *v1.Pod, nodeName string) *framework.Status {
	nodeInfoMap := s.handle.NodeInfoSnapshot().NodeInfoMap
	klog.V(3).Infof("The nodeName is %s, nodeInfo_map: %+v", nodeName, nodeInfoMap)
	return framework.NewStatus(framework.Success, "")
}
