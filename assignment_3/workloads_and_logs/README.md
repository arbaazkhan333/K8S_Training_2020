Steps used in this assignment:<BR/>
1) workload yamls creation.<BR/>
2) Running workloads<BR/>
   kubectl create -f pod/nginx_pod.yaml<BR/>
   kubectl create -f deployment/nginx_deployment.yaml<BR/>
   kubectl create -f replicationcontroller/replicationcontroller.yaml<BR/>
   kubectl create -f replicationsets/rs.yaml<BR/>
   kubectl create -f replicationsets/rs.yaml<BR/>
   kubectl create -f jobs/job.yaml job.batch/ping-job created<BR/>
   kubectl create -f statefulset/statefulset.yaml<BR/>
   kubectl create -f cronjobs/cronjob.yaml<BR/>
   <BR/>
3) Commands for logs:<BR/>
   kubectl logs kube-apiserver-kind-control-plane -n kube-system > kubeapi_log<BR/>
   kubectl logs kube-controller-manager-kind-control-plane -n kube-system > kubecontoller_log<BR/>
   kubectl logs kube-proxy-6rpfn -n kube-system > kubeproxy_log<BR/>
   kubectl logs kube-scheduler-kind-control-plane -n kube-system > kubescheduler_log<BR/>
   journalctl -u kubelet > kubelet_log
