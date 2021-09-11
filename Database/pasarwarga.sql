-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Sep 11, 2021 at 04:38 AM
-- Server version: 5.7.33
-- PHP Version: 7.4.19

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `pasarwarga`
--

-- --------------------------------------------------------

--
-- Table structure for table `articles`
--

CREATE TABLE `articles` (
  `id` int(50) NOT NULL,
  `title` varchar(250) DEFAULT NULL,
  `slug` varchar(250) DEFAULT NULL,
  `category_id` int(50) DEFAULT NULL,
  `content` varchar(250) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `articles`
--

INSERT INTO `articles` (`id`, `title`, `slug`, `category_id`, `content`, `created_at`, `updated_at`) VALUES
(2, 'kancil updated', 'kancil updated', 0, 'kancil updated', '2021-09-11 03:31:16', '2021-09-11 09:41:32'),
(3, 'kancil1', 'kancil1', 0, 'saasa1', '2021-09-11 09:14:59', '2021-09-11 09:14:59'),
(4, 'Mentimun', 'Mentimun', 0, 'Menntimun', '2021-09-11 09:48:35', '2021-09-11 09:48:35'),
(5, 'Jeruk', 'Jeruk', 23, 'Jeruk', '2021-09-11 09:53:08', '2021-09-11 09:53:08'),
(6, 'Singa', 'Singa', 25, 'Singa', '2021-09-11 10:45:47', '2021-09-11 10:45:47');

-- --------------------------------------------------------

--
-- Table structure for table `categoris`
--

CREATE TABLE `categoris` (
  `id` int(11) NOT NULL,
  `category_name` varchar(250) DEFAULT NULL,
  `category_slug` varchar(250) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `categoris`
--

INSERT INTO `categoris` (`id`, `category_name`, `category_slug`, `created_at`, `updated_at`) VALUES
(2, 'category pertama', 'category', '2021-09-10 20:59:57', '2021-09-10 20:59:57'),
(3, 'category pertama', 'category', '2021-09-10 21:22:09', '2021-09-10 21:22:09'),
(4, 'category pertama', 'category', '2021-09-10 21:23:55', '2021-09-10 21:23:55'),
(5, 'category pertama', 'category', '2021-09-10 21:35:11', '2021-09-10 21:35:11'),
(6, 'category pertama', 'category', '2021-09-10 21:35:44', '2021-09-10 21:35:44'),
(7, 'category kedua', 'category', '2021-09-10 21:35:54', '2021-09-10 21:35:54'),
(9, '', '', '2021-09-10 22:10:58', '2021-09-10 22:10:58'),
(10, '', '', '2021-09-10 22:11:20', '2021-09-10 22:11:20'),
(11, '', '', '2021-09-10 22:11:30', '2021-09-10 22:11:30'),
(12, '', '', '2021-09-10 22:25:36', '2021-09-10 22:25:36'),
(13, '', '', '2021-09-10 22:25:45', '2021-09-10 22:25:45'),
(14, '', '', '2021-09-10 22:25:48', '2021-09-10 22:25:48'),
(15, '', '', '2021-09-10 22:25:55', '2021-09-10 22:25:55'),
(16, '', '', '2021-09-10 22:26:34', '2021-09-10 22:26:34'),
(17, '', '', '2021-09-10 22:26:52', '2021-09-10 22:26:52'),
(19, '', 'category', '2021-09-10 22:28:08', '2021-09-10 22:28:08'),
(20, '', 'category', '2021-09-10 23:40:10', '2021-09-10 23:40:10'),
(23, 'buku updated', 'buku updated', '2021-09-11 07:52:50', '2021-09-11 07:53:50'),
(25, 'Hewan', 'Hewan', '2021-09-11 10:45:24', '2021-09-11 10:45:24');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `articles`
--
ALTER TABLE `articles`
  ADD PRIMARY KEY (`id`) USING BTREE;

--
-- Indexes for table `categoris`
--
ALTER TABLE `categoris`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `articles`
--
ALTER TABLE `articles`
  MODIFY `id` int(50) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `categoris`
--
ALTER TABLE `categoris`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=26;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
