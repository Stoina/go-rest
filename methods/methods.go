package rest

// Method exported
// Method
type Method int

// exported
const (
	UNDEFINED = -1
	GET       = 0
	POST      = 1
	PUT       = 2
	PATCH     = 3
	DELETE    = 4
	HEAD      = 5
	OPTIONS   = 6
	CONNECT   = 7
	TRACE     = 8
)

var methodNames = [9]string{
	"GET",
	"POST",
	"PUT",
	"PATCH",
	"DELETE",
	"HEAD",
	"OPTIONS",
	"CONNECTOR",
	"TRACE"}

// ParseFromString exported
// ParseFromString ...
func ParseFromString(restMethodToParse string) *Method {

	i := 0

	for arrayIndex, method := range methodNames {
		if method == restMethodToParse {
			i = arrayIndex
			break
		}
	}

	restMethod := Method(i)
	return &restMethod
}
