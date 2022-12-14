package config

import (
	"html/template"
	"log"

	"github.com/JoshiNikita17/udemyproject1/internal/models"
	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProduction  bool
	// Session       *scs.SessionManager
	Session  *scs.SessionManager
	MailChan chan models.MailData
}
