package repo

import (
	"log"
	"testing"

	"github.com/anandawira/anandapay/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserRepoTestSuite struct {
	suite.Suite

	DB   *gorm.DB
	repo domain.UserRepository
}

func (ts *UserRepoTestSuite) SetupSuite() {
	// Hardcore, later change to env variable
	dsn := "root:example@tcp(127.0.0.1:3306)/anandapay-test-user?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	ts.DB = db
	ts.repo = NewUserRepository(db)
}

func (ts *UserRepoTestSuite) SetupTest() {
	ts.DB.AutoMigrate(&domain.User{}, &domain.Wallet{})
}

func (ts *UserRepoTestSuite) TearDownTest() {
	ts.DB.Migrator().DropTable(&domain.User{}, &domain.Wallet{})
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
		user := domain.User{
			FullName:       "User1",
			Email:          "email1@gmail.com",
			HashedPassword: "hashedPassword1",
			IsVerified:     false,
		}

		err := ts.repo.Insert(user.FullName, user.Email, user.HashedPassword, user.IsVerified)
		assert.NoError(t, err)
	})

	ts.T().Run("It should not insert to the database if email already exist.", func(t *testing.T) {
		const email string = "duplicate@gmail.com"
		user1 := domain.User{
			FullName:       "User1",
			Email:          email,
			HashedPassword: "hashedPassword1",
			IsVerified:     false,
		}
		user2 := domain.User{
			FullName:       "User2",
			Email:          email,
			HashedPassword: "hashedPassword2",
			IsVerified:     false,
		}

		err := ts.repo.Insert(user1.FullName, user1.Email, user1.HashedPassword, user1.IsVerified)
		require.NoError(t, err)

		err = ts.repo.Insert(user2.FullName, user2.Email, user2.HashedPassword, user2.IsVerified)
		assert.Error(t, err)
	})
}

func (ts *UserRepoTestSuite) TestGetOne() {
	ts.T().Run("It should return user and error nil if record found", func(t *testing.T) {
		user := domain.User{
			FullName:       "User1",
			Email:          "email1@gmail.com",
			HashedPassword: "hashedPassword1",
			IsVerified:     false,
		}

		err := ts.repo.Insert(user.FullName, user.Email, user.HashedPassword, user.IsVerified)
		require.NoError(t, err)

		resUser, wallet, err := ts.repo.GetByEmail(user.Email)
		require.NoError(t, err)
		assert.Equal(t, user.Email, resUser.Email)
		assert.Equal(t, user.HashedPassword, resUser.HashedPassword)
		assert.Equal(t, resUser.ID, wallet.UserID)
	})

	ts.T().Run("It should return error if record not found", func(t *testing.T) {
		_, _, err := ts.repo.GetByEmail("noemail@gmail.com")
		assert.Error(t, err)
	})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}
