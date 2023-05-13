package main

// arabic + translation combnned
type CompleteSurah struct {
	Aarabic    Ayas
	Tranlaions []TranslatedVerses
	SurahInfo  ChapterInfo
}

// Verse arabic schema
type Verse struct {
	Id   int
	Text string `json:"text_uthmani"`
	Key  string `json:"verse_key"`
}

type Ayas struct {
	Verses []Verse
}

// Chaper

type ChapterInfo struct {
	Chapter SurahInformation
}

// surah info
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
