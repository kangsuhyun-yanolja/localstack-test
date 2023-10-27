import boto3


def main():
    s3 = boto3.client(
        "s3",
        endpoint_url="http://localhost.localstack.cloud:4566",
        region_name="ap-northeast-2",
        aws_access_key_id="dummy",
        aws_secret_access_key="dummy",
    )

    bucket_name = "my-bucket"
    prefix = "my-prefix/"

    response = s3.list_objects_v2(Bucket=bucket_name, Prefix=prefix)

    if "Contents" in response:
        for content in response["Contents"]:
            print(f"Object name: {content['Key']}")

        print(f"Total objects: {len(response['Contents'])}")
        # Total objects: 1

    else:
        print("No objects found.")


if __name__ == "__main__":
    main()
