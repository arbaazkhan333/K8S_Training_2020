Steps used in this assignment:
1) workload yamls creation.
2) Running workloads
   kubectl create -f pod/nginx_pod.yaml
   kubectl create -f deployment/nginx_deployment.yaml
   kubectl create -f replicationcontroller/replicationcontroller.yaml
   kubectl create -f replicationsets/rs.yaml
   kubectl create -f replicationsets/rs.yaml
   kubectl create -f jobs/job.yaml job.batch/ping-job created
   kubectl create -f statefulset/statefulset.yaml
   kubectl create -f cronjobs/cronjob.yaml
   
3) Commands for logs:
   kubectl logs kube-apiserver-kind-control-plane -n kube-system > kubeapi_log
   kubectl logs kube-controller-manager-kind-control-plane -n kube-system > kubecontoller_log
   kubectl logs kube-proxy-6rpfn -n kube-system > kubeproxy_log
   kubectl logs kube-scheduler-kind-control-plane -n kube-system > kubescheduler_log
   journalctl -u kubelet > kubelet_log