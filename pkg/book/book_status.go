package book

type BookStatus int

const (
	StatusAvailable BookStatus = iota
	StatusUnavailable
	StatusRented
)

func (s BookStatus) String() string {
	switch s {
	case StatusAvailable:
		return "available"
	case StatusUnavailable:
		return "unavailable"
	case StatusRented:
		return "rented"
	default:
		return "unknown"
	}
}
