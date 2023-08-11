create sequence "Goals_Id_seq"
    as integer;

alter sequence "Goals_Id_seq" owner to postgres;

create sequence "Tags_Id_seq"
    as integer;

alter sequence "Tags_Id_seq" owner to postgres;

create sequence "Goals_Tags_Goals_Id_seq"
    as integer;

alter sequence "Goals_Tags_Goals_Id_seq" owner to postgres;

create sequence "Goals_Tags_Tags_Id_seq"
    as integer;

alter sequence "Goals_Tags_Tags_Id_seq" owner to postgres;

create sequence user_id_seq
    as integer;

alter sequence user_id_seq owner to postgres;

create table tags
(
    "Id"      integer default nextval('"Tags_Id_seq"'::regclass) not null
        constraint "Tags_pkey"
            primary key,
    "TagName" varchar(50)                                        not null,
    "Color"   integer                                            not null
);

alter table tags
    owner to postgres;

alter sequence "Tags_Id_seq" owned by tags."Id";

create table users
(
    id        integer default nextval('user_id_seq'::regclass) not null
        constraint id
            primary key,
    firstname varchar                                          not null,
    lastname  varchar                                          not null,
    email     varchar                                          not null,
    add_date  date                                             not null
);

alter table users
    owner to postgres;

alter sequence user_id_seq owned by users.id;

create table goals
(
    id          integer default nextval('"Goals_Id_seq"'::regclass) not null
        constraint "Goals_pkey"
            primary key,
    title       varchar(50)                                         not null,
    description varchar(250)                                        not null,
    start_date  date                                                not null,
    relevancy   real                                                not null,
    pin         boolean                                             not null,
    done        boolean                                             not null,
    user_id     serial
        constraint goals_user_id_fk
            references users
            on delete cascade
);

alter table goals
    owner to postgres;

alter sequence "Goals_Id_seq" owned by goals.id;

create table goals_tags
(
    "Goals_Id" integer default nextval('"Goals_Tags_Goals_Id_seq"'::regclass) not null
        constraint "Goals_Tags_Goals_Id_fkey"
            references goals,
    "Tags_Id"  integer default nextval('"Goals_Tags_Tags_Id_seq"'::regclass)  not null
        constraint "Goals_Tags_Tags_Id_fkey"
            references tags
);

alter table goals_tags
    owner to postgres;

alter sequence "Goals_Tags_Goals_Id_seq" owned by goals_tags."Goals_Id";

alter sequence "Goals_Tags_Tags_Id_seq" owned by goals_tags."Tags_Id";

