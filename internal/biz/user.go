package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// Email represents the email verification subsystem
type Email struct {
	Address  string `bson:"address" json:"address"`   // Email address
	Verified bool   `bson:"verified" json:"verified"` // Verification status
}

// Status represents the account status
type Status struct {
	Disable bool   `bson:"disable" json:"disable"` // Whether the account is disabled
	Reason  string `bson:"reason" json:"reason"`   // Reason for disabling
}

// User represents a user in the system
type User struct {
	ID       bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // MongoDB ObjectID
	UserName string        `bson:"username" json:"username"`          // Unique username
	Password string        `bson:"password" json:"password"`          // PBKDF2-SHA256 encrypted password
	Date     string        `bson:"date" json:"date"`                  // Registration date
	IP       string        `bson:"ip" json:"ip"`                      // Registration IP address
	Browser  string        `bson:"browser" json:"browser"`            // Registration browser info
	Hash     string        `bson:"hash" json:"hash"`                  // SHA256 hash of username
	LastDate string        `bson:"lastDate" json:"lastDate"`          // Last login date
	LastIP   string        `bson:"lastIP" json:"lastIP"`              // Last login IP
	Email    Email         `bson:"email" json:"email"`                // Email verification subsystem
	Avatar   []byte        `bson:"avatar,omitempty" json:"avatar"`    // Avatar binary data (optional)
	Like     []int64       `bson:"like" json:"like"`                  // List of favorite resource IDs
	Group    []string      `bson:"group" json:"group"`                // User groups (default: ["user"])
	Status   *Status       `bson:"status,omitempty" json:"status"`    // Account status (optional)
}

type UserRepo interface {
	FindById(ctx context.Context, id string) (*User, error)
	FetchUserDataMap(ctx context.Context, userNames []string) (map[string]User, error)
	List(ctx context.Context, pageNo, pageSize int64, orderBy, filter string) ([]*User, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc UserUseCase) UserGroup(context context.Context, userNames []string) (map[string]User, error) {
	return uc.repo.FetchUserDataMap(context, userNames)
}
