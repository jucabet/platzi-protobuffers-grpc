package repository

import (
	"context"

	"github.com/jucabet/platzi-protobuffers-grpc/tree/main/models"
)

type Repository interface {
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	SetStudent(ctx context.Context, student *models.Student) error
	GetTest(ctx context.Context, id string) (*models.Test, error)
	SetTest(ctx context.Context, test *models.Test) error
	SetQuestion(ctx context.Context, question *models.Question) error
	GetQuestion(ctx context.Context, id string) (*models.Question, error)
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func SetStudent(ctx context.Context, studen *models.Student) error {
	return implementation.SetStudent(ctx, studen)
}

func GetStudent(ctx context.Context, id string) (*models.Student, error) {
	return implementation.GetStudent(ctx, id)
}

func GetTest(ctx context.Context, id string) (*models.Test, error) {
	return implementation.GetTest(ctx, id)
}

func SetTest(ctx context.Context, test *models.Test) error {
	return implementation.SetTest(ctx, test)
}

func GetQuestion(ctx context.Context, id string) (*models.Question, error) {
	return implementation.GetQuestion(ctx, id)
}

func SetQuestion(ctx context.Context, question *models.Question) error {
	return implementation.SetQuestion(ctx, question)
}
