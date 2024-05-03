package ecg_repository_postgres

import (
	"context"
	ecg_entity "tugas-sirs/internal/modules/ecg/model/entity"
	pkg_error "tugas-sirs/pkg/error"
)

// FindAllEcg retrieves all ECG records from the database.
// It takes the context as a parameter.
// If there are no ECG records found, it returns nil and a NotFound error.
// If an error occurs during the retrieval process, it returns nil and a BadRequest error.
// Otherwise, it returns a slice of Ecg pointers and nil.
func (ecgRepositoryPostgres *ecgRepositoryPostgres) FindAllEcg(ctx context.Context) (*[]*ecg_entity.Ecg, error) {
	ecgs := new([]*ecg_entity.Ecg)

	rows, err := ecgRepositoryPostgres.db.QueryxContext(ctx, FINDALLECGS)
	ecgRepositoryPostgres.db.SelectContext(ctx,ecgs,FINDALLECGS)
	if err != nil {
		return nil, pkg_error.NewBadRequest(err, "Failed to Get All Ecgs")
	} else if rows.Err() != nil {
		return nil,rows.Err()
	}

	return ecgs, nil
}
