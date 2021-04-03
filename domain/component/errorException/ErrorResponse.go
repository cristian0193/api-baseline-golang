package errorException

import (
	"ws-baseline-golang/domain/dto"
)

type ErrorResponse interface {
	SetError(string, dto.Response)
	Error() string
	Response() dto.Response
}
