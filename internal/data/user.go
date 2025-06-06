package data

import (
	"YYeTsBot-Go/internal/biz"
	"context"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/pbkdf2"
	"strconv"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (u *userRepo) FetchUserDataMap(ctx context.Context, userNames []string) (map[string]biz.User, error) {
	collection := u.data.db.Collection("users")
	filter := bson.M{"username": bson.M{"$in": userNames}}
	projection := bson.M{"username": 1, "group": 1, "avatar": 1}
	options := options.Find().SetProjection(projection)

	cursor, err := collection.Find(ctx, filter, options)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	userDataMap := make(map[string]biz.User)
	for cursor.Next(ctx) {
		var user biz.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		userDataMap[user.UserName] = user
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return userDataMap, nil
}

func (u *userRepo) Save(ctx context.Context, user *biz.User) (*biz.User, error) {
	collection := u.data.db.Collection("users")
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = result.InsertedID.(bson.ObjectID)
	return user, err
}

func (u *userRepo) FindById(ctx context.Context, id string) (*biz.User, error) {
	objectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	collection := u.data.db.Collection("users")
	filter := bson.M{"_id": objectId}
	var user biz.User
	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindUser implements biz.UserRepo.
func (u *userRepo) FindUser(ctx context.Context, username, password string) (*biz.User, error) {
	collection := u.data.db.Collection("users")
	filter := bson.M{
		"username": username,
		"deleted_at": bson.M{
			"$exists": false,
		},
	}
	var user biz.User
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	verify, err := VerifyPassword(password, user.Password)
	if !verify {
		return nil, errors.New("wrong password")
	}
	return &user, nil
}

// VerifyPassword verifies if the provided password matches the hashed value.
func VerifyPassword(password, hashed string) (bool, error) {
	parts := strings.Split(hashed, "$")
	if len(parts) != 4 || parts[0] != "pbkdf2-sha256" {
		return false, errors.New("invalid hashed password format")
	}

	iter, err := strconv.Atoi(parts[1])
	if err != nil {
		return false, fmt.Errorf("invalid iteration number: %w", err)
	}

	salt, err := base64.StdEncoding.DecodeString(parts[2])
	if err != nil {
		return false, fmt.Errorf("decode salt failed: %w", err)
	}

	storedHash, err := base64.StdEncoding.DecodeString(parts[3])
	if err != nil {
		return false, fmt.Errorf("decode stored hash failed: %w", err)
	}

	computedHash := pbkdf2.Key([]byte(password), salt, iter, len(storedHash), sha256.New)

	if subtle.ConstantTimeCompare(computedHash, storedHash) == 1 {
		return true, nil
	}

	return false, nil
}

// List implements biz.UserRepo.
func (u *userRepo) List(ctx context.Context, pageNo int64, pageSize int64, orderBy string, filter string) ([]*biz.User, error) {
	// Calculate skip value for pagination
	skip := (pageNo - 1) * pageSize

	// Create filter options
	filterOptions := options.Find().
		SetSkip(skip).
		SetLimit(pageSize)

	// Add sorting if orderBy is specified
	if orderBy != "" {
		var sortingKey string = "_id"
		var sortingValue int = 1

		// syntax foo, foo desc
		ordering := strings.SplitN(orderBy, " ", 2)

		sortingKey = ordering[0]
		if ordering[1] == "desc" {
			sortingValue = -1
		}

		filterOptions.SetSort(bson.D{{Key: sortingKey, Value: sortingValue}})
	}

	// Create filter query
	var filterQuery bson.D
	if filter != "" {
		// Parse the filter string into a MongoDB query
		// This is a simple implementation - you might want to enhance it based on your needs
		filterQuery = bson.D{
			{Key: "$or", Value: []bson.D{
				{{Key: "username", Value: bson.D{{Key: "$regex", Value: filter}}}},
				{{Key: "email.address", Value: bson.D{{Key: "$regex", Value: filter}}}},
			}},
		}
	}

	// Execute the query
	cursor, err := u.data.db.Collection("users").Find(ctx, filterQuery, filterOptions)
	if err != nil {
		u.log.Errorf("Failed to query users: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	// Decode results
	var users []*biz.User
	if err := cursor.All(ctx, &users); err != nil {
		u.log.Errorf("Failed to decode users: %v", err)
		return nil, err
	}

	return users, nil
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
