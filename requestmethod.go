package restserver

// Method exported
// Method
type RequestMethod int

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

// ParseRequestMethod exported
// ParseRequestMethod ...
func ParseToRequestMethod(restMethodToParse string) RequestMethod {

	i := UNDEFINED

	for arrayIndex, method := range methodNames {
		if method == restMethodToParse {
			i = arrayIndex
			break
		}
	}

	return RequestMethod(i)
}
