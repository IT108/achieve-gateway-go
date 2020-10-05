package serialization

type AppRequest struct {
	UserId    string
	ClientId  string
	RequestId int64
	Service   string
	Method    string
	Data      map[string]interface{}
	Error     string
}
