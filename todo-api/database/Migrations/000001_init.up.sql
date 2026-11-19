-- SQL dump generated using DBML (dbml.dbdiagram.io)
-- Database: MySQL
-- Generated at: 2026-10-30T13:02:48.615Z

CREATE TABLE `users` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `email` varchar(255) UNIQUE NOT NULL,
  `created_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `todos` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `description` text,
  `completed` boolean NOT NULL DEFAULT false,
  `is_archived` boolean NOT NULL DEFAULT false,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `todo_audit_logs` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT,
  `todo_id` bigint NOT NULL COMMENT 'Original TODO ID retained after deletion; intentionally no foreign key to todos',
  `user_id` bigint NOT NULL,
  `action` ENUM ('CREATED', 'UPDATED', 'DELETED', 'ARCHIVED', 'RESTORED', 'STATUS_CHANGED') NOT NULL,
  `created_at` datetime NOT NULL DEFAULT (now())
);

CREATE INDEX `idx_todos_archived_created` ON `todos` (`is_archived`, `created_at`);

CREATE INDEX `idx_todos_archived_completed` ON `todos` (`is_archived`, `completed`);

CREATE INDEX `idx_audit_todo_created` ON `todo_audit_logs` (`todo_id`, `created_at`);

CREATE INDEX `idx_audit_user` ON `todo_audit_logs` (`user_id`);

ALTER TABLE `todo_audit_logs` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE;
