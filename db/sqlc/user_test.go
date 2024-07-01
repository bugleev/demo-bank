package db

import (
	"context"
	"testing"

	"github.com/bugleev/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		Email:          util.RandomEmail(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	testUser := createRandomUser(t)
	userInDb, err := testQueries.GetUser(context.Background(), testUser.Username)

	require.NoError(t, err)
	require.NotEmpty(t, userInDb)

	require.Equal(t, userInDb.Username, testUser.Username)
	require.Equal(t, userInDb.Email, testUser.Email)
	require.Equal(t, userInDb.FullName, testUser.FullName)
	require.Equal(t, userInDb.HashedPassword, testUser.HashedPassword)
	require.Equal(t, userInDb.PasswordChangedAt, testUser.PasswordChangedAt)
	require.Equal(t, userInDb.CreatedAt, testUser.CreatedAt)
}
