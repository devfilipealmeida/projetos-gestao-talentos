
# Sistema de Gerenciamento de Talentos - Frontend

Este é o frontend do sistema de gerenciamento de talentos, desenvolvido com React, TypeScript, Material UI. O sistema permite cadastrar candidatos com informações de nome, status (ativo/inativo) e habilidades.

## Tecnologias Utilizadas

- **React**
- **TypeScript**
- **Material UI**

## Funcionalidades

- Cadastro de novos candidatos.
- Edição de dados dos candidatos, incluindo nome, status (ativo/inativo) e habilidades.
- Listagem de candidatos com filtros de status e habilidades.

## Requisitos

- **Docker** e **Docker Compose** instalados.

## Instruções para Rodar

Para rodar o frontend com Docker, siga os passos abaixo:

1. Navegue até a pasta do frontend:
   ```bash
   cd gestaotalentos-front
   ```

2. Execute o comando para construir e iniciar o contêiner:
   ```bash
   docker-compose up --build
   ```

3. O frontend estará disponível em `http://localhost:3000`.

## Variáveis de Ambiente

Certifique-se de configurar as variáveis de ambiente para apontar o frontend ao backend. Por exemplo, você pode adicionar o arquivo `.env` na raiz do projeto com:

```plaintext
REACT_APP_API_URL=http://localhost:8080
```
