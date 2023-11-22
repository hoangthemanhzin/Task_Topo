package models

const (
	GENERIC_ERROR           string = "GENERIC"
	N1N2_MSG_TRANSFER_ERROR string = "N1N2_TRANSFER"
)

type Error interface {
	GetType() string
	GetStatus() int32
	Error() string
}
