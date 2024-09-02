# Finance Manager API

![License](https://img.shields.io/badge/license-MIT-green)
![Build](https://img.shields.io/badge/build-passing-brightgreen)
![Version](https://img.shields.io/badge/version-1.0.0-blue)

## Sobre o Projeto

A **Finance Manager API** é uma API desenvolvida em Go utilizando o framework Echo para gerenciar finanças pessoais. Ela oferece funcionalidades para gerenciar usuários, contas, categorias de despesas e orçamentos, garantindo uma autenticação segura através de JSON Web Tokens (JWT).

## Funcionalidades

- **Usuários**: CRUD completo para gerenciar usuários.
- **Contas**: CRUD para gerenciar contas bancárias.
- **Categorias**: CRUD para gerenciar categorias de despesas.
- **Transações**: CRUD para registrar e gerenciar transações financeiras.
- **Orçamentos**: CRUD para gerenciar orçamentos, incluindo alocação e controle de despesas.
- **Autenticação JWT**: Proteção das rotas com JWT para garantir que apenas usuários autenticados possam acessar informações sensíveis.

## Tecnologias Utilizadas

- **Golang**: Linguagem de programação principal para desenvolvimento da API.
- **Echo**: Framework web utilizado para criar rotas e middleware de forma eficiente.
- **PostgreSQL**: Banco de dados utilizado para armazenar os dados financeiros.
- **Docker**: Utilizado para containerização da aplicação e do banco de dados.
- **JWT (JSON Web Tokens)**: Implementado para autenticação segura.

## Pré-requisitos

Antes de iniciar, certifique-se de ter o Docker e o Docker Compose instalados em sua máquina:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Como Executar

1. **Clone o repositório:**

   ```bash
   git clone https://github.com/seu-usuario/finance-manager-api.git
   cd finance-manager-api ´´´