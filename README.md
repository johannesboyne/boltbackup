#boltbackup.Do

A simple backup service for boltdb to S3.
![logo](boltbackup_logo.png)

##Usage

```go
db, err := bolt.Open("my2.db", 0600, nil)
if err != nil {
	log.Fatal(err)
}

cred     := aws.DefaultChainCredentials
s3Bucket := s3.New(&aws.Config{Region: "eu-central-1", Credentials: cred, LogLevel: 1})

boltbackup.Do(db, s3Bucket)
```

##License

MIT
