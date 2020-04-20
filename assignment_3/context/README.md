Commands used:
1) Create dev context:
kubectl config set-context dev --namespace=dev --user=kind-kind --cluster=kind-kind
2) Create qa context
kubectl config set-context qa --namespace=qa --user=kind-kind --cluster=kind-kind
3) Switch to dev context, create a pod and list it.
kubectl config use-context dev
kubectl apply -f <pod_yaml>
kubectl get pods (new pod is visible)
4) Switch to qa context and list pods
kubectl config use-context qa
kubectl get pods   (pod deployed in dev context not visible)