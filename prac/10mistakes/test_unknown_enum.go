package main

type Status unit32

// Bad ...
const (
	StatusOpen Status = iota
	StatusClosed
	StatusUnknown
)

type Request struct {
	ID        int    `json:"id"`
	Timestamp int    `json:"timestamp"`
	Status    Status `json:"status"`
}

// Good ...
const (
	StatusUnknown Status = iota
	StatusOpen
	StatusClosed
)

/* No status field: So the Status will be unknown
{
	"Id": 1235,
	"Timestamp": 1563362390
}
*/

func main() {
	// ... ------ ... -------... ------ ... -------
}
