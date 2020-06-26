package hacker_news

import (
	"strconv"
	"time"
)

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

func (t *Time) UnmarshalJSON(s []byte) (err error) {
	q, err := strconv.ParseInt(string(s), 10, 64)
	if err != nil {
		return err
	}

	*(*time.Time)(t) = time.Unix(q, 0)
	return nil
}

func (t Time) Unix() int64 {
	return time.Time(t).Unix()
}

func (t Time) Time() time.Time {
	return time.Time(t).UTC()
}

func (t Time) String() string {
	return t.Time().String()
}

type NewsItem struct {
	ID          int      `json:"id"`
	Deleted     bool     `json:"deleted"`
	Type        string   `json:"type"`
	By          string   `json:"by"`
	Time        Time     `json:"time"`
	Text        string   `json:"text"`
	Dead        bool     `json:"dead"`
	Parent      string   `json:"parent"`
	Poll        string   `json:"poll"`
	Kids        []int    `json:"kids"`
	URL         string   `json:"url"`
	Score       int      `json:"score"`
	Title       string   `json:"title"`
	Parts       []string `json:"parts"`
	Descendants int      `json:"descendants"`
}

func (i NewsItem) GetTitle() string {
	if len(i.Title) > 60 {
		return i.Title[:59] + "..."
	}
	return i.Title
}
