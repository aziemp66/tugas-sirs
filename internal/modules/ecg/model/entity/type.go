package ecg_entity

type Ecg struct {
	ID           string            `db:"id"`
	Name         string            `db:"name"`
	Age          int               `db:"age"`
	Sex          bool              `db:"sex"`
	HeartFailure HeartFailureClass `db:"heart_failure"`
}

type HeartFailureClass string
