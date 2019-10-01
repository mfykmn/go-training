package main

import (
	"log"
	"math"
	"strings"

	"go-aws-mock"
	"strconv"
	//"time"
)

var ddb *mock.DynamoDB

func main() {
	sess := mock.InitAWSSession()
	ddb = mock.NewDynamoDB(sess)

	//ExampleDynamoDB_Tables()
	//ExampleDynamoDB_Put()
	//ExampleDynamoDB_PutOver400KB()
	//ExampleDynamoDB_Get()
	//ExampleDynamoDB_Scan()
	//ExampleDynamoDB_Delete()

	// *** ExampleDynamoDB_ScanOver1MB ***
	//ExampleDynamoDB_CreateTable()
	ExampleDynamoDB_ScanOver1MB()
	//ExampleDynamoDB_DeleteTable()
}

func createDummyData(createCount int, ddb *mock.DynamoDB) {
	mib := strings.Repeat("a", int(math.Exp2(12.0))) // 4096Byte
	log.Print(len(mib))
	log.Println("バイト")

	for i := 0; i < createCount; i++ {
		if err := ddb.Put("DummyPartition", "Dummy"+strconv.Itoa(i), mib); err != nil {
			log.Println(err)
			return
		}
	}
}

func ExampleDynamoDB_ScanOver1MB() {
	log.Println("*** Scan 1MB以上の動作確認 ***")
	ddb.SetInfoForMakingQuery("kamo", "DummyPartition", "DummySort")
	ddb.InfoForMakingQuery.Column = "ScanKey"

	res, exclusiveStartKey, err := ddb.Scan("DummyPartition", "ScanKey", nil)
	log.Println(res)
	log.Println(exclusiveStartKey)
	log.Println(err)

	//time.Sleep(3 * time.Second)

	res, exclusiveStartKey, err = ddb.Scan("DummyPartition", "ScanKey", exclusiveStartKey)
	log.Println(res)
	log.Println(exclusiveStartKey)
	log.Println(err)

}

func ExampleDynamoDB_Tables() {
	log.Println("*** Tables 確認***")
	lists, err := ddb.GetTables()
	log.Println(lists)
	log.Println(err)
}

func ExampleDynamoDB_Put() {
	log.Println("*** Put 動作確認 ***")

	ddb.SetInfoForMakingQuery("Music", "Artist", "SongTitle")
	ddb.InfoForMakingQuery.Column = "song"

	err := ddb.Put("Aerosmith", "Jaded", "Just Push Play")
	log.Println(err)
}

func ExampleDynamoDB_PutOver400KB() {
	log.Println("*** Put 400KB以上のエラー確認***")
	mib := strings.Repeat("a", int(math.Exp2(20.0)))
	log.Print(len(mib))
	log.Println("バイト")

	ddb.SetInfoForMakingQuery("Music", "Artist", "SongTitle")
	ddb.InfoForMakingQuery.Column = "song"

	err := ddb.Put("Journey", "Open Arms", mib)
	log.Println(err)
}

func ExampleDynamoDB_Get() {
	log.Println("*** Get 動作確認 ***")

	ddb.SetInfoForMakingQuery("Music", "Artist", "SongTitle")
	ddb.InfoForMakingQuery.Column = "song"

	song, err := ddb.Get("Aerosmith", "Jaded")
	log.Println(song)
	log.Println(err)
}

func ExampleDynamoDB_Scan() {
	log.Println("*** Scan 動作確認 ***")

	ddb.SetInfoForMakingQuery("Music", "Artist", "SongTitle")

	albumTitles, _, err := ddb.Scan("hoge", "SongTitle", nil)
	log.Println(albumTitles)
	log.Println(err)
}

func ExampleDynamoDB_Delete() {
	log.Println("*** Delete 動作確認 ***")

	ddb.SetInfoForMakingQuery("Music", "Artist", "SongTitle")
	ddb.InfoForMakingQuery.Column = "Artist"

	err := ddb.Delete("Aerosmith", "Jaded")
	log.Println(err)

}

func ExampleDynamoDB_CreateTable() {
	log.Println("*** CreateTable 動作確認 ***")

	ddb.SetInfoForMakingQuery("kamo", "DummyPartition", "DummySort")

	err := ddb.CreateTable()
	log.Println(err)

}

func ExampleDynamoDB_DeleteTable() {
	log.Println("*** DeleteTable 動作確認 ***")

	ddb.SetInfoForMakingQuery("kamo", "DummyPartition", "DummySort")

	err := ddb.DeleteTable()
	log.Println(err)

}
