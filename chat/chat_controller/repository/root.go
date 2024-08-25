package repository

import (
	"chat_controller/config"
	"chat_controller/repository/kafka"
	"chat_controller/types/table"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type Repository struct {
	cfg   *config.Config
	db    *sql.DB
	Kafka *kafka.Kafka
}

const (
	room       = "chatting.room"
	chat       = "chatting.chat"
	serverInfo = "chatting.serverInfo"
)

func NewRepository(cfg *config.Config) (*Repository, error) {
	r := &Repository{cfg: cfg}
	var err error
	if r.db, err = sql.Open(cfg.DB.Database, cfg.DB.URL); err != nil {
		return nil, err
	} else if r.Kafka, err = kafka.NewKafka(cfg); err != nil {
		return nil, err
	} else {
		return r, nil
	}
}

func (r *Repository) GetAvailableServerList() ([]*table.ServerInfo, error) {
	qs := query([]string{"SELECT * FROM", serverInfo, "WHERE available = 1"})

	if cursor, err := r.db.Query(qs); err != nil {
		return nil, err
	} else {
		defer cursor.Close()

		var result []*table.ServerInfo

		for cursor.Next() {
			d := new(table.ServerInfo)

			if err = cursor.Scan(
				&d.IP,
				&d.Available,
			); err != nil {
				return nil, err
			} else {
				result = append(result, d)
			}
		}

		if len(result) == 0 {
			return []*table.ServerInfo{}, nil
		} else {
			return result, nil
		}
	}
}

func query(qs []string) string {
	return strings.Join(qs, " ") + ";"
}
