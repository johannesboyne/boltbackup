package boltbackup

import (
	"bufio"
	"bytes"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/s3"
	"github.com/boltdb/bolt"
)

func Do(db *bolt.DB, s3Bucket *s3.S3, bucketname string) error {
	err := db.View(func(tx *bolt.Tx) error {
		var b bytes.Buffer
		buffInMem := bufio.NewWriter(&b)
		_, err := tx.WriteTo(buffInMem)
		backupdate := time.Now()
		params := &s3.PutObjectInput{
			Bucket:               aws.String(bucketname),
			Key:                  aws.String(backupdate.Format("2006_01_02_15-04-05")),
			Body:                 bytes.NewReader(b.Bytes()),
			ACL:                  aws.String("authenticated-read"),
			ServerSideEncryption: aws.String("AES256"),
			Metadata: &map[string]*string{
				"bolt-backup": aws.String(backupdate.Format("2006/01/02/15:04:05")),
			},
		}
		_, err = s3Bucket.PutObject(params)
		return err
	})
	return err
}
