package book

type BookFormat int

const (
	FormatHardcover BookFormat = iota
	FormatPaperback
	FormatEbook
)

func (f BookFormat) String() string {
	switch f {
	case FormatHardcover:
		return "hardcover"
	case FormatPaperback:
		return "paperback"
	case FormatEbook:
		return "ebook"
	default:
		return "unknown"
	}
}
