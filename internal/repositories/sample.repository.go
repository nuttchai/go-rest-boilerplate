package repositories

import (
	"database/sql"

	dto "github.com/nuttchai/go-rest/internal/dto/sample"
	"github.com/nuttchai/go-rest/internal/models"
	"github.com/nuttchai/go-rest/internal/shared/console"
	"github.com/nuttchai/go-rest/internal/types"
	"github.com/nuttchai/go-rest/internal/utils/context"
	"github.com/nuttchai/go-rest/internal/utils/db"
)

var (
	SampleRepository ISampleRepository
)

func init() {
	SampleRepository = &DBModel{}
}

func (m *DBModel) Test() string {
	console.App.Log("Call Test Function in Repository!")
	return "test"
}

func (m *DBModel) GetSample(id string, filters ...*types.QueryFilter) (*models.Sample, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	baseQuery := "select * from samples where id = $1"
	baseArgs := []interface{}{id}

	query, args := db.BuildQueryWithFilter(baseQuery, baseArgs, filters...)
	row := m.SqlDB.QueryRowContext(ctx, query, args...)

	var sample models.Sample
	err := row.Scan(
		&sample.Id,
		&sample.Name,
	)

	return &sample, err
}

func (m *DBModel) CreateSample(sample *dto.CreateSampleDTO) (*models.Sample, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	query := `
		insert into samples (name, desc)
		values ($1, $2)
		returning *
	`
	row := m.SqlDB.QueryRowContext(ctx, query, sample.Name, sample.Desc)

	var newSample models.Sample
	if err := row.Scan(
		&newSample.Id,
		&newSample.Name,
		&newSample.Desc,
	); err != nil {
		return nil, err
	}

	return &newSample, nil
}

func (m *DBModel) UpdateSample(sample *dto.UpdateSampleDTO) (*models.Sample, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	query := `
		update samples set name = $1, desc = $2
		where id = $3
		returning *
	`
	row := m.SqlDB.QueryRowContext(ctx, query, sample.Name, sample.Desc, sample.Id)

	var updatedSample models.Sample
	if err := row.Scan(
		&updatedSample.Id,
		&updatedSample.Name,
		&updatedSample.Desc,
	); err != nil {
		return nil, err
	}

	return &updatedSample, nil
}

func (m *DBModel) DeleteSample(id string) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(3)
	defer cancel()

	query := `
		delete from samples 
		where id = $1
	`
	result, err := m.SqlDB.ExecContext(ctx, query, id)

	return result, err
}
