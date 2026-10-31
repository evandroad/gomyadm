# Gomyadm

O Gomyadm é uma aplicação para gerenciamento de bancos de dados, desenvolvida em Go e React/TypeScript. O projeto oferece uma interface desktop com Wails, uma API HTTP que serve a interface web e uma CLI.

## Estado atual

### Bancos suportados

- MySQL/MariaDB
- PostgreSQL
- SQLite (leitura e escrita)

O backend usa `database/sql` com drivers específicos para esses bancos. SQLite abre arquivos existentes para leitura e escrita, permite visualizar e alterar registros e criar, renomear e remover tabelas e colunas. Na edição de colunas existentes, por enquanto é possível apenas renomeá-las; modificar tipo e restrições exige uma futura reconstrução segura da tabela.

### Funcionalidades disponíveis

- Criar, editar, listar e remover conexões salvas.
- Conectar e desconectar de uma sessão de banco de dados.
- Listar e selecionar bancos de dados.
- Listar tabelas e visualizar dados e estrutura.
- Criar, editar e remover registros.
- Criar, editar e remover colunas.
- Executar SQL pela interface.
- Usar a API HTTP e a CLI para operações de gerenciamento.

As conexões são salvas em `gomyadm/connections.json`, dentro do diretório de configuração do usuário definido pelo sistema operacional.

## Tecnologias e arquitetura

### Backend

- Go
- Chi para roteamento HTTP
- `database/sql`
- MySQL: [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
- PostgreSQL: [pgx](https://github.com/jackc/pgx)
- SQLite: [modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite)
- Wails para a aplicação desktop
- Cobra para a CLI

O gerenciamento de conexões fica em `internal/db`; os drivers implementam o contrato em `internal/drivers`; API, serviços e modelos estão organizados em `internal/`.

### Frontend

- React
- TypeScript
- Vite
- Tailwind CSS e componentes inspirados no Shadcn

A aplicação web compilada é embutida no binário Go. O mesmo frontend é usado na aplicação Wails e no servidor HTTP.

### Executáveis e API

- **Desktop:** aplicação Wails.
- **Web/API:** servidor HTTP na porta `8181`, com interface web e endpoints sob `/api`.
- **CLI:** comandos de gerenciamento de conexões e bancos.
- **Swagger:** documentação da API disponível em `/swagger/index.html` no servidor.

## Desenvolvimento

O projeto usa Go e Bun. Alvos comuns estão definidos no `Makefile`:

```sh
make frontend   # instala dependências e compila o frontend
make run        # gera documentação Swagger e inicia o servidor
make wails-dev  # inicia a aplicação desktop em modo de desenvolvimento
make cli        # compila a CLI
```

Para gerar a documentação Swagger, instale a ferramenta:

```sh
go install github.com/swaggo/swag/cmd/swag@latest
```
