package filter

import (
	"k8s.io/kubernetes/pkg/scheduler/nodeinfo"
)

func FilterWithGPU(node *nodeinfo.NodeInfo) (bool, string) {
	var msg = ""
	if nodeFitGPU(node) {
		if nodeGPUHealthy(node) {
			return true, msg
		}
		return false, " GPU Unhealthy."
	}
	return false, " Not a GPU Node."
}

func nodeFitGPU(node *nodeinfo.NodeInfo) bool {
	if gpu, ok := node.Node().Labels["demoscheduler/GPU"]; ok {
		if gpu == "Yes" {
			return true
		}
	}
	return false
}

func nodeGPUHealthy(node *nodeinfo.NodeInfo) bool {
	if healthy, ok := node.Node().Labels["demoscheduler/Healthy"]; ok {
		if healthy == "Healthy" {
			return true
		}
	}
	return false
}
