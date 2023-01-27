package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/Rafaela314/instituto-maternar/util"
	"github.com/stretchr/testify/require"
)

func createRandomPlace(t *testing.T) Place {
	arg := CreatePlaceParams{
		Name:        util.RandomName(),
		Address:     sql.NullString{String: util.RandomName()},
		City:        util.RandomName(),
		State:       util.RandomName(),
		AverageRate: int32(util.RandInt(0, 5)),
	}

	place, err := testQueries.CreatePlace(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, place)

	require.Equal(t, arg.Name, place.Name)

	require.NotZero(t, place.ID)
	require.NotZero(t, place.CreatedAt)
	return place

}

func TestCreatePlace(t *testing.T) {
	createRandomPlace(t)
}
