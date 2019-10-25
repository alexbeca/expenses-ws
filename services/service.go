package services

import (
    "context"
    "time"
)

// Service provides some "date capabilities" to your application
type Service interface {
    Status(ctx context.Context) (string, error)
    Get(ctx context.Context) (string, error)
    Validate(ctx context.Context, date string) (bool, error)
}

type ExpService interface {
    Register(ctx context.Context) (string, error)
}

//Implementation of the interface starts here
type dateService struct{}

type expenseService struct{}

// NewService makes a new Service. Constructor
func NewService() Service {
    return dateService{}
}

func NewExpenseService() ExpService {
    return expenseService{}
}

// Status only tell us that our service is ok!
func (dateService) Status(ctx context.Context) (string, error) {
    return "ok", nil
}

// Get will return today's date
func (dateService) Get(ctx context.Context) (string, error) {
    now := time.Now()
    return now.Format("02/01/2006"), nil
}

// Validate will check if the date today's date
func (dateService) Validate(ctx context.Context, date string) (bool, error) {
    _, err := time.Parse("02/01/2006", date)
    if err != nil {
        return false, err
    }
    return true, nil
}

// Register expense (Methods implementation of the interface)
func (expenseService) Register(ctx context.Context, ) (string, error){
    return "ok", nil
}