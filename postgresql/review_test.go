package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateReview(t *testing.T) {

	doctor := createRandomDoctor(t)

	user := createRandomUser(t)

	place := createRandomPlace(t)

	arg := CreateReviewParams{
		Title:          "title",
		Content:        "conetnt",
		Date:           time.Now(),
		Classification: "e",
		UserID:         sql.NullInt32{Int32: int32(user.ID)},
		Amount:         sql.NullString{String: ""},
		PlaceID:        int32(place.ID),
		PlaceRate:      3,
		DoctorID:       int32(doctor.ID),
		DoctorRate:     4,
		OverallRate:    3,
	}

	review, err := testQueries.CreateReview(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, review)

	require.Equal(t, arg.Title, review.Title)

	require.NotZero(t, review.ID)
	require.NotZero(t, review.CreatedAt)
}
