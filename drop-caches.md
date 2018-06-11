# drop-caches

## Prerequisites

Make sure golang is installed on the server.

## Purpose

Drop the caches on a linux server on a regular basis at a specific point in time (40 seconds after start). Change the timings and overall number of runs accordingly to meet your needs.

Run in sync on server and on client, eg. on client run `siege` test for 30 seconds combined with 30 seconds of rest between invocations. Then drop caches on server at 40 seconds past the start of each run (taking approx 10 seconds).

## Execute

```
$ go run drop-caches.go
```
