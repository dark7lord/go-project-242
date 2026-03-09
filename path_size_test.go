package path_size

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPathSize(t *testing.T) {
	tests := []struct {
		name      string
		path      string
		recursive bool
		human     bool
		all       bool
		expected  string
		wantErr   bool
	}{
		{
			name:     "file tmnt.csv",
			path:     "./testdata/tmnt.csv",
			expected: "670B\ttmnt.csv",
		},
		{
			name:     "file one-piece.csv",
			path:     "./testdata/one-piece.csv",
			expected: "1681B\tone-piece.csv",
		},
		{
			name:      "directory testdata no recurse",
			path:      "./testdata",
			recursive: false,
			human:     false,
			all:       false,
			expected:  "2351B\ttestdata",
		},
		{
			name:      "directory testdata recurse no hidden",
			path:      "./testdata",
			recursive: true,
			human:     false,
			all:       false,
			expected:  "4713B\ttestdata",
		},
		{
			name:      "directory testdata recurse with hidden",
			path:      "./testdata",
			recursive: true,
			human:     false,
			all:       true,
			expected:  "9252B\ttestdata",
		},
		{
			name:      "directory testdata recurse human readable",
			path:      "./testdata",
			recursive: true,
			human:     true,
			all:       false,
			expected:  "4.6KB\ttestdata",
		},
		{
			name:      "non-existent file",
			path:      "./testdata/non-existent",
			wantErr:   true,
		},
		{
			name:      "file one-piece.csv with human",
			path:      "./testdata/one-piece.csv",
			human:     true,
			expected:  "1.6KB\tone-piece.csv",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := GetPathSize(tt.path, tt.recursive, tt.human, tt.all)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, actual)
			}
		})
	}
}