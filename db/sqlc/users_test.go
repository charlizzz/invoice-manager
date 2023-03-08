package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createTestUser(t *testing.T) Users {
	arg := CreateUserParams{
		FirstName: "Johnny",
		LastName:  "Billy Joe",
		Balance:   12345,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.Balance, user.Balance)

	require.NotZero(t, user.ID)

	return user
}

func TestCreateUser(t *testing.T) {
	createTestUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createTestUser(t)

	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.FirstName, user2.FirstName)
	require.Equal(t, user1.LastName, user2.LastName)
	require.Equal(t, user1.Balance, user2.Balance)
}

func TestUpdateUserBalance(t *testing.T) {
	user1 := createTestUser(t)

	arg := UpdateUserBalanceParams{
		ID:      user1.ID,
		Balance: 543,
	}

	user2, err := testQueries.UpdateUserBalance(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.FirstName, user2.FirstName)
	require.Equal(t, user1.LastName, user2.LastName)
	require.Equal(t, arg.Balance, user2.Balance)
}

func TestDeleteUser(t *testing.T) {
	user1 := createTestUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.ID)

	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}

func TestListUsers(t *testing.T) {
	var limit int = 2
	for i := 0; i < 10; i++ {
		createTestUser(t)
	}

	arg := ListUsersParams{
		Limit:  int32(limit),
		Offset: 2,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, limit)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}
