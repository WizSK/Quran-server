package translaion

type TransLations []Translation

type Translation struct {
	Translations []struct {
		Text string
	}

	Meta struct {
		Author          string `json:"author_name"`
		TranslationName string `json:"translation_name"`
		Filters         struct {
			ResourceId    int `json:"resource_id"`
			ChapterNumber string `json:"chapter_number"`
		}
	}
}

// translation dirs || translaion available
type AvailableTranslations []AvailableTranslation

type AvailableTranslation struct {
	Language     string   `json:"language"`
	TransLations []string `json:"translations"`
}
