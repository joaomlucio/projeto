package dtos

type ErrorResponse struct {
    FailedField string
    Tag         string
    Value       string
}