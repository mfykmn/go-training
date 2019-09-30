package mock

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type InfoForMakingQuery struct {
	Table               string
	PrimaryPartitionKey string
	PrimarySortKey      string
	Column              string // TODO
}

func NewInfoForMakingQuery(table, primaryPartitionKey, primarySortKey string) *InfoForMakingQuery {
	return &InfoForMakingQuery{
		Table:               table,
		PrimaryPartitionKey: primaryPartitionKey,
		PrimarySortKey:      primarySortKey,
	}
}

type DynamoDB struct {
	dynamoDB           *dynamodb.DynamoDB
	InfoForMakingQuery *InfoForMakingQuery
}

func NewDynamoDB(sess *session.Session) *DynamoDB {
	return &DynamoDB{dynamoDB: dynamodb.New(sess)}
}

func (d *DynamoDB) GetTables() ([]string, error) {
	result, err := d.dynamoDB.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		return nil, err
	}

	var tables []string
	for _, table := range result.TableNames {
		tables = append(tables, *table)
	}
	return tables, nil
}

func (d *DynamoDB) SetInfoForMakingQuery(table, primaryPartitionKey, primarySortKey string) {
	d.InfoForMakingQuery = NewInfoForMakingQuery(table, primaryPartitionKey, primarySortKey)
}

func (d *DynamoDB) Scan(primaryPartitionVal, targetKey string, exclusiveStartKey map[string]*dynamodb.AttributeValue) ([]string, map[string]*dynamodb.AttributeValue, error) {
	params := &dynamodb.ScanInput{
		TableName: aws.String(d.InfoForMakingQuery.Table),

		ProjectionExpression: aws.String(targetKey),

		ConsistentRead: aws.Bool(true),

		//ExpressionAttributeNames: map[string]*string{
		//	"Key": aws.String(d.InfoForMakingQuery.PrimaryPartitionKey),
		//},

		//ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
		//	"Key": { // Required
		//		S:    aws.String(primaryPartitionVal),
		//	},
		//},

		// Scanの開始位置
		ExclusiveStartKey: exclusiveStartKey,

		//Segment:       aws.Int64(1),
		//Select:        aws.String("Select"),
		//TotalSegments: aws.Int64(1),
	}

	resp, err := d.dynamoDB.Scan(params)
	if err != nil {
		return nil, nil, err
	}

	var resList []string
	for _, v := range resp.Items {
		resList = append(resList, *v[targetKey].S)
	}

	if resp.LastEvaluatedKey == nil {
		return resList, nil, nil
	}

	return resList, resp.LastEvaluatedKey, nil //TODO
}

func (d *DynamoDB) Get(primaryPartitionVal, primarySortVal string) (string, error) {
	params := &dynamodb.GetItemInput{
		TableName: aws.String(d.InfoForMakingQuery.Table),

		Key: map[string]*dynamodb.AttributeValue{
			d.InfoForMakingQuery.PrimaryPartitionKey: {
				S: aws.String(primaryPartitionVal),
			},
			d.InfoForMakingQuery.PrimarySortKey: {
				S: aws.String(primarySortVal),
			},
		},

		AttributesToGet: []*string{
			aws.String(d.InfoForMakingQuery.Column),
		},

		ConsistentRead:         aws.Bool(true),
		ReturnConsumedCapacity: aws.String("NONE"),
	}

	resp, err := d.dynamoDB.GetItem(params)
	if err != nil {
		return "", err
	}

	if resp.Item[d.InfoForMakingQuery.Column] != nil {
		return *resp.Item[d.InfoForMakingQuery.Column].S, nil
	} else {
		return "", fmt.Errorf("failed Get")
	}
}

func (d *DynamoDB) Put(primaryPartitionVal, primarySortVal, columnVal string) error {
	params := &dynamodb.PutItemInput{
		TableName: aws.String(d.InfoForMakingQuery.Table),

		Item: map[string]*dynamodb.AttributeValue{
			d.InfoForMakingQuery.PrimaryPartitionKey: {
				S: aws.String(primaryPartitionVal),
			},
			d.InfoForMakingQuery.PrimarySortKey: {
				S: aws.String(primarySortVal),
			},
			d.InfoForMakingQuery.Column: {
				S: aws.String(columnVal),
			},
		},
	}

	// PutItemの実行
	if _, err := d.dynamoDB.PutItem(params); err != nil {
		return err
	}
	return nil
}

func (d *DynamoDB) Delete(primaryPartitionVal, primarySortVal string) error {
	log.Println(d.InfoForMakingQuery.Column)
	params := &dynamodb.DeleteItemInput{
		TableName: aws.String(d.InfoForMakingQuery.Table),

		Key: map[string]*dynamodb.AttributeValue{
			d.InfoForMakingQuery.PrimaryPartitionKey: {
				S: aws.String(primaryPartitionVal),
			},
			d.InfoForMakingQuery.PrimarySortKey: {
				S: aws.String(primarySortVal),
			},
		},

		//返ってくるデータの種類
		//ReturnConsumedCapacity: aws.String("NONE"),
		//ReturnItemCollectionMetrics: aws.String("NONE"),
		//ReturnValues: aws.String("NONE"),

	}

	_, err := d.dynamoDB.DeleteItem(params)
	if err != nil {
		return err
	}

	return nil
}

func (d *DynamoDB) CreateTable() error {
	params := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String(d.InfoForMakingQuery.PrimaryPartitionKey),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String(d.InfoForMakingQuery.PrimarySortKey),
				AttributeType: aws.String("S"),
			},
		},

		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String(d.InfoForMakingQuery.PrimaryPartitionKey),
				KeyType:       aws.String("HASH"), // HASH=ハッシュキー
			},
			{
				AttributeName: aws.String(d.InfoForMakingQuery.PrimarySortKey),
				KeyType:       aws.String("RANGE"), // RANGE=レンジキー
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1), // 読み込みスループット
			WriteCapacityUnits: aws.Int64(1), // 書き込みスループット
		},
		TableName: aws.String(d.InfoForMakingQuery.Table), // テーブル名
	}

	_, err := d.dynamoDB.CreateTable(params)
	if err != nil {
		return err
	}

	return nil
}

func (d *DynamoDB) DeleteTable() error {
	params := &dynamodb.DeleteTableInput{
		TableName: aws.String(d.InfoForMakingQuery.Table),
	}
	_, err := d.dynamoDB.DeleteTable(params)
	if err != nil {
		return err
	}

	return nil
}
