-- Create Database
-- CREATE DATABASE GOTask;

-- Use Database
USE GOTask;
GO

SELECT * FROM Project;
GO
-- Tasks Table
-- Reset (for reruns)
-- DROP TABLE IF EXISTS Task;
-- DROP TABLE IF EXISTS Project;
-- DROP TABLE IF EXISTS Users;
-- GO

-- -- Users
-- CREATE TABLE Usersv2 (
--     id UNIQUEIDENTIFIER PRIMARY KEY,
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
-- Step 5: Insert rows into Employees table
-- INSERT INTO User (FirstName, LastName, Email, HireDate) VALUES
-- ('Amit', 'Sharma', 'amit.sharma@example.com', '2023-01-15'),
-- ('Priya', 'Singh', 'priya.singh@example.com', '2023-02-20'),
-- ('Rahul', 'Verma', 'rahul.verma@example.com', '2023-03-25');
-- GO

-- -- Step 6: Insert rows into Departments table
-- INSERT INTO Departments (DepartmentName) VALUES
-- ('Human Resources'),
-- ('Finance'),
-- ('Engineering');
-- GO

-- -- Step 7: Update a row in Employees table
-- UPDATE Employees
-- SET Email = 'amit.sharma123@example.com'
-- WHERE FirstName = 'Amit' AND LastName = 'Sharma';
-- GO

-- -- Step 8: Delete a row from Employees table
-- DELETE FROM Employees
-- WHERE FirstName = 'Rahul' AND LastName = 'Verma';
-- GO

-- Step 9: Select all rows from Employees table
-- SELECT * FROM Users;
-- GO