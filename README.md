# Gomyadm

O Gomyadm é uma aplicação para gerenciamento de bancos de dados, desenvolvida em Go e React/TypeScript. O projeto oferece uma interface desktop com Wails, uma API HTTP que serve a interface web e uma CLI.

## Estado atual

### Bancos suportados

- MySQL/MariaDB
- PostgreSQL

O backend usa `database/sql` com drivers específicos para esses bancos. **SQLite ainda não está funcional**: a interface apresenta a opção SQLite e um seletor de arquivo que mostra o nome escolhido, mas o caminho não é enviado ao backend e não existe driver SQLite registrado.

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

1. **Representar o caminho do arquivo:** adicionar um campo `FilePath` à configuração da conexão em Go e no tipo correspondente do frontend. Para a primeira versão, aceitar apenas arquivos existentes, evitando criar acidentalmente um banco vazio.
2. **Abrir o diálogo nativo de arquivo:** criar uma binding Wails para selecionar o arquivo e disponibilizar o caminho escolhido à tela de conexão. O seletor HTML atual não fornece o caminho local ao backend; se a seleção também for necessária no modo web/API, será preciso definir um fluxo separado, como upload do arquivo.
3. **Implementar o driver SQLite:** adicionar a dependência Go, registrar o driver e implementar primeiro conexão, listagem de tabelas e leitura de dados. Incluir testes com um arquivo SQLite temporário.
4. **Adaptar o gerenciador de conexão:** usar o caminho do arquivo na abertura e na comparação de conexões. Definir o banco ativo como `main` ou equivalente e não simular operações de criar, renomear ou excluir bancos — no SQLite, o próprio arquivo representa o banco.
5. **Conectar o fluxo completo e validar:** enviar o caminho selecionado junto à conexão, testar conexão, navegação e desconexão e exibir erros claros para arquivos inexistentes ou inválidos.
