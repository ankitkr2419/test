package db

import (
	"context"
	"errors"
	"fmt"
	"mylab/cpagent/responses"
	"strings"
	"time"

	logger "github.com/sirupsen/logrus"
)

const (
	insertCartridgeQuery = `INSERT INTO cartridges(
							id,
							type,
							description)
							VALUES %s `
	insertCartridgeWellsQuery = `INSERT INTO cartridge_wells(
							id,
							well_num,
							distance,
							height,
							volume)
							VALUES %s `
	onConflictDoNothing          = `ON CONFLICT DO NOTHING;`
	selectAllCartridgeQuery      = `SELECT c.*, count(cw.id) as wells_count FROM cartridge_wells cw LEFT JOIN cartridges c ON c.id=cw.id GROUP BY c.id`
	selectAllCartridgeWellsQuery = `SELECT *
							FROM cartridge_wells`
	selectCartridgeWellsQuery = `SELECT *
							FROM cartridge_wells WHERE id = $1`
	deleteCartridgeQuery      = `delete from cartridges where id = $1`
	countRecipeCartridgeQuery = `select count(*) from recipes where pos_cartridge_1 = $1 OR pos_cartridge_2 = $1`
	getCartridgeQuery         = `SELECT c.*, count(cw.id) as wells_count FROM cartridge_wells cw LEFT JOIN cartridges c ON c.id=cw.id WHERE c.id = $1 GROUP BY c.id`
)

type CartridgeType string

const (
	Cartridge1 CartridgeType = "cartridge_1"
	Cartridge2 CartridgeType = "cartridge_2"
)

type CartridgeWell struct {
	Cartridge      []Cartridge      `json:"cartridges"`
	CartridgeWells []CartridgeWells `json:"cartridge_wells"`
}
type Cartridge struct {
	ID          int64         `db:"id" json:"id"`
	Type        CartridgeType `db:"type" json:"type"`
	Description string        `db:"description" json:"description"`
	WellsCount  int64         `db:"wells_count" json:"wells_count"`
	CreatedAt   time.Time     `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time     `db:"updated_at" json:"updated_at"`
}

type CartridgeWells struct {
	ID        int64     `db:"id" json:"id"`
	WellNum   int64     `db:"well_num" json:"well_num"`
	Distance  float64   `db:"distance" json:"distance"`
	Height    float64   `db:"height" json:"height"`
	Volume    float64   `db:"volume" json:"volume"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (s *pgStore) InsertCartridge(ctx context.Context, cartridges []Cartridge, cartridgeWells []CartridgeWells) (err error) {

	go s.AddAuditLog(ctx, DBOperation, InitialisedState, CreateOperation, "", responses.CartridgeInitialisedState)

	stmt1 := makeCartridgeQuery(cartridges)
	stmt2 := makeCartridgeWellsQuery(cartridgeWells)

	_, err = s.db.Exec(
		stmt1,
	)
	defer func() {
		if err != nil {
			go s.AddAuditLog(ctx, DBOperation, ErrorState, CreateOperation, "", err.Error())
		} else {
			go s.AddAuditLog(ctx, DBOperation, CompletedState, CreateOperation, "", responses.CartridgeCompletedState)
		}
	}()
	if err != nil {
		logger.WithField("error in exec query", err.Error()).Error("Query Failed")
		return
	}

	_, err = s.db.Exec(
		stmt2,
	)
	if err != nil {
		logger.WithField("error in exec query", err.Error()).Error("Query Failed")
		return
	}
	return
}

func (s *pgStore) DeleteCartridge(ctx context.Context, id int64) (err error) {
	_, err = s.db.Exec(deleteCartridgeQuery, id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error deleting Cartridge data")
		return
	}

	return
}

func (s *pgStore) IsCartridgeSafeToDelete(ctx context.Context, id int64) (count []int64, err error) {
	err = s.db.Select(&count, countRecipeCartridgeQuery, id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error counting recipes related to Cartridge data")
		return
	}

	if count[0] != 0 {
		return count, errors.New("Not safe to delete")
	}
	return
}

func makeCartridgeQuery(cartridge []Cartridge) string {
	values := make([]string, 0, len(cartridge))

	for _, c := range cartridge {
		values = append(values, fmt.Sprintf("(%v, '%v', '%v')", c.ID, c.Type, c.Description))
	}

	stmt := fmt.Sprintf(insertCartridgeQuery,
		strings.Join(values, ","))

	stmt += onConflictDoNothing

	return stmt
}

func makeCartridgeWellsQuery(cartridgeWells []CartridgeWells) string {
	values := make([]string, 0, len(cartridgeWells))

	for _, c := range cartridgeWells {
		values = append(values, fmt.Sprintf("(%v, %v, %v, %v, %v)", c.ID, c.WellNum, c.Distance, c.Height, c.Volume))
	}

	stmt := fmt.Sprintf(insertCartridgeWellsQuery,
		strings.Join(values, ","))

	stmt += onConflictDoNothing
	return stmt
}

func (s *pgStore) ListCartridges(ctx context.Context) (cartridge []Cartridge, err error) {
	go s.AddAuditLog(ctx, DBOperation, InitialisedState, ShowOperation, "", responses.CartridgeListInitialisedState)

	err = s.db.Select(&cartridge, selectAllCartridgeQuery)
	defer func() {
		if err != nil {
			go s.AddAuditLog(ctx, DBOperation, ErrorState, ShowOperation, "", err.Error())
		} else {
			go s.AddAuditLog(ctx, DBOperation, CompletedState, ShowOperation, "", responses.CartridgeListCompletedState)
		}
	}()
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing cartridge details")
		return
	}
	return
}

func (s *pgStore) ListCartridgeWells() (cartridgeWells []CartridgeWells, err error) {
	err = s.db.Select(&cartridgeWells, selectAllCartridgeWellsQuery)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing cartridgeWells details")
		return
	}
	return
}

func (s *pgStore) ShowCartridgeWells(id int64) (cartridgeWells []CartridgeWells, err error) {
	err = s.db.Select(&cartridgeWells, selectCartridgeWellsQuery, id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing cartridgeWells details")
		return
	}
	return
}

func (s *pgStore) ShowCartridge(ctx context.Context, id int64) (dbCartridge Cartridge, err error) {
	err = s.db.Get(&dbCartridge, getCartridgeQuery, id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching cartridge")
		return
	}
	return
}
