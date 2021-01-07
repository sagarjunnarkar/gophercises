package cyoa

import (
	"encoding/json"
	"io"
)

// JsonStory bla
func JsonStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil
}

//Story json
type Story map[string]Chapter

// Chapter struct
type Chapter struct {
	Title     string   `json:"title"`
	Paragraph []string `json:"story"`
	Options   []Option `json:"options"`
}

// Option struct
type Option struct {
	Text   string `json:"text"`
	Chaper string `json:"arc"`
}
