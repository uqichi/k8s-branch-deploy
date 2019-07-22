# sample-k8s
small sample app, which consists of some containers, for testing kubernetes (EKS).

Start in localhost:

```shell
make up
```

Append the following line to `/etc/hosts`:

```
127.0.0.1 apple.hoge-dev.myproj.com
```

```shell
open http://apple.hoge-dev.myproj.com:8080
```

Check the console log on browser and server log.
