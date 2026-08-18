package models

import "database/sql"

type Connection struct {
	Config ConnectionConfig
	DB     *sql.DB
	DBs		[]string
}

type ConnectionConfig struct {
	ID        string   `json:"id" example:""`
	Name      string   `json:"name" example:""`
	Driver    string   `json:"driver" example:"mysql"`
	Host      string   `json:"host" example:"127.0.0.1"`
	Port      int      `json:"port" example:"3306"`
	Username  string   `json:"username" example:"admin"`
	Password  string   `json:"password" example:"senha123"`
	Database  string   `json:"database" example:""`
	Databases []string `json:"databases"`
}

type ConnectionResponse struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Driver   string    `json:"driver"`
	Host     string 	 `json:"host"`
	Port     int    	 `json:"port"`
	Database string    `json:"database"`
	Databases []string `json:"databases"`
}

func (c *Connection) Matches(cfg ConnectionConfig) bool {
	return c.Config.Driver == cfg.Driver &&
		c.Config.Host == cfg.Host &&
		c.Config.Port == cfg.Port &&
		c.Config.Username == cfg.Username &&
		c.Config.Database == cfg.Database
}