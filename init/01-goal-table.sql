
USE patienceBank_db;
CREATE TABLE IF NOT EXISTS goal-table (
    goal_id INT AUTO_INCREMENT PRIMARY KEY,
    goal_money INT NOT NULL,
    goal_count INT NOT NULL,
);

