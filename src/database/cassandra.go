package database

import (
	"context"
	"log"
	"time"

	"github.com/gocql/gocql"
	"github.com/j-04/cardhub/config"
	"github.com/j-04/cardhub/types"
)

const (
	GET_DECKS      string = "select id, name from cardhub.deck;"
	GET_DECK_BY_ID string = "select id, name from cardhub.deck where id = ?;"
)

type CassandraDatabase struct {
	session *gocql.Session
	_       struct{}
}

func (db *CassandraDatabase) Close() {
	db.session.Close()
}

func NewCassandraDatabse(config *config.Config) *CassandraDatabase {
	cluster := gocql.NewCluster(config.Database.Cassandra.Host)
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = time.Second * time.Duration(config.Cassandra.ConnectionTimeout)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username:              config.Database.Cassandra.Username,
		Password:              config.Database.Cassandra.Password,
		AllowedAuthenticators: []string{"com.instaclustr.cassandra.auth.InstaclustrPasswordAuthenticator"},
	}

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return &CassandraDatabase{
		session: session,
	}
}

func (db *CassandraDatabase) GetDecks(context context.Context) ([]*types.Deck, error) {
	var decks []*types.Deck = make([]*types.Deck, 0)
	scanner := db.session.Query(GET_DECKS).WithContext(context).Iter().Scanner()

	var (
		id   string
		name string
	)

	for scanner.Next() {
		scanner.Scan(&id, &name)
		decks = append(decks, &types.Deck{Id: id, Name: name})
	}

	if err := scanner.Err(); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return decks, nil
}

func (db *CassandraDatabase) GetDeck(context context.Context, deckId string) (*types.Deck, error) {
	var (
		id   string
		name string
	)

	err := db.session.Query(GET_DECK_BY_ID, deckId).Scan(&id, &name)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	deck := &types.Deck{
		Id:   id,
		Name: name,
	}

	return deck, nil
}

func (db *CassandraDatabase) PutWordsInDeck(context context.Context, words []types.Word, deckId string) error {
	return nil
}

func (db *CassandraDatabase) SaveDeck(context context.Context, deck types.Deck) error {
	return nil
}

func (db *CassandraDatabase) DeleteDeck(context context.Context, deckId string) error {
	return nil
}

func (db *CassandraDatabase) DeleteWordInDeck(context context.Context, deckId string, wordId string) error {
	return nil
}

func (db *CassandraDatabase) GetWords(context context.Context, pageSize int, pageNumber int) ([]types.Word, error) {
	return nil, nil
}

func (db *CassandraDatabase) SaveWords(context context.Context, words []types.Word) error {
	return nil
}

func (db *CassandraDatabase) UpdateWord(context context.Context, wordId string, newWord types.Word) error {
	return nil
}
