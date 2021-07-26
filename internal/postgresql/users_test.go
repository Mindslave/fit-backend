package postgresql

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Mindslave/fit-backend/internal/testdata"
)

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		FirstName: testdata.RandomName(),
		LastName: testdata.RandomName(),
		Goal: testdata.RandomName(),
		Coach: false,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.Coach, user.Coach)
	require.Equal(t, arg.Goal, user.Goal)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
}