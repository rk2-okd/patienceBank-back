
USE patienceBank_db;
CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(50) UNIQUE NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password_hash VARCHAR(255) NOT NULL,
  comment VARCHAR(255),
  auth_provider VARCHAR(50) NOT NULL DEFAULT 'local'
);
