package test

import (
	"github.com/jmoiron/sqlx"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"log"
	"micro_auth/internal/application"
	"micro_auth/internal/database"
	iLogger "micro_auth/internal/log"
	"net"
	"net/url"
	"os"
	"runtime"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/stretchr/testify/suite"
)

type BaseSuite struct {
	suite.Suite
	app *application.App
	log *iLogger.Logger
}

func (s *BaseSuite) SetupSuite() {
	s.log = iLogger.GetLogger()

	s.app = application.GetApp()
}

func (s *BaseSuite) App() *application.App {
	return s.app
}

// PostgresSuite struct for PostgreSQL Suite
type PostgresSuite struct {
	BaseSuite
	pgURL  *url.URL
	DBConn *sqlx.DB
}

var (
	pool      *dockertest.Pool
	resource  *dockertest.Resource
	logWaiter docker.CloseWaiter
)

func (s *PostgresSuite) SetupSuite() {
	s.BaseSuite.SetupSuite()

	var err error
	s.pgURL, err = url.Parse(os.Getenv("DATABASE_URL"))
	if err != nil {
		s.log.WithError(err).Fatal("Could not parse DATABASE_URL")
	}
	q := s.pgURL.Query()
	q.Add("sslmode", "disable")
	s.pgURL.RawQuery = q.Encode()

	pool, err = dockertest.NewPool("")
	if err != nil {
		s.log.WithError(err).Fatal("Could not connect to docker")
	}

	pw, _ := s.pgURL.User.Password()
	runOpts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "11.1-alpine",
		Env: []string{
			"POSTGRES_USER=" + s.pgURL.User.Username(),
			"POSTGRES_PASSWORD=" + pw,
			"POSTGRES_DB=" + s.pgURL.Path,
		},
	}

	resource, err = pool.RunWithOptions(&runOpts)
	if err != nil {
		s.log.WithError(err).Fatal("Could not start postgres container")
	}

	s.pgURL.Host = resource.Container.NetworkSettings.IPAddress

	// docker layer network is different on MacOS
	if runtime.GOOS == "darwin" {
		s.pgURL.Host = net.JoinHostPort(resource.GetBoundIP("5432/tcp"), resource.GetPort("5432/tcp"))
	}

	logWaiter, err = pool.Client.AttachToContainerNonBlocking(docker.AttachToContainerOptions{
		Container:    resource.Container.ID,
		OutputStream: log.Writer(),
		ErrorStream:  log.Writer(),
		Stderr:       false,
		Stdout:       false,
		Stream:       false,
	})
	if err != nil {
		s.log.WithError(err).Fatal("Could not connect to postgres container log output")
	}

	pool.MaxWait = 10 * time.Second
	err = pool.Retry(func() error {
		db, err := sqlx.Open("pgx", s.pgURL.String())
		if err != nil {
			return err
		}
		return db.Ping()
	})
	if err != nil {
		s.log.WithError(err).Fatal("Could not connect to postgres server")
	}
	runMigration(s.pgURL)
}

func (s *PostgresSuite) TearDownSuite() {
	err := resource.Expire(60)
	if err != nil {
		s.log.WithError(err).Error("Could not purge resource after 60 seconds")
	}
	err = pool.Purge(resource)
	if err != nil {
		s.log.WithError(err).Error("Could not purge resource")
	}

	err = logWaiter.Close()
	if err != nil {
		s.log.WithError(err).Error("Could not close container log")
	}
	err = logWaiter.Wait()
	if err != nil {
		s.log.WithError(err).Error("Could not wait for container log to close")
	}
}

func (s *PostgresSuite) SetupTest() {
	//var err error
	//database.DB, err = sqlx.Connect("pgx", s.pgURL.String())
	//if err != nil {
	//	s.log.WithError(err).Fatal("Could not connect to postgres")
	//}
	database.Open(s.pgURL.String())
	s.DBConn = database.DB
}

func (s *PostgresSuite) TearDownTest() {
	err := s.DBConn.Close()
	if err != nil {
		s.log.WithError(err).Fatal("Could not close postgres connection")
	}
}

type ApiClientSuite struct {
	PostgresSuite
	Client API
}

func (s *ApiClientSuite) SetupSuite() {
	s.PostgresSuite.SetupSuite()
	s.Client = API{t: s.T(), app: s.App()}
}
