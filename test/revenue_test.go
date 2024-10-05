package test

import (
	"go-backend-assessment/internal/calculations"
	"testing"
)

func TestCalculateRevenue(t *testing.T) {
	revenue, err := calculations.CalculateRevenue("2023-01-01", "2023-12-31")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if revenue <= 0 {
		t.Errorf("Expected revenue > 0, got %v", revenue)
	}
}
