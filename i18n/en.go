package i18n

func registerEn() {

	RegisterLocale("en", TranslationConfig{
		"register:subject":   "Welcome {{ .Username }}, please verify your email address!",
		"register:emailBody": loadTemplate("register-en"),
	})

}
