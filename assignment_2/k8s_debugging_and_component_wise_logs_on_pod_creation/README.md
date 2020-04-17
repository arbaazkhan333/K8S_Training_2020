Steps used in this assignment:
1) Pod yaml creation.
2) Running Pod
   kubectl create -f nginx.yaml
3) Commands for logs:
   kubectl logs kube-apiserver-kind-control-plane -n kube-system > kubeapi_log
   kubectl logs kube-controller-manager-kind-control-plane -n kube-system > kubecontoller_log
   kubectl logs kube-proxy-6rpfn -n kube-system > kubeproxy_log
   kubectl logs kube-scheduler-kind-control-plane -n kube-system > kubescheduler_log
   journalctl -u kubelet > kubelet_log