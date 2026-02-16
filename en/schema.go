package en

// The word that is the alternative form of another word.
// field `Word` contains the linked word, and `Extra` contains
// optional additional text.
type AltOf struct {
	Word  string  `json:"word" db:"INDEX"`
	Extra *string `json:"extra,omitempty"`
}

type LinkageData struct {
	Alt         *string  `json:"alt,omitempty"`
	English     *string  `json:"english,omitempty"` // DEPRECATED in favour of "translation"
	Translation string   `json:"translation"`
	Extra       *string  `json:"extra,omitempty"`
	Qualifier   *string  `json:"qualifier,omitempty"`
	RawTags     []string `json:"raw_tags,omitempty"`
	Roman       *string  `json:"roman,omitempty"`
	// Japanese Kanji and furigana
	Ruby      [][]string `json:"ruby,omitempty"`
	Sense     *string    `json:"sense,omitempty"`
	Tags      []string   `json:"tags,omitempty"`
	Taxonomic *string    `json:"taxonomic,omitempty"`
	Topics    []string   `json:"topics,omitempty"`
	Urls      []string   `json:"urls,omitempty"`
	Word      string     `json:"word" db:"INDEX"`
}

type ExampleData struct {
	Alt                    *string  `json:"alt,omitempty"`
	English                *string  `json:"english,omitempty"` // DEPRECATED in favour of "translation"
	Translation            *string  `json:"translation,omitempty"`
	BoldTranslationOffsets [][2]int `json:"bold_translation_offsets,omitempty"`
	// English-language parenthesized note from the beginning of a non-english example
	Note *string `json:"note,omitempty"`
	Ref  *string `json:"ref,omitempty"`
	// romanization (for some languages written in non-Latin scripts)
	Roman            *string  `json:"roman,omitempty"`
	BoldRomanOffsets [][2]int `json:"bold_roman_offsets,omitempty"`
	// Japanese Kanji and furigana
	Ruby [][]string `json:"ruby,omitempty"`
	// the example text
	Text            string   `json:"text"`
	BoldTextOffsets [][2]int `json:"bold_text_offsets,omitempty"`
	Tags            []string `json:"tags,omitempty"`
	RawTags         []string `json:"raw_tags,omitempty"`
}

type FormOf struct {
	Word  string  `json:"word" db:"INDEX"`
	Extra *string `json:"extra,omitempty"`
	Roman *string `json:"roman,omitempty"`
}

type LinkData [][]string

// It is the alias of `PlusObjTemplateData`.
type ExtraTemplateData struct {
	Tags    []string `json:"tags,omitempty"`
	Words   []string `json:"words,omitempty"`
	Meaning string   `json:"meaning"`
}

// python3: TemplateArgs = dict[Union[int, str], str]
//
// We will simply assume all keys are strings. In this way,
// fields can be more easily handled.

type TemplateArgs map[string]string

type TemplateData struct {
	// dictionary mapping argument names to their cleaned values. Positional
	// arguments have keys that are numeric strings, starting with "1".
	Args TemplateArgs `json:"args"`
	// the (cleaned) text the template expands to.
	Explansion string `json:"explansion"`
	// name of the template
	Name      string             `json:"name"`
	ExtraData *ExtraTemplateData `json:"extra_data,omitempty"`
}

type DescendantData struct {
	// Wiktionary language code
	LangCode string `json:"lang_code"`
	// Language name
	Lang        string           `json:"lang"`
	Word        string           `json:"word"`
	Roman       string           `json:"roman"`
	Tags        []string         `json:"tags,omitempty"`
	RawTags     []string         `json:"raw_tags,omitempty"`
	Descendants []DescendantData `json:"descendants,omitempty"`
	// Japanese Kanji and furigana
	Ruby  [][]string `json:"ruby,omitempty"`
	Sense string     `json:"sense,omitempty"`
}

type FormData struct {
	Form   string  `json:"form"`
	HeadNr int     `json:"head_nr"`
	Ipa    *string `json:"ipa,omitempty"`
	Roman  *string `json:"roman,omitempty"`
	// Japanese Kanji and furigana
	Ruby    [][]string `json:"ruby,omitempty"`
	Source  *string    `json:"source,omitempty"`
	Tags    []string   `json:"tags,omitempty"`
	RawTags []string   `json:"raw_tags,omitempty"`
	Topics  []string   `json:"topics,omitempty"`
}

type Hyphenation struct {
	Parts []string `json:"parts,omitempty"`
	Tags  []string `json:"tags,omitempty"`
}

type SoundData struct {
	// name of a sound file in WikiMedia Commons
	Audio *string `json:"audio,omitempty"`
	// IPA string associated with the audio file, generally giving IPA
	// transcription of what is in the
	AudioIpa *string `json:"audio-ipa,omitempty"`
	// English pronunciation respelling
	Enpr    *string `json:"enpr,omitempty"`
	Form    *string `json:"form,omitempty"`
	Hangeul *string `json:"hangeul,omitempty"`
	// list of homophones for the word.
	//
	// Note: a homophone is a word that is pronounced the same as another
	// word but differs in meaning **or** in spelling.
	//
	// This field is not documented in the official python TypedDict
	// models but added according to the project README.
	Homophone *string `json:"homophone,omitempty"`
	// list of hyphenations.
	//
	// Note: syllabification or syllabication, hyphenation, is the separation of a
	// word into syllables, whether spoken, written or signed.
	Hyphenation *string `json:"hyphenation,omitempty"`
	// International Phonetic Alphabet. /.../ or [...].
	Ipa *string `json:"ipa,omitempty"`
	// URL for an MP3 format sound file
	Mp3Url *string `json:"mp3_url,omitempty"`
	Note   *string `json:"note,omitempty"`
	// URL for an OGG Vorbis format sound file
	OggUrl *string `json:"ogg_url,omitempty"`
	Other  *string `json:"other,omitempty"`
	Rhymes *string `json:"rhymes,omitempty"`
	// other labels or context information attached to the pronunciation
	// entry (e.g., might indicate regional variant or dialect)
	Tags []string `json:"tags,omitempty"`
	// text associated with an audio file (often not very useful)
	Text   *string  `json:"text,omitempty"`
	Topics []string `json:"topics,omitempty"`
	// Chinese word pronunciation
	ZhPron *string `json:"zh-pron,omitempty"`
}

type TranslationData struct {
	// optional alternative form of the translation (e.g., in a different script)
	Alt *string `json:"alt,omitempty"`
	// Wiktionary's 2 or 3-letter language code for the language the language the
	// translation is for.
	LangCode string `json:"lang_code"`
	// Wiktionary's 2 or 3-letter language code for the language the language the
	// translation is for.
	//
	// DEPRECATED in favour of `lang_code
	Code *string `json:"code,omitempty"`
	// English text, generally clarifying the target sense of the translation.
	//
	// DEPRECATED in favour of `translation`
	English     *string `json:"english"`
	Translation string  `json:"translation"`
	// The language name that the translation is for.
	Lang string `json:"lang"`
	// optional text describing or commenting on the translation
	Note *string `json:"note,omitempty"`
	// optional romanization of the translation (when in non-Latin characters)
	Roman *string `json:"roman,omitempty"`
	// Optional sense indicating the meaning for which this is a translation
	// (this is a free-text string, and may not match any gloss exactly)
	//
	// P.S. I doubt there is grammar fault from the official documentation
	Sense *string `json:"sense,omitempty"`
	// optional list of qualifiers for the translations, e.g., gender
	Tags []string `json:"tags,omitempty"`
	// optional taxonomic name of an organism mentioned in the
	// translation.
	Taxonomic *string  `json:"taxonomic,omitempty"`
	Topics    []string `json:"topics,omitempty"`
	// the translation in the specified language (may be missing when `note`
	// is present)
	Word *string `json:"word,omitempty"`
}

// Xxyzz's East Asian etymology example data
type EtymologyExample struct {
	English     *string  `json:"english,omitempty"` // DEPRECATED in favour of `translation`
	Translation string   `json:"translation"`
	RawTags     []string `json:"raw_tags,omitempty"`
	Ref         *string  `json:"ref,omitempty"`
	Roman       *string  `json:"roman,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Text        *string  `json:"text,omitempty"`
	Type        *string  `json:"type,omitempty"`
}

type ReferenceData struct {
	Text string  `json:"text"`
	Refn *string `json:"refn,omitempty"`
}

type AttestationData struct {
	Date       string          `json:"date"`
	References []ReferenceData `json:"references,omitempty"`
}

type SenseData struct {
	// list of words that his sense is an inflected form of; this is a
	// list of dictionaries, with field `word` containing the linnked word
	// and optionally `extra` containing additional text.
	AltOf []AltOf `json:"alt_of,omitempty"`
	// sense-disambiguated antonym linkages for the word.
	Antonyms []LinkageData `json:"antonyms,omitempty"`
	// list of sense-disambiguated category names extracted from (a
	// subset) of the Category links on the page
	Categories []string `json:"categories,omitempty"`
	CompoundOf []AltOf  `json:"compound_of,omitempty"`
	// sense-disambiguated coordinate_terms linkages
	CoordinateTerms []LinkageData `json:"coordinate_terms,omitempty"`
	// list of usage examples, each example being a dictionary with `text`
	// field containing the example text, optional `ref` field containing a source reference,
	// optional `english` field containing English translation, optional `type" field containing
	// example type (currently `example` or `quotation` if present), optional
	// `roman` field containing romanization for some languages written in non-Latin scripts),
	// and optional (rare) `note` field contains English-language parenthesized note from the
	// beginning of a non-english example.
	Examples []ExampleData `json:"examples,omitempty"`
	FormOf   []FormOf      `json:"form_of,omitempty"`
	// list of gloss strings for the word sense (usually only one). This has
	// been cleaned, and should be straightforward text with no tagging.
	Glosses []string `json:"glosses,omitempty"`
	HeadNr  int      `json:"head_nr"`
	// sense-disambiguated linkages indicating being part of something
	// (not systematically encoded).
	Holonyms []LinkageData `json:"holonyms,omitempty"`
	// sense-disambiguated hypernym linkages for the word
	Hypernyms []LinkageData `json:"hypernyms,omitempty"`
	// sense-disambiguated hyponym linkages for the word
	Hyponyms  []LinkageData `json:"hyponyms,omitempty"`
	Instances []LinkageData `json:"instances,omitempty"`
	Links     []LinkData    `json:"links,omitempty"`
	// sense-disambiguated linkages indicating having a part (fairly rare)
	Meronyms  []LinkageData `json:"meronyms,omitempty"`
	Qualifier *string       `json:"qualifier,omitempty"`
	// list of gloss strings for the word sense, with less cleaning than
	// `glosses`. In particular, parenthesized parts that have been parsed from the gloss
	// into `tags` and `topics` are still present here. This version may be easier for
	// humans to interpret.
	RawGlosses []string `json:"raw_glosses,omitempty"`
	// sense-disambiguated related word linkages for the word.
	Related []LinkageData `json:"related,omitempty"`
	// list of textual indentifiers collected for the sense. If there is a
	// QID for the entry (e.g., Q123), those are stored in the wikidata field.
	Senseid []string `json:"senseid,omitempty"`
	// sense-disambiguated synonym linkages for the word
	Synonyms []LinkageData `json:"synonyms,omitempty"`
	// list of qualifiers and tags for the gloss. This is a list of strings,
	// and may include words such as "archaic", "colloquial", "present",
	// "participle", "plural", "feminine", and many others (new words may
	// appear arbitrarily).
	Tags      []string `json:"tags,omitempty"`
	Taxonomic *string  `json:"taxonomic,omitempty"`
	// list of sense-disambiguated topic names (kind of similar to
	// categories but determined differently).
	Topics []string `json:"topics,omitempty"`
	// list of QIDs (e.g., Q123) for the sense
	Wikidata []string `json:"wikidata,omitempty"`
	// linst of Wikipedia page titles (with optional language code prefix)
	Wikipedia    []string          `json:"wikipedia,omitempty"`
	Attestations []AttestationData `json:"attestations,omitempty"`
}

// Etymological information is stored under the `etymology_text`
// and `etymology_templates` keys in the word's data. When multiple
// part-of-speech are listed under the same etymology, the same data
// is copied to each part-of-speech entry under that etymology.
type WordData struct {
	Abbreviations []LinkageData `json:"abbreviations,omitempty"`
	// list of words that his sense is an alternative form of; this is a list
	// of dictionaries, with field `word` containing the linked word and
	// optionally `extra` containing additional text.
	AltOf []AltOf `json:"alt_of,omitempty"`
	// non-disambiguated antonym linkages for the word
	Antonyms []LinkageData `json:"antonyms,omitempty"`
	// list of non-disambiguated categories for the word
	Categories []string `json:"categories,omitempty"`
	// non-disambiguated coordinate term linkages for the word
	CoordinateTerms []LinkageData `json:"coordinate_terms,omitempty"`
	// non-disambiguated derived word linkages for the word
	Derived []LinkageData `json:"derived,omitempty"`
	// descendants of the word
	Descendants []DescendantData `json:"descendants,omitempty"`

	EtymologyExamples []EtymologyExample `json:"etymology_examples,omitempty"`
	// for words with multiple numbererd etymologies, this contains the number of the etymology
	// under which this entry appeared.
	EtymologyNumber *int `json:"etymology_number,omitempty"`
	// templates and their arguments and expansions from the etymology section.
	// This can be used to easily parse etymological relations. Certain common
	// templates that do not signify etymological relations are not included.
	//
	//
	// The `etymology_templates` field contains a list of templates from
	// the etymology section. Some common templates considered not relevent
	// for etymological information have been removed (e.g., `redlink category
	// and `isValidPageName`). The list also includes nested templates referenced
	// from templates directly used in the etymology description.
	EtymologyTemplates []TemplateData `json:"etymology_templates,omitempty"`
	// etymology section as cleaned text
	//
	// The `etymology_text` field contains the contents of the whole etymology
	// section cleaned into human-readable text (i.e., templates have been
	// expanded and HTML tags removed, among other things).
	EtymologyText *string  `json:"etymology_text,omitempty"`
	FormOf        []FormOf `json:"form_of,omitempty"`
	// list of inflected or alternative forms specified for the word (e.g.,
	// plural, comparative, superlative, roman script version). This is a list of dictionaries,
	// where each dictionary has a `form` key and a `tags` key. The `tags` identify what type
	// of form it is. It may also contain "ipa", "roman", and "source" fields. The form can be
	// "-" when the word is marked as not having that form (some of those will be word-specific,
	// while others are language-specific; post-processing can drop such forms when no word has
	// a value for that tag combination).
	Forms []FormData `json:"forms,omitempty"`
	// part-of-speech specific head tags for the word. This basically
	// just captures the templates (their name and arguments) as a list of
	// dictionaries. Most applications may want to ignore this.
	HeadTemplates []TemplateData `json:"head_templates,omitempty"`
	Holonyms      []LinkageData  `json:"holonyms,omitempty"`
	Hyphenation   []string       `json:"hyphenation,omitempty"` // Being deprecated
	Hyphenations  []Hyphenation  `json:"hyphenations,omitempty"`
	// non-disambiguated hypernym linkages for the word
	Hypernyms []LinkageData `json:"hypernyms,omitempty"`
	// non-disambiguated hyponym linkages for the word
	Hyponyms []LinkageData `json:"hyponyms,omitempty"`
	// conjugation and declension templates found for the word, as dictionaries.
	// These basically capture the language-specific inflection template for the word.
	// Note that for some languages inflection information is also contained in `head_templates`.
	// According to [wiktextract](github.com/tatuylonen/wiktextract), inflections
	// from the inflection tables will be parsed into forms, so there is usually no
	// need to use the `inflection_templates` data.
	InflectionTemplates []TemplateData `json:"inflection_templates,omitempty"`
	InfoTemplates       []TemplateData `json:"info_templates,omitempty"`
	Instances           []LinkageData  `json:"instances,omitempty"`
	// name of the language this word belongs to (e.g., `English`)
	Lang string `json:"lang" db:"INDEX"`
	// Wiktionary language code corresponding to `lang` key (e.g., `en`)
	LangCode       string `json:"lang_code" db:"INDEX"`
	LiteralMeaning string `json:"literal_meaning"`
	// non-disambiguated meronym linkages for the word
	Meronyms      []LinkageData `json:"meronyms,omitempty"`
	OriginalTitle string        `json:"original_title"`
	// part-of-speech, such as "noun", "verb", "adj", "adv", "pron", "determiner",
	// "prep" (preposition), "postp" (postposition), and many others.
	Pos       string        `json:"pos"`
	Proverbs  []LinkageData `json:"proverbs,omitempty"`
	Redirects []string      `json:"redirects,omitempty"`
	// non-ambiguated related word linkages for the word
	Related []LinkageData `json:"related,omitempty"`
	// list of word senses (dictionaries) for this word/part-of-speech
	Senses []SenseData `json:"senses,omitempty"`
	// list of dictionaries containing pronunciation, hyphenation, rhyming, and related
	// information. Each dictionary may have a `tags` key containing tags that clarify
	// what kind of form that entry is. Different types of information are stored in
	// different fields: `ipa` is
	// [IPA](https://en.wikipedia.org/wiki/International_Phonetic_Alphabet) pronunciation,
	// `enPR` is [enPR](https://en.wikipedia.org/wiki/Pronunciation_respelling_for_English)
	// pronunciation, `audio` is name of sound file in Wikimedia commons.
	Sounds []SoundData `json:"sounds,omitempty"`
	// non-disambiguated synonym linkages for the word
	Synonyms []LinkageData `json:"synonyms,omitempty"`
	// non-disambiguated translation entries
	Translations []TranslationData `json:"translations,omitempty"`
	Troponyms    []LinkageData     `json:"troponyms,omitempty"`
	// non-disambiguated Wikidata identifier
	Wikidata  []string `json:"wikidata,omitempty"`
	Wikipedia []string `json:"wikipedia,omitempty"`
	// the word form
	Word     string        `json:"word" db:"INDEX"`
	Anagrams []LinkageData `json:"anagrams,omitempty"`
}
