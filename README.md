# 📋 Task Manager REST API

<div align="center">

![Go](https://img.shields.io/badge/Go-1.27-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![Gorilla Mux](https://img.shields.io/badge/Gorilla_Mux-v1.8.1-5C2D91?style=for-the-badge&logo=go&logoColor=white)
![REST API](https://img.shields.io/badge/REST-API-FF6B6B?style=for-the-badge)

**Uma REST API simples e eficiente para gerenciamento de tarefas, construída com Go e PostgreSQL.**

</div>

---

## 📌 Sobre o Projeto

O **Task Manager REST API** é um projeto back-end desenvolvido em **Go (Golang)** com o objetivo de praticar a construção de APIs RESTful. A API permite criar, listar, atualizar e deletar tarefas, utilizando **PostgreSQL** como banco de dados e o roteador **Gorilla Mux** para o gerenciamento das rotas.

---

## 🛠️ Tecnologias Utilizadas

| Tecnologia | Descrição |
|---|---|
| **Go (Golang)** | Linguagem principal do projeto |
| **PostgreSQL** | Banco de dados relacional |
| **Gorilla Mux** | Roteador HTTP para Go |
| **godotenv** | Carregamento de variáveis de ambiente via `.env` |
| **lib/pq** | Driver PostgreSQL para Go |

---

## 🏗️ Estrutura do Projeto

```
TaskManager_RestAPI/
├── config/
│   └── db.go           # Configuração e conexão com o banco de dados
├── handlers/
│   └── task_handler.go # Handlers das rotas (CRUD)
├── models/
│   └── task.go         # Model da Task e SQL de criação da tabela
├── .env                # Variáveis de ambiente (não versionado)
├── .gitignore
├── go.mod
├── go.sum
└── main.go             # Ponto de entrada da aplicação
```

---

## 📦 Model — Task

```go
type Task struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Status      bool   `json:"status"`
}
```

---

## 🔌 Endpoints da API

Base URL: `http://localhost:8080`

| Método | Rota | Descrição |
|--------|------|-----------|
| `GET` | `/tasks` | Lista todas as tarefas |
| `POST` | `/tasks` | Cria uma nova tarefa |
| `PUT` | `/tasks/{id}` | Atualiza uma tarefa pelo ID |
| `DELETE` | `/tasks/{id}` | Deleta uma tarefa pelo ID |

---

## ⚙️ Como Rodar o Projeto

### Pré-requisitos

- [Go](https://go.dev/dl/) instalado (v1.21+)
- [PostgreSQL](https://www.postgresql.org/download/) instalado e rodando

### 1. Clone o repositório

```bash
git clone https://github.com/DevZank/TaskManager_RestAPI.git
cd TaskManager_RestAPI
```

### 2. Configure as variáveis de ambiente

Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=seu_usuario
DB_PASSWORD=sua_senha
DB_NAME=nome_do_banco
```

### 3. Instale as dependências

```bash
go mod tidy
```

### 4. Execute a aplicação

```bash
go run main.go
```

A API estará disponível em: **`http://localhost:8080`**

---

## 🧪 Exemplos de Uso

### ➕ POST — Criar Tarefa

**Request:**
```http
POST /tasks
Content-Type: application/json

{
    "title": "Estudar Go",
    "description": "Praticar construção de APIs REST com Go",
    "status": false
}
```

**Response** `201 Created`:
```json
{
    "id": 1,
    "title": "Estudar Go",
    "description": "Praticar construção de APIs REST com Go",
    "status": false
}
```

### 📄 GET — Listar Tarefas

**Request:**
```http
GET /tasks
```

**Response** `200 OK`:
```json
[
    {
        "id": 1,
        "title": "Estudar Go",
        "description": "Praticar construção de APIs REST com Go",
        "status": false
    }
]
```

### ✏️ PUT — Atualizar Tarefa

**Request:**
```http
PUT /tasks/1
Content-Type: application/json

{
    "title": "Estudar Go",
    "description": "Praticar construção de APIs REST com Go",
    "status": true
}
```

**Response** `200 OK`:
```json
{
    "id": 1,
    "title": "Estudar Go",
    "description": "Praticar construção de APIs REST com Go",
    "status": true
}
```

### 🗑️ DELETE — Deletar Tarefa

**Request:**
```http
DELETE /tasks/1
```

**Response** `204 No Content`

---

## 📸 Demonstração

> 💡 *Prints tirados utilizando [Postman](https://www.postman.com/) / [Insomnia](https://insomnia.rest/) / [Thunder Client](https://www.thunderclient.com/)*

### POST — Criando uma Tarefa
<!-- Adicione aqui o print do POST -->
> 🖼️ *[Insira o print aqui]*

---

### GET — Listando as Tarefas (após o POST)
<!-- Adicione aqui o print do GET inicial -->
> 🖼️ *[Insira o print aqui]*

---

### PUT — Atualizando uma Tarefa
<!-- Adicione aqui o print do PUT -->
> 🖼️ *[Insira o print aqui]*

---

### DELETE — Deletando uma Tarefa
<!-- Adicione aqui o print do DELETE -->
> 🖼️ *[Insira o print aqui]*

---

### GET FINAL — Estado Final das Tarefas
<!-- Adicione aqui o print do GET final mostrando o estado atual -->
> 🖼️ *[Insira o print aqui]*

---

## 📚 Aprendizados

- Construção de APIs RESTful com Go
- Conexão e operações com PostgreSQL usando `database/sql`
- Uso do roteador **Gorilla Mux** para definição de rotas com parâmetros
- Gerenciamento de variáveis de ambiente com **godotenv**
- Organização de projetos Go em pacotes (`config`, `models`, `handlers`)
- Tratamento de erros e respostas HTTP adequadas

---

## 👤 Autor

<div align="center">

**DevZank**

[![GitHub](https://img.shields.io/badge/GitHub-DevZank-181717?style=for-the-badge&logo=github)](https://github.com/DevZank)

</div>

---

<div align="center">
  <sub>Feito com 💙 e muito Go</sub>
</div>
