
# Sistema de Gerenciamento de Talentos - Backend

Este é o backend do sistema de gerenciamento de talentos, desenvolvido em Golang. Ele oferece APIs para gerenciar o cadastro, autenticação e as informações dos candidatos.

## Tecnologias Utilizadas

- **Golang**
- **Gorilla Mux** - para roteamento
- **PostgreSQL** (ou outro banco de dados configurado no Docker Compose)

## Endpoints

### Autenticação

- **`POST /register`** - Cadastra um novo usuário.
- **`POST /login`** - Realiza o login de um usuário.

### Gerenciamento de Candidatos

- **`GET /candidatos`** - Lista todos os candidatos (requer autenticação).
- **`POST /candidatos`** - Cria um novo candidato (requer autenticação).
- **`GET /candidatos/{id}`** - Busca um candidato pelo ID (requer autenticação).
- **`PUT /candidatos/{id}`** - Atualiza as informações de um candidato pelo ID (requer autenticação).
- **`DELETE /candidatos/{id}`** - Remove um candidato pelo ID (requer autenticação e permissões de administrador).

## Requisitos

- **Docker** e **Docker Compose** instalados.

## Instruções para Rodar

Para rodar o backend com Docker, siga os passos abaixo:

1. Navegue até a pasta do backend:
   ```bash
   cd gestaotalentos
   ```

2. Execute o comando para construir e iniciar o contêiner:
   ```bash
   docker-compose up --build
   ```

3. A API estará disponível em `http://localhost:8080`.

## Variáveis de Ambiente

O backend pode requerer variáveis de ambiente, como configurações do banco de dados e JWT. Configure um arquivo `.env` com as variáveis necessárias, por exemplo: (Já deixei o repo com as variáveis)

```plaintext
DB_HOST=db
DB_PORT=5432
DB_USER=usuario
DB_PASSWORD=senha
DB_NAME=gestaotalentos
JWT_SECRET=sua_chave_secreta
```

## Documentação da API

A documentação da API está disponível no Swagger após rodar o backend. Acesse `http://localhost:8080/swagger/index.html` para visualizar e interagir com os endpoints.
