Commands used:<BR/>
1) Create dev context:<BR/>
kubectl config set-context dev --namespace=dev --user=kind-kind --cluster=kind-kind<BR/>
2) Create qa context<BR/>
kubectl config set-context qa --namespace=qa --user=kind-kind --cluster=kind-kind<BR/>
3) Switch to dev context, create a pod and list it.<BR/>
kubectl config use-context dev<BR/>
kubectl apply -f <pod_yaml><BR/>
kubectl get pods (new pod is visible)<BR/>
4) Switch to qa context and list pods<BR/>
kubectl config use-context qa<BR/>
kubectl get pods   (pod deployed in dev context not visible)
