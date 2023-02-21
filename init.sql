# create databases
CREATE DATABASE IF NOT EXISTS `go_sample` COLLATE utf8mb4_general_ci;;

# create root user and grant rights
GRANT ALL PRIVILEGES ON *.* TO 'go_sample'@'%';