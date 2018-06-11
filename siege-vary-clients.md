# siege-vary-clients

## Prerequisites

Make sure golang is installed on the server.

## Edit siege-vary-clients.go

Edit `siege-vary-clients.go` as necessary to specify the correct range (line 41) and the correct input file with the presigned urls (line 23). 

## Execute range of tests

```
$ go run siege-vary-clients.go
```

## Results

The results are written to the `siege.log` file.
