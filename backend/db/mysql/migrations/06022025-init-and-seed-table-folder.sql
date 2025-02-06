CREATE TABLE folder (
    id int NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    parent_id int DEFAULT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX idx_folder_parent_id (parent_id),
    CONSTRAINT fk_folder_parent_id FOREIGN KEY (parent_id) REFERENCES folder (id)
);

INSERT INTO folder (name, parent_id) VALUES
    ('Folder 1', NULL),
    ('Folder 2', NULL),
    ('Folder 3', NULL),
    ('Sub Folder 2', 1),
    ('Sub Folder 3', 1),
    ('Sub Folder 4', 2);