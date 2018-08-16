-- MySQL dump 10.13  Distrib 8.0.11, for osx10.13 (x86_64)
--
-- Host: 127.0.0.1    Database: rotation_shifts
-- ------------------------------------------------------
-- Server version	5.7.23

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
 SET NAMES utf8mb4 ;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `overrides`
--

DROP TABLE IF EXISTS `overrides`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `overrides` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `rotation_id` int(11) unsigned NOT NULL,
  `date` date NOT NULL,
  `user_id` int(11) unsigned NOT NULL,
  `0000` tinyint(1) NOT NULL DEFAULT '0',
  `0100` tinyint(1) NOT NULL DEFAULT '0',
  `0200` tinyint(1) NOT NULL DEFAULT '0',
  `0300` tinyint(1) NOT NULL DEFAULT '0',
  `0400` tinyint(1) NOT NULL DEFAULT '0',
  `0500` tinyint(1) NOT NULL DEFAULT '0',
  `0600` tinyint(1) NOT NULL DEFAULT '0',
  `0700` tinyint(1) NOT NULL DEFAULT '0',
  `0800` tinyint(1) NOT NULL DEFAULT '0',
  `0900` tinyint(1) NOT NULL DEFAULT '0',
  `1000` tinyint(1) NOT NULL DEFAULT '0',
  `1100` tinyint(1) NOT NULL DEFAULT '0',
  `1200` tinyint(1) NOT NULL DEFAULT '0',
  `1300` tinyint(1) NOT NULL DEFAULT '0',
  `1400` tinyint(1) NOT NULL DEFAULT '0',
  `1500` tinyint(1) NOT NULL DEFAULT '0',
  `1600` tinyint(1) NOT NULL DEFAULT '0',
  `1700` tinyint(1) NOT NULL DEFAULT '0',
  `1800` tinyint(1) NOT NULL DEFAULT '0',
  `1900` tinyint(1) NOT NULL DEFAULT '0',
  `2000` tinyint(1) NOT NULL DEFAULT '0',
  `2100` tinyint(1) NOT NULL DEFAULT '0',
  `2200` tinyint(1) NOT NULL DEFAULT '0',
  `2300` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `rotation_id_date_0000` (`rotation_id`,`date`,`0000`),
  KEY `rotation_id_date_0100` (`rotation_id`,`date`,`0100`),
  KEY `rotation_id_date_0200` (`rotation_id`,`date`,`0200`),
  KEY `rotation_id_date_0300` (`rotation_id`,`date`,`0300`),
  KEY `rotation_id_date_0400` (`rotation_id`,`date`,`0400`),
  KEY `rotation_id_date_0500` (`rotation_id`,`date`,`0500`),
  KEY `rotation_id_date_0600` (`rotation_id`,`date`,`0600`),
  KEY `rotation_id_date_0700` (`rotation_id`,`date`,`0700`),
  KEY `rotation_id_date_0800` (`rotation_id`,`date`,`0800`),
  KEY `rotation_id_date_0900` (`rotation_id`,`date`,`0900`),
  KEY `rotation_id_date_1000` (`rotation_id`,`date`,`1000`),
  KEY `rotation_id_date_1100` (`rotation_id`,`date`,`1100`),
  KEY `rotation_id_date_1200` (`rotation_id`,`date`,`1200`),
  KEY `rotation_id_date_1300` (`rotation_id`,`date`,`1300`),
  KEY `rotation_id_date_1400` (`rotation_id`,`date`,`1400`),
  KEY `rotation_id_date_1500` (`rotation_id`,`date`,`1500`),
  KEY `rotation_id_date_1600` (`rotation_id`,`date`,`1600`),
  KEY `rotation_id_date_1700` (`rotation_id`,`date`,`1700`),
  KEY `rotation_id_date_1800` (`rotation_id`,`date`,`1800`),
  KEY `rotation_id_date_1900` (`rotation_id`,`date`,`1900`),
  KEY `rotation_id_date_2000` (`rotation_id`,`date`,`2000`),
  KEY `rotation_id_date_2100` (`rotation_id`,`date`,`2100`),
  KEY `rotation_id_date_2200` (`rotation_id`,`date`,`2200`),
  KEY `rotation_id_date_2300` (`rotation_id`,`date`,`2300`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `overrides_ibfk_2` FOREIGN KEY (`rotation_id`) REFERENCES `rotations` (`id`),
  CONSTRAINT `overrides_ibfk_3` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `rotation_details`
--

DROP TABLE IF EXISTS `rotation_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `rotation_details` (
  `rotation_id` int(11) unsigned NOT NULL,
  `order_id` int(11) unsigned NOT NULL,
  `shift_id` int(11) unsigned NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`rotation_id`,`order_id`,`shift_id`),
  KEY `rotation_id` (`rotation_id`),
  KEY `shift_id` (`shift_id`),
  CONSTRAINT `rotation_details_ibfk_1` FOREIGN KEY (`shift_id`) REFERENCES `shifts` (`id`),
  CONSTRAINT `rotation_details_ibfk_2` FOREIGN KEY (`rotation_id`) REFERENCES `rotations` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `rotations`
--

DROP TABLE IF EXISTS `rotations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `rotations` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `start_date` date NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `shifts`
--

DROP TABLE IF EXISTS `shifts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `shifts` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `user_id` int(11) unsigned NOT NULL,
  `monday_0000` tinyint(1) NOT NULL DEFAULT '0',
  `monday_0100` tinyint(1) NOT NULL DEFAULT '0',
  `monday_0200` tinyint(1) NOT NULL DEFAULT '0',
  `monday_0300` tinyint(1) NOT NULL DEFAULT '0',
  `monday_0400` tinyint(1) NOT NULL DEFAULT '0',
  `monday_0500` tinyint(1) NOT NULL DEFAULT '0',
  `monday_0600` tinyint(1) NOT NULL DEFAULT '0',
  `monday_0700` tinyint(1) NOT NULL DEFAULT '0',
  `monday_0800` tinyint(1) NOT NULL DEFAULT '0',
  `monday_0900` tinyint(1) NOT NULL DEFAULT '0',
  `monday_1000` tinyint(1) NOT NULL DEFAULT '0',
  `monday_1100` tinyint(1) NOT NULL DEFAULT '0',
  `monday_1200` tinyint(1) NOT NULL DEFAULT '0',
  `monday_1300` tinyint(1) NOT NULL DEFAULT '0',
  `monday_1400` tinyint(1) NOT NULL DEFAULT '0',
  `monday_1500` tinyint(1) NOT NULL DEFAULT '0',
  `monday_1600` tinyint(1) NOT NULL DEFAULT '0',
  `monday_1700` tinyint(1) NOT NULL DEFAULT '0',
  `monday_1800` tinyint(1) NOT NULL DEFAULT '0',
  `monday_1900` tinyint(1) NOT NULL DEFAULT '0',
  `monday_2000` tinyint(1) NOT NULL DEFAULT '0',
  `monday_2100` tinyint(1) NOT NULL DEFAULT '0',
  `monday_2200` tinyint(1) NOT NULL DEFAULT '0',
  `monday_2300` tinyint(1) NOT NULL DEFAULT '0',
  `tuesday_0000` tinyint(1) NOT NULL DEFAULT '0',
  `tuesday_0100` tinyint(1) NOT NULL DEFAULT '0',
  `tuesday_0200` tinyint(1) NOT NULL DEFAULT '0',
  `tuesday_0300` tinyint(1) NOT NULL DEFAULT '0',
  `tuesday_0400` tinyint(1) NOT NULL DEFAULT '0',
  `tuesday_0500` tinyint(1) NOT NULL DEFAULT '0',
  `tuesday_0600` tinyint(1) NOT NULL DEFAULT '0',
  `tuesday_0700` tinyint(1) NOT NULL DEFAULT '0',
  `tuesday_0800` tinyint(1) NOT NULL DEFAULT '0',
  `tuesday_0900` tinyint(1) NOT NULL DEFAULT '0',
  `tuesday_1000` tinyint(1) NOT NULL DEFAULT '0',
  `tuesday_1100` tinyint(1) NOT NULL DEFAULT '0',
  `tuesday_1200` tinyint(1) NOT NULL DEFAULT '0',
  `tuesday_1300` tinyint(1) NOT NULL DEFAULT '0',
  `tuesday_1400` tinyint(1) NOT NULL DEFAULT '0',
  `tuesday_1500` tinyint(1) NOT NULL DEFAULT '0',
  `tuesday_1600` tinyint(1) NOT NULL DEFAULT '0',
  `tuesday_1700` tinyint(1) NOT NULL DEFAULT '0',
  `tuesday_1800` tinyint(1) NOT NULL DEFAULT '0',
  `tuesday_1900` tinyint(1) NOT NULL DEFAULT '0',
  `tuesday_2000` tinyint(1) NOT NULL DEFAULT '0',
  `tuesday_2100` tinyint(1) NOT NULL DEFAULT '0',
  `tuesday_2200` tinyint(1) NOT NULL DEFAULT '0',
  `tuesday_2300` tinyint(1) NOT NULL DEFAULT '0',
  `wednesday_0000` tinyint(1) NOT NULL DEFAULT '0',
  `wednesday_0100` tinyint(1) NOT NULL DEFAULT '0',
  `wednesday_0200` tinyint(1) NOT NULL DEFAULT '0',
  `wednesday_0300` tinyint(1) NOT NULL DEFAULT '0',
  `wednesday_0400` tinyint(1) NOT NULL DEFAULT '0',
  `wednesday_0500` tinyint(1) NOT NULL DEFAULT '0',
  `wednesday_0600` tinyint(1) NOT NULL DEFAULT '0',
  `wednesday_0700` tinyint(1) NOT NULL DEFAULT '0',
  `wednesday_0800` tinyint(1) NOT NULL DEFAULT '0',
  `wednesday_0900` tinyint(1) NOT NULL DEFAULT '0',
  `wednesday_1000` tinyint(1) NOT NULL DEFAULT '0',
  `wednesday_1100` tinyint(1) NOT NULL DEFAULT '0',
  `wednesday_1200` tinyint(1) NOT NULL DEFAULT '0',
  `wednesday_1300` tinyint(1) NOT NULL DEFAULT '0',
  `wednesday_1400` tinyint(1) NOT NULL DEFAULT '0',
  `wednesday_1500` tinyint(1) NOT NULL DEFAULT '0',
  `wednesday_1600` tinyint(1) NOT NULL DEFAULT '0',
  `wednesday_1700` tinyint(1) NOT NULL DEFAULT '0',
  `wednesday_1800` tinyint(1) NOT NULL DEFAULT '0',
  `wednesday_1900` tinyint(1) NOT NULL DEFAULT '0',
  `wednesday_2000` tinyint(1) NOT NULL DEFAULT '0',
  `wednesday_2100` tinyint(1) NOT NULL DEFAULT '0',
  `wednesday_2200` tinyint(1) NOT NULL DEFAULT '0',
  `wednesday_2300` tinyint(1) NOT NULL DEFAULT '0',
  `thursday_0000` tinyint(1) NOT NULL DEFAULT '0',
  `thursday_0100` tinyint(1) NOT NULL DEFAULT '0',
  `thursday_0200` tinyint(1) NOT NULL DEFAULT '0',
  `thursday_0300` tinyint(1) NOT NULL DEFAULT '0',
  `thursday_0400` tinyint(1) NOT NULL DEFAULT '0',
  `thursday_0500` tinyint(1) NOT NULL DEFAULT '0',
  `thursday_0600` tinyint(1) NOT NULL DEFAULT '0',
  `thursday_0700` tinyint(1) NOT NULL DEFAULT '0',
  `thursday_0800` tinyint(1) NOT NULL DEFAULT '0',
  `thursday_0900` tinyint(1) NOT NULL DEFAULT '0',
  `thursday_1000` tinyint(1) NOT NULL DEFAULT '0',
  `thursday_1100` tinyint(1) NOT NULL DEFAULT '0',
  `thursday_1200` tinyint(1) NOT NULL DEFAULT '0',
  `thursday_1300` tinyint(1) NOT NULL DEFAULT '0',
  `thursday_1400` tinyint(1) NOT NULL DEFAULT '0',
  `thursday_1500` tinyint(1) NOT NULL DEFAULT '0',
  `thursday_1600` tinyint(1) NOT NULL DEFAULT '0',
  `thursday_1700` tinyint(1) NOT NULL DEFAULT '0',
  `thursday_1800` tinyint(1) NOT NULL DEFAULT '0',
  `thursday_1900` tinyint(1) NOT NULL DEFAULT '0',
  `thursday_2000` tinyint(1) NOT NULL DEFAULT '0',
  `thursday_2100` tinyint(1) NOT NULL DEFAULT '0',
  `thursday_2200` tinyint(1) NOT NULL DEFAULT '0',
  `thursday_2300` tinyint(1) NOT NULL DEFAULT '0',
  `friday_0000` tinyint(1) NOT NULL DEFAULT '0',
  `friday_0100` tinyint(1) NOT NULL DEFAULT '0',
  `friday_0200` tinyint(1) NOT NULL DEFAULT '0',
  `friday_0300` tinyint(1) NOT NULL DEFAULT '0',
  `friday_0400` tinyint(1) NOT NULL DEFAULT '0',
  `friday_0500` tinyint(1) NOT NULL DEFAULT '0',
  `friday_0600` tinyint(1) NOT NULL DEFAULT '0',
  `friday_0700` tinyint(1) NOT NULL DEFAULT '0',
  `friday_0800` tinyint(1) NOT NULL DEFAULT '0',
  `friday_0900` tinyint(1) NOT NULL DEFAULT '0',
  `friday_1000` tinyint(1) NOT NULL DEFAULT '0',
  `friday_1100` tinyint(1) NOT NULL DEFAULT '0',
  `friday_1200` tinyint(1) NOT NULL DEFAULT '0',
  `friday_1300` tinyint(1) NOT NULL DEFAULT '0',
  `friday_1400` tinyint(1) NOT NULL DEFAULT '0',
  `friday_1500` tinyint(1) NOT NULL DEFAULT '0',
  `friday_1600` tinyint(1) NOT NULL DEFAULT '0',
  `friday_1700` tinyint(1) NOT NULL DEFAULT '0',
  `friday_1800` tinyint(1) NOT NULL DEFAULT '0',
  `friday_1900` tinyint(1) NOT NULL DEFAULT '0',
  `friday_2000` tinyint(1) NOT NULL DEFAULT '0',
  `friday_2100` tinyint(1) NOT NULL DEFAULT '0',
  `friday_2200` tinyint(1) NOT NULL DEFAULT '0',
  `friday_2300` tinyint(1) NOT NULL DEFAULT '0',
  `saturday_0000` tinyint(1) NOT NULL DEFAULT '0',
  `saturday_0100` tinyint(1) NOT NULL DEFAULT '0',
  `saturday_0200` tinyint(1) NOT NULL DEFAULT '0',
  `saturday_0300` tinyint(1) NOT NULL DEFAULT '0',
  `saturday_0400` tinyint(1) NOT NULL DEFAULT '0',
  `saturday_0500` tinyint(1) NOT NULL DEFAULT '0',
  `saturday_0600` tinyint(1) NOT NULL DEFAULT '0',
  `saturday_0700` tinyint(1) NOT NULL DEFAULT '0',
  `saturday_0800` tinyint(1) NOT NULL DEFAULT '0',
  `saturday_0900` tinyint(1) NOT NULL DEFAULT '0',
  `saturday_1000` tinyint(1) NOT NULL DEFAULT '0',
  `saturday_1100` tinyint(1) NOT NULL DEFAULT '0',
  `saturday_1200` tinyint(1) NOT NULL DEFAULT '0',
  `saturday_1300` tinyint(1) NOT NULL DEFAULT '0',
  `saturday_1400` tinyint(1) NOT NULL DEFAULT '0',
  `saturday_1500` tinyint(1) NOT NULL DEFAULT '0',
  `saturday_1600` tinyint(1) NOT NULL DEFAULT '0',
  `saturday_1700` tinyint(1) NOT NULL DEFAULT '0',
  `saturday_1800` tinyint(1) NOT NULL DEFAULT '0',
  `saturday_1900` tinyint(1) NOT NULL DEFAULT '0',
  `saturday_2000` tinyint(1) NOT NULL DEFAULT '0',
  `saturday_2100` tinyint(1) NOT NULL DEFAULT '0',
  `saturday_2200` tinyint(1) NOT NULL DEFAULT '0',
  `saturday_2300` tinyint(1) NOT NULL DEFAULT '0',
  `sunday_0000` tinyint(1) NOT NULL DEFAULT '0',
  `sunday_0100` tinyint(1) NOT NULL DEFAULT '0',
  `sunday_0200` tinyint(1) NOT NULL DEFAULT '0',
  `sunday_0300` tinyint(1) NOT NULL DEFAULT '0',
  `sunday_0400` tinyint(1) NOT NULL DEFAULT '0',
  `sunday_0500` tinyint(1) NOT NULL DEFAULT '0',
  `sunday_0600` tinyint(1) NOT NULL DEFAULT '0',
  `sunday_0700` tinyint(1) NOT NULL DEFAULT '0',
  `sunday_0800` tinyint(1) NOT NULL DEFAULT '0',
  `sunday_0900` tinyint(1) NOT NULL DEFAULT '0',
  `sunday_1000` tinyint(1) NOT NULL DEFAULT '0',
  `sunday_1100` tinyint(1) NOT NULL DEFAULT '0',
  `sunday_1200` tinyint(1) NOT NULL DEFAULT '0',
  `sunday_1300` tinyint(1) NOT NULL DEFAULT '0',
  `sunday_1400` tinyint(1) NOT NULL DEFAULT '0',
  `sunday_1500` tinyint(1) NOT NULL DEFAULT '0',
  `sunday_1600` tinyint(1) NOT NULL DEFAULT '0',
  `sunday_1700` tinyint(1) NOT NULL DEFAULT '0',
  `sunday_1800` tinyint(1) NOT NULL DEFAULT '0',
  `sunday_1900` tinyint(1) NOT NULL DEFAULT '0',
  `sunday_2000` tinyint(1) NOT NULL DEFAULT '0',
  `sunday_2100` tinyint(1) NOT NULL DEFAULT '0',
  `sunday_2200` tinyint(1) NOT NULL DEFAULT '0',
  `sunday_2300` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`,`user_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `shifts_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `phone_number` varchar(13) NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-08-16 23:08:15
