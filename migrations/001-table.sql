-- Create Authenticate Table
create table auths(
    id uuid not null,
    email varchar(50) not null,
    password text not null,
    salt text not null,
    is_active boolean default false,
    role int default 0, -- 0 USER / 1 ORGANIZATION / 2 SPONSOR / 3 ADMIN
    primary key (id)
);

-- Create Profile Table
create table profiles(
    id uuid not null,
    auth_id uuid not null,
    type int default 0, -- 0 MAIN / 1 SUB
    primary key (id),
    CONSTRAINT auth_id_fk
    FOREIGN KEY (auth_id)
    REFERENCES auths (id)
)
