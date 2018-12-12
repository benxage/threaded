package database

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/go-pg/pg"
	"github.com/spf13/viper"
)

// Postgres is a wrapper around pg.DB
type Postgres struct {
	DB  *pg.DB
	Opt *pg.Options
}

// NewPostgres returns a postgres instance
func NewPostgres(filename string) *Postgres {
	dbo := readPostgresConfig(filename)
	return &Postgres{DB: pg.Connect(dbo), Opt: dbo}
}

// Close closes the database
func (db *Postgres) Close() {
	db.DB.Close()
}

// Insert TODO
func (db *Postgres) Insert() error {
	return nil
}

// PrintInfo prints the config info used
func (db *Postgres) PrintInfo() {
	fmt.Println("---Database Information---")
	log.Printf("Database: %s\n", db.Opt.Database)
	log.Printf("Network: %s\n", db.Opt.Network)
	log.Printf("Addr: %s\n", db.Opt.Addr)
	log.Printf("ApplicationName: %s\n", db.Opt.ApplicationName)
	log.Printf("User: %s\n", db.Opt.User)
	log.Printf("Password: %s\n", db.Opt.Password)
	log.Printf("DialTimeout: %s\n", db.Opt.DialTimeout.String())
	log.Printf("PoolSize: %d\n", db.Opt.PoolSize)
	log.Printf("PoolTimeout: %s\n", db.Opt.PoolTimeout.String())
	log.Printf("ReadTimeout: %s\n", db.Opt.ReadTimeout.String())
	log.Printf("WriteTimeout: %s\n", db.Opt.WriteTimeout.String())
	log.Printf("IdleTimeout: %s\n", db.Opt.IdleTimeout.String())
	log.Printf("IdleCheckFrequency: %s\n", db.Opt.IdleCheckFrequency.String())
	log.Printf("MinIdleConns: %d\n", db.Opt.MinIdleConns)
	log.Printf("MaxConnAge: %s\n", db.Opt.MaxConnAge.String())
	log.Printf("RetryStatementTimeout: %t\n", db.Opt.RetryStatementTimeout)
	log.Printf("MaxRetries: %d\n", db.Opt.MaxRetries)
	log.Printf("MaxRetryBackoff: %s\n", db.Opt.MaxRetryBackoff.String())
	log.Printf("MinRetryBackoff: %s\n", db.Opt.MinRetryBackoff.String())
}

func readPostgresConfig(filename string) *pg.Options {
	viper.AddConfigPath(".")
	viper.SetConfigName(filename)
	viper.SetDefault("database", "threaded")
	viper.ReadInConfig() // ignoring read error
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Database config changed at ", e.Name)
	})

	var dbo pg.Options
	viper.Unmarshal(&dbo)
	return &dbo
}
