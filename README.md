# siege4oss
siege performance tests for Object Storage Servers

## GET performance testing

```
$ seq -w 0 10 | parallel dd if=/dev/urandom of=5MB_{} bs=1048576 count=5 :::
```

```
$ mc share upload beast2/webserver/5MB.blob
```

```
$ siege -c 2 
```
