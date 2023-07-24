package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"go-starter-course/web/types"
	"os"
)

var DBAlreadyConnected = errors.New("Already connected to DB")

var (
	dbUsername = os.Getenv("DB_USERNAME")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName     = os.Getenv("DB_NAME")
)

type PositionEntity struct {
	id    int64
	label string
	types.Position
}

type App struct {
	DB *sql.DB
	// could add a router with some middleware and so on
	// define some methods for handlers
}

func InitApp() *App {
	return &App{}
}

func (app *App) ConnectToDB() error {
	if app.DB != nil {
		log.Warn().Msg("DB is already connected")
		return DBAlreadyConnected
	}

	log.Debug().Msg("Trying to create PostgreSQL database")
	var err error
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUsername, dbPassword, dbName)
	app.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Error().Msg("Couldn't connect to Database - check connection parameters")
		return err
	}
	err = app.DB.Ping()
	if err != nil {
		log.Error().Msg("Couldn't ping Database - check database status")
		return err
	}
	return nil
}

func (app *App) CloseDbConnection() {
	err := app.DB.Close()
	if err != nil {
		log.Warn().Msgf("Couldn't close database connection: %s", err)
	}
}

func (app *App) GetPositions() ([]PositionEntity, error) {
	rows, err := app.DB.Query("SELECT position_id, label, latitude, longitude, elevation FROM position")
	if err != nil {
		log.Error().Msgf("Couldn't query database: %s", err)
		return nil, err
	}
	defer rows.Close()

	var positions []PositionEntity
	for rows.Next() {
		var positionId int64
		var label string
		var latitude float64
		var longitude float64
		var elevation int
		err := rows.Scan(&positionId, &label, &latitude, &longitude, &elevation)
		if err != nil {
			log.Error().Msgf("Couldn't process rows: %s", err)
			return nil, err
		}
		positions = append(
			positions,
			PositionEntity{positionId, label,
				types.Position{latitude, longitude, elevation}})
	}
	err = rows.Err()
	if err != nil {
		log.Error().Msgf("Couldn't read rows: %s", err)
		return nil, err
	}
	return positions, nil
}

func (app *App) GetPositionByLabel(fetchedLabel string) (*PositionEntity, error) {
	return mapPositionEntity(app.DB.QueryRow(
		"SELECT position_id, label, latitude, longitude, elevation "+
			"FROM position WHERE label = $1",
		fetchedLabel))
}

func (app *App) GetPositionById(id int64) (*PositionEntity, error) {
	return mapPositionEntity(app.DB.QueryRow(
		"SELECT position_id, label, latitude, longitude, elevation "+
			"FROM position WHERE position_id = $1",
		id))
}

func mapPositionEntity(row *sql.Row) (*PositionEntity, error) {
	var positionId int64
	var label string
	var latitude float64
	var longitude float64
	var elevation int
	err := row.Scan(&positionId, &label, &latitude, &longitude, &elevation)
	switch {
	case err == sql.ErrNoRows:
		log.Debug().Msgf("There was no rows for get positions returned")
		return nil, err
	case err != nil:
		return nil, err
	}
	return &PositionEntity{positionId, label,
		types.Position{latitude, longitude, elevation}}, nil
}

func (app *App) CreateOrUpdatePosition(label string, position types.Position) (*PositionEntity, error) {
	tx, err := app.DB.Begin()
	if err != nil {
		return nil, err
	}
	row := tx.QueryRow("SELECT position_id FROM position WHERE label = $1", label)
	var positionId int64
	err = row.Scan(&positionId)
	if err == nil {
		// this is when positionId is in the table
		// do update here
		_, err := tx.Exec("UPDATE position SET latitude = $1, longitude = $2, elevation = $3 WHERE position_id = $4",
			position.Latitude, position.Longitude, position.Elevation, positionId)
		if err != nil {
			_ = tx.Rollback()
			return nil, err
		}
	} else if err == sql.ErrNoRows {
		// this is when we did a query and nothing was returned
		// do insert here
		stmt, err := tx.Prepare(
			"INSERT INTO position (label, latitude, longitude, elevation) VALUES ($1, $2, $3, $4) RETURNING position_id")
		if err != nil {
			_ = tx.Rollback()
			return nil, err
		}
		err = stmt.QueryRow(label, position.Latitude, position.Longitude, position.Elevation).Scan(&positionId)
		if err != nil {
			_ = tx.Rollback()
			return nil, err
		}
	} else {
		_ = tx.Rollback()
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return app.GetPositionById(positionId)
}

func main() {
	app := InitApp()
	err := app.ConnectToDB()
	switch {
	case err == DBAlreadyConnected:
		log.Info().Msg("DB is connected - ignoring reconnect")
	case err != nil:
		log.Fatal().Msg("We cannot use DB we exit...")
	}
	defer app.CloseDbConnection()

	_, err = app.GetPositions()
	if err != nil {
		log.Error().Msgf("GetPositions returned error: %s", err)
	}

	_, err = app.CreateOrUpdatePosition("home", types.Position{1, 1, 1})
	if err != nil {
		log.Error().Msgf("CreateOrUpdatePosition returned error: %s", err)
		return
	}
	position, err := app.CreateOrUpdatePosition("zero", types.Position{})
	if err != nil {
		log.Error().Msgf("CreateOrUpdatePosition returned error: %s", err)
		return
	}
	log.Info().Msgf("CreateOrUpdate for zero has %+v", *position)

	zeroPosition, err := app.GetPositionByLabel("zero")
	switch {
	case err == sql.ErrNoRows:
		log.Warn().Msg("GetPositionByLabel returned no rows - handle this differenly")
	case err != nil:
		log.Error().Msgf("GetPositionByLabel returned error: %s", err)
		return
	}
	log.Info().Msgf("GetPositionByLabel for zero has %+v", *zeroPosition)

	homePosition, err := app.GetPositionByLabel("home")
	switch {
	case err == sql.ErrNoRows:
		log.Warn().Msg("GetPositionByLabel returned no rows - handle this differently")
	case err != nil:
		log.Error().Msgf("GetPositionByLabel returned error: %s", err)
		return
	}
	log.Info().Msgf("GetPositionByLabel for zero has %+v", *homePosition)

	log.Info().Msg("Finished execution of our application...")
}
