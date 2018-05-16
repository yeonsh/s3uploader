# S3 uploader

This program uploads file(s) to S3.

```
$ ./s3uploader
NOTE: bucket and file name required

Usage: ./s3uploader [options] bucket_name filename

  -key string
    	Access Key ID
  -no-cache
    	Set no-cache to Cache Control
  -region string
    	Region (default "us-east-1")
  -secret string
    	Secret Access Key
```

# How to build

Download and install Go here: https://golang.org/

Install Go dep following instructions here: https://golang.github.io/dep/docs/installation.html

Clone this repository under $GOPATH/src.

Go into the cloned folder.

Install dependencies into `vendor` folder.

`$ dep ensure`

Build

`$ go build`

