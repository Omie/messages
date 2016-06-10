package message

import (
	"errors"
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

// validates if the object is okay to insert in the db
// returns error if the object is invalid
// returns nil if the object is valid
func (self *Message) Validate() error {
	if self.Text == "" {
		return ErrValidationError
	}
	return nil
}

// prepare and insert the message into the database
// returns error or nil
func Create(text string) (*Message, error) {
	db, err := db.GetDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	newMessage := &Message{
		Text: text,
		Uuid: uuid.NewV1().String(),
	}
	if err := newMessage.Validate(); err != nil {
		return nil, err
	}

	insertQuery := `INSERT INTO messages(uuid, text) VALUES(?, ?)`
	_, err = db.Exec(insertQuery, newMessage.Uuid, newMessage.Text)
	if err != nil {
		log.Error("Message.Create: ", err)
		return nil, err
	}

	return GetByUUID(newMessage.Uuid)
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
