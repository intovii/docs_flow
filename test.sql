-- getUser
SELECT id,
       name,
       phone_number,
       login,
       password_hash,
       role,
       ava_path
FROM users
WHERE id = 1;

-- GetDocuments
SELECT d.id,
       d.creator_id,
       d.status,
       d.name,
       d.file_path,
       d.burn_date_time,
       d.description,
       d.signing_status,
       ud.id as user_id
FROM document d
         INNER JOIN
     users_documents ud
     ON
         d.id = ud.document_id
WHERE (0 = 0 OR ud.user_id = 2)
  AND ('signed' = 'all' OR d.signing_status = 'signed')
ORDER BY d.burn_date_time DESC;


-- GetUserChats
SELECT c.id,
       u.name AS name
FROM chats c
         JOIN
     unnest(c.users) AS user_id ON user_id != c.users[1]
         JOIN
     users u ON u.id = user_id
WHERE 1 = ANY (c.users);

--GetUserByChatID
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
WHERE c.id = 1
  AND u.id != 2;

-- GetMessagesByChatID
SELECT id,
       sender_id,
       chat_id,
       message_text,
       burn_date_time,
       documents
FROM messages
WHERE chat_id = 1
ORDER BY messages.burn_date_time DESC;
SELECT id, name, file_path
FROM document
WHERE id = 1;

-- CreateMessage
INSERT INTO messages (sender_id,
                      chat_id,
                      message_text,
                      documents,
                      burn_date_time)
VALUES (1, 2, '$3', ARRAY [1, 2], '2024-12-22 10:07:00');

-- GetDocumentsTemplates
SELECT d.id,
       d.creator_id,
       d.status,
       d.name,
       d.file_path,
       d.burn_date_time,
       d.description,
       d.signing_status
FROM document d
WHERE d.status = 'template';

--CreateChatsUser
SELECT id
FROM users
WHERE id != 4;

INSERT INTO chats (users)
VALUES (ARRAY [3, 4]);

--CreateDocument
INSERT INTO document (creator_id,
                      status,
                      name,
                      file_path,
                      burn_date_time,
                      description,
                      signing_status)
VALUES (1, 'template', '$3', '$4', '2024-12-22 10:07:00', '$6', 'signed');

--CreateNotification
INSERT INTO notifications (text, burn_date)
VALUES ('$1', '2024-12-22 10:07:00')
RETURNING id;
INSERT INTO users_notifications
    (user_id, notification_id)
VALUES (1, 2);

SELECT *
FROM users_notifications;

--DeleteNotification
DELETE
FROM users_notifications
WHERE user_id = 1
  and notification_id = 1;

--GetNotifications
SELECT id, text, burn_date
FROM notifications
WHERE id IN (SELECT notification_id
             FROM users_notifications
             WHERE user_id = 1)
ORDER BY burn_date DESC;
