DROP TABLE IF EXISTS `Categories`;
CREATE TABLE `Categories` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Name` text DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `Categories` (`Id`, `Name`) VALUES
(1,	'Watersports'),
(2,	'Soccer');

DROP TABLE IF EXISTS `Products`;
CREATE TABLE `Products` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Name` text DEFAULT NULL,
  `Category` int(11) DEFAULT NULL,
  `Price` decimal(8,2) DEFAULT NULL,
  PRIMARY KEY (`Id`),
  KEY `Category` (`Category`),
  CONSTRAINT `Products_ibfk_1` FOREIGN KEY (`Category`) REFERENCES `Categories` (`Id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `Products` (`Id`, `Name`, `Category`, `Price`) VALUES
(1,	'Kayak',	1,	279.00),
(2,	'Lifejacket',	1,	48.95),
(3,	'Soccer Ball',	2,	19.50),
(4,	'Corner Flags',	2,	34.95);
