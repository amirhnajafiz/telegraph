package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ClientCollection = "clients"

type Client struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name,omitempty"`
	Pass  string             `bson:"pass,omitempty"`
	Chats []string           `bson:"chats,omitempty"`
}

func (s *Store) NewClient(ctx context.Context, client *Client) error {
	col := s.Database.Collection(ClientCollection)

	_, err := col.InsertOne(ctx, client)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetClient(ctx context.Context, name string) (Client, error) {
	filter := bson.D{{"name", name}}
	col := s.Database.Collection(ClientCollection)

	cursor, _ := col.Find(ctx, filter)
	defer cursor.Close(ctx)

	var client Client

	if cursor.Next(ctx) {
		if err := cursor.Decode(&client); err != nil {
			return client, err
		}
	}

	return client, nil
}

func (s *Store) AddChatToClient(ctx context.Context, name string, chat string) error {
	client, err := s.GetClient(ctx, name)
	if err != nil {
		return err
	}

	client.Chats = append(client.Chats, chat)

	filter := bson.D{{"name", name}}
	collection := s.Database.Collection(ClientCollection)

	if _, err := collection.UpdateOne(ctx, filter, client); err != nil {
		return err
	}

	return nil
}

func (s *Store) RemoveChatForClient(ctx context.Context, name string, chat string) error {
	client, err := s.GetClient(ctx, name)
	if err != nil {
		return err
	}

	for index, value := range client.Chats {
		if value == chat {
			client.Chats = append(client.Chats[:index], client.Chats[index+1:]...)

			break
		}
	}

	filter := bson.D{{"name", name}}
	collection := s.Database.Collection(ClientCollection)

	if _, err := collection.UpdateOne(ctx, filter, client); err != nil {
		return err
	}

	return nil
}

func (s *Store) RemoveClient(ctx context.Context, name string) error {
	filter := bson.D{{"name", name}}
	collection := s.Database.Collection(ClientCollection)

	if _, err := collection.DeleteOne(ctx, filter); err != nil {
		return err
	}

	return nil
}
