# siege4oss
siege performance tests for Object Storage Servers

## Prerequisites

Make sure the following are installed:
- siege (***NB Make sure to use a version 3.0.x -- not the latest***)
- parallel
- mc
- jq

## Read (GET) performance

1. Create sample blob 
```
$ dd if=/dev/urandom of=5MB bs=1048576 count=5
```

2. Create blobs on server
```
$ mc mb beast2/perf5mb
$ parallel mc cp 5MB beast2/perf5mb/5mb_{} ::: {1..100}
```

3. Create list of urls
```
$ mc share download --json beast2/perf5mb | jq -r '.share' > urls.txt
```

4. Test the server
```
$ siege -c 25 -i -b --time=10s --file="urls.txt" --log="siege.log"
** SIEGE 3.0.8
** Preparing 25 concurrent users for battle.
The server is now under siege...
Lifting the server siege... done.

Transactions:		       19826 hits
Availability:		      100.00 %
Elapsed time:		        9.46 secs
Data transferred:	    99130.00 MB
Response time:		        0.01 secs
Transaction rate:	     2095.77 trans/sec
Throughput:		    10478.86 MB/sec
Concurrency:		       24.94
Successful transactions:       19826
Failed transactions:	           0
Longest transaction:	        0.03
Shortest transaction:	        0.00
```

Change number of concurrent clients via `-c` parameter and/or the duration of the test via `--time`.

5. Clean up (optional)
```
$ mc rm --recursive --force beast2/perf5mb
```

## Write (POST) performance

```
$ mc share upload beast2/webserver/5MB.blob
```

```
$ siege -c 2 
```
