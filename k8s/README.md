## ポッドYAMLを書いて、ポッド作成

```shell
$ kubectl apply -f sample-pod.yaml 
pod "go-webapi-for-gke-study-pod" created
$
$ kubectl get pods
NAME                          READY     STATUS    RESTARTS   AGE
go-webapi-for-gke-study-pod   1/1       Running   0          3m
```

## ポッドYAMLを修正して、ポッド更新

```shell
$ kubectl apply -f sample-pod.yaml 
pod "go-webapi-for-gke-study-pod" configured
$
$ kubectl get pods
NAME                          READY     STATUS    RESTARTS   AGE
go-webapi-for-gke-study-pod   1/1       Running   1          9m
```

## ポッドを即時強制削除

```shell
$ kubectl delete -f sample-pod.yaml --grace-period 0 --force
warning: Immediate deletion does not wait for confirmation that the running resource has been terminated. The resource may continue to run on the cluster indefinitely.
pod "go-webapi-for-gke-study-pod" deleted
$
$ kubectl get pods
No resources found.
```
