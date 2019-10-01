package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	dlp "cloud.google.com/go/dlp/apiv2"
	dlppb "google.golang.org/genproto/googleapis/privacy/dlp/v2"
)

func main() {
	ctx := context.Background()

	projectID := "be-kamono-mafuyu"

	// Creates a DLP client.
	client, err := dlp.NewClient(ctx)
	if err != nil {
		log.Fatalf("error creating DLP client: %v", err)
	}
	defer client.Close()

	infoTypes := []string{"PHONE_NUMBER", "EMAIL_ADDRESS"}
	message := `
TO:hoge@hoge.com
こんにちわ、瀧川です。

080-0000-0000 までご連絡ください。
`

	inspectString(projectID, client, infoTypes, message)
	fmt.Println("-----------------------------")
	masking(projectID, client, infoTypes, message)
}

func inspectString(projectID string, client *dlp.Client, infoTypes []string, input string) {
	// 検出機密種別をdlppb.InfoTypeに変換
	var i []*dlppb.InfoType
	for _, it := range infoTypes {
		i = append(i, &dlppb.InfoType{Name: it})
	}

	// リクエストのデータを構築
	req := &dlppb.InspectContentRequest{
		Parent: "projects/" + projectID, // 必ずprojects/を前につける
		InspectConfig: &dlppb.InspectConfig{ // 検出条件など
			InfoTypes: i,
		},
		Item: &dlppb.ContentItem{ // 検出対象
			DataItem: &dlppb.ContentItem_Value{
				Value: input,
			},
		},
	}

	resp, err := client.InspectContent(context.Background(), req) // 検出を実行
	if err != nil {
		log.Fatal(err)
	}

	// 以下は結果表示のみ
	resultFormatter := func(inputStr string, result *dlppb.Finding) string {
		tmp := `
機密種別: %s
もっともらしさ(尤度): %s
範囲: %d ~ %d
検出文字列: %s
      `
		start := result.GetLocation().GetCodepointRange().GetStart()
		end := result.GetLocation().GetCodepointRange().GetEnd()
		return fmt.Sprintf(tmp,
			result.GetInfoType().GetName(),
			result.GetLikelihood(),
			start,
			end,
			string([]rune(input)[start:end]))
	}

	fmt.Printf("対象文字列: %s\n", input)
	for _, result := range resp.GetResult().GetFindings() { // 検出されたものはFinding
		fmt.Println(resultFormatter(input, result))
	}
}

func masking(projectID string, client *dlp.Client, infoTypes []string, input string) {
	// 検出機密種別をdlppb.InfoTypeに変換
	var i []*dlppb.InfoType
	for _, it := range infoTypes {
		i = append(i, &dlppb.InfoType{Name: it})
	}

	// リクエストのデータを構築
	req := &dlppb.InspectContentRequest{
		Parent: "projects/" + projectID, // 必ずprojects/を前につける
		InspectConfig: &dlppb.InspectConfig{ // 検出条件など
			InfoTypes: i,
		},
		Item: &dlppb.ContentItem{ // 検出対象
			DataItem: &dlppb.ContentItem_Value{
				Value: input,
			},
		},
	}

	resp, err := client.InspectContent(context.Background(), req) // 検出を実行
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(input)
	fmt.Println(mask(input, resp.GetResult().GetFindings(), 0))
}

func mask(input string, results []*dlppb.Finding, index int) (maskedStr string) {
	if len(results) <= index {
		return input
	}

	target := getMaskTarget(input, results[index])
	maskedStr = strings.Replace(input, target, strings.Repeat("*", len(target)), -1)

	index++
	return mask(maskedStr, results, index)
}

func getMaskTarget(input string, result *dlppb.Finding) string {
	start := result.GetLocation().GetCodepointRange().GetStart()
	end := result.GetLocation().GetCodepointRange().GetEnd()
	return string([]rune(input)[start:end])
}
