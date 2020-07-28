
# Scheduler FrameWork demo

This repo is a demo for the feature scheduler framework based on kubernetes 1.16.

Note that the feature scheduler framework in kubernetes 1.15 is in alpha, and this feature will not graduate to Beta in 1.17/1.18.

For the progress of the scheduler framework, please refer to the （[enhancements] (https://github.com/kubernetes/enhancements/issues/624))

## Get Started 

- Deploy
```
kubectl apply -f ./deploy/
```

- Build local
```
make local
```

- Build image
```
make image
```

## Deploy a pod using this scheduler

```
cat <<EOF | kubectl apply --filename -
apiVersion: apps/v1
kind: Deployment
metadata:
  name: busybox
spec:
  replicas: 1
  selector:
    matchLabels:
      app: busybox
  template:
    metadata:
      labels:
        app: busybox
    spec:
      schedulerName: scheduler-framework-demo
      terminationGracePeriodSeconds: 5
      containers:
      - image: busybox:latest
        imagePullPolicy: IfNotPresent
        name: busybox
        command: ["sleep", "3600"]
EOF
```

## How the dependencies inited for kubernetes 1.16.
pin all the dependencies to a matching level of k8s.io/kubernetes.
You can use v0.x.y for the published components that correspond to Kubernetes v1.x.y without needing to figure out the timestamps/hashes. 
This is available in 1.15.10+, 1.16.4+, and 1.17.0+. For Kubernetes 1.15, all the dependencies lists as below. 
```
require k8s.io/kubernetes v1.16.4
replace (
	k8s.io/api => k8s.io/api v0.16.4
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.16.4
	k8s.io/apimachinery => k8s.io/apimachinery v0.16.4
	k8s.io/apiserver => k8s.io/apiserver v0.16.4
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.16.4
	k8s.io/client-go => k8s.io/client-go v0.16.4
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.16.4
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.16.4
	k8s.io/code-generator => k8s.io/code-generator v0.16.4
	k8s.io/component-base => k8s.io/component-base v0.16.4
	k8s.io/cri-api => k8s.io/cri-api v0.16.4
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.16.4
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.16.4
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.16.4
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.16.4
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.16.4
	k8s.io/kubectl => k8s.io/kubectl v0.16.4
	k8s.io/kubelet => k8s.io/kubelet v0.16.4
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.16.4
	k8s.io/metrics => k8s.io/metrics v0.16.4
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.16.4
)
```