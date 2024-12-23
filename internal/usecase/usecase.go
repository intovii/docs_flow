package usecase

import (
	"context"
	"docs_flow/config"
	"docs_flow/internal/entities"
	"docs_flow/internal/repository/postgres"
	"time"

	protos "docs_flow/pkg/proto/gen/go"

	jwt "github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type Usecase struct {
	cfg  *config.ConfigModel
	log  *zap.Logger
	Repo *postgres.Repository
}

const (
	clientRole    = `CLIENT`
	adminRole     = `ADMIN`
	SuccessStatus = `Success`
	BadStatus     = `Fail`
)

func NewUsecase(logger *zap.Logger, Repo *postgres.Repository, cfg *config.ConfigModel) (*Usecase, error) {
	return &Usecase{
		cfg:  cfg,
		log:  logger,
		Repo: Repo,
	}, nil
}

//func (uc *Usecase) GetUserToken(ctx context.Context, user *protos.GetUserTokenRequest) (*protos.GetUserTokenResponse, error) {
//	userID, err := uc.Repo.GetUserIDbyPhoneOrEmail(ctx, user.Login, user.Login)
//	if err != nil {
//		uc.log.Error("failed to get user id", zap.Error(err))
//		return nil, err
//	}
//
//	// u, err := uc.authViaYClients(ctx, userID)
//	// if err != nil {
//	// 	uc.log.Error("failed to create user token", zap.Error(err))
//	// 	return nil, err
//	// }
//
//	u := &entities.User{
//		ID: userID,
//		// Name:  User.Data.Name,
//		// Phone: User.Data.Phone,
//		// Login: User.Data.Phone,
//		// Email: User.Data.Email,
//	}
//
//	return uc.createUserTokens(u)
//}

//func (uc *Usecase) CreateUser(ctx context.Context) (*protos.GetUserTokenResponse, error) {
//	// createUser, err := uc.YClient.Register(ctx, user)
//	// if err != nil {
//	// 	uc.log.Error("failed to create user in yclients", zap.Error(err))
//	// 	return nil, err
//	// }
//
//	// u := &entities.User{
//	// 	ID:    createUser.Data.ID,
//	// 	Name:  createUser.Data.Name,
//	// 	Phone: createUser.Data.Phone,
//	// 	Login: createUser.Data.Phone,
//	// 	Email: createUser.Data.Email,
//	// }
//
//	// err = uc.Repo.CreateUser(ctx, u)
//	// if err != nil {
//	// 	uc.log.Error("failed to create user in db", zap.Error(err))
//	// 	return nil, err
//	// }
//
//	// //todo create user in db
//
//	// return uc.createUserTokens(u)
//	return nil, nil
//}

//func (uc *Usecase) createUserTokens(u *entities.User) (*protos.GetUserTokenResponse, error) {
//	accessToken, err := uc.createAccessToken(u)
//	if err != nil {
//		uc.log.Error("failed to get access token", zap.Error(err))
//		return nil, err
//	}

//refreshToken, err := uc.createRefreshToken(u.ID)
//if err != nil {
//	uc.log.Error("failed to get refresh token", zap.Error(err))
//	return nil, err
//}
//
//return &protos.GetUserTokenResponse{
//	AccessToken:  accessToken,
//	RefreshToken: refreshToken,
//}, nil
//}

//func (uc *Usecase) createAccessToken(user *entities.User) (string, error) {
//	token := jwt.New(jwt.SigningMethodHS256)
//	claims := token.Claims.(jwt.MapClaims)
//	claims["id"] = user.ID
//	//claims["user_token"] = user.UserToken
//	claims["name"] = user.Name
//	claims["phone"] = user.Phone
//	claims["login"] = user.Login
//	claims["email"] = user.Email
//	claims["avatar"] = user.Avatar
//	claims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()
//
//	tokenString, err := token.SignedString([]byte(uc.cfg.Secret))
//	if err != nil {
//		return "", err
//	}
//
//	return tokenString, nil
//}

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

func (uc *Usecase) GetProfileData(ctx context.Context, userID int64) (*protos.GetProfileDataResponse, error) {
	user, err := uc.Repo.GetUser(ctx, userID)
	if err != nil {
		uc.log.Error("Fail to get profile data(user): ", zap.Error(err))
		return nil, err
	}
	documents, err := uc.Repo.GetDocuments(ctx, userID, "all")
	if err != nil {
		uc.log.Error("Fail to get profile data(documents): ", zap.Error(err))
		return nil, err
	}
	return &protos.GetProfileDataResponse{
		User: &protos.User{
			ID:           user.ID,
			Name:         user.Name,
			Login:        user.Login,
			PasswordHash: user.PasswordHash,
			Role:         user.Role,
			AvaPath:      user.AvaPath,
		},

		Documents: entities.ConvertDocumentsToProto(documents),
	}, nil
}
func (uc *Usecase) GetMessengerChat(ctx context.Context, chatID int64, userID int64) (*protos.GetMessangerChatResponse, error) {
	messages, err := uc.Repo.GetMessagesByChatID(ctx, chatID)
	if err != nil {
		uc.log.Error("Fail to get messenger chat(messages): ", zap.Error(err))
		return nil, err
	}
	receiver, err := uc.Repo.GetUserBy(ctx, chatID)
	if err != nil {
		uc.log.Error("Fail to get messenger chat(messages): ", zap.Error(err))
		return nil, err
	}
	return nil, nil
}
func (uc *Usecase) GetMessengerChats(ctx context.Context, userID int64) (*protos.GetMessangerChatsResponse, error) {
	return nil, nil
}
func (uc *Usecase) SendMessage(ctx context.Context, message *protos.Message, userID int64) (*protos.SendMessageResponse, error) {
	return nil, nil
}
func (uc *Usecase) GetDocsSectionData(ctx context.Context, userID int64) (*protos.GetDocsSectionDataResponse, error) {
	return nil, nil
}

func (uc *Usecase) CreateDocument(ctx context.Context, document *protos.Document) (*protos.CreateDocumentResponse, error) {
	return nil, nil
}
func (uc *Usecase) GetDocuments(ctx context.Context, userId int64) (*protos.CreateDocumentResponse, error) {
	return nil, nil
}
func (uc *Usecase) SignDocument(ctx context.Context, documentID int64) (*protos.SignDocumentResponse, error) {
	return nil, nil
}
func (uc *Usecase) CreateNotification(ctx context.Context, notification *protos.Notification) (*protos.CreateNotificationResponse, error) {
	return nil, nil
}
func (uc *Usecase) DeleteNotification(ctx context.Context, notification *protos.Notification) (*protos.DeleteNotificationResponse, error) {
	return nil, nil
}
func (uc *Usecase) GetNotifications(ctx context.Context, userID int64) (*protos.GetNotificationsResponse, error) {
	return nil, nil
}

// func (uc *Usecase) GetDocumentsTemplates(ctx context.Context) {
//
// }
