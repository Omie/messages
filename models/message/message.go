package message

import (
	"errors"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/satori/go.uuid"

	"github.com/omie/messages/lib/db"
)

// type that represents Message entity
type Message struct {
	Id        int       `json:"id" db:"id"`
	Uuid      string    `json:"uuid" db:"uuid"`
	Text      string    `json:"text" db:"text"`
	CreatedOn time.Time `json:"created_on" db:"created_on"`
}

// Predefined error instance so that controller can easily figure out
// which type of error GetByUUID() returned
var ErrMessageNotFound = errors.New("Message Not Found")
var ErrValidationError = errors.New("Message Text cannot be empty")

// prepare and insert the message into the database
// returns error or nil
func Create(text string) (*Message, error) {
	if strings.TrimSpace(text) == "" {
		return nil, ErrValidationError
	}

	db, err := db.GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	newUUID := uuid.NewV1().String()
	insertQuery := `INSERT INTO messages(uuid, text) VALUES(?, ?)`
	_, err = db.Exec(insertQuery, newUUID, text)
	if err != nil {
		log.Error("Message.Create: ", err)
		return nil, err
	}

	return GetByUUID(newUUID)
}

// Retrive a message using given uuid
// param: uuid: string: uuid of the message to retrive
// returns Message object on success
// returns error in case of error
func GetByUUID(uuid string) (*Message, error) {
	db, err := db.GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	row, err := db.Queryx("SELECT id, uuid, text, created_on FROM messages WHERE uuid = ?", uuid)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	if row.Next() {
		var message Message
		err = row.StructScan(&message)
		if err != nil {
			return nil, err
		}
		return &message, nil
	}

	return nil, ErrMessageNotFound
}
