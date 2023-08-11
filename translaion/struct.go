package translaion

type TransLations []Translation

type Translation struct {
	Translations []struct {
		Text string
	}

	Meta struct {
		Author          string `json:"author_name"`
		TranslationName string `json:"translation_name"`
	}
}

// translation dirs || translaion available
type AvailableTranslations []AvailableTranslation

type AvailableTranslation struct {
	Language     string   `json:"language"`
	TransLations []string `json:"translations"`
}
