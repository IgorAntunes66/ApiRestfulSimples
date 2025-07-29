Tasklist API

Descrição

Esta é uma API RESTful simples para gerenciar uma lista de tarefas (tasklist). Ela permite que os usuários criem, leiam, atualizem e excluam tarefas. O projeto foi desenvolvido em Go e utiliza um banco de dados PostgreSQL para armazenar os dados.

Tecnologias Utilizadas

    Go: Versão 1.23.0

    PostgreSQL: Banco de dados relacional para armazenamento das tarefas

    Docker & Docker Compose: Para criar um ambiente de desenvolvimento e produção containerizado e de fácil configuração

    Chi: Roteador HTTP leve e rápido para Go

    pgx: Driver e toolkit para PostgreSQL em Go

Como Executar o Projeto

Pré-requisitos

    Docker

    Docker Compose

Passos

Clone o repositório:

    git clone git@github.com:IgorAntunes66/ApiRestfulSimples.git
    cd ApiRestfulSimples

Inicie os containers:

O docker-compose.yml irá configurar e iniciar o container da aplicação e o container do banco de dados PostgreSQL.

    docker-compose up -d

A aplicação estará disponível em http://localhost:8080.

Pare os containers:

Para parar e remover os containers, redes e volumes, execute:

    docker-compose down

Endpoints da API

A API fornece os seguintes endpoints para gerenciar as tarefas:

Tarefas

    Listar todas as tarefas

        GET /tasks

        Descrição: Retorna uma lista com todas as tarefas cadastradas.

        Exemplo de Resposta:
        JSON

    [
        {
            "ID": 1,
            "Title": "Comprar pão",
            "Description": "Ir na padaria da esquina.",
            "Done": false
        },
        {
            "ID": 2,
            "Title": "Estudar Go",
            "Description": "Fazer o curso da Udemy.",
            "Done": true
        }
    ]

Buscar uma tarefa por ID

    GET /tasks/{ID}

    Descrição: Retorna uma tarefa específica com base no seu ID.

    Exemplo de Resposta:
    JSON

    {
        "ID": 1,
        "Title": "Comprar pão",
        "Description": "Ir na padaria da esquina.",
        "Done": false
    }

Criar uma nova tarefa

    POST /tasks

    Descrição: Cria uma nova tarefa.

    Exemplo de Corpo da Requisição:

      {
          "Title": "Fazer café",
          "Description": "Usar o pó novo.",
          "Done": false
      }

Exemplo de Resposta: 201 Created
JSON

    {
        "ID": 3,
        "Title": "Fazer café",
        "Description": "Usar o pó novo.",
        "Done": false
    }

Atualizar uma tarefa

    PUT /tasks/{ID}

    Descrição: Atualiza uma tarefa existente.

    Exemplo de Corpo da Requisição:

    {
        "Title": "Fazer um bom café",
        "Description": "Usar o pó novo e o coador de pano.",
        "Done": true
    }

    Exemplo de Resposta: 200 OK

Deletar uma tarefa

    DELETE /tasks/{ID}

    Descrição: Deleta uma tarefa com base no seu ID.

    Exemplo de Resposta: 204 No Content