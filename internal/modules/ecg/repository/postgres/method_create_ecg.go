package ecg_repository_postgres

import (
	"context"
	ecg_entity "tugas-sirs/internal/modules/ecg/model/entity"
	pkg_error "tugas-sirs/pkg/error"

	"github.com/google/uuid"
)

// CreateEcg creates a new ECG record in the database.
// It takes the name, age, sex, and heart failure class as parameters
// and returns the ID of the newly created record.
// If an error occurs during the creation process, it returns an error.
func (ecgRepositoryPostgres *ecgRepositoryPostgres) CreateEcg(ctx context.Context, name string, age int, sex bool, heartFailure ecg_entity.HeartFailureClass) (string, error) {
	id := uuid.NewString()

	_, err := ecgRepositoryPostgres.db.ExecContext(ctx, query_CREATE_ECG, id, name, age, sex, string(heartFailure))
	if err != nil {
		return "", pkg_error.NewBadRequest(err, "Failed to Create Ecg")
	}

	return id, err
}
