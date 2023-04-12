package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	// password := RandomString(14)

	hashedPassword, err := HashPassword("younessyoun")
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	fmt.Printf(hashedPassword)

	err = CheckPassword(hashedPassword, "younessyoun")
	require.NoError(t, err)
}
