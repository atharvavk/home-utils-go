package models

type ErrorCode int

const (
	NO_SUCH_USER = iota + 1000
	GEYSER_ALREADY_ON
	GEYSER_OFF_OR_INVALID_USER
  INVALID_REQUEST
)

const (
	NO_SUCH_USER_MESSAGE       = "NO SUCH USER"
	GEYSER_ALREADY_ON_MESSAGE  = "GEYSER ALREADY ON"
	GEYSER_OFF_OR_INVALID_USER_MESSAGE = "GEYSER OFF OR INVALID USER"
  INVALID_REQUEST_MESSAGE = "INVALID REQUEST"
)

type BadReqErr struct {
	Code   ErrorCode `json:"code"`
	Messge string    `json:"message"`
}

func (bre BadReqErr) Error() string {
	return bre.Messge
}

func NewBadReqError(code ErrorCode) BadReqErr {
	errorMessage := ""

	switch code {
	case NO_SUCH_USER:
		errorMessage = NO_SUCH_USER_MESSAGE
	case GEYSER_ALREADY_ON:
		errorMessage = GEYSER_ALREADY_ON_MESSAGE
	case GEYSER_OFF_OR_INVALID_USER:
		errorMessage = GEYSER_OFF_OR_INVALID_USER_MESSAGE
  case INVALID_REQUEST:
    errorMessage = INVALID_REQUEST_MESSAGE
  default:
    panic("invalid error code")
	}

  return BadReqErr{
    Code: code,
    Messge: errorMessage,
  }
}
