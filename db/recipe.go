package db

import (
	"context"
	"mylab/cpagent/responses"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

const (
	defaultRecipeTime  = 900
	getRecipeQuery     = `SELECT * FROM recipes WHERE id = $1`
	selectRecipesQuery = `SELECT * FROM recipes`
	deleteRecipeQuery  = `DELETE FROM recipes WHERE id = $1`
	createRecipeQuery  = `INSERT INTO recipes (
						name,
						description,
						pos_1,
						pos_2,
						pos_3,
						pos_4,
						pos_5,
						pos_6,
						pos_7,
						pos_cartridge_1,
						pos_9,
						pos_cartridge_2,
						pos_11,
						total_time,
						is_published)
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) RETURNING id`
	updateRecipeQuery = `UPDATE recipes SET (
						name,
						description,
						pos_1,
						pos_2,
						pos_3,
						pos_4,
						pos_5,
						pos_6,
						pos_7,
						pos_cartridge_1,
						pos_9,
						pos_cartridge_2,
						pos_11,
						total_time,
						is_published,
						updated_at) = ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) WHERE id = $17`
)

type Recipe struct {
	ID                 uuid.UUID `db:"id" json:"id"`
	Name               string    `db:"name" json:"name" validate:"required"`
	Description        string    `db:"description" json:"description"`
	Position1          *int64    `db:"pos_1" json:"pos_1"`
	Position2          *int64    `db:"pos_2" json:"pos_2"`
	Position3          *int64    `db:"pos_3" json:"pos_3"`
	Position4          *int64    `db:"pos_4" json:"pos_4"`
	Position5          *int64    `db:"pos_5" json:"pos_5"`
	Position6          *int64    `db:"pos_6" json:"pos_6"`
	Position7          *int64    `db:"pos_7" json:"pos_7"`
	Cartridge1Position *int64    `db:"pos_cartridge_1" json:"pos_cartridge_1"`
	Position9          *int64    `db:"pos_9" json:"pos_9"`
	Cartridge2Position *int64    `db:"pos_cartridge_2" json:"pos_cartridge_2"`
	Position11         *int64    `db:"pos_11" json:"pos_11"`
	ProcessCount       int64     `db:"process_count" json:"process_count"`
	IsPublished        bool      `db:"is_published" json:"is_published"`
	TotalTime          int64     `db:"total_time" json:"total_time"`
	CreatedAt          time.Time `db:"created_at" json:"created_at"`
	UpdatedAt          time.Time `db:"updated_at" json:"updated_at"`
}

func (s *pgStore) ShowRecipe(ctx context.Context, id uuid.UUID) (dbRecipe Recipe, err error) {
	err = s.db.Get(&dbRecipe, getRecipeQuery, id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching recipe")
		return
	}
	return
}

func (s *pgStore) ListRecipes(ctx context.Context) (dbRecipe []Recipe, err error) {
	err = s.db.Select(&dbRecipe, selectRecipesQuery)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching recipes")
		return
	}
	return
}

func (s *pgStore) CreateRecipe(ctx context.Context, r Recipe) (createdRecipe Recipe, err error) {
	var lastInsertID uuid.UUID

	if r.TotalTime == 0 {
		r.TotalTime = defaultRecipeTime
	}

	err = s.db.QueryRow(
		createRecipeQuery,
		r.Name,
		r.Description,
		r.Position1,
		r.Position2,
		r.Position3,
		r.Position4,
		r.Position5,
		r.Position6,
		r.Position7,
		r.Cartridge1Position,
		r.Position9,
		r.Cartridge2Position,
		r.Position11,
		r.TotalTime,
		r.IsPublished,
	).Scan(&lastInsertID)

	if err != nil {
		logger.WithField("err", err.Error()).Error("Error creating Recipe")
		return
	}

	err = s.db.Get(&createdRecipe, getRecipeQuery, lastInsertID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error in getting Recipe")
		return
	}
	return
}

func (s *pgStore) DeleteRecipe(ctx context.Context, id uuid.UUID) (err error) {
	result, err := s.db.Exec(deleteRecipeQuery, id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error deleting recipe")
		return
	}

	c, _ := result.RowsAffected()
	// check row count as no error is returned when row not found for update
	if c == 0 {
		logger.Errorln(responses.RecipeIDInvalidError)
		return responses.RecipeIDInvalidError
	}
	return
}

func (s *pgStore) UpdateRecipe(ctx context.Context, r Recipe) (err error) {

	if r.TotalTime == 0 {
		r.TotalTime = defaultRecipeTime
	}

	result, err := s.db.Exec(
		updateRecipeQuery,
		r.Name,
		r.Description,
		r.Position1,
		r.Position2,
		r.Position3,
		r.Position4,
		r.Position5,
		r.Position6,
		r.Position7,
		r.Cartridge1Position,
		r.Position9,
		r.Cartridge2Position,
		r.Position11,
		r.TotalTime,
		r.IsPublished,
		time.Now(),
		r.ID,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating recipe")
		return
	}

	c, _ := result.RowsAffected()
	// check row count as no error is returned when row not found for update
	if c == 0 {
		logger.Errorln(responses.RecipeIDInvalidError)
		return responses.RecipeIDInvalidError
	}
	return
}
