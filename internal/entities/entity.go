package entities

import (
	"database/sql"
	protos "docs_flow/pkg/proto/gen/go"
	"time"
)

//type User struct {
//	ID    int
//	Mail  string
//	Phone string
//	Role  string
//	Token string
//}

func NewNullString(s string) *sql.NullString {
	if len(s) == 0 {
		return &sql.NullString{}
	} else {
		return &sql.NullString{
			String: s,
			Valid:  true,
		}
	}
}

type User struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	PhoneNumber  string `json:"phone_number"`
	Login        string `json:"login"`
	PasswordHash string `json:"password_hash"`
	Role         string `json:"role"`
	AvaPath      string `json:"ava_path"`
}

type UserDTO struct {
	ID           int64          `json:"id" db:"id"`
	Name         sql.NullString `json:"name" db:"name"`
	PhoneNumber  sql.NullString `json:"phone_number" db:"phone_number"`
	Login        sql.NullString `json:"login" db:"login"`
	PasswordHash sql.NullString `json:"password-hash" db:"password-hash"`
	Role         sql.NullString `json:"role" db:"role"`
	AvaPath      sql.NullString `json:"ava_path" db:"ava_path"`
}

func (c *User) ConvertToDTO() *UserDTO {
	dto := new(UserDTO)
	dto.ID = c.ID
	dto.Name = *NewNullString(c.Name)
	dto.PhoneNumber = *NewNullString(c.PhoneNumber)
	dto.Login = *NewNullString(c.Login)
	dto.PasswordHash = *NewNullString(c.PasswordHash)
	dto.Role = *NewNullString(c.Role)
	dto.AvaPath = *NewNullString(c.AvaPath)
	return dto
}

func (dto *UserDTO) FromDTOConvert() *User {
	c := new(User)
	c.ID = dto.ID
	c.PhoneNumber = dto.PhoneNumber.String
	c.Name = dto.Name.String
	c.Login = dto.Login.String
	c.PasswordHash = dto.PasswordHash.String
	c.Role = dto.Role.String
	c.AvaPath = dto.AvaPath.String

	return c
}

type Document struct {
	ID            int64
	CreatorID     int64
	Status        string
	Name          string
	FilePath      string
	BurnDateTime  string
	Description   string
	SigningStatus string
}

type DocumentDTO struct {
	ID            int64          `json:"id" db:"id"`
	CreatorID     int64          `json:"creator_id" db:"creator_id"`
	Status        sql.NullString `json:"status" db:"status"`
	Name          sql.NullString `json:"name" db:"name"`
	FilePath      sql.NullString `json:"file_path" db:"file_path"`
	BurnDateTime  sql.NullString `json:"burn_date_time" db:"burn_date_time"`
	Description   sql.NullString `json:"description" db:"description"`
	SigningStatus sql.NullString `json:"signing_status" db:"signing_status"`
}

func (c *Document) ConvertToDTO() *DocumentDTO {
	dto := new(DocumentDTO)
	dto.ID = c.ID
	dto.CreatorID = c.CreatorID
	dto.Status = *NewNullString(c.Status)
	dto.Name = *NewNullString(c.Name)
	dto.FilePath = *NewNullString(c.FilePath)
	dto.BurnDateTime = *NewNullString(c.BurnDateTime)
	dto.Description = *NewNullString(c.Description)
	dto.SigningStatus = *NewNullString(c.SigningStatus)
	return dto
}
func ConvertDocumentsToProto(documents []*Document) []*protos.Document {
	var protoDocuments []*protos.Document
	for _, doc := range documents {
		protoDoc := &protos.Document{
			ID:            doc.ID,
			CreatorID:     doc.CreatorID,
			Status:        doc.Status,
			Name:          doc.Name,
			FilePath:      doc.FilePath,
			BurnDateTime:  doc.BurnDateTime,
			Description:   doc.Description,
			SigningStatus: doc.SigningStatus,
		}
		protoDocuments = append(protoDocuments, protoDoc)
	}
	return protoDocuments
}
func (dto *DocumentDTO) FromDTOConvert() *Document {
	c := new(Document)
	c.ID = dto.ID
	c.CreatorID = dto.CreatorID
	c.Status = dto.Status.String
	c.Name = dto.Name.String
	c.FilePath = dto.FilePath.String
	c.BurnDateTime = dto.BurnDateTime.String
	c.Description = dto.Description.String
	c.SigningStatus = dto.SigningStatus.String
	return c
}

type Message struct {
	ID          int64
	SenderID    int64
	ChatID      int64
	MessageText string
	//DocumentsIDs []int64
	Documents    []*DocumentMessage
	BurnDateTime string
}

type MessageDTO struct {
	ID           int64          `json:"id" db:"id"`
	SenderID     int64          `json:"sender_id" db:"sender_id"`
	ChatID       int64          `json:"chat_id" db:"chat_id"`
	MessageText  sql.NullString `json:"message_text" db:"message_text"`
	Documents    []*DocumentMessage
	BurnDateTime sql.NullTime `json:"burn_date_time" db:"burn_date_time"`
}

func (m *Message) ConvertToDTO() *MessageDTO {
	dto := new(MessageDTO)
	dto.ID = m.ID
	dto.SenderID = m.SenderID
	dto.ChatID = m.ChatID
	dto.MessageText = *NewNullString(m.MessageText) // предполагается, что NewNullString работает правильно
	dto.BurnDateTime = sql.NullTime{Valid: m.BurnDateTime != "", Time: time.Time{}}
	//dto.DocumentsIDs = pq.Int64Array(m.DocumentsIDs)
	dto.Documents = m.Documents

	// Если BurnDateTime имеет значение (не пустая строка), его нужно привести к времени
	if m.BurnDateTime != "" {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", m.BurnDateTime) // Используем формат времени, который вам нужен
		if err == nil {
			dto.BurnDateTime = sql.NullTime{Valid: true, Time: parsedTime}
		}
	}

	return dto
}

func (dto *MessageDTO) FromDTOConvert() *Message {
	m := new(Message)
	m.ID = dto.ID
	m.SenderID = dto.SenderID
	m.ChatID = dto.ChatID
	m.MessageText = dto.MessageText.String
	//m.DocumentsIDs = []int64(dto.DocumentsIDs)
	m.Documents = dto.Documents

	// Преобразуем поле BurnDateTime
	if dto.BurnDateTime.Valid {
		m.BurnDateTime = dto.BurnDateTime.Time.Format("2006-01-02 15:04:05") // Форматируем время в строку
	} else {
		m.BurnDateTime = "" // Если BurnDateTime невалидно, оставляем пустую строку
	}

	return m
}

type Notification struct {
	ID       int64
	Text     string
	BurnDate string
}

type NotificationDTO struct {
	ID       int64          `json:"id" db:"id"`
	Text     sql.NullString `json:"text" db:"text"`
	BurnDate sql.NullTime   `json:"burn_date" db:"burn_date"`
}

func (n *Notification) ConvertToDTO() *NotificationDTO {
	dto := new(NotificationDTO)
	dto.ID = n.ID
	dto.Text = *NewNullString(n.Text)
	// Если BurnDateTime имеет значение (не пустая строка), его нужно привести к времени
	if n.BurnDate != "" {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", n.BurnDate) // Используем формат времени, который вам нужен
		if err == nil {
			dto.BurnDate = sql.NullTime{Valid: true, Time: parsedTime}
		}
	}

	return dto
}

func (dto *NotificationDTO) FromDTOConvert() *Notification {
	n := new(Notification)
	n.ID = dto.ID
	n.Text = dto.Text.String
	// Преобразуем поле BurnDateTime
	if dto.BurnDate.Valid {
		n.BurnDate = dto.BurnDate.Time.Format("2006-01-02 15:04:05") // Форматируем время в строку
	} else {
		n.BurnDate = "" // Если BurnDateTime невалидно, оставляем пустую строку
	}
	return n
}

type Chat struct {
	ID   int64
	Name string
}

type ChatDTO struct {
	ID   int64          `json:"id" db:"id"`
	Name sql.NullString `json:"name" db:"name"`
}

func (c *Chat) ConvertToDTO() *ChatDTO {
	dto := new(ChatDTO)
	dto.ID = c.ID
	dto.Name = *NewNullString(c.Name)
	//return &ChatDTO{
	//	ID: c.ID,
	//	Name: c.Name,
	//}
	return dto
}

func (dto *ChatDTO) FromDTOConvert() *Chat {
	c := new(Chat)
	c.ID = dto.ID
	c.Name = dto.Name.String
	return c
}

type DocumentMessage struct {
	ID       int64
	Name     string
	FilePath string
}
type DocumentMessageDTO struct {
	ID       int64          `json:"id" db:"id"`
	Name     sql.NullString `json:"Name" db:"Name"`
	FilePath sql.NullString `json:"file_path" db:"file_path"`
}

func (c *DocumentMessage) ConvertToDTO() *DocumentMessageDTO {
	dto := new(DocumentMessageDTO)
	dto.ID = c.ID
	dto.Name = *NewNullString(c.Name)
	dto.FilePath = *NewNullString(c.FilePath)

	return dto
}

func (dto *DocumentMessageDTO) FromDTOConvert() *DocumentMessage {
	c := new(DocumentMessage)
	c.ID = dto.ID
	c.Name = dto.Name.String
	c.FilePath = dto.FilePath.String
	return c
}
