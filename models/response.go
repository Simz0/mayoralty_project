package models

type SuccessResponse struct {
	Success string
	Message string
	Data    DataFormat
}

type FailResponse struct {
	Error string
}

type DataFormat struct {
	Data    string
	IntData int
}
