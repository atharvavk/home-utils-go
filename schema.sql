CREATE TABLE `geyser_status` (
  `is_on` tinyint(1) NOT NULL DEFAULT '0',
  `action_by` varchar(100) NOT NULL,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  KEY `geyser_status_residents_FK` (`action_by`),
  CONSTRAINT `geyser_status_residents_FK` FOREIGN KEY (`action_by`) REFERENCES `residents` (`key`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `residents` (
  `key` varchar(100) NOT NULL,
  `display_name` varchar(100) NOT NULL,
  PRIMARY KEY (`key`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE `geyser_history` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `action` varchar(100) NOT NULL,
  `resident_key` varchar(100) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `geyser_history_residents_FK` (`resident_key`),
  CONSTRAINT `geyser_history_residents_FK` FOREIGN KEY (`resident_key`) REFERENCES `residents` (`key`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
