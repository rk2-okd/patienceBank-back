USE patienceBank_db;

CREATE TABLE IF NOT EXISTS records (
    gaman_id INT AUTO_INCREMENT PRIMARY KEY,
    gaman_thing VARCHAR(255) NOT NULL,
    gaman_minutes INT NOT NULL,
    gaman_day DATE NOT NULL,
    user_id INT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL
);
