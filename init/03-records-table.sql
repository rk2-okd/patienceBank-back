USE patienceBank_db;

CREATE TABLE IF NOT EXISTS records (
    workout_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NULL,
    trained_part TINYINT NOT NULL,
    workout_duration INT NOT NULL,
    workout_date DATE NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL
);
