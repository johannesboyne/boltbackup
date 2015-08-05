package main

import (
	"log"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/s3"
	"github.com/boltdb/bolt"
	"github.com/johannesboyne/boltbackup"
)

func main() {
	db, err := bolt.Open("my2.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	cred := aws.DefaultChainCredentials
	s3Bucket := s3.New(&aws.Config{Region: "eu-central-1", Credentials: cred, LogLevel: 1})
	err = boltbackup.Do(db, s3Bucket)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
