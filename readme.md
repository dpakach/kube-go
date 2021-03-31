## kube go
Testing out deploying simple go application in local kubernetes cluster with minikube

### Commands available
- `make build` - Build the go binary
- `make docker-build` - Build the docker image for the app
- `make docker-push` - Push docker image to docker hub (change the image name to push to your docker hub)
- `make docker-run` - Run the docker image


### Running the app in kubernetes with minikube
To run the app in kubernetes cluster use following commands
- First start local kubernetes cluster with minikube
```
minikube start
```

- Then create the configmap and secret
```
kubectl apply -f kube-go-secret.yaml
kubectl apply -f kube-go-config.yaml
```

- Then Start the deployment with
```
kubectl apply -f kube-go-deployment.yaml
```
This will run the app in kubernetes cluster with 2 replicas. 
In order to change it, modify `kube-go-deployment.yaml` and rerun the above command.

- To access the server from the browser start the service with
```
minikube service kube-go-service
```
minikube will automatically start your browser in the correct ip and port

### Minikube
#### Installing

https://minikube.sigs.k8s.io/docs/start/

#### Commands
Some of the commands available in minikube
```
minikube start
minikube status
minikube service kube-go-service
```


### Kubectl

#### Installing
https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/

#### commands
Some of the commands available in kubectl
```
kubectl get deployments
kubectl create deployment <deployment-name> --image=<image> [options]
kubectl delete deployment <deployment-name>

kubectl edit deployment <deployment-name>

kubectl get pods
kubectl delete pods

kubectl get replicaset
kubectl delete replicaset

kubectl get services
kubectl create services
kubectl delete services

kubectl logs <pod-name>

kubectl exec -it <pod-name> -- <command>

kubectl apply -f <config-file>
```
