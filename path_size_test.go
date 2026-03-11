package code

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPathSize_Files(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		human    bool
		expected string
	}{
		{
			name:     "returns file size for tmnt.csv without human readable",
			path:     "./testdata/tmnt.csv",
			expected: "670B\ttmnt.csv",
		},
		{
			name:     "returns file size for one-piece.csv without human readable",
			path:     "./testdata/one-piece.csv",
			expected: "1681B\tone-piece.csv",
		},
		{
			name:     "returns file size for one-piece.csv in human readable format",
			path:     "./testdata/one-piece.csv",
			human:    true,
			expected: "1.6KB\tone-piece.csv",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := GetPathSize(tt.path, false, tt.human, false)
			require.NoError(t, err)
			require.Equal(t, tt.expected, actual)
		})
	}
}

func TestGetPathSize_Directories(t *testing.T) {
	tests := []struct {
		name      string
		path      string
		recursive bool
		human     bool
		all       bool
		expected  string
	}{
		{
			name:      "returns directory size for testdata without recursion",
			path:      "./testdata",
			recursive: false,
			expected:  "2351B\ttestdata",
		},
		{
			name:      "returns directory size for testdata recursively without hidden files",
			path:      "./testdata",
			recursive: true,
			expected:  "4713B\ttestdata",
		},
		{
			name:      "returns directory size for testdata recursively with hidden files",
			path:      "./testdata",
			recursive: true,
			all:       true,
			expected:  "9252B\ttestdata",
		},
		{
			name:      "returns directory size for testdata recursively in human readable format",
			path:      "./testdata",
			recursive: true,
			human:     true,
			expected:  "4.6KB\ttestdata",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := GetPathSize(tt.path, tt.recursive, tt.human, tt.all)
			require.NoError(t, err)
			require.Equal(t, tt.expected, actual)
		})
	}
}

func TestGetPathSize_Errors(t *testing.T) {
	t.Run("returns an error for a non-existent file", func(t *testing.T) {
		_, err := GetPathSize("./testdata/non-existent", false, false, false)
		require.Error(t, err)
	})
}
