# auth-golang-react

An authentication developed using Golang for backend and React for frontend.

## Prerequisites

- [Go](https://go.dev/)
- [Node.js](https://nodejs.org/)
- [PostgreSQL](https://www.postgresql.org/)
- [CompileDaemon](https://github.com/githubnemo/CompileDaemon)

## Getting Started

Clone the repository:
```bash
git clone https://github.com/pxiaccount/auth-golang-react.git
cd auth-golang-react
```

### Backend Setup

1. Change directory to backend:
```bash
cd auth-golang-react/backend/
```

2. Create a `.env` file in the backend directory:
```env
PORT=8080                  # Your backend port
SECRET="your-secret-key"   # JWT secret key
DATABASE="host=localhost user=postgres password=yourpassword dbname=authdb port=5432 sslmode=disable"
```

3. Start [PostgreSQL](https://www.postgresql.org/) server

4. Run the backend server with [CompileDaemon](https://github.com/githubnemo/CompileDaemon):
```bash
compiledaemon -command="./auth-golang-react"
```

### React

1. Change directory to frontend:
```bash
cd auth-golang-react/frontend
```

2. Install dependencies:
```bash
npm i
```

3. Start the development server:
```bash
npm run dev
```