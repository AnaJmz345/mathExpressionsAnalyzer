# Math Expression Analyzer

This was a school project consisting of a client-server application that analyzes mathematical expressions and validates them in real time using a Pushdown Automaton (PDA). The backend is written in Go, and the frontend is built with React + Vite.

## This project was made by:
- César Morán Macías (@cesarmrn)
- Ana Paola Jiménez Sedano (@AnaJmz345)
- Ana Elena Velasco García (@Anaaveg17)

---
## What does it validate?
- Balanced parenthesis `()` and brackets `[]`
- You cannot nest `[]` inside `()`, but you can nest `()` inside `[]`
-Supported operators:
  - `*` for multiplications
  - `**` for power (exponentiation) 
  - `/` for divisions
  - `+` for additions
  - `-` for subtractions
- Accepts real numbers `(0, 0.1, 1.22, -1.23)`
- Any other symbol outside the supported ones will be rejected
---
## Previous requirements

### Needed software

#### Backend (Go)
- Go 1.21 o superior
- Internet connection (only needed the first time to install dependencies)

#### Frontend (React + Vite)
- Node.js 18.x or superior
- npm (Node Package Manager)

---

## Steps to run the project

### 1. Clone or download the project

```bash
git clone https://github.com/CodersSquad-Classes/ct-math-exp-analyzer-ui-cesaryana
cd ct-math-exp-analyzer-ui-cesaryana
```

## Start the server (Go)

In a terminal open the folder `backend`.

Execute:

```bash
go run math-exp-analyzer.go
```
If you encounter issues, try installing the required libraries:

```bash
go mod init
go get github.com/gin-gonic/gin
go get github.com/gin-contrib/cors
```

## Start the backend

In another terminal open the folder `frontend` and go to the folder `ui` with cd .

Execute:

```bash
npm install
npm run dev
```
Open the browser in the next link
```bash
http://localhost:5173
```

*The interface allows you to write math expressions and see their validation results in real time.*


