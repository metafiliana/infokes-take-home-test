CREATE TABLE file (
    id int NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    folder_id int NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX idx_file_folder_id (folder_id),
    CONSTRAINT fk_file_folder_id FOREIGN KEY (folder_id) REFERENCES folder (id)
);

INSERT INTO file (name, folder_id) VALUES
    ('File 1', 1),
    ('File 2', 1),
    ('File 3', 2),
    ('File 4', 3),
    ('File 5', 4),
    ('File 6', 5),
    ('File 7', 7);