create table tb_account (
  id integer primary key auto_increment,
  document_number varchar(14) not null unique,
  credit_limit float default 0
)

insert into tb_account (document_number) values
('117.988.224-58');


create table tb_operation_type (
  id integer primary key auto_increment,
  description varchar(100)
)


insert into tb_operation_type (description) values
('COMPRA A VISTA'),
('COMPRA PARCELADA'),
('SAQUE'),
('PAGAMENTO');

create table tb_transaction (
  id integer primary key auto_increment,
  amount float not null,
  account_id integer not null,
  operation_type_id integer not null,
  event_date date,
  foreign key fk_transaction_account(account_id) references tb_account(id),
  foreign key fk_transaction_operation_type(operation_type_id) references tb_operation_type(id)
)

insert into tb_transaction (amount, account_id, operation_type_id, event_date)
values
(100, 1, 4, NOW());

update tb_account set credit_limit=100 where id=1;

