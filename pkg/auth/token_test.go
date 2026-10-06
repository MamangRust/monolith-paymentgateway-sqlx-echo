package auth

import (
	"errors"
	"testing"
	"time"

	mock_auth "github.com/MamangRust/monolith-payment-gateway-pkg/auth/mocks"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

const secretKey = "mySecretKey"

func TestGenerateToken_Success(t *testing.T) {
	mgr, err := NewManager(secretKey)
	assert.NoError(t, err)

	tokenStr, err := mgr.GenerateToken(123, "gateway")
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenStr)
}

func TestValidateToken_Success(t *testing.T) {
	mgr, _ := NewManager(secretKey)

	tokenStr, err := mgr.GenerateToken(456, "apigateway")
	assert.NoError(t, err)

	userID, err := mgr.ValidateToken(tokenStr)
	assert.NoError(t, err)
	assert.Equal(t, "456", userID)
}

func TestValidateToken_InvalidToken(t *testing.T) {
	mgr, _ := NewManager(secretKey)

	badToken := "this.is.invalid.token"
	_, err := mgr.ValidateToken(badToken)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse token")
}

func TestValidateToken_Expired(t *testing.T) {
	mgr, _ := NewManager(secretKey)

	expiredTime := time.Now().Add(-1 * time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expiredTime),
		Subject:   "999",
		Audience:  []string{"expired"},
	})

	tokenStr, err := token.SignedString([]byte(secretKey))
	assert.NoError(t, err)

	userID, err := mgr.ValidateToken(tokenStr)

	assert.Error(t, err)
	assert.Equal(t, "", userID)
	assert.True(t, errors.Is(err, ErrTokenExpired))
}

func TestValidateToken_Expired_WithMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMgr := mock_auth.NewMockTokenManager(ctrl)

	mockMgr.EXPECT().
		ValidateToken(gomock.Any()).
		Return("", ErrTokenExpired)

	userID, err := mockMgr.ValidateToken("expired.jwt.token")
	assert.Error(t, err)
	assert.Equal(t, "", userID)
	assert.True(t, errors.Is(err, ErrTokenExpired))
}
