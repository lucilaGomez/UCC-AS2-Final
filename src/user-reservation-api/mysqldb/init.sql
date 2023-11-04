CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(300) NOT NULL,
    last_name VARCHAR(300) NOT NULL,
    dni VARCHAR(8) NOT NULL,
    email VARCHAR(300) NOT NULL UNIQUE,
    password VARCHAR(300) NOT NULL,
    role VARCHAR(10) NOT NULL
);

INSERT INTO users (name, last_name, dni, email, password, role) VALUES
('Lucila', 'Gomez', '44755008', 'luligomez.ar@gmail.com', '$2a$10$EwlJ7rPRJPSpKtpsXchYoOs0.YpG7KAlfr42RmjECDFMelR9ICFQW', 'Admin'), -- Password: admin
('Franco', 'Tallone', '43273250', 'francotallone1109@gmail.com', '$2a$10$vnzBIZ0rRDOWX96L/jLhguXPLYG/gAFTbsAMzB/8RtzB5VDuh0jKq', 'Customer') -- Password: pass
