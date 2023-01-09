package services

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Rafaela314/instituto-maternar/models"
)

const (
	//
	maxConnections         = 1   // Allow 1 connections (per application instance)
	connMaxLifetimeMinutes = 0   // Active connections can live forever.
	connMaxIdleMinutes     = 240 // But if they are idle, remove them after 4 hours.
)

// Subscriber item subscription interface
type Subscriber interface {
	CreateSubscription(reviewID int64) (int, error)
}

// SubscriptionService service to send data to subscription database
type SubscriptionService struct {
	Conn   *sql.DB
	Config *models.AppConfig
}

// NewSubscriptionService create a service for sending data to database
func NewSubscriptionService(config *models.AppConfig) (Subscriber, error) {

	subscriptionService := &SubscriptionService{
		Config: config,
	}

	dbConnection, err := subscriptionService.ConnectDb()
	if err != nil {
		return nil, err
	}

	subscriptionService.Conn = dbConnection

	return subscriptionService, nil
}

// ConnectDb connects to postgres database
func (ss *SubscriptionService) ConnectDb() (*sql.DB, error) {

	dbName := ss.Config.DBName
	dbUser := ss.Config.DBUser
	dbHost := ss.Config.DBHost
	dbPort := ss.Config.DBPort
	dbPassword := ss.Config.DBPassword

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxConnections)
	db.SetConnMaxIdleTime(time.Duration(connMaxIdleMinutes) * time.Minute)
	db.SetConnMaxLifetime(time.Duration(connMaxLifetimeMinutes) * time.Minute)

	return db, err
}

// CreateSubscription calls add data procedure
func (ss *SubscriptionService) CreateSubscription(reviewID int64) (int, error) {

	var upsertResult int

	/*
			err = ss.GetConnectionStatus(span)
			if err != nil {
				subspan.Errorf("error getting db connection status: " + err.Error())
				return 0, err
			}

		// Call the stored procedure for subscription check
		// subspan.Info("calling the stored procedure")
		rows, err := ss.Conn.Query("select upsertsubscription($1,$2,$3,$4);", productID, itemID, siteID, platformID)
		if err != nil {
			return 0, err
		}

		defer func() {
			err = rows.Close()
			if err != nil {
				fmt.Sprintf("Error closing rows: %w", err)

			}
		}()
	*/
	return upsertResult, nil
}
