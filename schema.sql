-- SQLite Schema for Cloudflare R1
PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS `users` (
  `id` INTEGER PRIMARY KEY AUTOINCREMENT,
  `name` varchar(255) NOT NULL,
  `last_name` varchar(255) NOT NULL,
  `email` varchar(191) NOT NULL,
  `password` varchar(255) NOT NULL,
  `remember_token` varchar(100) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS `users_email_unique` ON `users` (`email`);

CREATE TABLE IF NOT EXISTS `blood_pressures` (
  `id` INTEGER PRIMARY KEY AUTOINCREMENT,
  `user_id` integer NOT NULL,
  `systolic` integer NOT NULL,
  `diastolic` integer NOT NULL,
  `pulse` integer NOT NULL,
  `reading_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS `blood_pressures_user_id_index` ON `blood_pressures` (`user_id`);

CREATE TABLE IF NOT EXISTS `weights` (
  `id` INTEGER PRIMARY KEY AUTOINCREMENT,
  `user_id` integer NOT NULL,
  `weight` double NOT NULL,
  `entered_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS `weights_user_id_index` ON `weights` (`user_id`);
CREATE INDEX IF NOT EXISTS `weights_weight_index` ON `weights` (`weight`);

CREATE TABLE IF NOT EXISTS `sessions` (
  `id` TEXT PRIMARY KEY,
  `user_id` INTEGER,
  `ip_address` TEXT,
  `user_agent` TEXT,
  `payload` TEXT NOT NULL,
  `last_activity` INTEGER NOT NULL,
  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS `sessions_user_id_index` ON `sessions` (`user_id`);
CREATE INDEX IF NOT EXISTS `sessions_last_activity_index` ON `sessions` (`last_activity`);

CREATE TABLE IF NOT EXISTS `migrations` (
  `id` INTEGER PRIMARY KEY AUTOINCREMENT,
  `migration` varchar(191) NOT NULL,
  `batch` integer NOT NULL
);

CREATE TABLE IF NOT EXISTS `password_resets` (
  `email` varchar(191) NOT NULL,
  `token` varchar(191) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL
);

CREATE INDEX IF NOT EXISTS `password_resets_email_index` ON `password_resets` (`email`);
