package database

import (
	"projeto-modelo/internal/entity"
	"projeto-modelo/internal/infra/database"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserDBTestSuite struct {
	suite.Suite
	db     *gorm.DB
	userDB *database.User
}

func (suite *UserDBTestSuite) SetupTest() {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	suite.Require().NoError(err, "failed to connect to database")

	err = db.AutoMigrate(&entity.User{})
	suite.Require().NoError(err, "failed to migrate database")

	suite.db = db
	suite.userDB = database.NewUser(db)
}

func (suite *UserDBTestSuite) TearDownTest() {
	if suite.db != nil {
		sqlDB, err := suite.db.DB()
		if err == nil {
			sqlDB.Close()
		}
	}
}

func (suite *UserDBTestSuite) createTestUser(name, email, password string) *entity.User {
	user, err := entity.NewUser(name, email, password)
	suite.Require().NoError(err, "failed to create user entity")

	err = suite.userDB.Create(user)
	suite.Require().NoError(err, "failed to save user to database")

	return user
}

func TestUserDBTestSuite(t *testing.T) {
	suite.Run(t, new(UserDBTestSuite))
}

func (suite *UserDBTestSuite) TestCreateUser() {
	user, err := entity.NewUser("Tote Araujo", "tote@tapp.dev.br", "password")
	assert.NoError(suite.T(), err)

	err = suite.userDB.Create(user)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), user.ID)
	assert.Equal(suite.T(), "Tote Araujo", user.Name)
	assert.Equal(suite.T(), "tote@tapp.dev.br", user.Email)
	assert.NotEmpty(suite.T(), user.Password)

	// Verify user was saved in database
	var userFound entity.User
	err = suite.db.First(&userFound, "id = ?", user.ID).Error
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), user.ID, userFound.ID)
	assert.Equal(suite.T(), user.Name, userFound.Name)
	assert.Equal(suite.T(), user.Email, userFound.Email)
	assert.NotEmpty(suite.T(), userFound.Password)
}

func (suite *UserDBTestSuite) TestFindByEmail() {
	// Create a test user
	createdUser := suite.createTestUser("Tote Araujo", "tote@tapp.dev.br", "password")

	// Find the user by email
	foundUser, err := suite.userDB.FindByEmail(createdUser.Email)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), createdUser.ID, foundUser.ID)
	assert.Equal(suite.T(), createdUser.Name, foundUser.Name)
	assert.Equal(suite.T(), createdUser.Email, foundUser.Email)
	assert.NotEmpty(suite.T(), foundUser.Password)

	// Test finding non-existent user
	nonExistentUser, err := suite.userDB.FindByEmail("nonexistent@email.com")
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), nonExistentUser)
}

func (suite *UserDBTestSuite) TestFindByID() {
	// Create a test user
	createdUser := suite.createTestUser("Tote Araujo", "tote@tapp.dev.br", "password")

	// Find the user by ID
	foundUser, err := suite.userDB.FindByID(createdUser.ID.String())
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), createdUser.ID, foundUser.ID)
	assert.Equal(suite.T(), createdUser.Name, foundUser.Name)
	assert.Equal(suite.T(), createdUser.Email, foundUser.Email)
	assert.NotEmpty(suite.T(), foundUser.Password)

	// Test finding non-existent user
	nonExistentUser, err := suite.userDB.FindByID("non-existent-id")
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), nonExistentUser)
}

func (suite *UserDBTestSuite) TestUpdateUser() {
	// Create a test user
	user := suite.createTestUser("Tote Araujo", "tote@tapp.dev.br", "password")

	// Update the user
	user.Name = "Tote Updated"
	user.Email = "tote.updated@tapp.dev.br"

	err := suite.userDB.Update(user)
	assert.NoError(suite.T(), err)

	// Verify the update
	updatedUser, err := suite.userDB.FindByID(user.ID.String())
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Tote Updated", updatedUser.Name)
	assert.Equal(suite.T(), "tote.updated@tapp.dev.br", updatedUser.Email)
}

func (suite *UserDBTestSuite) TestDeleteUser() {
	// Create a test user
	user := suite.createTestUser("Tote Araujo", "tote@tapp.dev.br", "password")

	// Delete the user
	err := suite.userDB.Delete(user.ID.String())
	assert.NoError(suite.T(), err)

	// Verify the user was deleted
	deletedUser, err := suite.userDB.FindByID(user.ID.String())
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), deletedUser)
}

func (suite *UserDBTestSuite) TestFindAllUsers() {
	// Create multiple test users
	suite.createTestUser("User 1", "user1@test.com", "password1")
	suite.createTestUser("User 2", "user2@test.com", "password2")
	suite.createTestUser("User 3", "user3@test.com", "password3")

	// Test finding all users
	users, err := suite.userDB.FindAll(1, 10, "id asc")
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), users, 3)

	// Test pagination
	users, err = suite.userDB.FindAll(1, 2, "id asc")
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), users, 2)
}
