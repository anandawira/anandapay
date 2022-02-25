package repo

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/anandawira/anandapay/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserRepoTestSuite struct {
	suite.Suite

	DB   *gorm.DB
	repo model.UserRepository
}

func (ts *UserRepoTestSuite) SetupSuite() {
	// Hardcore, later change to env variable
	dsn := "root:example@tcp(127.0.0.1:3306)/anandapay-test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	ts.DB = db
	ts.repo = NewUserRepository(db)
}

func (ts *UserRepoTestSuite) SetupTest() {
	ts.DB.Migrator().DropTable(&model.User{})
	ts.DB.AutoMigrate(&model.User{})
}

func (ts *UserRepoTestSuite) TearDownSuite() {
	conn, err := ts.DB.DB()
	if err != nil {
		log.Fatal("Database not found")
	}
	conn.Close()
}

func (ts *UserRepoTestSuite) TestInsert() {
	ts.T().Run("It should insert to the database.", func(t *testing.T) {
		user := model.User{
			FullName:       "User1",
			Email:          "email1@gmail.com",
			HashedPassword: "hashedPassword1",
			IsVerified:     false,
		}

		err := ts.repo.Insert(context.TODO(), user.FullName, user.Email, user.HashedPassword, user.IsVerified)
		assert.NoError(t, err)
	})

	ts.T().Run("It should not insert to the database if email already exist.", func(t *testing.T) {
		user := model.User{
			FullName:       "User2",
			Email:          "email1@gmail.com",
			HashedPassword: "hashedPassword2",
			IsVerified:     false,
		}
		err := ts.repo.Insert(context.TODO(), user.FullName, user.Email, user.HashedPassword, user.IsVerified)
		fmt.Println(err)
		// Check error
		assert.Error(t, err)
	})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}