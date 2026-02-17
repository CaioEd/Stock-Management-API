# Stock-Management-API
Go API used in a stock management web application.

This is a Go project that uses a MySQL database with docker.

## Prerequisites

Before getting started, ensure you have the following installed:

- Go installed on your machine.
- Docker installed on your machine.


## Go Project Setup

1. **Clone the Repository**:
   - Clone this repository to your local machine using the following command:
     ```bash
     https://github.com/CaioEd/Stock-Management.git
     cd your-repository
     ```

2. **Install all dependencies**:
      - To ensure all required dependencies for the project are installed, run the following command:
         ```bash
         go mod tidy
         or
         make install (if you have Make installed)
         ```

3. **Run MySQL**:
   - With Docker installed:
     ```bash
     docker-compose up -d
     or
     make database
     ```

4. **.Env configs**:
   - Create a .env file on the root diretory and:
      - Create a JWT_SECRET
      - Database informations: User, Secret, Host, Port, DB name

## Running the Project

1. **Start the Go Application**:
   - After setting up the database and installing dependencies, run the project using:
     ```bash
     go run main.go
     or
     make run
     ```

2. **Access the Application**:
   - Follow any additional instructions provided in the project to interact with the application, such as accessing an API endpoint or a frontend interface.

## Contribution

Feel free to open issues or create pull requests if you encounter any problems or have suggestions for improvement.


## Postman collection (Mercado Livre - anúncios de teste)

Foi adicionada uma collection pronta em `postman/mercadolivre-test-items.collection.json` com fluxo para:

1. Gerar `access_token` via `refresh_token`.
2. Validar usuário autenticado (`/users/me`).
3. Descobrir categoria sugerida (`/domain_discovery/search`).
4. Criar anúncio com dados aleatórios (título, preço, quantidade, SKU e descrição).
5. Consultar o anúncio criado.

### Arquivos

- Collection: `postman/mercadolivre-test-items.collection.json`
- Environment: `postman/mercadolivre-test-items.environment.json`

### Como usar

1. Importe a collection e o environment no Postman.
2. Preencha as variáveis:
   - `app_id`
   - `client_secret`
   - `refresh_token`
3. (Opcional) Ajuste `category_id`, `listing_type_id` e demais variáveis.
4. Execute os requests na ordem numérica da collection.

> Observação: Dependendo da categoria, o Mercado Livre pode exigir atributos adicionais. Se ocorrer erro de validação, consulte o retorno da API e inclua os campos exigidos no request de criação do item.
