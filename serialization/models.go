package serialization

type AppRequest struct {
	UserId    string
	ClientId  string
	RequestId int64
	Service   string
	Data      string
	Error     string
}
