create sequence user__id_seq
    as integer;

alter sequence user__id_seq owner to postgres;

create sequence user__id_seq1
    as integer;

alter sequence user__id_seq1 owner to postgres;

create table if not exists task
(
    _id          integer default nextval('user__id_seq'::regclass) not null,
    title        varchar(150)                                      not null,
    class        varchar(10)                                       not null,
    date         timestamp                                         not null,
    content      varchar(255),
    user_id      integer,
    contempled   boolean default false,
    satisfactory boolean default false,
    obs          varchar(255)
);

alter table task
    owner to postgres;

alter sequence user__id_seq owned by task._id;

create table if not exists "user"
(
    _id      integer default nextval('user__id_seq1'::regclass) not null,
    username varchar(20)                                        not null,
    password varchar(32)                                        not null
);

alter table "user"
    owner to postgres;

alter sequence user__id_seq1 owned by "user"._id;
