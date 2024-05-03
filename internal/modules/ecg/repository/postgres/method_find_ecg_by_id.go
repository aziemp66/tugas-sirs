package ecg_repository_postgres

import (
	"context"
	"database/sql"
	ecg_entity "tugas-sirs/internal/modules/ecg/model/entity"
	pkg_error "tugas-sirs/pkg/error"
)

// FindEcgByID retrieves an ECG record from the database by its ID.
// It takes the context and the ID of the ECG record as parameters.
// If the ECG record is not found, it returns nil and a NotFound error.
// If an error occurs during the retrieval process, it returns nil and a BadRequest error.
// Otherwise, it returns the Ecg pointer and nil.
func (ecgRepositoryPostgres *ecgRepositoryPostgres) FindEcgByID(ctx context.Context, id string) (*ecg_entity.Ecg, error) {
	ecg := new(ecg_entity.Ecg)

	row := ecgRepositoryPostgres.db.QueryRowContext(ctx, query_FIND_ECG_BY_ID, id)
	if err := row.Err(); err != nil {
		if err == sql.ErrNoRows {
			return nil, pkg_error.NewNotFound(err, "Ecg Not Found")
		}
		return nil, pkg_error.NewBadRequest(err, "Error while getting Ecg by id")
	}

	err := row.Scan(ecg)
	if err != nil {
		return nil, pkg_error.NewBadRequest(err, "Error Parsing Ecg Data")
	}

	return ecg, nil
}
