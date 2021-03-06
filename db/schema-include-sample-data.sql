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
-- Table structure for table `additionals`
--

DROP TABLE IF EXISTS `additionals`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `additionals` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `rotation_id` int(11) unsigned NOT NULL,
  `date` datetime NOT NULL,
  `hour` datetime NOT NULL,
  `user_id` int(11) unsigned NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `user_id` (`user_id`),
  KEY `overrides_ibfk_2` (`rotation_id`),
  CONSTRAINT `additionals_ibfk_1` FOREIGN KEY (`rotation_id`) REFERENCES `rotations` (`id`),
  CONSTRAINT `additionals_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `additionals`
--

LOCK TABLES `additionals` WRITE;
/*!40000 ALTER TABLE `additionals` DISABLE KEYS */;
INSERT INTO `additionals` VALUES (2,'add1',4,'2018-09-11 00:00:00','0000-01-01 13:00:00',12,'2018-09-11 04:16:28','2018-09-11 04:16:37');
/*!40000 ALTER TABLE `additionals` ENABLE KEYS */;
UNLOCK TABLES;

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
  `date` datetime NOT NULL,
  `hour` datetime NOT NULL,
  `user_id` int(11) unsigned NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `user_id` (`user_id`),
  KEY `overrides_ibfk_2` (`rotation_id`),
  CONSTRAINT `overrides_ibfk_2` FOREIGN KEY (`rotation_id`) REFERENCES `rotations` (`id`),
  CONSTRAINT `overrides_ibfk_3` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `overrides`
--

LOCK TABLES `overrides` WRITE;
/*!40000 ALTER TABLE `overrides` DISABLE KEYS */;
INSERT INTO `overrides` VALUES (14,'test',4,'2018-09-11 00:00:00','0000-01-01 13:00:00',14,'2018-09-11 03:55:21','2018-09-11 03:55:21'),(15,'test2',4,'2018-09-11 00:00:00','0000-01-01 13:00:00',13,'2018-09-11 03:57:53','2018-09-11 04:07:39');
/*!40000 ALTER TABLE `overrides` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `reduces`
--

DROP TABLE IF EXISTS `reduces`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `reduces` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `rotation_id` int(11) unsigned NOT NULL,
  `date` datetime NOT NULL,
  `hour` datetime NOT NULL,
  `user_id` int(11) unsigned NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `user_id` (`user_id`),
  KEY `overrides_ibfk_2` (`rotation_id`),
  CONSTRAINT `reduces_ibfk_1` FOREIGN KEY (`rotation_id`) REFERENCES `rotations` (`id`),
  CONSTRAINT `reduces_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `reduces`
--

LOCK TABLES `reduces` WRITE;
/*!40000 ALTER TABLE `reduces` DISABLE KEYS */;
INSERT INTO `reduces` VALUES (3,'add1',4,'2018-09-11 00:00:00','0000-01-01 13:00:00',13,'2018-09-11 04:22:48','2018-09-11 04:22:48');
/*!40000 ALTER TABLE `reduces` ENABLE KEYS */;
UNLOCK TABLES;

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
-- Dumping data for table `rotation_details`
--

LOCK TABLES `rotation_details` WRITE;
/*!40000 ALTER TABLE `rotation_details` DISABLE KEYS */;
INSERT INTO `rotation_details` VALUES (4,1,12,'2018-07-14 02:04:51','2018-07-14 02:05:11'),(4,1,15,'2018-07-14 02:05:16','2018-07-14 02:05:16'),(4,1,18,'2018-07-14 02:05:20','2018-07-14 02:05:20'),(4,2,13,'2018-07-14 02:05:34','2018-07-14 02:05:34'),(4,2,16,'2018-07-14 02:05:38','2018-07-14 02:05:38'),(4,2,19,'2018-07-14 02:05:48','2018-07-14 02:05:48'),(4,3,14,'2018-07-14 02:05:54','2018-07-14 02:05:54'),(4,3,17,'2018-07-14 02:06:01','2018-07-14 02:06:01'),(4,3,20,'2018-07-14 02:06:12','2018-07-14 02:06:12');
/*!40000 ALTER TABLE `rotation_details` ENABLE KEYS */;
UNLOCK TABLES;

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
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rotations`
--

LOCK TABLES `rotations` WRITE;
/*!40000 ALTER TABLE `rotations` DISABLE KEYS */;
INSERT INTO `rotations` VALUES (4,'sample','2018-08-06','2018-07-05 15:21:04','2018-08-16 08:17:33');
/*!40000 ALTER TABLE `rotations` ENABLE KEYS */;
UNLOCK TABLES;

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
-- Dumping data for table `shifts`
--

LOCK TABLES `shifts` WRITE;
/*!40000 ALTER TABLE `shifts` DISABLE KEYS */;
INSERT INTO `shifts` VALUES (12,'sample_user1_1',12,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0,0,0,0,0,'2018-07-05 15:09:02','2018-08-16 08:17:58'),(13,'sample_user1_2',12,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,'2018-07-05 15:11:22','2018-08-16 08:18:02'),(14,'sample_user1_3',12,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,'2018-07-05 15:15:51','2018-08-16 08:18:04'),(15,'sample_user2_1',14,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,'2018-07-05 15:15:51','2018-08-16 08:18:09'),(16,'sample_user2_2',14,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0,0,0,0,0,'2018-07-05 15:09:02','2018-08-16 08:18:12'),(17,'sample_user2_3',14,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,'2018-07-05 15:11:22','2018-08-16 08:18:15'),(18,'sample_user3_1',13,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,'2018-07-05 15:11:22','2018-08-16 08:18:20'),(19,'sample_user3_2',13,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,'2018-07-05 15:15:51','2018-08-16 08:18:22'),(20,'sample_user3_3',13,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0,0,0,0,0,'2018-07-05 15:09:02','2018-08-16 08:18:26');
/*!40000 ALTER TABLE `shifts` ENABLE KEYS */;
UNLOCK TABLES;

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

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (12,'user1','+000000000001','2018-07-05 15:05:29','2018-08-16 08:17:12'),(13,'user2','+000000000002','2018-07-05 15:05:31','2018-08-16 08:17:15'),(14,'user3','+000000000003','2018-07-05 15:05:32','2018-08-16 08:17:17');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-09-11 13:54:31
