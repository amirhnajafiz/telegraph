package message

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var Collection = "messages"

type Message struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	From string             `bson:"from,omitempty"`
	To   string             `bson:"to,omitempty"`
	Msg  string             `bson:"msg,omitempty"`
	Date time.Time          `bson:"date,omitempty"`
}
