package db

import (
	"context"
	"testing"

	"github.com/Rafaela314/instituto-maternar/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		Username: util.RandomName(),
		Email:    util.RandomName(),
		Password: util.RandomName(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
}
