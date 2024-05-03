package ecg_repository_postgres

import (
	"context"
	"database/sql"
	ecg_entity "tugas-sirs/internal/modules/ecg/model/entity"
	pkg_error "tugas-sirs/pkg/error"
)

// FindAllEcg retrieves all ECG records from the database.
// It takes the context as a parameter.
// If there are no ECG records found, it returns nil and a NotFound error.
// If an error occurs during the retrieval process, it returns nil and a BadRequest error.
// Otherwise, it returns a slice of Ecg pointers and nil.
func (ecgRepositoryPostgres *ecgRepositoryPostgres) FindAllEcg(ctx context.Context, tx *sql.Tx) (*[]*ecg_entity.Ecg, error) {
	ecgs := new([]*ecg_entity.Ecg)

	rows, err := tx.QueryContext(ctx, FINDALLECGS)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, pkg_error.NewNotFound(err, "Cannot Find Any Ecgs")
		}
		return nil, pkg_error.NewBadRequest(err, "Failed to Get All Ecgs")
	}

	for rows.Next() {
		ecg := new(ecg_entity.Ecg)
		err = rows.Scan(ecg)
		if err != nil {
			return nil, pkg_error.NewBadRequest(err, "Cannot Get Ecg Rows")
		}

		*ecgs = append(*ecgs, ecg)
	}

	return ecgs, nil
}
