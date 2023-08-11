package surah

// this is a modified vertion from  "https://api.quran.com" version
type SurahByWord struct {
	Verses []struct {
		Id           int
		VserseNumber int    `json:"verse_number"`
		VerseKey     string `json:"verse_key"`
		Words        []struct {
			Id             int
			Position       int
			TextUthmani    string `json:"text_uthmani"`
			QpcUthmaniHafs string `json:"qpc_uthmani_hafs"`
			Text           string
			Translation    struct {
				Text         string
				LanguageName string `json:"language_name"`
				LanguageId   int    `json:"language_id"`
			}
		}
	}
}
