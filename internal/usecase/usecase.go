package usecase

import (
	"context"
	"time"
	"zavad/config"
	"zavad/internal/entities"
	"zavad/internal/repository/postgres"
	protos "zavad/pkg/proto/gen/go"

	jwt "github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type Usecase struct {
	cfg     *config.ConfigModel
	log     *zap.Logger
	Repo    *postgres.Repository
}

const (
	clientRole    = `CLIENT`
	adminRole     = `ADMIN`
	SuccessStatus = `Success`
	BadStatus     = `Fail`
)

func NewUsecase(logger *zap.Logger, Repo *postgres.Repository, cfg *config.ConfigModel) (*Usecase, error) {
	return &Usecase{
		cfg:     cfg,
		log:     logger,
		Repo:    Repo,
	}, nil
}

func (uc *Usecase) GetUserToken(ctx context.Context, user *protos.GetUserTokenRequest) (*protos.GetUserTokenResponse, error) {
	userID, err := uc.Repo.GetUserIDbyPhoneOrEmail(ctx, user.Login, user.Login)
	if err != nil {
		uc.log.Error("failed to get user id", zap.Error(err))
		return nil, err
	}

	// u, err := uc.authViaYClients(ctx, userID)
	// if err != nil {
	// 	uc.log.Error("failed to create user token", zap.Error(err))
	// 	return nil, err
	// }

	u := &entities.User{
		ID:    userID,
		// Name:  User.Data.Name,
		// Phone: User.Data.Phone,
		// Login: User.Data.Phone,
		// Email: User.Data.Email,
	}

	return uc.createUserTokens(u)
}

func (uc *Usecase) CreateUser(ctx context.Context) (*protos.GetUserTokenResponse, error) {
	// createUser, err := uc.YClient.Register(ctx, user)
	// if err != nil {
	// 	uc.log.Error("failed to create user in yclients", zap.Error(err))
	// 	return nil, err
	// }

	// u := &entities.User{
	// 	ID:    createUser.Data.ID,
	// 	Name:  createUser.Data.Name,
	// 	Phone: createUser.Data.Phone,
	// 	Login: createUser.Data.Phone,
	// 	Email: createUser.Data.Email,
	// }

	// err = uc.Repo.CreateUser(ctx, u)
	// if err != nil {
	// 	uc.log.Error("failed to create user in db", zap.Error(err))
	// 	return nil, err
	// }

	// //todo create user in db

	// return uc.createUserTokens(u)
	return nil, nil
}

func (uc *Usecase) createUserTokens(u *entities.User) (*protos.GetUserTokenResponse, error) {
	accessToken, err := uc.createAccessToken(u)
	if err != nil {
		uc.log.Error("failed to get access token", zap.Error(err))
		return nil, err
	}

	refreshToken, err := uc.createRefreshToken(u.ID)
	if err != nil {
		uc.log.Error("failed to get refresh token", zap.Error(err))
		return nil, err
	}

	return &protos.GetUserTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (uc *Usecase) createAccessToken(user *entities.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	//claims["user_token"] = user.UserToken
	claims["name"] = user.Name
	claims["phone"] = user.Phone
	claims["login"] = user.Login
	claims["email"] = user.Email
	claims["avatar"] = user.Avatar
	claims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()

	tokenString, err := token.SignedString([]byte(uc.cfg.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (uc *Usecase) createRefreshToken(userID int) (string, error) {
	Claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Minute * 60 * 24 * 365 * 10).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	refreshToken.Header["kid"] = "signing"

	signedString, err := refreshToken.SignedString([]byte(uc.cfg.Secret))
	if err != nil {
		return "", err
	}

	return signedString, nil
}


func (uc *Usecase) GetProfileData() {
	
}
func (uc *Usecase) GetMessangerChat() {
	
}
func (uc *Usecase) GetMessangerChats() {
	
}
func (uc *Usecase) SendMessage() {
	
}
func (uc *Usecase) GetDocsSectionData() {
	
}
func (uc *Usecase) GetDocumentsTemplates() {
	
}
func (uc *Usecase) CreateDocument() {
	
}
func (uc *Usecase) GetDocuments() {
	
}
func (uc *Usecase) SignDocument() {
	
}
func (uc *Usecase) CreateNotification() {
	
}
func (uc *Usecase) DeleteNotification() {
	
}
func (uc *Usecase) GetNotifications() {
	
}