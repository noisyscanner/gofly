-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               10.4.12-MariaDB - mariadb.org binary distribution
-- Server OS:                    Win64
-- HeidiSQL Version:             10.2.0.5599
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- Dumping database structure for ivapi
CREATE DATABASE IF NOT EXISTS `ivapi_1` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `ivapi_1`;

-- Dumping structure for table ivapi.conjugations
CREATE TABLE IF NOT EXISTS `conjugations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `pronoun_id` int(10) unsigned NOT NULL,
  `tense_id` int(10) unsigned NOT NULL,
  `conjugation` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `normalisedConjugation` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `verb_id` int(10) unsigned NOT NULL DEFAULT 0,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `conjugations_verb_id_tense_id_pronoun_id_unique` (`verb_id`,`tense_id`,`pronoun_id`),
  KEY `conjugations_verb_id_index` (`verb_id`),
  KEY `conjugations_updated_at_index` (`updated_at`),
  KEY `conjugations_tense_id_foreign` (`tense_id`),
  KEY `conjugations_pronoun_id_foreign` (`pronoun_id`),
  CONSTRAINT `conjugations_pronoun_id_foreign` FOREIGN KEY (`pronoun_id`) REFERENCES `pronouns` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `conjugations_tense_id_foreign` FOREIGN KEY (`tense_id`) REFERENCES `tenses` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `conjugations_verb_id_foreign` FOREIGN KEY (`verb_id`) REFERENCES `verbs` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=80747 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table ivapi.languages
CREATE TABLE IF NOT EXISTS `languages` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `code` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `lang` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `locale` varchar(5) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `enabled` tinyint(1) NOT NULL DEFAULT 1,
  `hasReflexives` tinyint(1) NOT NULL DEFAULT 0,
  `hasHelpers` tinyint(1) NOT NULL DEFAULT 0,
  `minVersion` double NOT NULL DEFAULT 1,
  PRIMARY KEY (`id`),
  UNIQUE KEY `languages_code_lang_locale_unique` (`code`,`lang`,`locale`),
  KEY `languages_code_index` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table ivapi.migrations
CREATE TABLE IF NOT EXISTS `migrations` (
  `migration` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `batch` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table ivapi.password_resets
CREATE TABLE IF NOT EXISTS `password_resets` (
  `email` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `token` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  KEY `password_resets_email_index` (`email`),
  KEY `password_resets_token_index` (`token`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table ivapi.pronouns
CREATE TABLE IF NOT EXISTS `pronouns` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `lang_id` int(10) unsigned NOT NULL,
  `identifier` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `displayName` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `order` int(11) NOT NULL,
  `reflexive` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `pronouns_lang_id_foreign` (`lang_id`),
  CONSTRAINT `pronouns_lang_id_foreign` FOREIGN KEY (`lang_id`) REFERENCES `languages` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=45 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table ivapi.tenses
CREATE TABLE IF NOT EXISTS `tenses` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `lang_id` int(10) unsigned NOT NULL,
  `identifier` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `displayName` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `order` int(11) NOT NULL,
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `tenses_lang_id_foreign` (`lang_id`),
  CONSTRAINT `tenses_lang_id_foreign` FOREIGN KEY (`lang_id`) REFERENCES `languages` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=50 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table ivapi.users
CREATE TABLE IF NOT EXISTS `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `email` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `password` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `remember_token` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `isAdmin` tinyint(1) NOT NULL,
  `api_token` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `users_email_unique` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- Data exporting was unselected.

-- Dumping structure for table ivapi.verbs
CREATE TABLE IF NOT EXISTS `verbs` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `lang_id` int(10) unsigned NOT NULL,
  `infinitive` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `normalisedInfinitive` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `english` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `helperID` int(10) unsigned DEFAULT 0,
  `isHelper` tinyint(1) NOT NULL DEFAULT 0,
  `updated_at` timestamp NULL DEFAULT NULL,
  `isReflexive` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `verbs_lang_id_index` (`lang_id`),
  KEY `verbs_updated_at_index` (`updated_at`),
  KEY `verbs_helperid_foreign` (`helperID`),
  CONSTRAINT `verbs_helperid_foreign` FOREIGN KEY (`helperID`) REFERENCES `verbs` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
  CONSTRAINT `verbs_lang_id_foreign` FOREIGN KEY (`lang_id`) REFERENCES `languages` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=43635 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- Data exporting was unselected.

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;

