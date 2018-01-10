# siege4oss
siege performance tests for Object Storage Servers

## Prerequisites

Make sure the following are installed:
- siege
- parallel
- jq

## Read (GET) performance

1. Create sample blob 
```
$ dd if=/dev/urandom of=5MB bs=1048576 count=5
```

2. Create blobs on server

```
$ mc mb play/perf5mb
$ parallel mc cp 5MB play/perf5mb/5mb_{} ::: {1..100}
```

3. Create list of urls

```
$ mc share download --json play/perf5mb | jq -r '.share' > urls.txt
```

4. Test the server

```
$ siege -c 20 -i -b --time=10s --file="urls.txt" --log="siege.log"
```

Change number of concurrent clients via `-c` parameter and/or the duration of the test via `--time`.


## Write (POST) performance

```
$ mc share upload beast2/webserver/5MB.blob
```

```
$ siege -c 2 
```
