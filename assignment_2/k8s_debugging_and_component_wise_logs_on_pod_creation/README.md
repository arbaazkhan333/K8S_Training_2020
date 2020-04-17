Steps used in this assignment:
1) Pod yaml creation.
2) Running Pod
   kubectl create -f nginx.yaml
3) Commands for logs:
   kubectl logs kube-apiserver-kind-control-plane
   kubectl logs kube-controller-manager-kind-control-plane
   kubectl logs kube-proxy-6rpfn
   kubectl logs kube-scheduler-kind-control-plane
   journalctl -u kubelet