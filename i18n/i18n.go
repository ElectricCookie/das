package i18n

import (
	"html/template"

	"bytes"

	"fmt"

	"github.com/ElectricCookie/das-cms/assets"
)

var requiredLocales = []string{
	"register:subject",
	"register:emailBody",
}

// Translation translates all the things
type Translation map[string]*template.Template

// TranslationConfig is used to initalize a translation
type TranslationConfig map[string]string

var locales map[string]Translation

// LoadTranslations loads the translations
func LoadTranslations() {

	locales = map[string]Translation{}

	registerDe()
	registerEn()
}

// RegisterLocale registers a possible locale in the server
func RegisterLocale(lang string, register TranslationConfig) {

	for _, item := range requiredLocales {
		if _, ok := register[item]; !ok {
			panic("Missing translation for " + item)
		}
	}

	translation := Translation{}

	for key, item := range register {

		t := template.New(lang + ":" + key)

		t, err := t.Parse(item)

		if err != nil {
			panic(err)
		}

		translation[key] = t

	}

	locales[lang] = translation

}

// GetTemplate is used to fetch a translation
func GetTemplate(lang string, name string) *template.Template {
	if locales[lang] == nil {
		panic("Invalid language requested: " + lang)
	}
	return locales[lang][name]
}

func loadTemplate(templateName string) string {
	content, err := assets.Asset("assets/files/" + templateName + ".html")
	if err != nil {
		panic(err)
	}
	return string(content)
}

// GetTranslator returns a translator bound to a language
func GetTranslator(language string) func(name string, params interface{}) string {
	return func(name string, params interface{}) string {
		var b bytes.Buffer
		t := GetTemplate(language, name)
		err := t.Execute(&b, params)
		if err != nil {
			fmt.Println(err)
		}

		return b.String()
	}
}
