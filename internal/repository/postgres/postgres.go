package postgres

import (
	"context"
	"docs_flow/config"
	"docs_flow/internal/entities"
	//_ "docs_flow/internal/entities"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
	"go.uber.org/zap"
	"time"
	// protos "docs_flow/pkg/proto/gen/go"
)

type Repository struct {
	ctx context.Context
	log *zap.Logger
	cfg *config.ConfigModel
	DB  *pgxpool.Pool
}

func NewRepository(log *zap.Logger, cfg *config.ConfigModel, ctx context.Context) (*Repository, error) {
	return &Repository{
		ctx: ctx,
		log: log,
		cfg: cfg,
	}, nil
}

func (r *Repository) OnStart(_ context.Context) error {
	connectionUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		r.cfg.Postgres.Host,
		r.cfg.Postgres.Port,
		r.cfg.Postgres.User,
		r.cfg.Postgres.Password,
		r.cfg.Postgres.DBName,
		r.cfg.Postgres.SSLMode)
	pool, err := pgxpool.Connect(r.ctx, connectionUrl)
	if err != nil {
		return err
	}
	r.DB = pool
	return nil
}

func (r *Repository) OnStop(_ context.Context) error {
	r.DB.Close()
	return nil
}

const qGetUserID = `
select 
    id 
from 
    users 
`

func (r *Repository) GetUserIDbyPhoneOrEmail(ctx context.Context, phone, email string) (id int, err error) {
	err = r.DB.QueryRow(ctx, qGetUserID, phone, email).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

const qCreateUser = `
INSERT INTO users (id, name) VALUES ($1, $2)
`

//func (r *Repository) CreateUser(ctx context.Context, user *entities.User) error {
//	_, err := r.DB.Exec(ctx, qCreateUser, user.ID, user.Name, user.Email, user.Phone)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

const queryGetUser = `
SELECT 
    id, 
    name, 
    phone_number, 
    login, 
    password_hash, 
    role, 
    ava_path
FROM 
    users
WHERE 
    id = $1;

`

func (r *Repository) GetUser(ctx context.Context, userID int64) (*entities.User, error) {
	userDto := entities.UserDTO{}
	err := r.DB.QueryRow(ctx, queryGetUser, userID).Scan(
		&userDto.ID,
		&userDto.Name,
		&userDto.Login,
		&userDto.PasswordHash,
		&userDto.Role,
		&userDto.AvaPath,
	)
	if err != nil {
		return nil, err
	}
	return userDto.FromDTOConvert(), nil
}

const queryGetDocuments = `
SELECT 
    d.id, 
    d.creator_id, 
    d.status, 
    d.name, 
    d.file_path, 
    d.burn_date_time, 
    d.description, 
    d.signing_status
FROM 
    document d
INNER JOIN 
    users_documents ud 
ON 
    d.id = ud.document_id
WHERE 
    ($1 = 0 OR ud.user_id = $1) AND ($2 = 'all' OR d.signing_status = $2)
ORDER BY 
    d.burn_date_time DESC;
`

// GetDocuments
// Если userID == 0, то этот параметр не учитывается.
// Если signingStatus == "all", то этот параметр не учитывается.
// Всегда самые новые документы будут первыми в результатах.
func (r *Repository) GetDocuments(ctx context.Context, userID int64, signingStatus string) ([]*entities.Document, error) {
	rows, err := r.DB.Query(ctx, queryGetDocuments, userID, signingStatus)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Освобождаем ресурсы после завершения работы с rows

	var documents []*entities.Document
	for rows.Next() {
		documentDto := entities.DocumentDTO{}
		er := rows.Scan(
			&documentDto.ID,
			&documentDto.CreatorID,
			&documentDto.Status,
			&documentDto.Name,
			&documentDto.FilePath,
			&documentDto.BurnDateTime,
			&documentDto.Description,
			&documentDto.SigningStatus,
		)
		if er != nil {
			return nil, er
		}
		document := documentDto.FromDTOConvert()
		documents = append(documents, document)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return documents, nil
}

const queryGetUserByChatID = `
SELECT u.id,
       u.name,
       u.phone_number,
       u.login,
       u.role,
       u.ava_path
FROM chats c
         JOIN
     users u
     ON
         u.id = ANY (c.users)
WHERE c.id = $1
  AND u.id != $2;
`

func (r *Repository) GetUserByChatID(ctx context.Context, chatID, userID int64) (*entities.User, error) {
	row := r.DB.QueryRow(ctx, queryGetUserByChatID, chatID, userID)

	user := new(entities.User)
	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Login, &user.Role, &user.AvaPath)
	if err != nil {
		return nil, err
	}

	return user, nil

}

const queryGetChats = `
SELECT
    c.id,
    u.name AS name
FROM
    chats c
JOIN
    unnest(c.users) AS user_id ON user_id != c.users[1]
JOIN
    users u ON u.id = user_id
WHERE
    $1 = ANY(c.users);
`

func (r *Repository) GetUserChats(ctx context.Context, userID int64) ([]*entities.Chat, error) {
	rows, err := r.DB.Query(ctx, queryGetChats, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chats []*entities.Chat

	for rows.Next() {
		chatDto := &entities.ChatDTO{}

		err := rows.Scan(&chatDto.ID, &chatDto.Name)
		if err != nil {
			return nil, err
		}
		chats = append(chats, chatDto.FromDTOConvert())
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return chats, nil
}

const queryGetMessagesByChatID = `
SELECT id,
       sender_id,
       chat_id,
       message_text,
       burn_date_time,
       documents
FROM messages
WHERE chat_id = $1
ORDER BY messages.burn_date_time DESC;
	`

const queryGetDocument = `SELECT 
    id, 
    name, 
    file_path 
FROM document 
WHERE id = $1`

func (r *Repository) GetMessagesByChatID(ctx context.Context, chatID int64) ([]*entities.Message, error) {
	rows, err := r.DB.Query(ctx, queryGetMessagesByChatID, chatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*entities.Message
	for rows.Next() {
		messageDTO := entities.MessageDTO{}
		var documentsIDs []int64
		err := rows.Scan(
			&messageDTO.ID,
			&messageDTO.SenderID,
			&messageDTO.ChatID,
			&messageDTO.MessageText,
			&messageDTO.BurnDateTime,
			&documentsIDs,
		)
		if err != nil {
			return nil, err
		}
		message := messageDTO.FromDTOConvert()
		var documentsMessage []*entities.DocumentMessage
		for _, documentID := range documentsIDs {
			documentMessageDto := &entities.DocumentMessageDTO{}
			err = r.DB.QueryRow(ctx, queryGetDocument, documentID).Scan(
				&documentMessageDto.ID,
				&documentMessageDto.Name,
				&documentMessageDto.FilePath,
			)
			if err != nil {
				return nil, err
			}
			documentsMessage = append(documentsMessage, documentMessageDto.FromDTOConvert())
		}
		message.Documents = documentsMessage
		messages = append(messages, message)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return messages, nil
}

const queryCreateMessage = `
		INSERT INTO messages (
			sender_id,
			chat_id,
			message_text,
			documents,
			burn_date_time
		) 
		VALUES ($1, $2, $3, $4, $5) 
	`

func (r *Repository) CreateMessage(ctx context.Context, message *entities.Message) (bool, error) {
	messageDto := message.ConvertToDTO()
	var documentsIDs []int64
	for _, document := range messageDto.Documents {
		documentsIDs = append(documentsIDs, document.ID)
	}
	_, err := r.DB.Query(ctx, queryCreateMessage,
		messageDto.SenderID,
		messageDto.ChatID,
		messageDto.MessageText,
		pq.Int64Array(documentsIDs),
		messageDto.BurnDateTime,
	)

	if err != nil {
		return false, err
	}

	return true, nil
}

const queryGetDocumentsTemplates = `
SELECT 
    d.id, 
    d.creator_id, 
    d.status, 
    d.name, 
    d.file_path, 
    d.burn_date_time, 
    d.description, 
    d.signing_status
FROM 
    document d
WHERE 
    d.status = 'template';
`

func (r *Repository) GetDocumentsTemplates(ctx context.Context) ([]*entities.Document, error) {
	rows, err := r.DB.Query(ctx, queryGetDocumentsTemplates)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Освобождаем ресурсы после завершения работы с rows

	var documents []*entities.Document
	for rows.Next() {
		documentDto := entities.DocumentDTO{}
		er := rows.Scan(
			&documentDto.ID,
			&documentDto.CreatorID,
			&documentDto.Status,
			&documentDto.Name,
			&documentDto.FilePath,
			&documentDto.BurnDateTime,
			&documentDto.Description,
			&documentDto.SigningStatus,
		)
		if er != nil {
			return nil, er
		}
		document := documentDto.FromDTOConvert()
		documents = append(documents, document)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return documents, nil
}

const queryGetAllIDs = `SELECT
    id 
FROM 
    users 
WHERE 
    id != $1;
`

const queryCreateChat = `INSERT INTO chats (
   users
)
VALUES ($1);
`

func (r *Repository) CreateChatsUser(ctx context.Context, userID int64) (bool, error) {
	// Получаем всех пользователей, кроме текущего
	rows, err := r.DB.Query(ctx, queryGetAllIDs, userID)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	// Для каждого пользователя создаем чат с текущим пользователем
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return false, err
		}

		// Создаем новый чат для текущего пользователя и другого пользователя
		users := []int64{userID, id}
		_, err := r.DB.Query(ctx, queryCreateChat, pq.Array(users))
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

const queryCreateDocument = `INSERT INTO document (
    creator_id,
    status,
    name,
    file_path,
    burn_date_time,
    description,
    signing_status
) 
VALUES ($1, $2, $3, $4, $5, $6, $7);`

func (r *Repository) CreateDocument(ctx context.Context, document *entities.Document) (bool, error) {
	documentDto := document.ConvertToDTO()
	_, err := r.DB.Query(ctx, queryCreateDocument,
		documentDto.CreatorID,
		documentDto.Status,
		documentDto.Name,
		documentDto.FilePath,
		documentDto.BurnDateTime,
		documentDto.Description,
		documentDto.SigningStatus,
	)
	if err != nil {
		return false, err
	}
	return true, err
}

const queryCreateNotification = `INSERT INTO notifications (
   text, burn_date
   )
VALUES ($1, $2)
RETURNING id`

const queryCreateUserNotification = `INSERT INTO users_notifications 
    (user_id, notification_id)
VALUES ($1, $2)` // Для создания связи между пользователем и уведомлением

func (r *Repository) CreateNotification(ctx context.Context, notification *entities.Notification) (bool, error) {
	if notification.BurnDate == "" {
		notification.BurnDate = time.Now().Format(time.RFC3339) // Преобразуем текущее время в строку
	}
	notificationDto := notification.ConvertToDTO()
	var notificationID int64

	err := r.DB.QueryRow(ctx, queryCreateNotification,
		notificationDto.Text,
		notificationDto.BurnDate,
	).Scan(&notificationID)
	if err != nil {
		return false, err
	}

	rows, err := r.DB.Query(ctx, "SELECT id FROM users")
	if err != nil {
		return false, err
	}
	defer rows.Close()

	users := make([]int64, 0, 16) // Здесь будет список всех пользователей, для которых создадим связь

	for rows.Next() {
		var userID int64
		if err := rows.Scan(&userID); err != nil {
			return false, err
		}
		users = append(users, userID)
	}

	for _, userID := range users {
		_, err := r.DB.Exec(ctx, queryCreateUserNotification, userID, notificationID)
		if err != nil {
			return false, err
		}
	}

	return true, err
}

const queryDeleteNotification = `DELETE FROM users_notifications
WHERE user_id = $1 and notification_id = $1;
`

func (r *Repository) DeleteNotification(ctx context.Context, userID int64, notificationID int64) (bool, error) {
	// Выполняем запрос на удаление
	result, err := r.DB.Exec(ctx, queryDeleteNotification, userID, notificationID)
	if err != nil {
		return false, err
	}

	// Проверяем, был ли затронут хотя бы один ряд (удалена ли запись)
	rowsAffected := result.RowsAffected()
	if rowsAffected > 0 {
		return true, nil // Удаление успешно
	}

	return false, nil // Уведомление не найдено для удаления
}

const queryGetNotifications = `
	SELECT id, text, burn_date
	FROM notifications
	WHERE id IN (
		SELECT notification_id 
		FROM users_notifications 
		WHERE user_id = $1
	)
	ORDER BY burn_date DESC;
`

func (r *Repository) GetNotifications(ctx context.Context, userID int64) ([]*entities.Notification, error) {
	rows, err := r.DB.Query(ctx, queryGetNotifications, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []*entities.Notification
	for rows.Next() {
		notificationDto := entities.NotificationDTO{}
		err := rows.Scan(
			&notificationDto.ID,
			&notificationDto.Text,
			&notificationDto.BurnDate)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, notificationDto.FromDTOConvert())
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notifications, nil
}
