-- phpMyAdmin SQL Dump
-- version 4.6.6deb5
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Sep 05, 2019 at 05:27 PM
-- Server version: 10.1.40-MariaDB-0ubuntu0.18.04.1
-- PHP Version: 7.2.19-0ubuntu0.18.04.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `petshop`
--

-- --------------------------------------------------------

--
-- Table structure for table `reports`
--

CREATE TABLE `reports` (
  `id` varchar(22) NOT NULL,
  `month` varchar(11) NOT NULL,
  `week` int(11) NOT NULL,
  `units_sold` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `reports`
--

INSERT INTO `reports` (`id`, `month`, `week`, `units_sold`) VALUES
('September1', 'September', 1, 82);

-- --------------------------------------------------------

--
-- Table structure for table `stock`
--

CREATE TABLE `stock` (
  `ID` int(11) NOT NULL,
  `animal` varchar(30) NOT NULL,
  `quantity_ready` int(8) NOT NULL,
  `quantity_available` int(8) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `stock`
--

INSERT INTO `stock` (`ID`, `animal`, `quantity_ready`, `quantity_available`) VALUES
(1, 'Birds', 15, 30),
(2, 'Cats', 10, 30),
(3, 'Dogs', 5, 15);

-- --------------------------------------------------------

--
-- Table structure for table `testing_reports`
--

CREATE TABLE `testing_reports` (
  `id` varchar(22) NOT NULL,
  `month` varchar(11) NOT NULL,
  `week` int(11) NOT NULL,
  `units_sold` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `testing_stock`
--

CREATE TABLE `testing_stock` (
  `ID` int(11) NOT NULL,
  `animal` varchar(30) NOT NULL,
  `quantity_ready` int(8) NOT NULL,
  `quantity_available` int(8) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `testing_stock`
--

INSERT INTO `testing_stock` (`ID`, `animal`, `quantity_ready`, `quantity_available`) VALUES
(1, 'bird', 22000, 22000),
(2, 'cat', 21879, 21879),
(3, 'dog', 22000, 22000);

-- --------------------------------------------------------

--
-- Table structure for table `testing_transactions_log`
--

CREATE TABLE `testing_transactions_log` (
  `date` datetime NOT NULL,
  `log` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `testing_transactions_log`
--

INSERT INTO `testing_transactions_log` (`date`, `log`) VALUES
('2019-09-04 04:09:00', 'cat bought X 22'),
('2019-09-04 19:48:36', 'cat bought X 22'),
('2019-09-04 20:02:02', 'Bird bought X 2'),
('2019-09-05 13:18:00', 'Cat bought X 33'),
('2019-09-05 13:19:57', 'Dog bought X 4443'),
('2019-09-05 13:30:58', 'Cat bought X 2'),
('2019-09-05 13:32:43', 'Bird bought X 2'),
('2019-09-05 13:36:27', 'Cat bought X 2'),
('2019-09-05 14:08:57', 'Cat bought X 3'),
('2019-09-05 14:13:11', 'Dog bought X 2'),
('2019-09-05 14:39:39', 'Cat bought X 2'),
('2019-09-05 14:39:53', 'Cat bought X 2'),
('2019-09-05 14:46:30', 'Cat bought X 3'),
('2019-09-05 14:46:46', 'Dog bought X 22'),
('2019-09-05 14:49:51', 'Cat bought X 22'),
('2019-09-05 14:49:58', 'Dog bought X 22'),
('2019-09-05 14:50:09', 'Cat bought X 44'),
('2019-09-05 14:50:48', 'Cat bought X 222'),
('2019-09-05 14:51:31', 'Dog bought X 1'),
('2019-09-05 14:55:11', 'Cat bought X 33'),
('2019-09-05 14:55:48', 'Cat bought X 22'),
('2019-09-05 14:57:24', 'Dog bought X 333'),
('2019-09-05 15:16:11', 'Cat bought X 2'),
('2019-09-05 15:18:41', 'Dog bought X 33'),
('2019-09-05 15:19:21', 'Cat bought X 2'),
('2019-09-05 15:20:54', 'Cat bought X 20'),
('2019-09-05 15:30:43', 'Cat bought X 33'),
('2019-09-05 15:30:57', 'Cat bought X 1'),
('2019-09-05 15:33:41', 'Bird bought X 2'),
('2019-09-05 15:40:22', 'Bird bought X 1'),
('2019-09-05 15:47:44', 'Cat bought X 1'),
('2019-09-05 15:47:54', 'Bird bought X 1'),
('2019-09-05 15:49:24', 'Cat bought X 2'),
('2019-09-05 15:50:17', 'Cat bought X 5'),
('2019-09-05 15:51:03', 'Cat bought X 2'),
('2019-09-05 15:55:28', 'Cat bought X 2'),
('2019-09-05 15:59:05', 'Cat bought X 2'),
('2019-09-05 16:01:52', 'Cat bought X 2'),
('2019-09-05 16:47:00', 'Cat bought X 1'),
('2019-09-05 17:11:08', 'Cat bought X 1'),
('2019-09-05 17:11:12', 'Dog bought X 1'),
('2019-09-05 17:11:31', 'Bird bought X 1');

-- --------------------------------------------------------

--
-- Table structure for table `transactions_log`
--

CREATE TABLE `transactions_log` (
  `date` datetime NOT NULL,
  `log` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `transactions_log`
--

INSERT INTO `transactions_log` (`date`, `log`) VALUES
('2019-09-04 04:09:00', 'cat bought X 22'),
('2019-09-04 19:48:36', 'cat bought X 22'),
('2019-09-04 19:56:37', 'cat bought X 15');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `reports`
--
ALTER TABLE `reports`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `stock`
--
ALTER TABLE `stock`
  ADD PRIMARY KEY (`ID`);

--
-- Indexes for table `testing_reports`
--
ALTER TABLE `testing_reports`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `testing_stock`
--
ALTER TABLE `testing_stock`
  ADD PRIMARY KEY (`ID`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `stock`
--
ALTER TABLE `stock`
  MODIFY `ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
--
-- AUTO_INCREMENT for table `testing_stock`
--
ALTER TABLE `testing_stock`
  MODIFY `ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
