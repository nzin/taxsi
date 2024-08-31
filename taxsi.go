package taxsi

import (
	"context"
	"net/http"
	"os"
	//"github.com/rs/zerolog/log"
)

type Rule struct {
	Rule   string
	Action string
}

// Config the plugin configuration.
type Config struct {
	RulesFiles []string `json:"RulesFiles"`
	Rules      []Rule   `json:"Rules"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		RulesFiles: make([]string, 0),
		Rules:      make([]Rule, 0),
	}
}

// Demo a Demo plugin.
type Taxsi struct {
	next http.Handler
	// headers  map[string]string
	name string
}

// New created a new Demo plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {

	return &Taxsi{
		next: next,
		name: name,
	}, nil
}

func (t *Taxsi) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	//log.With().Str("taxsi", t.name)
	os.Stdout.WriteString("taxsi log\n")

	t.next.ServeHTTP(rw, req)
}
