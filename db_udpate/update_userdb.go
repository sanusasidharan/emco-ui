package main

import (
	context2 "context"
	"encoding/json"
	"fmt"
	pkgerrors "github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
	"os"
	"sort"
	"strings"
)

type DefaultUserEntry struct {
	Provider    string `json:"provider"`
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	FirstName   string `json:"firstName"`
	Tenant      string `json:"tenant"`
	Role        string `json:"role"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

const (
	DBName = "rbac_userdb"
)

func CreateUser(dbName string, ip string) error {
	clientOptions := options.Client()
	clientOptions.ApplyURI(ip)
	clientOptions.SetAuth(options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		AuthSource:    dbName,
		Username:      os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
		Password:      os.Getenv("MONGO_INITDB_ROOT_PASSWORD")})
	mongoClient, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Printf("Failed to get mongo client err %s", err.Error())
		return err
	}
	err = mongoClient.Connect(context.Background())
	if err != nil {
		log.Printf("Failed to connect to mongo client err %s", err.Error())
		return err
	}
	defer func(mongoClient *mongo.Client, ctx context2.Context) {
		err := mongoClient.Disconnect(ctx)
		if err != nil {
			log.Printf("Failed to connect to mongo client err %s", err.Error())
			return
		}
	}(mongoClient, context.Background())

	dbLocal := mongoClient.Database("rbac_userdb")
	r := dbLocal.RunCommand(context.Background(), bson.D{{"createUser", os.Getenv("DB_EMCOUI_USERNAME")},
		{"pwd", os.Getenv("DB_EMCOUI_PASSWORD")},
		{"roles", []bson.M{{"role": "readWrite", "db": "rbac_userdb"},
			{"role": "readWrite", "db": "middleend"}}},
	})
	// UPdate the user entry: If we helm delete and install the app
	// the secret will get refreshed, hence the user needs to be updated
	// with new secret.
	if (r.Err() != nil && strings.Contains(r.Err().Error(), "already exists")) {
		r := dbLocal.RunCommand(context.Background(), bson.D{{"updateUser", os.Getenv("DB_EMCOUI_USERNAME")},
			{"pwd", os.Getenv("DB_EMCOUI_PASSWORD")},
			{"roles", []bson.M{{"role": "readWrite", "db": "rbac_userdb"},
				{"role": "readWrite", "db": "middleend"}}},
		})
		if r.Err() != nil {
			log.Printf("Failed to Update User EMCOUI %s", err.Error())
			return err

		}
		return nil
	}
	if r.Err() != nil {
		log.Printf("Failed to create User EMCOUI %s", err.Error())
		return err
	}
	return nil
}

func main() {
	URI := os.ExpandEnv("mongodb://${MONGODB_HOST}")
	// Create DB users for middleend and authservice databases,
	// if the env variable dbauthEnable is set to true
	if len(os.Getenv("MONGO_INITDB_ROOT_USERNAME")) > 0 && len(os.Getenv("MONGO_INITDB_ROOT_PASSWORD")) > 0 {
		err := CreateUser("admin", URI)
		if err != nil {
			fmt.Println("Failed to connect to mongo", err)
			return
		}
		log.Printf("Created user for emcoui")
	}
	log.Printf("Connect to DB client without auth")

	err := CreateDBClient("mongo", DBName, URI)
	if err != nil {
		fmt.Println("Failed to connect to mongo", err)
		return
	}

	userData := DefaultUserEntry{
		Provider:    "amcop",
		ID:          "admin@enterprise.com",
		DisplayName: "Admin",
		FirstName:   "Admin",
		Tenant:      "default",
		Role:        "admin",
		Email:       "admin@enterprise.com",
		Password:    "$2a$10$SaNe/etlslsSqW3adzO9Zuzic8okeEZxaS6/6oACjQFzO9CU8IRfW",
	}

	err = DBconn.Insert("users", userData, nil, "userData", userData)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Printf("Created default Admin user in rbac_userdb")
	return
}

// MongoStore is the interface which implements the db.Store interface
type MongoStore struct {
	db *mongo.Database
}

// Key interface
type Key interface {
}

// DBconn variable of type Store
var DBconn Store

// Store Interface which implements the data store functions
type Store interface {
	Insert(coll string, key Key, query interface{}, tag string, data interface{}) error
}

// NewMongoStore Return mongo client
func NewMongoStore(name string, store *mongo.Database, svcEp string) (Store, error) {
	if store == nil {
		ip := svcEp
		clientOptions := options.Client()
		clientOptions.ApplyURI(ip)
		if len(os.Getenv("DB_EMCOUI_USERNAME")) > 0 && len(os.Getenv("DB_EMCOUI_PASSWORD")) > 0 {
			clientOptions.SetAuth(options.Credential{
				AuthMechanism: "SCRAM-SHA-256",
				AuthSource:    "rbac_userdb",
				Username:      os.Getenv("DB_EMCOUI_USERNAME"),
				Password:      os.Getenv("DB_EMCOUI_PASSWORD")})
		}
		mongoClient, err := mongo.NewClient(clientOptions)
		if err != nil {
			return nil, err
		}

		err = mongoClient.Connect(context.Background())
		if err != nil {
			return nil, err
		}
		store = mongoClient.Database(name)
	}
	return &MongoStore{
		db: store,
	}, nil
}

// CreateDBClient creates the DB client, currently only mongo
func CreateDBClient(dbType string, dbName string, svcEp string) error {
	var err error
	switch dbType {
	case "mongo":
		DBconn, err = NewMongoStore(dbName, nil, svcEp)
	default:
		fmt.Println(dbType + "DB not supported")
	}
	return err
}

func (m *MongoStore) updateFilter(key interface{}) (primitive.M, error) {

	var n map[string]string
	st, err := json.Marshal(key)
	if err != nil {
		return primitive.M{}, pkgerrors.Errorf("Error Marshalling key: %s", err.Error())
	}
	err = bson.UnmarshalExtJSON(st, true, &n)
	if err != nil {
		return primitive.M{}, pkgerrors.Errorf("Error Unmarshalling key to Bson Map: %s", err.Error())
	}
	p := make(bson.M, len(n))
	for k, v := range n {
		p[k] = v
	}
	filter := bson.M{
		"$set": p,
	}
	return filter, nil
}

func (m *MongoStore) createKeyField(key interface{}) (string, error) {

	var n map[string]string
	st, err := json.Marshal(key)
	if err != nil {
		return "", pkgerrors.Errorf("Error Marshalling key: %s", err.Error())
	}
	err = bson.UnmarshalExtJSON(st, true, &n)
	if err != nil {
		return "", pkgerrors.Errorf("Error Unmarshalling key to Bson Map: %s", err.Error())
	}
	var keys []string
	for k := range n {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	s := "{"
	for _, k := range keys {
		s = s + k + ","
	}
	s = s + "}"
	return s, nil
}

func (m *MongoStore) findFilter(key Key) (primitive.M, error) {

	var bsonMap bson.M
	st, err := json.Marshal(key)
	if err != nil {
		return primitive.M{}, pkgerrors.Errorf("Error Marshalling key: %s", err.Error())
	}
	err = bson.UnmarshalExtJSON(st, true, &bsonMap)
	if err != nil {
		return primitive.M{}, pkgerrors.Errorf("Error Unmarshalling key to Bson Map: %s", err.Error())
	}
	filter := bson.M{
		"$and": []bson.M{bsonMap},
	}
	return filter, nil
}

var decodeBytes = func(sr *mongo.SingleResult) (bson.Raw, error) {
	return sr.DecodeBytes()
}

// Insert is used to insert/add element to a document
func (m *MongoStore) Insert(coll string, key Key, query interface{}, tag string, data interface{}) error {

	c := m.db.Collection(coll)
	ctx := context.Background()

	filter, err := m.findFilter(key)
	if err != nil {
		return err
	}
	// Create and add key tag
	s, err := m.createKeyField(key)
	if err != nil {
		return err
	}
	_, err = decodeBytes(
		c.FindOneAndUpdate(
			ctx,
			filter,
			bson.D{
				{"$set", bson.D{
					{tag, data},
					{"key", s},
				}},
			},
			options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)))

	if err != nil {
		return pkgerrors.Errorf("Error updating master table: %s", err.Error())
	}
	if query == nil {
		return nil
	}

	// Update to add Query fields
	update, err := m.updateFilter(query)
	if err != nil {
		return err
	}
	_, err = c.UpdateOne(
		ctx,
		filter,
		update)

	if err != nil {
		return pkgerrors.Errorf("Error updating Query fields: %s", err.Error())
	}
	return nil
}
