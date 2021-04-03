package errorException

import (
	"ws-baseline-golang-v2/domain/dto"
)

type ErrorResponse interface {
	SetError(string, dto.Response)
	Error() string
	Response() dto.Response
}
