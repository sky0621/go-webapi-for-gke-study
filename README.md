# go-webapi-for-gke-study

## env

### # OS
<pre>
$ cat /etc/os-release 
NAME="Ubuntu"
VERSION="17.10 (Artful Aardvark)"
</pre>

### # Golang
<pre>
$ go version
go version go1.10 linux/amd64
</pre>

### # Docker
<pre>
$ sudo docker version
[sudo] koge のパスワード: 
Client:
 Version:           18.06.1-ce
 API version:       1.38
 Go version:        go1.10.3
 Git commit:        e68fc7a
 Built:             Tue Aug 21 17:25:03 2018
 OS/Arch:           linux/amd64
 Experimental:      false

Server:
 Engine:
  Version:          18.06.1-ce
  API version:      1.38 (minimum version 1.12)
  Go version:       go1.10.3
  Git commit:       e68fc7a
  Built:            Tue Aug 21 17:23:27 2018
  OS/Arch:          linux/amd64
  Experimental:     false
</pre>

### # gcloud
<pre>
$ gcloud version
Google Cloud SDK 216.0.0
app-engine-go 
app-engine-java 1.9.64
app-engine-python 1.9.75
beta 2018.07.16
bigtable 
bq 2.0.34
cbt 
cloud-build-local 
cloud-datastore-emulator 2.0.2
cloud_sql_proxy 
container-builder-local 
core 2018.09.07
docker-credential-gcr 
gsutil 4.33
kubectl 2018.08.17
pubsub-emulator 2018.02.02
</pre>

## web framework

### echo
https://echo.labstack.com/
