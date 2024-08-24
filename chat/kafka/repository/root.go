package repository

import (
	"chat-kafka/config"
	"chat-kafka/repository/kafka"
	"chat-kafka/types/schema"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
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

func (r *Repository) ServerSet(ip string, available bool) error {
	_, err := r.db.Exec("INSERT INTO chatting.serverInfo(`ip`, `available`) VALUES (?, ?) ON DUPLICATE KEY UPDATE `available` = VALUES(`available`)", ip, available)
	return err
}

func (r *Repository) RoomList() ([]*schema.Room, error) {
	qs := query([]string{"SELECT * FROM", room})

	if cursor, err := r.db.Query(qs); err != nil {
		return nil, err
	} else {
		defer cursor.Close()

		var result []*schema.Room

		for cursor.Next() {
			d := new(schema.Room)
			if err = cursor.Scan(&d.ID, &d.Name, &d.CreatedAt, &d.UpdatedAt); err != nil {
				return nil, err
			} else {
				result = append(result, d)
			}
		}

		if len(result) == 0 {
			return []*schema.Room{}, nil
		} else {
			return result, nil
		}
	}
}

func (r *Repository) MakeRoom(name string) error {
	_, err := r.db.Exec("INSERT INTO chatting.room(name) VALUES(?)", name)
	return err
}

func (r *Repository) Room(name string) (*schema.Room, error) {
	d := new(schema.Room)
	qs := query([]string{"SELECT * FROM", room, "WHERE name = ?"})

	err := r.db.QueryRow(qs, name).Scan(
		&d.ID,
		&d.Name,
		&d.CreatedAt,
		&d.UpdatedAt,
	)

	if err = noResult(err); err != nil {
		return nil, err
	} else {
		return nil, nil
	}
}

func (r *Repository) GetChatList(roomName string) ([]*schema.Chat, error) {
	qs := query([]string{"SELECT * FROM", chat, "WHERE room = ? ORDER BY `when` DESC LIMIT 10"})

	if cursor, err := r.db.Query(qs, roomName); err != nil {
		return nil, err
	} else {
		defer cursor.Close()

		var result []*schema.Chat

		for cursor.Next() {
			d := new(schema.Chat)
			if err = cursor.Scan(&d.ID, &d.Room, &d.Name, &d.Message, &d.When); err != nil {
				return nil, err
			} else {
				result = append(result, d)
			}
		}

		if len(result) == 0 {
			return []*schema.Chat{}, nil
		} else {
			return result, nil
		}
	}
}

func (r *Repository) InsertChatting(user string, message string, roomName string) error {
	log.Println("Insert Chatting Using WSS", "from", user, "message", message, "room", roomName)
	_, err := r.db.Exec("INSERT INTO chatting.chat(room, name, message) VALUES(?,?,?)", roomName)

	return err
}

func query(qs []string) string {
	return strings.Join(qs, " ") + ";"
}

func noResult(err error) error {
	if strings.Contains(err.Error(), "sql: no rows in result set") {
		return nil
	} else {
		return err
	}
}
