package migrations

import (
	"context"
	"fmt"
	"testing"

	"github.com/G-Research/unicorn-history-server/internal/log"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/G-Research/unicorn-history-server/test/config"
	"github.com/G-Research/unicorn-history-server/test/database"
)

type GoMigrateIntTest struct {
	suite.Suite
	tp   *database.TestPostgresContainer
	pool *pgxpool.Pool
	cfg  *database.InstanceConfig
}

func (ts *GoMigrateIntTest) SetupSuite() {
	ctx := context.Background()
	cfg := database.InstanceConfig{
		User:     "test",
		Password: "test",
		DBName:   "template",
		Host:     "localhost",
		Port:     15439,
	}

	tp, err := database.NewTestPostgresContainer(ctx, cfg)
	require.NoError(ts.T(), err)
	ts.tp = tp
	ts.pool = tp.Pool(ctx, ts.T(), &cfg)
	ts.cfg = &cfg
}

func (ts *GoMigrateIntTest) TearDownSuite() {
	err := ts.tp.Container.Terminate(context.Background())
	require.NoError(ts.T(), err)
}

func (ts *GoMigrateIntTest) TestGoMigrate() {
	ctx := context.Background()
	log.Init(config.GetTestLogConfig())

	cfg := config.GetTestPostgresConfig()
	cfg.Host = ts.cfg.Host
	cfg.Port = ts.cfg.Port
	cfg.Username = ts.cfg.User
	cfg.Password = ts.cfg.Password
	cfg.DbName = ts.cfg.DBName
	cfg.Schema = "test"

	_, err := ts.pool.Exec(ctx, fmt.Sprintf("CREATE SCHEMA %s", cfg.Schema))
	require.NoError(ts.T(), err)

	setSearchPathQuery := fmt.Sprintf("SET search_path TO %s", cfg.Schema)
	_, err = ts.pool.Exec(ctx, setSearchPathQuery)
	require.NoError(ts.T(), err)

	m, err := New(cfg, "../../../migrations")
	require.NoError(ts.T(), err)

	applied, err := m.Up()
	assert.Truef(ts.T(), applied, "expected up migrations to be applied for the first run")
	require.NoError(ts.T(), err)

	applied, err = m.Up()
	assert.Falsef(ts.T(), applied, "expected no up migrations to be applied for the second run")
	require.NoError(ts.T(), err)

	applied, err = m.Down()
	assert.Truef(ts.T(), applied, "expected down migrations to be applied for the first run")
	require.NoError(ts.T(), err)

	applied, err = m.Down()
	assert.Falsef(ts.T(), applied, "expected no down migrations to be applied for the second run")
	require.NoError(ts.T(), err)
}

func TestGoMigrateIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}
	suite.Run(t, new(GoMigrateIntTest))
}
