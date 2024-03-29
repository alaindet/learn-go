package models

import (
	"testing"

	"snippetbox.dev/internal/assert"
)

func TestUserModelExists(t *testing.T) {

	// This is flagged as a "long" test since it's interacting with the database
	// Skip this if "go test -short" is executed
	if testing.Short() {
		t.Skip("models: skipping integration test")
	}

	testCases := []struct {
		name     string
		userID   int
		expected bool
	}{
		{
			name:     "Valid ID",
			userID:   1,
			expected: true,
		},
		{
			name:     "Zero ID",
			userID:   0,
			expected: false,
		},
		{
			name:     "Non-existent ID",
			userID:   123,
			expected: false,
		},
	}

	// It's safe to create just one db connection, as this test only reads data
	db := newTestDB(t)

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			m := NewUserModel(db)
			result, err := m.Exists(testCase.userID)
			assert.Equal(t, result, testCase.expected)
			assert.NilError(t, err)
		})
	}
}
