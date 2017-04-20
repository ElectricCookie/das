package i18n

func registerDe() {

	RegisterLocale("de", TranslationConfig{
		"register:subject":   "Willkommen {{ .Username }}, bitte bestätigen Sie Ihre E-Mail Adresse!",
		"register:emailBody": loadTemplate("register-de"),
	})

}
