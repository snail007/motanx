package motanx

import "github.com/weibocom/motan-go/core"

var (
	authFilter *AuthFilter
)

type AuthFilter struct {
	authToken string
	next      core.EndPointFilter
}

func NewAuthFilter(authToken string) *AuthFilter {
	return &AuthFilter{authToken: authToken}
}

func GetAuthFilter(authToken string) *AuthFilter {
	if authFilter == nil {
		authFilter = NewAuthFilter(authToken)
	}
	return authFilter
}

func (t *AuthFilter) GetIndex() int {
	return 1
}

func (t *AuthFilter) GetName() string {
	return "auth"
}

func (t *AuthFilter) NewFilter(url *core.URL) core.Filter {
	return authFilter
}

func (t *AuthFilter) Filter(caller core.Caller, request core.Request) core.Response {
	if t.authToken != "" && request.GetAttachment("token") != t.authToken {
		return core.BuildExceptionResponse(request.GetRequestID(),
			&core.Exception{ErrCode: 500, ErrMsg: "auth fail", ErrType: core.ServiceException})
	}
	response := t.GetNext().Filter(caller, request)
	return response
}

func (t *AuthFilter) HasNext() bool {
	return t.next != nil
}

func (t *AuthFilter) SetNext(nextFilter core.EndPointFilter) {
	t.next = nextFilter
}

func (t *AuthFilter) GetNext() core.EndPointFilter {
	return t.next
}

func (t *AuthFilter) GetType() int32 {
	return core.EndPointFilterType
}
