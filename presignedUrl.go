
package main

import (
	"fmt"
	"hash/crc32"
	"log"
	"time"
	"net/url"

	"github.com/minio/minio-go"
)

func hashOrder(key string, cardinality int) []int {
	if cardinality <= 0 {
		// Returns an empty int slice for cardinality < 0.
		return nil
	}

	nums := make([]int, cardinality)
	keyCrc := crc32.Checksum([]byte(key), crc32.IEEETable)

	start := (int(keyCrc % uint32(cardinality)) &^ 3) + 3
	for i := 1; i <= cardinality; i++ {
		nums[i-1] = 1 + ((start + i) % cardinality)
	}
	return nums
}


func main() {
	for i := 1; i <= 5000; i++ {
		objectName := fmt.Sprintf("test240mb.obj_%d", i)
		ho := hashOrder(objectName, 16)
		firstServer := int(0)
		if ho[0] == 1 {
			firstServer = 1
		} else if ho[4] == 1 {
			firstServer = 2
		} else if ho[8] == 1 {
			firstServer = 3 
		} else if ho[12] == 1 {
			firstServer = 4
		}
		//firstServer := ((ho[0] - 1) / 4) + 1
		// fmt.Println(i, ":", firstServer)
		presignedUrl(firstServer, "temp1", objectName)
	}

}

func presignedUrl(node int, bucketName, objectName string) {
	server := fmt.Sprintf("10.0.1.%d:9000", node)
	s3Client, err := minio.New(server, "minio", "minio123", false)
	if err != nil {
		log.Fatalln(err)
	}

	// Set request parameters
	reqParams := make(url.Values)

	// Generate presigned get object url.
	presignedURL, err := s3Client.PresignedGetObject(bucketName, objectName, time.Duration(24*3600)*time.Second, reqParams)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(presignedURL)
}

