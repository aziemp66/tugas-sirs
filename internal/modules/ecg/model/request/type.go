package ecg_request

type CreateEcgRequest struct {
	Name         string `json:"name"`
	Age          int    `json:"age"`
	Sex          bool   `json:"sex"`
	HeartFailure int    `json:"heart_failure"`
}

type UpdateEcgByID struct {
	Name         string `json:"name"`
	Age          int    `json:"age"`
	Sex          bool   `json:"sex"`
	HeartFailure int    `json:"heart_failure"`
}
