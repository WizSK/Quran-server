package main

// arabic + translation combnned
type CompleteSurah struct {
	Aarabic           Ayas
	Translaions       []TranslatedVerses
	BanglaTranslation TranslatedVerses // local lang of mine.
	SurahInfo         ChapterInfo
}

// Verse arabic schema
type Verse struct {
	Id int
	// for the black () in the text
	Text string `json:"text_imlaei"`
	// Text string `json:"text_uthmani","text_imlaei"`
	Key string `json:"verse_key"`
}

type Ayas struct {
	Verses []Verse
}

// Chaper
type ChapterInfo struct {
	Chapter SurahInformation
}

// surah info
type SurahIndexPrefixed struct {
	Prefixes    []string
	ChaptersIdx ChaptersIdx
}

type ChaptersIdx struct {
	Chapters []Sura
}

type Sura struct {
	Name    string `json:"name_simple"`
	AraName string `json:"name_arabic"`
	Id      int
}

type SurahInformation struct {
	Id          int
	Name        string `json:"name_simple"`
	Bismillah   bool   `json:"bismillah_pre"`
	VersesCount int    `json:"verses_count"`
	Place       string `json:"revelation_place"`
}

// Translations
type TranslatedVerses struct {
	Translations []TranslatedVerse
	Meta         AboutTrnaslation
}

type TranslatedVerse struct {
	Text string
}

type AboutTrnaslation struct {
	Author      string `json:"author_name"`
	Translation string `json:"translation_name"`
}

// word by word
type CompleteSurahWordByWord struct {
	Aarabic           Ayas
	Translaions       []TranslatedVerses
	BanglaTranslation TranslatedVerses // local lang of mine.
	// Aarabic           Ayas
	// Translaions       []TranslatedVerses
	// BanglaTranslation TranslatedVerses // local lang of mine.
	SurahInfo       ChapterInfo
	WordByWordArray []WordByWord
}

type WordByWord struct {
	Verses []WordVerse
}

type WordVerse struct {
	Id       int
	VerseKey string `json:"verse_key"`
	Words    []Word
	IndexForTrans int `json:",omitempty"`
}

type Word struct {
	Position    int
	Text        string
	Translation Translation
}

type Translation struct {
	Text string
}
