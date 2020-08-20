package factory

type MediaType int

const (
	NewsPaper MediaType = iota
	Megazine
)

type Media interface {
	Publish()
}

type NewsPaper struct{}

func (n *NewsPaper) Publish() {

}

type Megazine struct{}

func (m *Megazine) Publish() {

}

func CreateMedia(m MediaType) Media {
	switch m {
	case NewsPaper:
		return new(NewsPaper)
	case Megazine:
		return new(Megazine)
	default:
		return nil
	}
}
