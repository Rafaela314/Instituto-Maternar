package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// Store provides all functions to execute db queries and transactions
type Store interface {
	Querier
	ReviewTx(ctx context.Context, arg CreateReviewParams) (ReviewTxResult, error)
}

// SQLStore provides all functions to execute db queries and transactions
type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new Store
func NewStore(db *sql.DB) Store {
	return &SQLStore{db: db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// ReviewTxParams contains the input parameter of the transfer transaction
type ReviewTxParams struct {
	ID             int64     `json:"id"`
	UserID         string    `json:"user_id"`
	Title          string    `json:"title" validate:"required"`
	Content        string    `json:"content" validate:"required"`
	Date           time.Time `json:"date"`
	Classification string    `json:"classification"` // public or private
	Amount         float64   `json:"amount"`
	Insurance      string    `json:"insurance"`
	OverallRate    int32     `json:"overall_rate"`
	PlaceID        int       `json:"place_id"`
	PlaceRate      int32     `json:"place_rate"`
	DoctorID       int       `json:"doctor_id"`
	DoctorRate     int32     `json:"doctor_rate"`
	MidwifeID      int       `json:"midwife"`
	MidwifeRate    int32     `json:"midwife_rate"`
	DoulaID        int       `json:"doula_id"`
	DoulaRate      int32     `json:"doula_rate"`
	Team           string    `json:"team"`
	TeamRate       int32     `json:"team_rate"`
	UpdatedAt      time.Time `json:"updated_at"`
	CreatedAt      time.Time `json:"created_at"`
	Image          string    `json:"image"`
}

// ReviewTxResult is the result of the transfer transaction
type ReviewTxResult struct {
	Review Review `json:"review"`
}

// ReviewTx performs a creation of review and updates scores from both doctor and place
func (store *SQLStore) ReviewTx(ctx context.Context, arg CreateReviewParams) (ReviewTxResult, error) {
	var result ReviewTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		review, err := q.CreateReview(ctx, CreateReviewParams{
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
		})

		if err != nil {
			return err
		}

		result.Review = review

		doctor, err := q.GetDoctor(ctx, int64(result.Review.DoctorID))
		if err != nil {
			return err
		}

		doctorReviews, err := q.CountDoctorReviews(ctx, int32(doctor.ID))
		if err != nil {
			return err
		}

		doctor.AverageRate = (doctor.AverageRate + result.Review.DoctorRate) / int32(doctorReviews)

		err = q.UpdateDoctor(ctx, ParseDoctorToUpdateParams(doctor))
		if err != nil {
			return err
		}

		place, err := q.GetPlace(ctx, int64(result.Review.PlaceID))
		if err != nil {
			return err
		}

		placeReviews, err := q.CountPlaceReviews(ctx, int32(place.ID))
		if err != nil {
			return err
		}

		place.AverageRate = (place.AverageRate + result.Review.PlaceRate) / int32(placeReviews)

		err = q.UpdateDoctor(ctx, ParseDoctorToUpdateParams(doctor))
		if err != nil {
			return err
		}

		return nil

	})

	return result, err

}
