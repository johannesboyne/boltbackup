#boltbackup.Do

A simple backup service for [BoltDB](https://github.com/boltdb/bolt) backups to Amazon [S3](https://github.com/aws/aws-sdk-go).
![logo](boltbackup_logo.png)

##Usage

```go
db, err := bolt.Open("my2.db", 0600, nil)
if err != nil {
	log.Fatal(err)
}

cred     := aws.DefaultChainCredentials
s3Bucket := s3.New(&aws.Config{Region: "eu-central-1", Credentials: cred, LogLevel: 1})

boltbackup.Do(db, s3Bucket, "bucketname")
```

##License

MIT
