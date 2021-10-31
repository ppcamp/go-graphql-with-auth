package auth_test

// import (
// 	"testing"

// 	"github.com/eventials/vlab-baby-app-api/internal/helpers/validators"
// 	loginmodels "github.com/eventials/vlab-baby-app-api/internal/models/login"
// 	"github.com/eventials/vlab-baby-app-api/internal/services/auth"
// 	redismks "github.com/eventials/vlab-baby-app-api/internal/services/redis/mocks"
// 	"github.com/eventials/vlab-baby-app-api/internal/services/sms"
// 	"github.com/eventials/vlab-baby-app-api/internal/utils"
// 	"github.com/go-playground/validator/v10"
// 	"github.com/go-redis/redis/v8"
// 	"github.com/stretchr/testify/mock"
// 	"github.com/stretchr/testify/suite"
// )

// type AuthService struct {
// 	suite.Suite

// 	redis       *redismks.MockedRedisService
// 	authService auth.AuthService
// 	phone       string
// 	code        string
// 	validator   *validator.Validate
// }

// func TestPostUserSMS(t *testing.T) {
// 	suite.Run(t, new(AuthService))
// }

// func (t *AuthService) SetupTest() {
// 	t.redis = utils.Must(redismks.NewRedisService(nil)).(*redismks.MockedRedisService)
// 	t.redis.GetCode = "123456"
// 	t.authService = *auth.NewAuthService(t.redis)
// 	t.phone = "(11) 91112-8834"
// 	t.code = t.redis.GetCode

// 	t.validator = validators.RegisterAndGetValidators()
// }

// func (t *AuthService) mockSetRedis(phone string, err error) {
// 	key := sms.FormatSMSKey(utils.RemovePhoneEspecialChars(phone))
// 	t.redis.On("Get", key, mock.AnythingOfType("*string")).Return(err)
// }

// func (t *AuthService) TestLoginPayloadShouldError() {
// 	assert := t.Require()

// 	payload := &loginmodels.LoginPayload{
// 		PhoneNumber: "(11) 1122-3344",
// 		AccessCode:  t.code,
// 	}
// 	t.mockSetRedis(payload.PhoneNumber, nil)

// 	err := t.validator.Struct(payload)
// 	assert.Error(err)

// 	payload.PhoneNumber = "(11) 91122-3344"
// 	payload.AccessCode = "1234"
// 	assert.Error(err)
// }

// func (t *AuthService) TestRedisShouldError() {
// 	assert := t.Require()
// 	payload := &loginmodels.LoginPayload{
// 		PhoneNumber: t.phone,
// 		AccessCode:  t.code,
// 	}

// 	t.mockSetRedis(payload.PhoneNumber, redis.ErrClosed)
// 	_, err := t.authService.ApiAuthenticator(payload)

// 	assert.ErrorIs(err, auth.ErrKeyValue)
// }

// func (t *AuthService) TestShouldLoginSuccessfully() {
// 	assert := t.Require()
// 	payload := &loginmodels.LoginPayload{
// 		PhoneNumber: t.phone,
// 		AccessCode:  t.code,
// 	}

// 	t.mockSetRedis(payload.PhoneNumber, nil)
// 	_, err := t.authService.ApiAuthenticator(payload)

// 	assert.ErrorIs(err, nil)
// }
