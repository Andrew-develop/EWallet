alter table
    if exists "transactions" drop constraint fk_from_wallet;

alter table
    if exists "transactions" drop constraint fk_to_wallet;

drop table if exists "transactions";