
USE patienceBank_db;
CREATE TABLE IF NOT EXISTS record-table (
    gaman_id INT AUTO_INCREMENT PRIMARY KEY,
    gaman_thing VARCHAR(255) NOT NULL,
    gaman_money INT NOT NULL,
    gaman_day DATE NOT NULL
);

