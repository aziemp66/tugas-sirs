package ecg_repository_postgres

import (
	"context"
	"errors"
	ecg_entity "tugas-sirs/internal/modules/ecg/model/entity"
	pkg_error "tugas-sirs/pkg/error"
)

// UpdateEcgByID updates an existing ECG record in the database by its ID.
// It takes the context, ID of the ECG record, and updated fields (name, age, sex, and heartFailure) as parameters.
// If the update operation is successful, it returns nil.
// If an error occurs during the update process, it returns a BadRequest error.
func (ecgRepositoryPostgres *ecgRepositoryPostgres) UpdateEcgByID(ctx context.Context, id string, name string, age int, sex bool, heartFailure ecg_entity.HeartFailureClass) error {
	result, err := ecgRepositoryPostgres.db.ExecContext(ctx, UPDATEEGCBYID, id, name, age, sex, string(heartFailure))
	if err != nil {
		return pkg_error.NewBadRequest(err, "Error Updating Ecg By ID")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return errors.New("not one of ecg rows are affected")
	}

	return nil
}
