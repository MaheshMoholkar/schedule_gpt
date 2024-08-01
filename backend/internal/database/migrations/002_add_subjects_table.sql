-- +goose Up
CREATE TABLE subjects (
    id INT IDENTITY(1,1) PRIMARY KEY,
    teacher_name NVARCHAR(100),
    subject_name NVARCHAR(100),
    occurrence INT
);

-- +goose Down
DROP TABLE subjects;
