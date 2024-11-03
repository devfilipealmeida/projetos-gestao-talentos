
# 🧑‍💼 Talent Management Application

Uma aplicação para gestão de talentos, permitindo o cadastro, consulta e atualização de informações de candidatos. A aplicação utiliza uma estrutura moderna e escalável com **React** + **Material UI** no frontend, **Golang** no backend e **PostgreSQL** como banco de dados.

## 🚀 Tecnologias Utilizadas

- **Frontend**: React + Material UI
- **Backend**: Golang
- **Banco de Dados**: PostgreSQL
- **Containerização**: Docker e Docker Compose

## 📦 Instalação e Execução

Para iniciar a aplicação localmente, você precisa ter o [Docker](https://www.docker.com/get-started) instalado. Siga os passos abaixo:

1. Clone o repositório:
   ```bash
   git clone https://github.com/devfilipealmeida/projetos-gestao-talentos.git
   cd projetos-gestao-talentos
   ```

2. Inicie os containers com o Docker Compose:
   ```bash
   docker-compose up --build
   ```

Isso iniciará os serviços da aplicação, que incluirão:

- **Frontend**: disponível em `http://localhost:3000`
- **Backend**: disponível em `http://localhost:8080`
- **Banco de Dados**: PostgreSQL rodando na porta `5432`

---

## 📖 Endpoints da API

Aqui estão os principais endpoints da API para gerenciar candidatos.

### Autenticação

| Método | Endpoint         | Descrição                      |
|--------|-------------------|--------------------------------|
| POST   | `/register`      | Cria um novo usuário          |
| POST   | `/login`         | Autentica o usuário           |

### Endpoints de Candidatos

| Método | Endpoint             | Descrição                                  |
|--------|-----------------------|--------------------------------------------|
| GET    | `/candidatos`        | Retorna a lista de todos os candidatos     |
| POST   | `/candidatos`        | Cria um novo candidato                     |
| GET    | `/candidatos/{id}`    | Retorna as informações de um candidato específico |
| PUT    | `/candidatos/{id}`    | Atualiza as informações de um candidato específico |
| DELETE | `/candidatos/{id}`    | Deleta um candidato específico            |

### Exemplo de Payload

#### POST `/candidatos`

```json
{
  "nome": "João Silva",
  "email": "joao.silva@example.com",
  "telefone": "123456789",
  "cargo": "Desenvolvedor"
}
```

---

## 📂 Estrutura de Pastas

A estrutura principal do projeto é organizada da seguinte forma:

```
.
├── gestaotalentos-front/                # Diretório do frontend em React
│   ├── Dockerfile
│   └── ...
├── gestaotalentos/                # Diretório do backend em Go
│   ├── Dockerfile
│   └── ...
├── docker-compose.yml         # Configuração do Docker Compose
└── README.md                  # Documentação do projeto
```

---

## 🛠️ Funcionalidades Futuras

- **Filtros Avançados**: Filtragem de candidatos por habilidades, experiência, e localização.
- **Relatórios**: Geração de relatórios com gráficos e estatísticas dos candidatos.
- **Notificações**: Envio de notificações para candidatos e recrutadores.

---

## 📄 Licença

Este projeto está sob a licença MIT.

---

Sinta-se à vontade para contribuir, sugerir melhorias e abrir issues. Agradecemos seu interesse no projeto!