package biz

import (
	"YYeTsBot-Go/internal/middleware"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v5"
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
	FindUser(ctx context.Context, username, password string) (*User, error)
	Save(ctx context.Context, user *User) (*User, error)
	FetchUserDataMap(ctx context.Context, userNames []string) (map[string]User, error)
	List(ctx context.Context, pageNo, pageSize int64, orderBy, filter string) ([]*User, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

// 默认密码安全参数
const (
	defaultSaltSize   = 16
	defaultIterations = 29000
	keyLength         = 32
)

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc UserUseCase) UserGroup(context context.Context, userNames []string) (map[string]User, error) {
	return uc.repo.FetchUserDataMap(context, userNames)
}

func (uc UserUseCase) Login(ctx context.Context, username, password string) (*User, error) {
	passwordHash, err := PasswordHash(password)

	user, err := uc.repo.FindUser(ctx, username, passwordHash)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return user, nil
	}

	user = &User{
		UserName: username,
		Password: passwordHash,
		Date:     time.Now().Format(time.DateTime),
		Hash:     fmt.Sprintf("%x", sha256.Sum256([]byte(username))),
		Group:    []string{"user"},
		IP:       ctx.Value(middleware.ClientIPKey).(string),
	}

	user, err = uc.repo.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (uc UserUseCase) FindById(ctx context.Context, id string) (*User, error) {
	return uc.repo.FindById(ctx, id)
}

// PasswordHash 生成密码哈希
func PasswordHash(password string) (string, error) {
	salt := make([]byte, defaultSaltSize)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("generate salt failed: %w", err)
	}

	hash := pbkdf2.Key([]byte(password), salt, defaultIterations, keyLength, sha256.New)

	saltEncoded := base64.StdEncoding.EncodeToString(salt)
	hashEncoded := base64.StdEncoding.EncodeToString(hash)

	return fmt.Sprintf("pbkdf2-sha256$%d$%s$%s", defaultIterations, saltEncoded, hashEncoded), nil
}

func (uc UserUseCase) GenerateToken(user *User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": user.ID.Hex(),
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(24 * time.Hour).Unix(),
		"sub": user.UserName,
	})

	secretKey := []byte("your-secret-key")
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(secretKey)

	return tokenString, err
}
