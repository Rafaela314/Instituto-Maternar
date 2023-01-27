package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Rafaela314/instituto-maternar/util"
	"github.com/stretchr/testify/require"
)

func createRandomDoctor(t *testing.T) Doctor {
	arg := CreateDoctorParams{
		Name:        util.RandomName(),
		Crm:         util.RandomName(),
		AverageRate: 3,
	}

	doctor, err := testQueries.CreateDoctor(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, doctor)

	require.Equal(t, arg.Name, doctor.Name)

	require.NotZero(t, doctor.ID)
	require.NotZero(t, doctor.CreatedAt)

	return doctor
}
func TestCreateDoctor(t *testing.T) {
	createRandomDoctor(t)

}

func TestGetDoctor(t *testing.T) {
	doctor := createRandomDoctor(t)

	doctor1, err := testQueries.GetDoctor(context.Background(), doctor.ID)
	require.NoError(t, err)
	require.NotEmpty(t, doctor)

	require.Equal(t, doctor.ID, doctor1.ID)
	require.Equal(t, doctor.Name, doctor1.Name)
	require.Equal(t, doctor.Crm, doctor1.Crm)
	require.WithinDuration(t, doctor.CreatedAt.Time, doctor1.CreatedAt.Time, time.Second)
}

func TestUpdateDoctor(t *testing.T) {

	doctor := createRandomDoctor(t)

	arg := UpdateDoctorParams{
		ID:   doctor.ID,
		Name: util.RandomName(),
		Crm:  util.RandomName(),
	}

	err := testQueries.UpdateDoctor(context.Background(), arg)
	require.NoError(t, err)

	doctor1, err := testQueries.GetDoctor(context.Background(), doctor.ID)
	require.NoError(t, err)
	require.NotEmpty(t, doctor1)

	require.Equal(t, doctor.ID, doctor1.ID)
	require.Equal(t, arg.Name, doctor1.Name)
	require.Equal(t, arg.Crm, doctor1.Crm)

}

func TestDeleteDoctor(t *testing.T) {
	doctor := createRandomDoctor(t)

	err := testQueries.DeleteDoctor(context.Background(), doctor.ID)
	require.NoError(t, err)

	doctor1, err := testQueries.GetDoctor(context.Background(), doctor.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, doctor1)
}

func TestListDoctors(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomDoctor(t)
	}

	arg := ListDoctorsParams{
		Limit:  5,
		Offset: 5,
	}

	doctors, err := testQueries.ListDoctors(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, doctors, 5)

}
