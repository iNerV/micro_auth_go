package test

import (
	"github.com/sirupsen/logrus"
	"net/url"
	"path/filepath"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func runMigration(pgURL *url.URL) {
	log := logrus.New()
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Join(filepath.Dir(b), "./..")
	m, err := migrate.New(
		"file://"+basePath+"/database/migrations",
		pgURL.String(),
	)
	if err != nil {
		log.WithError(err).Fatalln("Could not create migration")
	}
	if err := m.Up(); err != nil {
		log.WithError(err).Fatalln("Could not migrate")
	}
}
