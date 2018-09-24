## sample-pod.yaml

### ポッドYAMLを書いて、ポッド作成

```shell
$ kubectl apply -f sample-pod.yaml 
pod "go-webapi-for-gke-study-pod" created
$
$ kubectl get pods
NAME                          READY     STATUS    RESTARTS   AGE
go-webapi-for-gke-study-pod   1/1       Running   0          3m
```

### ポッドYAMLを修正して、ポッド更新

```shell
$ kubectl apply -f sample-pod.yaml 
pod "go-webapi-for-gke-study-pod" configured
$
$ kubectl get pods
NAME                          READY     STATUS    RESTARTS   AGE
go-webapi-for-gke-study-pod   1/1       Running   1          9m
```

### ポッドを即時強制削除

```shell
$ kubectl delete -f sample-pod.yaml --grace-period 0 --force
warning: Immediate deletion does not wait for confirmation that the running resource has been terminated. The resource may continue to run on the cluster indefinitely.
pod "go-webapi-for-gke-study-pod" deleted
$
$ kubectl get pods
No resources found.
```

## sample-multi-resource-manifest.yaml

### デプロイメントとサービスを同じYAMLに書いて適用

```shell
$ kubectl apply -f sample-multi-resource-manifest.yaml
deployment.apps "order1-deployment" created
service "order-2-deployment" created
$
$ kubectl get deploy
NAME                DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
order1-deployment   3         3         3            0           22s
```

### デプロイメントに誤りがあったので修正して再度適用

```shell
$ kubectl apply -f sample-multi-resource-manifest.yaml
deployment.apps "order1-deployment" configured
service "order-2-deployment" unchanged
$
$ kubectl get deploy
NAME                DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
order1-deployment   3         4         3            2           1m
$ kubectl get deploy
NAME                DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
order1-deployment   3         3         3            3           1m
$
$ kubectl get svc
NAME                 TYPE           CLUSTER-IP     EXTERNAL-IP     PORT(S)          AGE
order-2-deployment   LoadBalancer   10.63.252.74   【IPアドレス】   8080:30947/TCP   1m
$
$ kubectl get pods
NAME                                 READY     STATUS    RESTARTS   AGE
order1-deployment-585885f756-b5dlk   1/1       Running   0          9m
order1-deployment-585885f756-dr7m2   1/1       Running   0          9m
order1-deployment-585885f756-vkkb2   1/1       Running   0          9m
```
