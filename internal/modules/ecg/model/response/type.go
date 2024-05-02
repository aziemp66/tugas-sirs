package ecg_response

type EcgResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Age          int    `json:"age"`
	Sex          bool   `json:"sex"`
	HeartFailure string `json:"heart_failure"`
}
