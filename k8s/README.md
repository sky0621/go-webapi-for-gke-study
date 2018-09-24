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

### リソース使用量の確認

```shell
$ kubectl top node
NAME                                              CPU(cores)   CPU%      MEMORY(bytes)   MEMORY%
gke-my-cluster-1-min-default-pool-ea807978-8d2t   31m          3%        383Mi           64%
gke-my-cluster-1-min-default-pool-ea807978-hpqc   40m          4%        416Mi           70%
gke-my-cluster-1-min-default-pool-ea807978-ss7p   33m          3%        402Mi           67%
```

### Pod上でのコマンド実行

```shell
$ kubectl exec -it sample-label1 /bin/sh
#
```

### Podのログ確認

```shell
$ kubectl logs order1-deployment-bbd48b68f-nxld6

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v3.3.6
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:80
{"level":"info","ts":1537795575.036393,"caller":"go-webapi-for-gke-study/main.go:36","msg":"INFO LEVEL with severity","severity":"INFO"}
{"level":"warn","ts":1537795575.0502615,"caller":"go-webapi-for-gke-study/main.go:37","msg":"WARN LEVEL with severity","severity":"WARN"}
{"level":"error","ts":1537795575.0502732,"caller":"go-webapi-for-gke-study/main.go:38","msg":"ERROR LEVEL with severity","severity":"ERROR","stacktrace":"main.main.func1\n\t/go/src/github.com/sky0621/go-webapi-for-gke-study/main.go:38\ngithub.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo.(*Echo).Add.func1\n\t/go/src/github.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/echo.go:480\ngithub.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/middleware.RequestIDWithConfig.func1.1\n\t/go/src/github.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/middleware/request_id.go:57\ngithub.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/middleware.RecoverWithConfig.func1.1\n\t/go/src/github.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/middleware/recover.go:78\ngithub.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo.(*Echo).ServeHTTP\n\t/go/src/github.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/echo.go:583\nnet/http.serverHandler.ServeHTTP\n\t/usr/local/go/src/net/http/server.go:2619\nnet/http.(*conn).serve\n\t/usr/local/go/src/net/http/server.go:1801"}
{"level":"info","ts":1537795772.0383852,"caller":"go-webapi-for-gke-study/main.go:36","msg":"INFO LEVEL with severity","severity":"INFO"}
{"level":"warn","ts":1537795772.0384548,"caller":"go-webapi-for-gke-study/main.go:37","msg":"WARN LEVEL with severity","severity":"WARN"}
{"level":"error","ts":1537795772.038462,"caller":"go-webapi-for-gke-study/main.go:38","msg":"ERROR LEVEL with severity","severity":"ERROR","stacktrace":"main.main.func1\n\t/go/src/github.com/sky0621/go-webapi-for-gke-study/main.go:38\ngithub.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo.(*Echo).Add.func1\n\t/go/src/github.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/echo.go:480\ngithub.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/middleware.RequestIDWithConfig.func1.1\n\t/go/src/github.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/middleware/request_id.go:57\ngithub.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/middleware.RecoverWithConfig.func1.1\n\t/go/src/github.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/middleware/recover.go:78\ngithub.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo.(*Echo).ServeHTTP\n\t/go/src/github.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/echo.go:583\nnet/http.serverHandler.ServeHTTP\n\t/usr/local/go/src/net/http/server.go:2619\nnet/http.(*conn).serve\n\t/usr/local/go/src/net/http/server.go:1801"}
{"level":"info","ts":1537795773.4908075,"caller":"go-webapi-for-gke-study/main.go:36","msg":"INFO LEVEL with severity","severity":"INFO"}
{"level":"warn","ts":1537795773.49086,"caller":"go-webapi-for-gke-study/main.go:37","msg":"WARN LEVEL with severity","severity":"WARN"}
{"level":"error","ts":1537795773.4908679,"caller":"go-webapi-for-gke-study/main.go:38","msg":"ERROR LEVEL with severity","severity":"ERROR","stacktrace":"main.main.func1\n\t/go/src/github.com/sky0621/go-webapi-for-gke-study/main.go:38\ngithub.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo.(*Echo).Add.func1\n\t/go/src/github.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/echo.go:480\ngithub.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/middleware.RequestIDWithConfig.func1.1\n\t/go/src/github.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/middleware/request_id.go:57\ngithub.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/middleware.RecoverWithConfig.func1.1\n\t/go/src/github.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/middleware/recover.go:78\ngithub.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo.(*Echo).ServeHTTP\n\t/go/src/github.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/echo.go:583\nnet/http.serverHandler.ServeHTTP\n\t/usr/local/go/src/net/http/server.go:2619\nnet/http.(*conn).serve\n\t/usr/local/go/src/net/http/server.go:1801"}
{"level":"info","ts":1537795774.7320976,"caller":"go-webapi-for-gke-study/main.go:36","msg":"INFO LEVEL with severity","severity":"INFO"}
{"level":"warn","ts":1537795774.7321355,"caller":"go-webapi-for-gke-study/main.go:37","msg":"WARN LEVEL with severity","severity":"WARN"}
{"level":"error","ts":1537795774.7321706,"caller":"go-webapi-for-gke-study/main.go:38","msg":"ERROR LEVEL with severity","severity":"ERROR","stacktrace":"main.main.func1\n\t/go/src/github.com/sky0621/go-webapi-for-gke-study/main.go:38\ngithub.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo.(*Echo).Add.func1\n\t/go/src/github.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/echo.go:480\ngithub.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/middleware.RequestIDWithConfig.func1.1\n\t/go/src/github.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/middleware/request_id.go:57\ngithub.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/middleware.RecoverWithConfig.func1.1\n\t/go/src/github.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/middleware/recover.go:78\ngithub.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo.(*Echo).ServeHTTP\n\t/go/src/github.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/echo.go:583\nnet/http.serverHandler.ServeHTTP\n\t/usr/local/go/src/net/http/server.go:2619\nnet/http.(*conn).serve\n\t/usr/local/go/src/net/http/server.go:1801"}
{"level":"info","ts":1537795776.0714352,"caller":"go-webapi-for-gke-study/main.go:36","msg":"INFO LEVEL with severity","severity":"INFO"}
{"level":"warn","ts":1537795776.0715005,"caller":"go-webapi-for-gke-study/main.go:37","msg":"WARN LEVEL with severity","severity":"WARN"}
{"level":"error","ts":1537795776.071509,"caller":"go-webapi-for-gke-study/main.go:38","msg":"ERROR LEVEL with severity","severity":"ERROR","stacktrace":"main.main.func1\n\t/go/src/github.com/sky0621/go-webapi-for-gke-study/main.go:38\ngithub.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo.(*Echo).Add.func1\n\t/go/src/github.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/echo.go:480\ngithub.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/middleware.RequestIDWithConfig.func1.1\n\t/go/src/github.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/middleware/request_id.go:57\ngithub.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/middleware.RecoverWithConfig.func1.1\n\t/go/src/github.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/middleware/recover.go:78\ngithub.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo.(*Echo).ServeHTTP\n\t/go/src/github.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/echo.go:583\nnet/http.serverHandler.ServeHTTP\n\t/usr/local/go/src/net/http/server.go:2619\nnet/http.(*conn).serve\n\t/usr/local/go/src/net/http/server.go:1801"}
{"level":"info","ts":1537795777.3998318,"caller":"go-webapi-for-gke-study/main.go:36","msg":"INFO LEVEL with severity","severity":"INFO"}
{"level":"warn","ts":1537795777.3998837,"caller":"go-webapi-for-gke-study/main.go:37","msg":"WARN LEVEL with severity","severity":"WARN"}
{"level":"error","ts":1537795777.3998911,"caller":"go-webapi-for-gke-study/main.go:38","msg":"ERROR LEVEL with severity","severity":"ERROR","stacktrace":"main.main.func1\n\t/go/src/github.com/sky0621/go-webapi-for-gke-study/main.go:38\ngithub.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo.(*Echo).Add.func1\n\t/go/src/github.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/echo.go:480\ngithub.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/middleware.RequestIDWithConfig.func1.1\n\t/go/src/github.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/middleware/request_id.go:57\ngithub.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/middleware.RecoverWithConfig.func1.1\n\t/go/src/github.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/middleware/recover.go:78\ngithub.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo.(*Echo).ServeHTTP\n\t/go/src/github.com/sky0621/go-webapi-for-gke-study/vendor/github.com/labstack/echo/echo.go:583\nnet/http.serverHandler.ServeHTTP\n\t/usr/local/go/src/net/http/server.go:2619\nnet/http.(*conn).serve\n\t/usr/local/go/src/net/http/server.go:1801"}
```
