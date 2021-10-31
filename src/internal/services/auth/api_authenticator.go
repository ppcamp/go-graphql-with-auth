package auth

// import (
// 	"errors"

// 	loginmodels "github.com/eventials/vlab-baby-app-api/internal/models/login"
// 	"github.com/eventials/vlab-baby-app-api/internal/services/jwt"
// 	"github.com/eventials/vlab-baby-app-api/internal/services/redis"
// 	"github.com/eventials/vlab-baby-app-api/internal/services/sms"
// 	"github.com/eventials/vlab-baby-app-api/internal/utils"
// 	"github.com/sirupsen/logrus"
// )

// var (
// 	ErrInvalidPayload      = errors.New("invalid payload")
// 	ErrKeyValue            = errors.New("invalid payload infos")
// 	ErrInvalidLoginPayload = errors.New("invalid login payload")
// )

// type AuthService struct {
// 	redisService redis.RedisService
// }

// func NewAuthService(r redis.RedisService) *AuthService {
// 	return &AuthService{redisService: r}
// }

// // ApiAuthenticator does the login checkup.
// // If everything occurrs well, it'll returns the JWT session to be encoded into
// // JWT
// func (a *AuthService) ApiAuthenticator(login *loginmodels.LoginPayload) (jwt.Session, error) {
// 	login.PhoneNumber = utils.RemovePhoneEspecialChars(login.PhoneNumber)
// 	log := logrus.WithFields(logrus.Fields{
// 		"phone_number": login.PhoneNumber,
// 		"code":         login.AccessCode,
// 	})

// 	var code string
// 	err := a.redisService.Get(sms.FormatSMSKey(login.PhoneNumber), &code)
// 	if err != nil {
// 		log.WithError(err).Error("failed to get the key-value from redis")
// 		return jwt.BLANK_SESSION, ErrKeyValue
// 	}

// 	if code == login.AccessCode {
// 		log.Debug("code valid, returning a valid jwt token")
// 		return jwt.Session{
// 			LoginSession: jwt.LoginSession{PhoneNumber: login.PhoneNumber},
// 		}, nil
// 	}

// 	return jwt.BLANK_SESSION, ErrInvalidPayload
// }
