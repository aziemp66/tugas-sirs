package ecg_repository_postgres

const (
	CREATEECG     = "INSERT INTO ecgs (id,name,age,sex,heart_failure) VALUES ($1,$2,$3,$4,$5)"
	FINDECGBYID   = "SELECT id,name,age,sex,heart_failure FROM ecgs WHERE id = $1"
	FINDALLECGS   = "SELECT id,name,age,sex,heart_failure FROM ecgs"
	UPDATEEGCBYID = "UPDATE ecgs set name = $2, age = $3, sex = $4, heart_failure = $5 WHERE id = $1"
	DELETEECG     = "DELETE FROM ecgs WHERE id = $1"
)
