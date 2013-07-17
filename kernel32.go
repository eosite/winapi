// Copyright 2013 The winapi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package winapi

import (
	"syscall"
	"unicode/utf16"
)

var (
	modKernel32         = syscall.NewLazyDLL("Kernel32.dll")
	procGetLastError    = modKernel32.NewProc("GetLastError")
	procGetLocaleInfo   = modKernel32.NewProc("GetLocaleInfoW")
	procGetModuleHandle = modKernel32.NewProc("GetModuleHandleW")
	procLoadString      = modUser32.NewProc("LoadStringW")
)

func GetLastError() uint {
	ret, _, _ := procGetLastError.Call()

	return uint(ret)
}

func GetLocaleInfo(lcid LCID, lctype LCTYPE) []uint16 {
	buf := make([]uint16, 256)
	ret, _, _ := procGetLocaleInfo.Call(Ptr(lcid), Ptr(lctype), Ptr(buf), Ptr(256))

	if ret > 0 {
		return buf[:ret]
	} else {
		return nil
	}
}

func GetModuleHandle(moduleName string) HMODULE {
	var param uintptr = 0
	if moduleName != "" {
		param = GoStringToPtr(moduleName)
	}

	ret, _, _ := procGetModuleHandle.Call(param)

	return HMODULE(ret)
}

func LoadString(inst HINSTANCE, id uint) (ret string) {
	text := make([]uint16, 1024)
	r, _, _ := procLoadString.Call(Ptr(inst), Ptr(id), Ptr(text), 1024)

	if int(r) <= 0 {
		ret = ""
	} else {
		ret = string(utf16.Decode(text[0:r]))
	}

	return
}

type (
	LCID   uint32
	LCTYPE uint32
)

// Predefined locale ids
const (
	LOCALE_CUSTOM_DEFAULT     LCID = 0x0c00
	LOCALE_CUSTOM_UI_DEFAULT  LCID = 0x1400
	LOCALE_CUSTOM_UNSPECIFIED LCID = 0x1000
	LOCALE_INVARIANT          LCID = 0x007f
	LOCALE_USER_DEFAULT       LCID = 0x0400
	LOCALE_SYSTEM_DEFAULT     LCID = 0x0800
)

// Predefined LCType ids
const (
	LOCALE_SISO3166CTRYNAME  LCTYPE = 0x5a
	LOCALE_SISO3166CTRYNAME2 LCTYPE = 0x68
	LOCALE_SISO639LANGNAME   LCTYPE = 0x59
	LOCALE_SISO639LANGNAME2  LCTYPE = 0x67
)

// Predefined LANGID
const (
	LANG_NEUTRAL             LANGID = 0x0
	LANG_INVARIANT           LANGID = 0x7f
	LANG_AFRIKAANS           LANGID = 0x36
	LANG_ALBANIAN            LANGID = 0x1c
	LANG_ALSATIAN            LANGID = 0x84
	LANG_AMHARIC             LANGID = 0x5e
	LANG_ARABIC              LANGID = 0x1
	LANG_ARMENIAN            LANGID = 0x2b
	LANG_ASSAMESE            LANGID = 0x4d
	LANG_AZERI               LANGID = 0x2c
	LANG_BASHKIR             LANGID = 0x6d
	LANG_BASQUE              LANGID = 0x2d
	LANG_BELARUSIAN          LANGID = 0x23
	LANG_BENGALI             LANGID = 0x45
	LANG_BRETON              LANGID = 0x7e
	LANG_BOSNIAN             LANGID = 0x1a
	LANG_BOSNIAN_NEUTRAL     LANGID = 0x781a
	LANG_BULGARIAN           LANGID = 0x2
	LANG_CATALAN             LANGID = 0x3
	LANG_CHINESE             LANGID = 0x4
	LANG_CHINESE_SIMPLIFIED  LANGID = 0x4
	LANG_CHINESE_TRADITIONAL LANGID = 0x7c04
	LANG_CORSICAN            LANGID = 0x83
	LANG_CROATIAN            LANGID = 0x1a
	LANG_CZECH               LANGID = 0x5
	LANG_DANISH              LANGID = 0x6
	LANG_DARI                LANGID = 0x8c
	LANG_DIVEHI              LANGID = 0x65
	LANG_DUTCH               LANGID = 0x13
	LANG_ENGLISH             LANGID = 0x9
	LANG_ESTONIAN            LANGID = 0x25
	LANG_FAEROESE            LANGID = 0x38
	LANG_FARSI               LANGID = 0x29
	LANG_FILIPINO            LANGID = 0x64
	LANG_FINNISH             LANGID = 0xb
	LANG_FRENCH              LANGID = 0xc
	LANG_FRISIAN             LANGID = 0x62
	LANG_GALICIAN            LANGID = 0x56
	LANG_GEORGIAN            LANGID = 0x37
	LANG_GERMAN              LANGID = 0x7
	LANG_GREEK               LANGID = 0x8
	LANG_GREENLANDIC         LANGID = 0x6f
	LANG_GUJARATI            LANGID = 0x47
	LANG_HAUSA               LANGID = 0x68
	LANG_HEBREW              LANGID = 0xd
	LANG_HINDI               LANGID = 0x39
	LANG_HUNGARIAN           LANGID = 0xe
	LANG_ICELANDIC           LANGID = 0xf
	LANG_IGBO                LANGID = 0x70
	LANG_INDONESIAN          LANGID = 0x21
	LANG_INUKTITUT           LANGID = 0x5d
	LANG_IRISH               LANGID = 0x3c
	LANG_ITALIAN             LANGID = 0x10
	LANG_JAPANESE            LANGID = 0x11
	LANG_KANNADA             LANGID = 0x4b
	LANG_KASHMIRI            LANGID = 0x60
	LANG_KAZAK               LANGID = 0x3f
	LANG_KHMER               LANGID = 0x53
	LANG_KICHE               LANGID = 0x86
	LANG_KINYARWANDA         LANGID = 0x87
	LANG_KONKANI             LANGID = 0x57
	LANG_KOREAN              LANGID = 0x12
	LANG_KYRGYZ              LANGID = 0x40
	LANG_LAO                 LANGID = 0x54
	LANG_LATVIAN             LANGID = 0x26
	LANG_LITHUANIAN          LANGID = 0x27
	LANG_LOWER_SORBIAN       LANGID = 0x2e
	LANG_LUXEMBOURGISH       LANGID = 0x6e
	LANG_MACEDONIAN          LANGID = 0x2f
	LANG_MALAY               LANGID = 0x3e
	LANG_MALAYALAM           LANGID = 0x4c
	LANG_MALTESE             LANGID = 0x3a
	LANG_MANIPURI            LANGID = 0x58
	LANG_MAORI               LANGID = 0x81
	LANG_MAPUDUNGUN          LANGID = 0x7a
	LANG_MARATHI             LANGID = 0x4e
	LANG_MOHAWK              LANGID = 0x7c
	LANG_MONGOLIAN           LANGID = 0x50
	LANG_NEPALI              LANGID = 0x61
	LANG_NORWEGIAN           LANGID = 0x14
	LANG_OCCITAN             LANGID = 0x82
	LANG_ORIYA               LANGID = 0x48
	LANG_PASHTO              LANGID = 0x63
	LANG_PERSIAN             LANGID = 0x29
	LANG_POLISH              LANGID = 0x15
	LANG_PORTUGUESE          LANGID = 0x16
	LANG_PUNJABI             LANGID = 0x46
	LANG_QUECHUA             LANGID = 0x6b
	LANG_ROMANIAN            LANGID = 0x18
	LANG_ROMANSH             LANGID = 0x17
	LANG_RUSSIAN             LANGID = 0x19
	LANG_SAMI                LANGID = 0x3b
	LANG_SANSKRIT            LANGID = 0x4f
	LANG_SCOTTISH_GAELIC     LANGID = 0x91
	LANG_SERBIAN             LANGID = 0x1a
	LANG_SERBIAN_NEUTRAL     LANGID = 0x7c1a
	LANG_SINDHI              LANGID = 0x59
	LANG_SINHALESE           LANGID = 0x5b
	LANG_SLOVAK              LANGID = 0x1b
	LANG_SLOVENIAN           LANGID = 0x24
	LANG_SOTHO               LANGID = 0x6c
	LANG_SPANISH             LANGID = 0xa
	LANG_SWAHILI             LANGID = 0x41
	LANG_SWEDISH             LANGID = 0x1d
	LANG_SYRIAC              LANGID = 0x5a
	LANG_TAJIK               LANGID = 0x28
	LANG_TAMAZIGHT           LANGID = 0x5f
	LANG_TAMIL               LANGID = 0x49
	LANG_TATAR               LANGID = 0x44
	LANG_TELUGU              LANGID = 0x4a
	LANG_THAI                LANGID = 0x1e
	LANG_TIBETAN             LANGID = 0x51
	LANG_TIGRIGNA            LANGID = 0x73
	LANG_TSWANA              LANGID = 0x32
	LANG_TURKISH             LANGID = 0x1f
	LANG_TURKMEN             LANGID = 0x42
	LANG_UIGHUR              LANGID = 0x80
	LANG_UKRAINIAN           LANGID = 0x22
	LANG_UPPER_SORBIAN       LANGID = 0x2e
	LANG_URDU                LANGID = 0x20
	LANG_UZBEK               LANGID = 0x43
	LANG_VIETNAMESE          LANGID = 0x2a
	LANG_WELSH               LANGID = 0x52
	LANG_WOLOF               LANGID = 0x88
	LANG_XHOSA               LANGID = 0x34
	LANG_YAKUT               LANGID = 0x85
	LANG_YI                  LANGID = 0x78
	LANG_YORUBA              LANGID = 0x6a
	LANG_ZULU                LANGID = 0x35
)

const (
	SUBLANG_NEUTRAL                             LANGID = 0x0
	SUBLANG_DEFAULT                             LANGID = 0x1
	SUBLANG_SYS_DEFAULT                         LANGID = 0x2
	SUBLANG_CUSTOM_DEFAULT                      LANGID = 0x3
	SUBLANG_CUSTOM_UNSPECIFIED                  LANGID = 0x4
	SUBLANG_UI_CUSTOM_DEFAULT                   LANGID = 0x5
	SUBLANG_AFRIKAANS_SOUTH_AFRICA              LANGID = 0x1
	SUBLANG_ALBANIAN_ALBANIA                    LANGID = 0x1
	SUBLANG_ALSATIAN_FRANCE                     LANGID = 0x1
	SUBLANG_AMHARIC_ETHIOPIA                    LANGID = 0x1
	SUBLANG_ARABIC_SAUDI_ARABIA                 LANGID = 0x1
	SUBLANG_ARABIC_IRAQ                         LANGID = 0x2
	SUBLANG_ARABIC_EGYPT                        LANGID = 0x3
	SUBLANG_ARABIC_LIBYA                        LANGID = 0x4
	SUBLANG_ARABIC_ALGERIA                      LANGID = 0x5
	SUBLANG_ARABIC_MOROCCO                      LANGID = 0x6
	SUBLANG_ARABIC_TUNISIA                      LANGID = 0x7
	SUBLANG_ARABIC_OMAN                         LANGID = 0x8
	SUBLANG_ARABIC_YEMEN                        LANGID = 0x9
	SUBLANG_ARABIC_SYRIA                        LANGID = 0xa
	SUBLANG_ARABIC_JORDAN                       LANGID = 0xb
	SUBLANG_ARABIC_LEBANON                      LANGID = 0xc
	SUBLANG_ARABIC_KUWAIT                       LANGID = 0xd
	SUBLANG_ARABIC_UAE                          LANGID = 0xe
	SUBLANG_ARABIC_BAHRAIN                      LANGID = 0xf
	SUBLANG_ARABIC_QATAR                        LANGID = 0x10
	SUBLANG_ARMENIAN_ARMENIA                    LANGID = 0x1
	SUBLANG_ASSAMESE_INDIA                      LANGID = 0x1
	SUBLANG_AZERI_LATIN                         LANGID = 0x1
	SUBLANG_AZERI_CYRILLIC                      LANGID = 0x2
	SUBLANG_BASHKIR_RUSSIA                      LANGID = 0x1
	SUBLANG_BASQUE_BASQUE                       LANGID = 0x1
	SUBLANG_BELARUSIAN_BELARUS                  LANGID = 0x1
	SUBLANG_BENGALI_INDIA                       LANGID = 0x1
	SUBLANG_BENGALI_BANGLADESH                  LANGID = 0x2
	SUBLANG_BOSNIAN_BOSNIA_HERZEGOVINA_LATIN    LANGID = 0x5
	SUBLANG_BOSNIAN_BOSNIA_HERZEGOVINA_CYRILLIC LANGID = 0x8
	SUBLANG_BRETON_FRANCE                       LANGID = 0x1
	SUBLANG_BULGARIAN_BULGARIA                  LANGID = 0x1
	SUBLANG_CATALAN_CATALAN                     LANGID = 0x1
	SUBLANG_CHINESE_TRADITIONAL                 LANGID = 0x1
	SUBLANG_CHINESE_SIMPLIFIED                  LANGID = 0x2
	SUBLANG_CHINESE_HONGKONG                    LANGID = 0x3
	SUBLANG_CHINESE_SINGAPORE                   LANGID = 0x4
	SUBLANG_CHINESE_MACAU                       LANGID = 0x5
	SUBLANG_CORSICAN_FRANCE                     LANGID = 0x1
	SUBLANG_CZECH_CZECH_REPUBLIC                LANGID = 0x1
	SUBLANG_CROATIAN_CROATIA                    LANGID = 0x1
	SUBLANG_CROATIAN_BOSNIA_HERZEGOVINA_LATIN   LANGID = 0x4
	SUBLANG_DANISH_DENMARK                      LANGID = 0x1
	SUBLANG_DARI_AFGHANISTAN                    LANGID = 0x1
	SUBLANG_DIVEHI_MALDIVES                     LANGID = 0x1
	SUBLANG_DUTCH                               LANGID = 0x1
	SUBLANG_DUTCH_BELGIAN                       LANGID = 0x2
	SUBLANG_ENGLISH_US                          LANGID = 0x1
	SUBLANG_ENGLISH_UK                          LANGID = 0x2
	SUBLANG_ENGLISH_AUS                         LANGID = 0x3
	SUBLANG_ENGLISH_CAN                         LANGID = 0x4
	SUBLANG_ENGLISH_NZ                          LANGID = 0x5
	SUBLANG_ENGLISH_EIRE                        LANGID = 0x6
	SUBLANG_ENGLISH_SOUTH_AFRICA                LANGID = 0x7
	SUBLANG_ENGLISH_JAMAICA                     LANGID = 0x8
	SUBLANG_ENGLISH_CARIBBEAN                   LANGID = 0x9
	SUBLANG_ENGLISH_BELIZE                      LANGID = 0xa
	SUBLANG_ENGLISH_TRINIDAD                    LANGID = 0xb
	SUBLANG_ENGLISH_ZIMBABWE                    LANGID = 0xc
	SUBLANG_ENGLISH_PHILIPPINES                 LANGID = 0xd
	SUBLANG_ENGLISH_INDIA                       LANGID = 0x10
	SUBLANG_ENGLISH_MALAYSIA                    LANGID = 0x11
	SUBLANG_ENGLISH_SINGAPORE                   LANGID = 0x12
	SUBLANG_ESTONIAN_ESTONIA                    LANGID = 0x1
	SUBLANG_FAEROESE_FAROE_ISLANDS              LANGID = 0x1
	SUBLANG_FILIPINO_PHILIPPINES                LANGID = 0x1
	SUBLANG_FINNISH_FINLAND                     LANGID = 0x1
	SUBLANG_FRENCH                              LANGID = 0x1
	SUBLANG_FRENCH_BELGIAN                      LANGID = 0x2
	SUBLANG_FRENCH_CANADIAN                     LANGID = 0x3
	SUBLANG_FRENCH_SWISS                        LANGID = 0x4
	SUBLANG_FRENCH_LUXEMBOURG                   LANGID = 0x5
	SUBLANG_FRENCH_MONACO                       LANGID = 0x6
	SUBLANG_FRISIAN_NETHERLANDS                 LANGID = 0x1
	SUBLANG_GALICIAN_GALICIAN                   LANGID = 0x1
	SUBLANG_GEORGIAN_GEORGIA                    LANGID = 0x1
	SUBLANG_GERMAN                              LANGID = 0x1
	SUBLANG_GERMAN_SWISS                        LANGID = 0x2
	SUBLANG_GERMAN_AUSTRIAN                     LANGID = 0x3
	SUBLANG_GERMAN_LUXEMBOURG                   LANGID = 0x4
	SUBLANG_GERMAN_LIECHTENSTEIN                LANGID = 0x5
	SUBLANG_GREEK_GREECE                        LANGID = 0x1
	SUBLANG_GREENLANDIC_GREENLAND               LANGID = 0x1
	SUBLANG_GUJARATI_INDIA                      LANGID = 0x1
	SUBLANG_HAUSA_NIGERIA_LATIN                 LANGID = 0x1
	SUBLANG_HEBREW_ISRAEL                       LANGID = 0x1
	SUBLANG_HINDI_INDIA                         LANGID = 0x1
	SUBLANG_HUNGARIAN_HUNGARY                   LANGID = 0x1
	SUBLANG_ICELANDIC_ICELAND                   LANGID = 0x1
	SUBLANG_IGBO_NIGERIA                        LANGID = 0x1
	SUBLANG_INDONESIAN_INDONESIA                LANGID = 0x1
	SUBLANG_INUKTITUT_CANADA                    LANGID = 0x1
	SUBLANG_INUKTITUT_CANADA_LATIN              LANGID = 0x2
	SUBLANG_IRISH_IRELAND                       LANGID = 0x2
	SUBLANG_ITALIAN                             LANGID = 0x1
	SUBLANG_ITALIAN_SWISS                       LANGID = 0x2
	SUBLANG_JAPANESE_JAPAN                      LANGID = 0x1
	SUBLANG_KANNADA_INDIA                       LANGID = 0x1
	SUBLANG_KASHMIRI_SASIA                      LANGID = 0x2
	SUBLANG_KASHMIRI_INDIA                      LANGID = 0x2
	SUBLANG_KAZAK_KAZAKHSTAN                    LANGID = 0x1
	SUBLANG_KHMER_CAMBODIA                      LANGID = 0x1
	SUBLANG_KICHE_GUATEMALA                     LANGID = 0x1
	SUBLANG_KINYARWANDA_RWANDA                  LANGID = 0x1
	SUBLANG_KONKANI_INDIA                       LANGID = 0x1
	SUBLANG_KOREAN                              LANGID = 0x1
	SUBLANG_KYRGYZ_KYRGYZSTAN                   LANGID = 0x1
	SUBLANG_LAO_LAO                             LANGID = 0x1
	SUBLANG_LATVIAN_LATVIA                      LANGID = 0x1
	SUBLANG_LITHUANIAN                          LANGID = 0x1
	SUBLANG_LOWER_SORBIAN_GERMANY               LANGID = 0x2
	SUBLANG_LUXEMBOURGISH_LUXEMBOURG            LANGID = 0x1
	SUBLANG_MACEDONIAN_MACEDONIA                LANGID = 0x1
	SUBLANG_MALAY_MALAYSIA                      LANGID = 0x1
	SUBLANG_MALAY_BRUNEI_DARUSSALAM             LANGID = 0x2
	SUBLANG_MALAYALAM_INDIA                     LANGID = 0x1
	SUBLANG_MALTESE_MALTA                       LANGID = 0x1
	SUBLANG_MAORI_NEW_ZEALAND                   LANGID = 0x1
	SUBLANG_MAPUDUNGUN_CHILE                    LANGID = 0x1
	SUBLANG_MARATHI_INDIA                       LANGID = 0x1
	SUBLANG_MOHAWK_MOHAWK                       LANGID = 0x1
	SUBLANG_MONGOLIAN_CYRILLIC_MONGOLIA         LANGID = 0x1
	SUBLANG_MONGOLIAN_PRC                       LANGID = 0x2
	SUBLANG_NEPALI_INDIA                        LANGID = 0x2
	SUBLANG_NEPALI_NEPAL                        LANGID = 0x1
	SUBLANG_NORWEGIAN_BOKMAL                    LANGID = 0x1
	SUBLANG_NORWEGIAN_NYNORSK                   LANGID = 0x2
	SUBLANG_OCCITAN_FRANCE                      LANGID = 0x1
	SUBLANG_ORIYA_INDIA                         LANGID = 0x1
	SUBLANG_PASHTO_AFGHANISTAN                  LANGID = 0x1
	SUBLANG_PERSIAN_IRAN                        LANGID = 0x1
	SUBLANG_POLISH_POLAND                       LANGID = 0x1
	SUBLANG_PORTUGUESE                          LANGID = 0x2
	SUBLANG_PORTUGUESE_BRAZILIAN                LANGID = 0x1
	SUBLANG_PUNJABI_INDIA                       LANGID = 0x1
	SUBLANG_QUECHUA_BOLIVIA                     LANGID = 0x1
	SUBLANG_QUECHUA_ECUADOR                     LANGID = 0x2
	SUBLANG_QUECHUA_PERU                        LANGID = 0x3
	SUBLANG_ROMANIAN_ROMANIA                    LANGID = 0x1
	SUBLANG_ROMANSH_SWITZERLAND                 LANGID = 0x1
	SUBLANG_RUSSIAN_RUSSIA                      LANGID = 0x1
	SUBLANG_SAMI_NORTHERN_NORWAY                LANGID = 0x1
	SUBLANG_SAMI_NORTHERN_SWEDEN                LANGID = 0x2
	SUBLANG_SAMI_NORTHERN_FINLAND               LANGID = 0x3
	SUBLANG_SAMI_LULE_NORWAY                    LANGID = 0x4
	SUBLANG_SAMI_LULE_SWEDEN                    LANGID = 0x5
	SUBLANG_SAMI_SOUTHERN_NORWAY                LANGID = 0x6
	SUBLANG_SAMI_SOUTHERN_SWEDEN                LANGID = 0x7
	SUBLANG_SAMI_SKOLT_FINLAND                  LANGID = 0x8
	SUBLANG_SAMI_INARI_FINLAND                  LANGID = 0x9
	SUBLANG_SANSKRIT_INDIA                      LANGID = 0x1
	SUBLANG_SCOTTISH_GAELIC                     LANGID = 0x1
	SUBLANG_SERBIAN_BOSNIA_HERZEGOVINA_LATIN    LANGID = 0x6
	SUBLANG_SERBIAN_BOSNIA_HERZEGOVINA_CYRILLIC LANGID = 0x7
	SUBLANG_SERBIAN_MONTENEGRO_LATIN            LANGID = 0xb
	SUBLANG_SERBIAN_MONTENEGRO_CYRILLIC         LANGID = 0xc
	SUBLANG_SERBIAN_SERBIA_LATIN                LANGID = 0x9
	SUBLANG_SERBIAN_SERBIA_CYRILLIC             LANGID = 0xa
	SUBLANG_SERBIAN_CROATIA                     LANGID = 0x1
	SUBLANG_SERBIAN_LATIN                       LANGID = 0x2
	SUBLANG_SERBIAN_CYRILLIC                    LANGID = 0x3
	SUBLANG_SINDHI_INDIA                        LANGID = 0x1
	SUBLANG_SINDHI_PAKISTAN                     LANGID = 0x2
	SUBLANG_SINDHI_AFGHANISTAN                  LANGID = 0x2
	SUBLANG_SINHALESE_SRI_LANKA                 LANGID = 0x1
	SUBLANG_SOTHO_NORTHERN_SOUTH_AFRICA         LANGID = 0x1
	SUBLANG_SLOVAK_SLOVAKIA                     LANGID = 0x1
	SUBLANG_SLOVENIAN_SLOVENIA                  LANGID = 0x1
	SUBLANG_SPANISH                             LANGID = 0x1
	SUBLANG_SPANISH_MEXICAN                     LANGID = 0x2
	SUBLANG_SPANISH_MODERN                      LANGID = 0x3
	SUBLANG_SPANISH_GUATEMALA                   LANGID = 0x4
	SUBLANG_SPANISH_COSTA_RICA                  LANGID = 0x5
	SUBLANG_SPANISH_PANAMA                      LANGID = 0x6
	SUBLANG_SPANISH_DOMINICAN_REPUBLIC          LANGID = 0x7
	SUBLANG_SPANISH_VENEZUELA                   LANGID = 0x8
	SUBLANG_SPANISH_COLOMBIA                    LANGID = 0x9
	SUBLANG_SPANISH_PERU                        LANGID = 0xa
	SUBLANG_SPANISH_ARGENTINA                   LANGID = 0xb
	SUBLANG_SPANISH_ECUADOR                     LANGID = 0xc
	SUBLANG_SPANISH_CHILE                       LANGID = 0xd
	SUBLANG_SPANISH_URUGUAY                     LANGID = 0xe
	SUBLANG_SPANISH_PARAGUAY                    LANGID = 0xf
	SUBLANG_SPANISH_BOLIVIA                     LANGID = 0x10
	SUBLANG_SPANISH_EL_SALVADOR                 LANGID = 0x11
	SUBLANG_SPANISH_HONDURAS                    LANGID = 0x12
	SUBLANG_SPANISH_NICARAGUA                   LANGID = 0x13
	SUBLANG_SPANISH_PUERTO_RICO                 LANGID = 0x14
	SUBLANG_SPANISH_US                          LANGID = 0x15
	SUBLANG_SWAHILI_KENYA                       LANGID = 0x1
	SUBLANG_SWEDISH                             LANGID = 0x1
	SUBLANG_SWEDISH_FINLAND                     LANGID = 0x2
	SUBLANG_SYRIAC_SYRIA                        LANGID = 0x1
	SUBLANG_TAJIK_TAJIKISTAN                    LANGID = 0x1
	SUBLANG_TAMAZIGHT_ALGERIA_LATIN             LANGID = 0x2
	SUBLANG_TAMIL_INDIA                         LANGID = 0x1
	SUBLANG_TATAR_RUSSIA                        LANGID = 0x1
	SUBLANG_TELUGU_INDIA                        LANGID = 0x1
	SUBLANG_THAI_THAILAND                       LANGID = 0x1
	SUBLANG_TIBETAN_PRC                         LANGID = 0x1
	SUBLANG_TIGRIGNA_ERITREA                    LANGID = 0x2
	SUBLANG_TSWANA_SOUTH_AFRICA                 LANGID = 0x1
	SUBLANG_TURKISH_TURKEY                      LANGID = 0x1
	SUBLANG_TURKMEN_TURKMENISTAN                LANGID = 0x1
	SUBLANG_UIGHUR_PRC                          LANGID = 0x1
	SUBLANG_UKRAINIAN_UKRAINE                   LANGID = 0x1
	SUBLANG_UPPER_SORBIAN_GERMANY               LANGID = 0x1
	SUBLANG_URDU_PAKISTAN                       LANGID = 0x1
	SUBLANG_URDU_INDIA                          LANGID = 0x2
	SUBLANG_UZBEK_LATIN                         LANGID = 0x1
	SUBLANG_UZBEK_CYRILLIC                      LANGID = 0x2
	SUBLANG_VIETNAMESE_VIETNAM                  LANGID = 0x1
	SUBLANG_WELSH_UNITED_KINGDOM                LANGID = 0x1
	SUBLANG_WOLOF_SENEGAL                       LANGID = 0x1
	SUBLANG_XHOSA_SOUTH_AFRICA                  LANGID = 0x1
	SUBLANG_YAKUT_RUSSIA                        LANGID = 0x1
	SUBLANG_YI_PRC                              LANGID = 0x1
	SUBLANG_YORUBA_NIGERIA                      LANGID = 0x1
	SUBLANG_ZULU_SOUTH_AFRICA                   LANGID = 0x1
)

// MakeLongID build LANGID from primary and sub.
func MakeLongID(p, s LANGID) LANGID {
	return LANGID(p | (s << 10))
}

// PrimaryLangID get primary part of LANGID
func PrimaryLangID(id LANGID) LANGID {
	return id & 0x3ff
}

// SubLangID get sub part of LANGID
func SubLangID(id LANGID) LANGID {
	return id >> 10
}
