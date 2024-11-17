// main_test.go
package main

import (
	"testing"

	mocks "github.com/akharb1192/LearningAgain/mocks" // Import the mocks package
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// TestInsertUser_Success tests InsertUser function with successful execution.
func TestInsertUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock of db.DBInterface
	mockDB := mocks.NewMockDBInterface(ctrl)

	// Mock result of Exec method
	mockResult := mocks.NewMockResultInterface(ctrl)

	// Expect Exec to be called with the correct query and parameters
	mockDB.EXPECT().
		Exec("INSERT INTO users (name, email) VALUES (?, ?)", "Jane Doe", "jane.doe@example.com").
		Return(mockResult, nil).
		Times(1)

	// Mock the LastInsertId method to return a specific value
	mockResult.EXPECT().
		LastInsertId().
		Return(int64(123), nil).
		Times(1)

	// Initialize DB with the mock instance
	InitDB(mockDB)

	// Call InsertUser function and assert the result
	newUserID, err := InsertUser("Jane Doe", "jane.doe@example.com")
	assert.NoError(t, err)
	assert.Equal(t, int64(123), newUserID)
}

// TestInsertUser_Error tests InsertUser function when Exec returns an error.
func TestInsertUser_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock of db.DBInterface
	mockDB := mocks.NewMockDBInterface(ctrl)

	// Mock an error being returned from Exec
	mockDB.EXPECT().
		Exec("INSERT INTO users (name, email) VALUES (?, ?)", "Jane Doe", "jane.doe@example.com").
		Return(nil, assert.AnError). // Return a mock error
		Times(1)

	// Initialize DB with the mock instance
	InitDB(mockDB)

	// Call InsertUser function and assert the error is handled
	newUserID, err := InsertUser("Jane Doe", "jane.doe@example.com")
	assert.Error(t, err)
	assert.Equal(t, int64(0), newUserID)
}

// TestInsertUser_LastInsertId_Error tests InsertUser when LastInsertId returns an error.
func TestInsertUser_LastInsertId_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock of db.DBInterface
	mockDB := mocks.NewMockDBInterface(ctrl)

	// Mock result of Exec method
	mockResult := mocks.NewMockResultInterface(ctrl)

	// Expect Exec to be called with the correct query and parameters
	mockDB.EXPECT().
		Exec("INSERT INTO users (name, email) VALUES (?, ?)", "Jane Doe", "jane.doe@example.com").
		Return(mockResult, nil).
		Times(1)

	// Mock the LastInsertId method to return an error
	mockResult.EXPECT().
		LastInsertId().
		Return(int64(0), assert.AnError). // Return a mock error
		Times(1)

	// Initialize DB with the mock instance
	InitDB(mockDB)

	// Call InsertUser function and assert the error is handled
	newUserID, err := InsertUser("Jane Doe", "jane.doe@example.com")
	assert.Error(t, err)
	assert.Equal(t, int64(0), newUserID)
}
