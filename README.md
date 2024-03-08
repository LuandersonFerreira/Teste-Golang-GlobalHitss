# Case GlobalHitss

> Olá nesse projeto temos como objetivo demonstrar conhecimentos de desenvolvimento em Golang, 
  o sistema consiste em um CRUD de usuários com a implementação de uma fila para inserção de dados no banco de dados.

## 💻 Pré-requisitos

Antes de começar, verifique se você atendeu aos seguintes requisitos:

- Utilizamos a versão `<v1.22.0>` do Go.
- Para fazer o setup do banco de dados certifique-se de que você tenha o docker desktop instalado em sua máquina;

## 🚀 Instalando <Projeto>

Para instalar o <sistema> basta ter os requisitos acima e fazer o clone do projeto na sua máquina e rode os seguintes comandos: 
- `make up` --> para subir o banco de dados;
- `docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3.12-management` --> para subir o rabbitMQ;


## ☕ Usando <Sistema>

Para usar <Sistema>, siga estas etapas:

dentro da raiz, rode o seguinte comando

```
make runapi
```
Aqui você já estará com o BackEnd rodando na sua máquina na porta 9000.

## ☕ Documentação do sistema<Sistema>
[Você pode acessar toda a documentação da API através desse link](https://documenter.getpostman.com/view/14355244/2sA2xh2CgE)

## 🚀 decisões do projeto

Utilizei alguns conceitos para tornar o projeto mais escalável: 

- UUID ao invés de Id serial, como identificador primário decidi utilizar uuid baseado no artigo a seguir: https://espigah.medium.com/id-vs-uuid-vs-publickey-5f7e19455b90
  com o id universal e unico(uuid) adicionamos uma camada a mais de proteção para a nossa Api evitando que usuários possam varrer nossa base de dados e pegar alguns dados que podem ser sensíveis, além disso tornamos a validação por chave primária bem mais robusta e sólida;
- Montagem das configs de acordo com um arquivo toml, para processos onde utilizamos nuvem e dockerização da aplicação sempre é importante poder controlar isso e maneira externa ao código, para isso decidi construir as configs do projeto a partir de um arquivo toml.