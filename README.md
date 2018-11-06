# proposal-istio-grpc

## Run
### Minikube
[Istio/Installation with Helm](https://istio.io/docs/setup/kubernetes/helm-install/)

```bash
$ curl -L https://git.io/getLatestIstio | sh -

$ minikube start --memory=8192 --cpus=4 --kubernetes-version=v1.10.0 --vm-driver=virtualbox
```

### Istio
```bash
$ watch kubectl get all
$ watch kubectl get all -n kube-system
$ watch kubectl get all -n istio-system
```

```bash
$ kubectl apply -f istio-1.0.2/install/kubernetes/helm/helm-service-account.yaml

$ helm init --service-account tiller

$ helm install istio-1.0.2/install/kubernetes/helm/istio --name istio --namespace istio-system --set grafana.enabled=true
```

### Grafana
```bash
$ kubectl get svc grafana -n istio-system -o='jsonpath={.spec.ports}'
[map[name:http protocol:TCP port:3000 targetPort:3000]]
$ kubectl -n istio-system get pod -l app=grafana -o jsonpath='{.items[0].metadata.name}'
grafana-75485f89b9-pb4mz
$ kubectl -n istio-system port-forward $(kubectl -n istio-system get pod -l app=grafana -o jsonpath='{.items[0].metadata.name}') 3000:3000
...
$ open http://localhost:3000
```

### Bookinfo Sample
```bash
$ kubectl label namespace default istio-injection=enabled
$ kubectl apply -f istio-1.0.2/samples/bookinfo/platform/kube/bookinfo.yaml
$ kubectl apply -f istio-1.0.2/samples/bookinfo/networking/bookinfo-gateway.yaml
```

```bash
$ minikube ip
192.168.99.100
$ kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].nodePort}'
31380
$ open http://192.168.99.100:31380/productpage
```

### Example Ping Service
```bash
$ kubectl label namespace default istio-injection=enabled

$ kubectl apply -f ping.yaml
```

```bash
$ minikube ip
192.168.99.100
$ kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].nodePort}'
31380
$ example-ping-service-client -h 192.168.99.100 -p 31380
```

### Troubleshooting
#### kube-controller-manager-minikube CrashLoopBackOff
[Follow these instructions to prepare Minikube for Istio](https://istio.io/docs/setup/kubernetes/platform-setup/minikube/)
```
$ minikube start --memory=8192 --cpus=4 --kubernetes-version=v1.10.0 \
    --extra-config=controller-manager.cluster-signing-cert-file="/var/lib/localkube/certs/ca.crt" \
    --extra-config=controller-manager.cluster-signing-key-file="/var/lib/localkube/certs/ca.key" \
    --vm-driver=virtualbox

$ kubectl get pods --all-namespaces
NAMESPACE     NAME                               READY   STATUS             RESTARTS   AGE
kube-system   etcd-minikube                      1/1     Running            0          20s
kube-system   kube-addon-manager-minikube        1/1     Running            0          1m
kube-system   kube-apiserver-minikube            1/1     Running            0          42s
kube-system   kube-controller-manager-minikube   0/1     CrashLoopBackOff   2          1m
kube-system   kube-scheduler-minikube            1/1     Running            0          30s
```

Fix kube-controller-manager-minikube status is _CrashLoopBackOff_ or _Error_.

```
minikube start --memory=8192 --cpus=4 --kubernetes-version=v1.10.0 --vm-driver=virtualbox
```
