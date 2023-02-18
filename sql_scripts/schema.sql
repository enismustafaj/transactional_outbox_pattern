
-- Drop and recreate DB
drop database if exists transactional_outbox;
create database transactional_outbox;

-- Drop and create User's table
drop table if exists user_table;
create table user_table (
    UserId int not null auto_increment,
    FirstName varchar(255),
    LastName varchar(255),
    Age int,
    Email varchar(255)
    primary key (UserId)
);

-- Create Outbox Table
drop table if exists outbox_table;
create table outbox_table (
    EventId int auto_increment,
    Event varchar(255),
    EntityId int
    primary key (EventId)
) 