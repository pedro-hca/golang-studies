package media

type Catalogable interface {
	SetMovie(title string, rating Rating, boxOffice float32)
	GetMovie() Movie
	SetRating(newRating Rating)
	SetTitle(newTitle string)
	SetboxOffice(newBoxOffice float32)
}

type Movie struct {
	title     string
	rating    Rating
	boxOffice float32
}

type Rating string

const (
	R    = "R (Restricted)"
	G    = "G (General audiences)"
	PG   = "PG (Parental Guidance)"
	PG13 = "PG13 (Parental Caution)"
	NC17 = "NC17 (No children under 17)"
)

func NewMovie(title string, rating Rating, boxOffice float32) *Movie {
	return &Movie{
		title:     title,
		rating:    rating,
		boxOffice: boxOffice,
	}
}

func (m *Movie) SetMovie(title string, rating Rating, boxOffice float32) {
	m.title = title
	m.rating = rating
	m.boxOffice = boxOffice
}

func (m *Movie) GetMovie() Movie {
	return *m
}

func (m *Movie) SetTitle(title string) {
	m.title = title
}
func (m *Movie) SetRating(rating Rating) {
	m.rating = rating
}
func (m *Movie) SetboxOffice(boxOffice float32) {
	m.boxOffice = boxOffice
}
