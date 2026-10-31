# Gomyadm

O Gomyadm é uma aplicação para gerenciamento de bancos de dados, desenvolvida em Go e React/TypeScript. O projeto oferece uma interface desktop com Wails, uma API HTTP que serve a interface web e uma CLI.

## Estado atual

### Bancos suportados

- MySQL/MariaDB
- PostgreSQL
- SQLite (conexão e visualização em modo somente leitura)

O backend usa `database/sql` com drivers específicos para esses bancos. O suporte SQLite inicial abre arquivos existentes, lista tabelas e permite visualizar dados e estrutura; alterações de bancos, tabelas, colunas e registros não estão disponíveis.

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
- SQLite: [modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite) (somente leitura nesta versão)
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

## Próximos passos: suporte SQLite no backend

1. **Representar o caminho do arquivo:** concluído com o campo `FilePath` na configuração Go e nos tipos do frontend. O gerenciador valida se o caminho existe e aponta para um arquivo regular.
2. **Abrir o diálogo nativo de arquivo:** concluído na aplicação desktop com uma binding Wails. A seleção no modo web/API requer um fluxo separado, como upload do arquivo.
3. **Implementar o driver SQLite:** concluído para conexão com arquivos existentes, listagem de tabelas, leitura de dados e estrutura. As operações de alteração ainda não estão disponíveis.
4. **Adaptar o gerenciador de conexão:** concluído; o caminho é validado e usado para abrir o arquivo existente em modo somente leitura, e `main` representa o banco ativo. Operações de criar, renomear ou excluir bancos não são simuladas.
5. **Conectar o fluxo completo e validar:** concluído para o desktop; conexão, navegação e leitura são cobertas por teste com arquivo temporário. A conexão falha para caminhos inexistentes e arquivos inválidos.

## Status do suporte SQLite

- [x] Adicionar `FilePath` aos modelos Go e TypeScript.
- [x] Permitir selecionar um arquivo SQLite pelo diálogo nativo da aplicação Wails.
- [x] Implementar o driver SQLite para conexão e leitura de tabelas e dados.
- [x] Integrar a abertura do arquivo ao gerenciador e validar com arquivo temporário.
- [ ] Definir um fluxo de seleção para o modo web/API, se necessário.
