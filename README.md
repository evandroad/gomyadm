# Gomyadm

O Gomyadm é uma aplicação para gerenciamento de bancos de dados, desenvolvida em Go e React/TypeScript. O projeto oferece uma interface desktop com Wails, uma API HTTP que serve a interface web e uma CLI.

## Estado atual

### Bancos suportados

- MySQL/MariaDB
- PostgreSQL

O backend usa `database/sql` com drivers específicos para esses bancos. **SQLite ainda não está funcional**: a interface desktop já permite selecionar um arquivo local e envia o caminho na configuração da conexão, mas o backend ainda não tem driver SQLite registrado.

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

1. **Representar o caminho do arquivo:** concluído com o campo `FilePath` na configuração Go e nos tipos do frontend. A validação do arquivo ainda faz parte da implementação do driver.
2. **Abrir o diálogo nativo de arquivo:** concluído na aplicação desktop com uma binding Wails. A seleção no modo web/API requer um fluxo separado, como upload do arquivo.
3. **Implementar o driver SQLite:** adicionar a dependência Go, registrar o driver e implementar primeiro conexão, listagem de tabelas e leitura de dados. Incluir testes com um arquivo SQLite temporário.
4. **Adaptar o gerenciador de conexão:** a comparação já considera o caminho; falta usar esse caminho para abrir o SQLite. Definir o banco ativo como `main` ou equivalente e não simular operações de criar, renomear ou excluir bancos — no SQLite, o próprio arquivo representa o banco.
5. **Conectar o fluxo completo e validar:** abrir no gerenciador o caminho já recebido, testar conexão, navegação e desconexão e exibir erros claros para arquivos inexistentes ou inválidos.

## Status do suporte SQLite

- [x] Adicionar `FilePath` aos modelos Go e TypeScript.
- [x] Permitir selecionar um arquivo SQLite pelo diálogo nativo da aplicação Wails.
- [ ] Implementar o driver SQLite e as operações iniciais de conexão e leitura.
- [ ] Integrar a abertura do arquivo ao gerenciador e validar o fluxo completo.
- [ ] Definir um fluxo de seleção para o modo web/API, se necessário.
