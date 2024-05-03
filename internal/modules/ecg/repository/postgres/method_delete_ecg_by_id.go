package ecg_repository_postgres

import (
	"context"
	pkg_error "tugas-sirs/pkg/error"
)

// DeleteEcgByID deletes an ECG record from the database based on its ID.
// It takes the context and the ID of the ECG record as parameters.
// If the deletion is successful, it returns nil.
// If an error occurs during the deletion process, it returns an error.
func (ecgRepositoryPostgres *ecgRepositoryPostgres) DeleteEcgByID(ctx context.Context, id string) error {
	_, err := ecgRepositoryPostgres.db.ExecContext(ctx, query_DELETE_ECG_BY_ID,id)
	if err != nil {
		return pkg_error.NewBadRequest(err, "Failed to Delete Ecg")
	}

	return nil
}
