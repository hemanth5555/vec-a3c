package store

import (
	"context"
	"errors"
)

// Store is an interface that defines the methods for interacting with the data store.
type Store interface {
	GetQueueSize(ctx context.Context) (int, error)
	SetQueueSize(ctx context.Context, size int) error
	UpdateQueueSize(ctx context.Context, size int) error
}

// StoreError is a custom error type for store-related errors. It includes the original error and a status code.
type StoreError struct {
	Err        error
	StatusCode int
}

// Error returns the error message of the wrapped error.
func (e *StoreError) Error() string {
	return e.Err.Error()
}

// Unwrap returns the original wrapped error.
func (e *StoreError) Unwrap() error {
	return e.Err
}

// ErrNotFound is a sentinel error for when a requested item is not found in the store.
var ErrNotFound = errors.New("not found")
