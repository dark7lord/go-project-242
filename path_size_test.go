package path_size

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetPathSize_File1(t *testing.T) {
	actual, _ := GetPathSize("./testdata/tmnt.csv", false, false, false)
	expected := "670B\ttmnt.csv"
	require.Equal(t, actual, expected)
}

func TestGetPathSize_File2(t *testing.T) {
	actual, _ := GetPathSize("./testdata/one-piece.csv", false, false, false)
	expected := "1681B\tone-piece.csv"
	require.Equal(t, actual, expected)
}

func TestGetPathSize_Dir(t *testing.T) {
	actual, _ := GetPathSize("./testdata", false, false, false)
	expected := "2351B\ttestdata"
	require.Equal(t, actual, expected)
}
