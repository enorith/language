package language

import (
	"fmt"
	"sync"

	"github.com/enorith/supports/str"
)

// DefaultLanguage default language
var DefaultLanguage = "en"

var languageData *languages

type languageValues struct {
	keyValues map[string]string
	mu        sync.RWMutex
}

func (lv languageValues) get(id string) (string, bool) {
	lv.mu.RLock()
	defer lv.mu.RUnlock()
	v, b := lv.keyValues[id]

	return v, b
}

type language struct {
	languageData map[string]languageValues
	mu           sync.RWMutex
}

func (l language) get(lang string) (languageValues, bool) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	v, b := l.languageData[lang]

	return v, b
}

func (l *language) set(lang string, values languageValues) {
	l.mu.Lock()
	l.languageData[lang] = values
	l.mu.Unlock()
}

type languages struct {
	languages map[string]*language // key : languages
	mu        sync.RWMutex
}

func (ls *languages) get(key string) (*language, bool) {
	ls.mu.RLock()
	defer ls.mu.RUnlock()
	v, b := ls.languages[key]

	return v, b
}

func (ls *languages) set(key string, lang *language) {
	ls.mu.Lock()
	if ls.languages == nil {
		ls.languages = map[string]*language{}
	}
	ls.languages[key] = lang
	ls.mu.Unlock()
}

// Translate with giving language
func Translate(key, id, lang string, params ...map[string]string) (string, error) {

	l, ok := languageData.get(key)

	if !ok {
		return "", fmt.Errorf("not found for language key [%s]", key)
	}

	lv, ok := l.get(lang)

	if !ok {
		return "", fmt.Errorf("not found language [%s] of [%s]", lang, key)
	}

	if s, ok := lv.get(id); ok {
		var vars map[string]string

		if len(params) > 0 {
			vars = params[0]
		}

		return str.ReplaceVar(s, vars), nil
	}

	return "", fmt.Errorf("not found language [%s] of [%s] id: %s", lang, key, id)
}

// T translate with default language
func T(key, id string, params ...map[string]string) (string, error) {
	return Translate(key, id, DefaultLanguage, params...)
}

// Register language data
func Register(key, lang string, data map[string]string) {

	v := languageValues{data, sync.RWMutex{}}
	if le, ok := languageData.get(key); ok {
		le.set(lang, v)
		languageData.set(key, le)
	} else {
		l := &language{map[string]languageValues{}, sync.RWMutex{}}
		l.set(lang, v)
		languageData.set(key, l)
	}
}

func init() {
	languageData = &languages{}
}
