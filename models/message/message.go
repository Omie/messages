package message

import (
	"errors"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/satori/go.uuid"

	"github.com/omie/messages/lib/db"
)

// type that represents Message entity
// Note: using tags we can control the encode-decode behavior of the object
// Since in this particular case we would like to only present uuid-as-id,
// I have tagged it as "id"
// tagging with "-" makes it ignore the field during json encode
type Message struct {
	Id        int       `json:"-" db:"id"`
	Uuid      string    `json:"id" db:"uuid"`
	Text      string    `json:"-" db:"text"`
	CreatedOn time.Time `json:"-" db:"created_on"`
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
func (self *Message) Create() error {
	db, err := db.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()

	self.Uuid = uuid.NewV1().String()
	if err := self.Validate(); err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO messages(uuid, text) VALUES(?, ?)")
	if err != nil {
		log.Error("Message.Create: ", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(self.Uuid, self.Text)
	if err != nil {
		log.Error("Message.Create: ", err)
		return err
	}
	return nil
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

	rows, err := db.Query("SELECT id, uuid, text, created_on FROM messages WHERE uuid = ?", uuid)
	if err != nil {
		log.Error("GetByUUID: ", err)
		return nil, err
	}
	defer rows.Close()

	var message Message
	if rows.Next() {
		message = Message{}
		rows.Scan(&message.Id, &message.Uuid, &message.Text, &message.CreatedOn)
	} else {
		return nil, ErrMessageNotFound
	}
	return &message, nil
}
