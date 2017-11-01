# demo
 Demo test from Iris

### How to use
A web server project written by golang.

In a docker environment, pull the latest code to $GOPATH, and run below command to build a docker image
```console
>cd $GOPATH/demo
>docker build -t demo .
```
Then run below command check the image generated
```console
>docker image ls -a
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
demo                latest              1835adb7b8ae        14 minutes ago      235 MB
```
Start service by running demo docker image
```console
>docker run -p 5000:5000 demo
```
