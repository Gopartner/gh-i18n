package i18n

import (
	"encoding/json"
	"embed"
	"fmt"
	"os"
	"path"
	"sync"
)

//go:embed locales/*.json
var localeFS embed.FS

var (
	mu      sync.RWMutex
	active  string
	locales map[string]map[string]string
)

func Init() error {
	mu.Lock()
	defer mu.Unlock()

	locales = make(map[string]map[string]string)
	active = "en"

	entries, err := localeFS.ReadDir("locales")
	if err != nil {
		return fmt.Errorf("failed to read locales directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if path.Ext(entry.Name()) != ".json" {
			continue
		}

		data, err := localeFS.ReadFile(path.Join("locales", entry.Name()))
		if err != nil {
			continue
		}

		var locale struct {
			Language string            `json:"language"`
			Messages map[string]string `json:"messages"`
		}
		if err := json.Unmarshal(data, &locale); err != nil {
			continue
		}

		locales[locale.Language] = locale.Messages
	}

	if _, ok := locales["en"]; !ok {
		locales["en"] = make(map[string]string)
	}

	if lang := os.Getenv("GH_LANG"); lang != "" {
		if _, ok := locales[lang]; ok {
			active = lang
		}
	}

	return nil
}

func T(key string) string {
	mu.RLock()
	defer mu.RUnlock()

	if msgs, ok := locales[active]; ok {
		if msg, ok := msgs[key]; ok {
			return msg
		}
	}

	if msgs, ok := locales["en"]; ok {
		if msg, ok := msgs[key]; ok {
			return msg
		}
	}

	return key
}

func SetLocale(lang string) {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := locales[lang]; ok {
		active = lang
	}
}

func ActiveLocale() string {
	mu.RLock()
	defer mu.RUnlock()
	return active
}
