package ecg_repository_postgres

const (
	query_CREATE_ECG     = "INSERT INTO ecgs (id,name,age,sex,heart_failure) VALUES ($1,$2,$3,$4,$5)"
	query_FIND_ECG_BY_ID   = "SELECT id,name,age,sex,heart_failure FROM ecgs WHERE id = $1"
	query_FIND_ALL_EGCS= "SELECT id,name,age,sex,heart_failure FROM ecgs"
	query_UPDATE_ECG_BY_ID= "UPDATE ecgs set name = $2, age = $3, sex = $4, heart_failure = $5 WHERE id = $1"
	query_DELETE_ECG_BY_ID= "DELETE FROM ecgs WHERE id = $1"
)
