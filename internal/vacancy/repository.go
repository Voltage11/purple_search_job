package vacancy

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type VacancyRepository struct {
	DbPool *pgxpool.Pool
	CustomLogger *zerolog.Logger
}

func NewVacancyRepository(dbPool *pgxpool.Pool, customLogger *zerolog.Logger) *VacancyRepository{
	return &VacancyRepository{
		DbPool: dbPool,
		CustomLogger: customLogger,
	}
}

func (r *VacancyRepository) CountAll() (int) {
	query := "SELECT count(*) FROM vacancies"
	var count int
	r.DbPool.QueryRow(context.Background(), query).Scan(&count)
	return count
}

func (r *VacancyRepository) GetAll(limit int, offset int) ([]Vacancy, error) {
	query := "SELECT * FROM vacancies ORDER BY createdat LIMIT @limit OFFSET @offset"
	args := pgx.NamedArgs{
		"limit": limit,
		"offset": offset,
	}
	rows, err := r.DbPool.Query(context.Background(), query, args)
	if (err != nil) {
		return nil, err
	}
	vacancies, err := pgx.CollectRows(rows, pgx.RowToStructByName[Vacancy])
	if (err != nil) {
		return nil, err
	}
	return vacancies, nil
}

func (r *VacancyRepository) addVacancy(form VacancyCreateForm) error {
	query := `INSERT INTO vacancies (email, role, company, salary, type, location, createdat) VALUES (@email, @role, @company, @salary, @type, @location, @createdat)`
	args := pgx.NamedArgs{
		"email": form.Email,
		"role": form.Role,
		"company": form.Salary,
		"salary": form.Salary,
		"type": form.Type,
		"location": form.Location,
		"createdat": time.Now(),
	}
	_, err := r.DbPool.Exec(context.Background(), query, args)
	if (err != nil) {
		return fmt.Errorf("невозможно создать вакансию: %w", err)
	}
	return nil
}