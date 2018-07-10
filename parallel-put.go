/*
 * Minio Cloud Storage (C) 2017 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Change this value to test with a different object size.
const defaultObjectSize = 48 * 1024 * 1024

// Uploads all the inputs objects in parallel, upon any error this function panics.
func parallelUploads(objectNames []string, data []byte, conc int) {

	var wg sync.WaitGroup
	inputs := make(chan string)

	// Start one go routine per CPU
	for i := 0; i < conc; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			uploadWorker(inputs, data)
		}()
	}

	// Push objects onto input channel
	go func() {

		for _, objectName := range objectNames {
			inputs <- objectName
		}

		// Close input channel
		close(inputs)
	}()

	// Wait for workers to complete
	wg.Wait()
}

// Worker routine for uploading an object
func uploadWorker(inputs <-chan string, data []byte) {

	for name := range inputs {
		if err := uploadBlob(data, name); err != nil {
			panic(err)
		}
	}
}

// uploadBlob does an upload to the S3/Minio server
func uploadBlob(data []byte, objectName string) error {
	parts := strings.Split(objectName, "-")
	num, _ := strconv.Atoi(parts[2])
	endpoint := os.Getenv("ENDPOINT")
	if num % 2 == 0 {
		//endpoint = endpoint[:len(endpoint)-1] + "0"
	} else {
		//endpoint = endpoint[:len(endpoint)-1] + "1"
	}
	// fmt.Println(endpoint)
	credsUp := credentials.NewStaticCredentials(os.Getenv("ACCESSKEY"), os.Getenv("SECRETKEY"), "")
	sessUp := session.New(aws.NewConfig().
		WithCredentials(credsUp).
		WithRegion("us-east-1").
		WithEndpoint(endpoint).
		WithS3ForcePathStyle(true))

	uploader := s3manager.NewUploader(sessUp, func(u *s3manager.Uploader) {
		u.PartSize = 64 * 1024 * 1024 // 64MB per part
	})
	var err error
	_, err = uploader.Upload(&s3manager.UploadInput{
		Body:   bytes.NewReader(data),
		Bucket: aws.String(os.Getenv("BUCKET")),
		Key:    aws.String(objectName),
	})

	return err
}

var (
	objectSize = flag.Int("size", defaultObjectSize, "Size of the object to upload.")
)

func main() {
	flag.Parse()

	totalstr := os.Getenv("TOTAL")
	total, err := strconv.Atoi(totalstr)
	if err != nil {
		log.Fatalln(err)
	}

	concurrency := os.Getenv("CONCURRENCY")
	conc, err := strconv.Atoi(concurrency)
	if err != nil {
		log.Fatalln(err)
	}
	concurrencyEnd := os.Getenv("CONCURRENCY_END")
	concEnd, err := strconv.Atoi(concurrencyEnd)
	if err != nil {
		concEnd = conc
	}
	concurrencyIncr := os.Getenv("CONCURRENCY_INCR")
	concIncr, err := strconv.Atoi(concurrencyIncr)
	if err != nil {
		concIncr = 1
	}

	for c := conc; c <= concEnd; c += concIncr {
		if c > conc {
			fmt.Println("Sleeping...")
			time.Sleep(15 * time.Second)
		}
		fmt.Println("Testing with concurrency", c)
		run(total, c)
	}
}

func run(total, concurrency int) {

	utime := time.Now().UnixNano()
	var objectNames []string
	for i := 0; i < total; i++ {
		objectNames = append(objectNames, fmt.Sprintf("object-%d-%d", utime, i+1))
	}

	var data = bytes.Repeat([]byte("a"), *objectSize)

	start := time.Now().UTC()
	parallelUploads(objectNames, data, concurrency)

	totalSize := total * *objectSize
	elapsed := time.Since(start)
	fmt.Println("Elapsed time :", elapsed)
	seconds := float64(elapsed) / float64(time.Second)
	objsPerSec := float64(total)/seconds
	mbytePerSec := float64(totalSize)/seconds/1024/1024
	fmt.Printf("Speed        : %4.0f objs/sec\n", objsPerSec)
	fmt.Printf("Bandwidth    : %4.0f MByte/sec\n", mbytePerSec)
	fmt.Printf("%d, %f, %f, %f\n", concurrency, seconds, objsPerSec, mbytePerSec)
}
