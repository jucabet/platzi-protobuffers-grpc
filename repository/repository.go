package repository

import (
	"context"

	"github.com/jucabet/platzi-protobuffers-grpc/models"
)

type Repository interface {
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	SetStudent(ctx context.Context, student *models.Student) error
	GetTest(ctx context.Context, id string) (*models.Test, error)
	SetTest(ctx context.Context, test *models.Test) error
	SetQuestion(ctx context.Context, question *models.Question) error
	SetEnrollment(ctx context.Context, enroll *models.Enrollment) error
	GetStudentsPerTest(ctx context.Context, id string) ([]*models.Student, error)
	GetQuestionsPerTest(ctx context.Context, testId string) ([]*models.Question, error)
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

func SetQuestion(ctx context.Context, question *models.Question) error {
	return implementation.SetQuestion(ctx, question)
}

func SetEnrollment(ctx context.Context, enroll *models.Enrollment) error {
	return implementation.SetEnrollment(ctx, enroll)
}

func GetStudentsPerTest(ctx context.Context, id string) ([]*models.Student, error) {
	return implementation.GetStudentsPerTest(ctx, id)
}

func GetQuestionsPerTest(ctx context.Context, testId string) ([]*models.Question, error) {
	return implementation.GetQuestionsPerTest(ctx, testId)
}
