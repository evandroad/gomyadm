# Gomyadm

O Gomyadm é uma GUI web para gerenciamento de bancos de dados, desenvolvida em Go e distribuída como um único binário executável.
O projeto nasceu inspirado em ferramentas como Adminer e phpMyAdmin, buscando uma alternativa com visual moderno, menor complexidade de instalação e melhor experiência de uso.
Diferente de aplicações desktop como DBeaver ou MySQL Workbench, o Gomyadm roda diretamente no navegador, sem depender de Java ou instalações pesadas.

## Objetivos do projeto
- Interface moderna e leve
- Distribuição em binário único
- Execução local sem dependências complexas
- Frontend embutido no executável
- Fácil instalação e compartilhamento
- Suporte a múltiplos bancos de dados
- API simples e extensível

## Instalar swagger
```
go install github.com/swaggo/swag/cmd/swag@latest
```

## Stack inicial

### Backend

* Go
* router [Chi](https://github.com/go-chi/chi)
* Bibliotecas Go
  * MySQL: [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql?utm_source=chatgpt.com)
  * PostgreSQL: [pgx](https://github.com/jackc/pgx?utm_source=chatgpt.com)

### Frontend

* React + Vite
* TypeScript
* UI
  * [Tailwind CSS](https://tailwindcss.com)
  * [Shadcn](https://ui.shadcn.com/)

### Banco de Dados
* MySQL/MariaDB
* PostgreSQL

---

# A arquitetura
O frontend será compilado e embutido diretamente no binário Go utilizando embed, permitindo que toda a aplicação seja distribuída como um único executável.

## Estrutura

```txt
/backend
  /internal
    /db
    /drivers
    /api
    /services
/frontend
```

---

# Modelagem de conexão

Exemplo:

```go
type Driver interface {
  Connect(config ConnectionConfig) error
  Disconnect() error

  ListDatabases() ([]Database, error)
  ListTables(database string) ([]Table, error)
  DescribeTable(database, table string) ([]Column, error)

  Query(query string, limit int) (*QueryResult, error)

  BeginTx() error
}
```

Depois:

```go
type MySQLDriver struct {}
type PostgresDriver struct {}
```

---

## Connection Manager

Tipo:

```txt
[ Nova conexão ]

Nome: Local MySQL
Host:
Porta:
Usuário:
Senha:
SSL:
Testar conexão
Salvar
```

---

# Como salvar conexões

No início:

```json
~/.gomyadm/connections.json
```

---

# Modelo de conexão

```go
type ConnectionConfig struct {
  ID       string
  Name     string
  Driver   string

  Host     string
  Port     int

  Username string
  Password string

  Database string
  SSLMode  string
}
```

---

## Sessão — manter pool vivo

```txt
ConnectionManager
    -> mantém pools vivos
    -> cacheados por ID
```

Algo assim:

```go
map[string]*sql.DB
```

---

# API inicialmente

## Conexões

```http
GET    /api/connections
POST   /api/connections
DELETE /api/connections/:id
POST   /api/connections/:id/test
```

## Schema

```http
GET /api/connections/:id/databases
GET /api/connections/:id/tables
GET /api/connections/:id/table/:name
```

## Query

```http
POST /api/query
```

Body:

```json
{
  "connectionId": "abc",
  "query": "SELECT * FROM users LIMIT 100"
}
```

---

# Query execution

No início:

## suporte SOMENTE:

* SELECT
* INSERT
* UPDATE
* DELETE

---

# Uma arquitetura MUITO boa

## Driver Registry

Exemplo:

```go
var drivers = map[string]DriverFactory{}
```

Registro:

```go
Register("mysql", NewMySQLDriver)
Register("postgres", NewPostgresDriver)
```

---

# Frontend

## Layout

```txt
Sidebar esquerda
    Conexões
    Databases
    Tables

Centro
    SQL Editor

Bottom
    Results Grid
```

---

# Para o editor SQL
* textarea simples
---

# Feature roadmap

## Fase 1

* conexões
* listar schema
* executar queries
* grid de resultados

## Fase 2

* tabs
* histórico SQL
* favoritos
* export CSV

## Fase 3

* autocomplete
* explain analyze
* schema diff
* snippets

## Fase 4

* SSH tunnel
* Redis
* Mongo
* migrations
* users/roles

---

# MVP

## Features

- [x] salvar conexão
- [x] conectar
- [x] listar tabelas
- [x] executar SELECT
- [x] mostrar grid
- [] Criar tabela
- [] Renomear tabela
- [] Excluir tabela
- [x] Criar coluna
- [x] Editar coluna
- [x] Excluir coluna
- [] Visualizar índices
- [x] Executar SQL livre