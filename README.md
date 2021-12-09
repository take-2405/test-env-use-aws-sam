# test-env-use-aws-sam
AWS  SAMを利用してローカルのテスト環境を構築する

## 実行方法
1. リポジトリをクローン
2. DynamoDB localの準備  

dockerでdynamodb localを起動
```
#(初回のみ)
aws configure set aws_access_key_id dummy     --profile local
aws configure set aws_secret_access_key dummy --profile local
aws configure set region ap-northeast-1       --profile local

docker network create lambda-local-test
docker run -d --network lambda-local-test --name dynamoTest -p 8000:8000 amazon/dynamodb-local:1.12.0 -jar DynamoDBLocal.jar -sharedDb

// 停止する場合
// docker stop dynamoTest
```
テーブルを作成  
 該当ファイル：./testdata/db-local.json

dynamodb-localにテーブルを作成  
```
aws dynamodb --profile local --endpoint-url http://localhost:8000 create-table --cli-input-json file://./testdata/db-local.json
```

3. SAMを起動  

```
sam local start-api --docker-network lambda-local-test
```

4. もう一つターミナルを立ち上げる

5. dynamodbの中身を確認  
```
aws dynamodb scan --table-name local_company_table --profile local --endpoint-url http://localhost:8000
{
    "Items": [],
    "Count": 0,
    "ScannedCount": 0,
    "ConsumedCapacity": null
}
```

6. lambdaにリクエスト送信(データ挿入)
```
curl localhost:3000/apigw
```
7. dynamodbの中身を確認
```
aws dynamodb scan --table-name local_company_table --profile local --endpoint-url http://localhost:8000
{
    "Items": [
        {
            "rate": {
                "N": 5467470
            },
            "timestamp": {
                "N": 1638997225
            }
        }
    ],
    "Count": 1,
    "ScannedCount": 1,
    "ConsumedCapacity": null
}
```
# 作成者  
若松丈人
