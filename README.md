# Problem

Making a `ListObjectsV2` request to a local S3 bucket through Go AWS SDK V2 returns an empty response. The unexpected behavior is that LocalStack logs it as an `s3.ListBuckets` request, whereas it should log as `s3.ListObjectsV2`.

# Steps to reproduce

1. Run LocalStack
2. Create a local S3 bucket

```shell
aws --endpoint-url=http://localhost:4566 \
  s3api create-bucket \
  --bucket my-bucket \
  --region ap-northeast-2 \
  --create-bucket-configuration LocationConstraint=ap-northeast-2
```

3. Upload a file to the bucket

```shell
aws --endpoint-url=http://localhost:4566 s3 cp sample.jsonl s3://my-bucket/my-prefix/sample.jsonl
```

4. Run Go code

```shell
cd go
go run test.go
```

The result will be:

```shell
Total objects:  0
```

The log of LocalStack will be:

```shell
INFO --- [   asgi_gw_0] localstack.request.aws     : AWS s3.ListBuckets => 200
```

5. Run Python code (optional)

When making a request using Python code, the operation works as expected.

```shell
cd python
python test.py
```

The result will be:

```shell
Object name: my-prefix/sample.jsonl
Total objects: 1
```

The log of LocalStack will be:

```shell
INFO --- [   asgi_gw_1] localstack.request.aws     : AWS s3.ListObjectsV2 => 200
```

## Versions

LocalStack

```
localstack==2.3.2
localstack-client==2.3
localstack-core==2.3.2
localstack-ext==2.3.2
```

OS

```
macOS Monterey 12.6.3
```
