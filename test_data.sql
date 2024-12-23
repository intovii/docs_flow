-- Создаем пользователей
INSERT INTO users (name, phone_number, login, password_hash, role, ava_path) VALUES
('Alice', '1234567890', 'alice_login', 'password_hash_1', 'user', '/avatars/alice.jpg'),
('Bob', '9876543210', 'bob_login', 'password_hash_2', 'user', '/avatars/bob.jpg'),
('Charlie', '1122334455', 'charlie_login', 'password_hash_3', 'user', '/avatars/charlie.jpg');

-- Создаем документы
INSERT INTO document (creator_id, status, name, file_path, burn_date_time, description, signing_status) VALUES
(1, 'draft', 'Document 1', '/files/doc1.pdf', '2024-12-22 10:00:00', 'Description of document 1', 'signed'),
(1, 'active', 'Document 2', '/files/doc2.pdf', '2024-12-22 12:00:00', 'Description of document 2', 'unsigned'),
(2, 'draft', 'Document 3', '/files/doc3.pdf', '2024-12-23 10:00:00', 'Description of document 3', 'signed');

-- Создаем связи пользователей с документами
INSERT INTO users_documents (user_id, document_id, task_name, can_edit) VALUES
(1, 1, 'Task 1', TRUE),
(2, 1, 'Task 2', FALSE),
(1, 2, 'Task 3', TRUE),
(3, 3, 'Task 4', FALSE);

-- Создаем чаты
INSERT INTO chats (users) VALUES
(ARRAY[1, 2]),  -- Chat between Alice and Bob
(ARRAY[1, 3]),  -- Chat between Alice and Charlie
(ARRAY[2, 3]);  -- Chat between Bob and Charlie

-- Создаем сообщения в чатах
INSERT INTO messages (sender_id, chat_id, message_text, documents, burn_date_time) VALUES
-- (1, 1, 'Hello Bob', ARRAY[1], '2024-12-22 10:05:00'),
(1, 1, 'How are you', ARRAY[1, 2, 3], '2024-12-22 10:07:00');
-- (2, 1, 'Hi Alice', ARRAY[1], '2024-12-22 10:06:00'),
-- (1, 2, 'Hello Charlie', ARRAY[2], '2024-12-22 12:05:00'),
-- (3, 2, 'Hi Alice', ARRAY[2], '2024-12-22 12:06:00'),
-- (2, 3, 'Hello Bob and Charlie', ARRAY[3], '2024-12-23 10:05:00');

-- Создаем уведомления
INSERT INTO notifications (text, burn_date) VALUES
('You have a new message', '2024-12-22 10:00:00'),
('Document 1 has been signed', '2024-12-22 12:00:00'),
('Document 2 is active', '2024-12-23 10:00:00');

-- Связываем уведомления с пользователями
INSERT INTO users_notifications (user_id, notification_id) VALUES
(1, 1),
(2, 2),
(3, 3);
