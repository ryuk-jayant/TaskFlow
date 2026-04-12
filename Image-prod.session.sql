-- Create Database
-- CREATE DATABASE GOTask;

-- Use Database
USE GOTask;
GO



-- Tasks Table
-- Reset (for reruns)
-- DROP TABLE IF EXISTS Task;
-- DROP TABLE IF EXISTS Project;
-- DROP TABLE IF EXISTS Users;
-- GO

-- Users
-- CREATE TABLE Users (
--     id UNIQUEIDENTIFIER PRIMARY KEY DEFAULT NEWID(),
--     name NVARCHAR(100) NOT NULL,
--     email NVARCHAR(255) NOT NULL UNIQUE,
--     password NVARCHAR(255) NOT NULL,
--     created_at DATETIME2 DEFAULT GETDATE()
-- );
-- GO

-- -- Projects
-- CREATE TABLE Project (
--     id UNIQUEIDENTIFIER PRIMARY KEY DEFAULT NEWID(),
--     name NVARCHAR(150) NOT NULL,
--     description NVARCHAR(MAX),
--     owner_id UNIQUEIDENTIFIER NOT NULL,
--     created_at DATETIME2 DEFAULT GETDATE(),

--     CONSTRAINT FK_Project_Users FOREIGN KEY (owner_id)
--         REFERENCES Users(id)
-- );
-- GO

-- -- Tasks
-- CREATE TABLE Task (
--     id UNIQUEIDENTIFIER PRIMARY KEY DEFAULT NEWID(),
--     title NVARCHAR(150) NOT NULL,
--     description NVARCHAR(MAX),
--     status NVARCHAR(20) CHECK (status IN ('todo', 'in_progress', 'done')),
--     priority NVARCHAR(10) CHECK (priority IN ('low', 'medium', 'high')),
--     project_id UNIQUEIDENTIFIER NOT NULL,
--     assignee_id UNIQUEIDENTIFIER NULL,
--     due_date DATE,
--     created_at DATETIME2 DEFAULT GETDATE(),
--     updated_at DATETIME2 DEFAULT GETDATE(),

--     CONSTRAINT FK_Task_Project FOREIGN KEY (project_id)
--         REFERENCES Project(id),

--     CONSTRAINT FK_Task_Users FOREIGN KEY (assignee_id)
--         REFERENCES Users(id)
-- );
-- GO

-- SELECT * FROM Users;
-- GO
SELECT * FROM Task;
GO