create table transactions (
    id uuid primary key,
    from_wallet uuid not null,
    to_wallet uuid not null,
    amount decimal not null,
    date_time timestamp not null
);

alter table "transactions"
    add constraint fk_from_wallet foreign key (from_wallet) references wallets (id) on delete no action on update no action;

alter table "transactions"
    add constraint fk_to_wallet foreign key (to_wallet) references wallets (id) on delete no action on update no action;