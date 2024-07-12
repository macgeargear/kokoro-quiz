package quiz

import (
	proto "github.com/macgeargear/kokoro-go-proto/proto/quiz/quiz/v1"
)

type Service interface {
	proto.QuizServiceServer
}

type serviceImpl struct {
	proto.UnimplementedQuizServiceServer
	repo Repository
}

func NewService(repo Repository) Service {
	return &serviceImpl{repo: repo}
}
