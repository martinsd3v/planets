# Projeto Planetas Star Wars:

Esse projeto de trata de uma POC utilizando os conceitos de Clean Arch, Hexagonal Arch, Clean Code, DDD, e SOLID.

O principal motivo por escolher arquitetura hexagonal foi para deixar o desacoplamento, e divisão de responsabilidades bem claras.

Optei por trabalhar com dois "Domínios" sendo "users" um gerenciamento de usuário simples com autenticação JWT e "planets" com endpoints públicos para **listar planetas** e endpoints de `create`, `update`, e `delete` privados protegidos por autenticação JWT.

Para facilitar os testes adicionei no repositório uma pasta contendo os arquivos do **POSTMAN** pasta importar os arquivos presentes nesta no POSTMAN e sucesso!

## Cobertura por testes:

Para verificar a cobertura de testes basta rodar o comando: `make test-cover`

![Coverage test](https://github.com/martinsd3v/planets/blob/main/test-coverage.png?raw=true)

# Rodando o Projeto

- Para rodar o projeto é necessário ter o docker, e docker-compose instalado.
- Com docker devidamente instalado basta rodar o comando `docker-compose up`
- Url Base: http://localhost:9099 "Porta pode ser alterada em `config.yml` "
- Por padrão quando a aplicação sobe via docker compose um usuário é adicionado:
  - Email: emailteste@gmail.com
  - Senha: 123456
  - OBS: Utilizar esses dados para se autenticar

---

## Rotas de usuários:

**Autenticação/login:**

- POST: `{{urlBase}}`/users/auth
- BODY: `{"email": "emailteste@gmail.com","password": "123456"}`
- RESPONSE:

```
{
    "status": 200,
    "code": 100020,
    "message": "Authenticate success",
    "data": "TOKEN HERE"
}
```

---

**Cadastro de usuários:**

- POST: `{{urlBase}}`/users
- HEADER: `Authorization: Bearer {{Token}}`
- BODY: `{"name": "...","email": "...","password":"..."}`
- RESPONSE OK:

```
{
    "status": 200,
    "code": 100006,
    "message": "Success",
    "data": {
        "uuid": "...",
        "name": "...",
        "email": "..."
    }
}
```

- RESPONSE WITH ERROR:

```
{
    "status": 400,
    "code": 100021,
    "message": "Validation failed",
    "fields": [
        {
            "code": 100000,
            "message": "Already exists",
            "field": "email"
        }
    ]
}
```

---

**Lista de usuários:**

- GET: `{{urlBase}}`/users
- HEADER: `Authorization: Bearer {{Token}}`
- RESPONSE:

```
{
    "status": 200,
    "code": 100006,
    "message": "Success",
    "data": [
        {
            "uuid": "...",
            "name": "...",
            "email": "..."
        },
        {
            "uuid": "...",
            "name": "...",
            "email": "..."
        }
    ]
}
```

---

**Lista usuário por UUID:**

- GET: `{{urlBase}}`/users/`{{userID}}`
- HEADER: `Authorization: Bearer {{Token}}`
- RESPONSE:

```
{
    "status": 200,
    "code": 100006,
    "message": "Success",
    "data": {
        "uuid": "...",
        "name": "...",
        "email": "..."
    }
}
```

---

**Atualiza usuário por UUID:**

- PATCH: `{{urlBase}}`/users/`{{userID}}`
- HEADER: `Authorization: Bearer {{Token}}`
- BODY: `{"name": "...","email": "...","password":"..."}`
- RESPONSE OK:

```
{
    "status": 200,
    "code": 100006,
    "message": "Success",
    "data": {
        "uuid": "...",
        "name": "...",
        "email": "..."
    }
}
```

- RESPONSE WITH ERROR:

```
{
    "status": 400,
    "code": 100021,
    "message": "Validation failed",
    "fields": [
        {
            "code": 100000,
            "message": "Already exists",
            "field": "email"
        }
    ]
}
```

---

**Deleta usuário por UUID:**

- DELETE: `{{urlBase}}`/users/`{{userID}}`
- HEADER: `Authorization: Bearer {{Token}}`
- RESPONSE OK:

```
{
    "status": 200,
    "code": 100006,
    "message": "Success"
}
```

- RESPONSE WITH ERROR:

```
{
    "status": 500,
    "code": 100017,
    "message": "Unable to delete record"
}
```

## Rotas de planetas:

**Cadastro de planetas:**

- POST: `{{urlBase}}`/planets
- HEADER: `Authorization: Bearer {{Token}}`
- BODY: `{"name": "...","terrain": "...","climate":"..."}`
- RESPONSE OK:

```
{
    "status": 200,
    "code": 100006,
    "message": "Success",
    "data": {
        "uuid": "...",
        "name": "...",
        "terrain": "...",
        "climate": "...",
        "films": "..."
    }
}
```

- RESPONSE WITH ERROR:

```
{
    "status": 400,
    "code": 100021,
    "message": "Validation failed",
    "fields": [
        {
            "code": 100001,
            "message": "Required",
            "field": "climate"
        }
    ]
}
```

---

**Lista de planetas:** Na listagem de planetas é possivel aplicar busca por nome, sendo opcional. Se nao informar irá listar todos.

- GET: `{{urlBase}}`/planets?name=`{{planetName}}`
- RESPONSE:

```
{
    "status": 200,
    "code": 100006,
    "message": "Success",
    "data": [
        {
            "uuid": "...",
            "name": "...",
            "terrain": "...",
            "climate": "...",
            "films": "..."
        },
        {
            "uuid": "...",
            "name": "...",
            "terrain": "...",
            "climate": "...",
            "films": "..."
        }
    ]
}
```

---

**Lista planeta por UUID:**

- GET: `{{urlBase}}`/planets/`{{userID}}`
- RESPONSE:

```
{
    "status": 200,
    "code": 100006,
    "message": "Success",
    "data": {
        "uuid": "...",
        "name": "...",
        "terrain": "...",
        "climate": "...",
        "films": "..."
    }
}
```

---

**Atualiza planeta por UUID:**

- PATCH: `{{urlBase}}`/planets/`{{userID}}`
- HEADER: `Authorization: Bearer {{Token}}`
- BODY: `{"name": "...","terrain": "...","climate":"..."}`
- RESPONSE OK:

```
{
    "status": 200,
    "code": 100006,
    "message": "Success",
    "data": {
        "uuid": "...",
        "name": "...",
        "terrain": "...",
        "climate": "...",
        "films": "..."
    }
}
```

- RESPONSE WITH ERROR:

```
{
    "status": 400,
    "code": 100021,
    "message": "Validation failed",
    "fields": [
        {
            "code": 100001,
            "message": "Required",
            "field": "climate"
        }
    ]
}
```

---

**Deleta planeta por UUID:**

- DELETE: `{{urlBase}}`/planets/`{{userID}}`
- HEADER: `Authorization: Bearer {{Token}}`
- RESPONSE OK:

```
{
    "status": 200,
    "code": 100006,
    "message": "Success"
}
```

- RESPONSE WITH ERROR:

```
{
    "status": 500,
    "code": 100017,
    "message": "Unable to delete record"
}
```

## Dinâmica dos erros possíveis:

Para facilitar e padronizar os retornos em cada endpoint criei um padrão mensagens e codigos.

| Codigo |                   Mensagem                    |
| ------ | :-------------------------------------------: |
| 100000 |                Already exists                 |
| 100001 |                   Required                    |
| 100002 |                    Invalid                    |
| 100003 |                Invalid e-mail                 |
| 100004 |                 Invalid date                  |
| 100005 | Password must be between 6 and 40 characters. |
| 100006 |                    Success                    |
| 100007 |                   Not found                   |
| 100008 |                     Error                     |
| 100009 |          Record successfully created          |
| 100010 |          Record successfully updated          |
| 100011 |          Record successfully deleted          |
| 100012 |            Unable to create record            |
| 100013 |            Unable to update record            |
| 100014 |            Unable to delete record            |
| 100015 |             Unable to list record             |
| 100016 |              Authenticate failed              |
| 100017 |             Authenticate success              |
| 100018 |               Validation failed               |
| 100019 |              Endpoint not found               |
| 100020 |                  Unexpected                   |
