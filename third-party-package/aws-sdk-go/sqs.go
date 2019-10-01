package mock

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

const queueUrl = ""

type SimpleQueueService struct {
	sqs *sqs.SQS
}

func NewSimpleQueueService(sess *session.Session) *SimpleQueueService {
	return &SimpleQueueService{sqs: sqs.New(sess)}
}

func (s *SimpleQueueService) SendMessage(message string) error {
	params := &sqs.SendMessageInput{
		MessageBody: aws.String(message),
		QueueUrl:    aws.String(queueUrl),
	}

	resp, err := s.sqs.SendMessage(params)
	if err != nil {
		return err
	}
	log.Println(resp)
	return nil
}

func (s *SimpleQueueService) ReceiveMessage() (string, error) {
	params := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueUrl),
		MaxNumberOfMessages: aws.Int64(1),
		WaitTimeSeconds:     aws.Int64(20),
	}

	resp, err := s.sqs.ReceiveMessage(params)
	if err != nil {
		return "", err
	}
	if len(resp.Messages) == 0 {
		return "", fmt.Errorf("wait time expaired")
	}

	msg := resp.Messages[0]
	if err := s.deleteMessage(msg); err != nil {
		return "", err
	}
	return *msg.Body, nil
}

func (s *SimpleQueueService) deleteMessage(msg *sqs.Message) error {
	params := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueUrl),
		ReceiptHandle: aws.String(*msg.ReceiptHandle),
	}

	if _, err := s.sqs.DeleteMessage(params); err != nil {
		return err
	}
	return nil
}
