package service

import (
	"context"
	"time"

	"user-api/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// ---------------- AGE CALCULATION ----------------
func CalculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()

	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}

// ---------------- CREATE ----------------
func (s *UserService) Create(ctx context.Context, name string, dob string) (interface{}, error) {
	return s.repo.Create(ctx, name, dob)
}

// ---------------- GET BY ID ----------------
func (s *UserService) Get(ctx context.Context, id int32) (map[string]interface{}, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.Dob.Time.Format("2006-01-02"),
		"age":  CalculateAge(user.Dob.Time),
	}, nil
}

// ---------------- LIST ALL ----------------
func (s *UserService) List(ctx context.Context) ([]map[string]interface{}, error) {
	users, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	for _, u := range users {
		result = append(result, map[string]interface{}{
			"id":   u.ID,
			"name": u.Name,
			"dob":  u.Dob.Time.Format("2006-01-02"),
			"age":  CalculateAge(u.Dob.Time),
		})
	}

	return result, nil
}

// ---------------- UPDATE ----------------
func (s *UserService) Update(ctx context.Context, id int32, name string, dob string) (interface{}, error) {
	return s.repo.Update(ctx, id, name, dob)
}

// ---------------- DELETE ----------------
func (s *UserService) Delete(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}
