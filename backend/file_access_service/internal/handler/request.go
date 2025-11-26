package handler

type ErrorResp struct {
	Error string `json:"error"`
}

type AddFileReq struct{
	FileID string `json:"file_id" validate:"required"`
	UserID string `json:"user_id" validate:"required"`
}

type AddAccess struct{
	FileID string `json:"file_id" validate:"required"`
	UserID string `json:"user_id" validate:"required"`
}