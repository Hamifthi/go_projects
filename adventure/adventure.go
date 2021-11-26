package adventure

type Option struct {
	Text string `jsonReader:"text"`
	Arc  string `jsonReader:"arc"`
}

type Story struct {
	Title   string   `jsonReader:"title"`
	Story   []string `jsonReader:"story"`
	Options []Option `jsonReader:"options"`
}

type Reader interface {
	readInput(filepath string) map[string]Story
}
