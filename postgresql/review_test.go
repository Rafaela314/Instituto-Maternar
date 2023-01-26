package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateReview(t *testing.T) {

	arg := CreateReviewParams{
		Title:          "title",
		Content:        "conetnt",
		Date:           time.Now(),
		Classification: "e",
		UserID:         sql.NullInt32{Int32: 5},
		Amount:         sql.NullString{String: ""},
		PlaceID:        4,
		PlaceRate:      3,
		DoctorID:       2,
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
