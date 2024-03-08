create table users (
    uuid varchar(255) not null,
    nome varchar(255) not null,
    sobrenome varchar(255) not null,
    contato varchar(255),
    endereco varchar(255),
    nascimento date,
    cpf varchar(14)
);