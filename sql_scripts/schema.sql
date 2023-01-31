
-- Drop and recreate DB
drop database if exists transactional_outbox;
create database transactional_outbox;

-- Drop and create User's table
drop table if exists user_table;
create table user_table (
    UserId int,
    FirstName varchar(255),
    LastName varchar(255),
    primary key (UserId)
);

-- Create Outbox Table
drop table if exists outbox_table
create table outbox_table (
    EventId int,
    Event varchar(255),
    primary key (EventId)
)