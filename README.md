# auth-golang-react
An authentication developed using Golang for backend and React for frontend.

## Installation

Clone the repository:
    ```
    git clone https://github.com/pxiaccount/auth-golang-react.git
    ```

### Golang

1. Change directory in to backend:

    ```bash
    cd auth-golang-react/backend
    ```

2. Create .env files containing:

```bash
PORT= // Insert port here
SECRET="" // Insert your password
DATABASE="host=YOURHOST user=USERNAME password=PASSWORD dbname=YOURDBNAME port=YOURDBPORT sslmode=disable"
```

3. Run your <a href="https://www.postgresql.org/">PostgreSQL</a> server

4. Run the server using <a href="https://github.com/githubnemo/CompileDaemon">CompileDaemon</a> by <a href="https://github.com/githubnemo">githubnemo</a>

```bash
    compiledaemon -command="./auth-golang-react"
```

### React
1. Change directory in to client:

    ```bash
    cd auth-golang-react/frontend
    ```

2. Install packages:

    ```bash
    npm i
    ```

3. Run the development server:

    ```bash
    npm run dev
    ```