package puppy

type Error struct {
	Message  string
	Code     int
	CausedBy string
}

// Not being able to do map consts really trips me.
// There must be a better way.
const (
	Err400BadRequest     = 400
	Err404NotFound       = 404
	Err409Conflict       = 409
	Err500InternalError  = 500
	Err501NotImplemented = 501
)

const (
	Err400BadRequestS     = "400 Bad Request"
	Err404NotFoundS       = "404 Not Found"
	Err409ConflictS       = "409 Conflict"
	Err500InternalErrorS  = "500 Internal Server Error"
	Err501NotImplementedS = "501 Not Implemented"
)

func (e Error) Error() string {
	switch e.Code {
	case Err400BadRequest:
		e.Message = Err400BadRequestS
	case Err404NotFound:
		e.Message = Err404NotFoundS
	case Err409Conflict:
		e.Message = Err409ConflictS
	case Err500InternalError:
		e.Message = Err500InternalErrorS
	default:
		e.Code = Err501NotImplemented
		e.Message = Err501NotImplementedS
	}
	return e.Message
}

func NewError(i int) error {
	var e Error
	e.Code = i
	e.CausedBy = ""
	e.Message = e.Error()
	return e
}
