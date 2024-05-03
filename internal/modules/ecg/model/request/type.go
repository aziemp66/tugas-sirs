package ecg_request

type CreateEcgRequest struct {
	Name         string `json:"name" binding:"required"`
	Age          int    `json:"age" binding:"required"`
	Sex          bool   `json:"sex" binding:"required"`
	HeartFailure int    `json:"heart_failure" binding:"required"`
}

type UpdateEcgByID struct {
	Name         string `json:"name" binding:"required"`
	Age          int    `json:"age" binding:"required"`
	Sex          bool   `json:"sex" binding:"required"`
	HeartFailure int    `json:"heart_failure" binding:"required"`
}
