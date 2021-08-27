package db

import (
	"context"
	"fmt"
	"mylab/cpagent/responses"
	"strings"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

const (
	insertDyeQuery1 = `INSERT INTO dyes(
				name,
				position)
				VALUES %s `
	insertDyeQuery2 = `ON CONFLICT DO NOTHING;`

	getDyes         = `SELECT * FROM dyes`
	getDyeByIDQuery = `SELECT * FROM dyes WHERE id = $1`
)

type Dye struct {
	ID       uuid.UUID `db:"id"`
	Name     string    `db:"name"`
	Position int       `db:"position"`
}

func (s *pgStore) InsertDyes(ctx context.Context, dyes []Dye) (DBdyes []Dye, err error) {

	stmt := makeDyeQuery(dyes)

	_, err = s.db.Exec(
		stmt,
	)
	if err != nil {
		logger.WithField("error in exec query", err.Error()).Error("Query Failed")
		return
	}

	err = s.db.Select(&DBdyes, getDyes)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing dyes")
		return
	}
	return
}
func (s *pgStore) ShowDye(ctx context.Context, dyeID uuid.UUID) (DBdye Dye, err error) {

	err = s.db.Get(&DBdye, getDyeByIDQuery, dyeID)
	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.DyeDBFetchError)
		return
	}
	return

}

// prepare bulk insert query statement
func makeDyeQuery(dye []Dye) string {

	values := make([]string, 0, len(dye))

	for _, d := range dye {
		values = append(values, fmt.Sprintf("('%v', %v)", d.Name, d.Position))
	}

	stmt := fmt.Sprintf(insertDyeQuery1,
		strings.Join(values, ","))

	stmt += insertDyeQuery2

	return stmt
}
