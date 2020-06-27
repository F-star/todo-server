package commonError

type commonErr struct {
	Err int    `json:"err"`
	Msg string `json:"msg"`
}

var (
	NotFound      = commonErr{Err: 1, Msg: "not found"}
	AlreadyExists = commonErr{Err: 2, Msg: "already exits"}
	Unauthorized  = commonErr{Err: 3, Msg: "unauthorized"}
	DBError       = commonErr{Err: 4, Msg: "db operations error"}
)
