CREATE TABLE targets
(
    id varchar(36) default(lower(hex(randomblob(4))) || '-' || lower(hex(randomblob(2))) || '-4' || substr(lower(hex(randomblob(2))),2) || '-' || substr('89ab',abs(random()) % 4 + 1, 1) || substr(lower(hex(randomblob(2))),2) || '-' || lower(hex(randomblob(6)))),
    name varchar(50),
    dye_id uuid,
    template_id uuid,
    threshold float,
    primary key ("id")
);
