# HOWTO.md – Math Expression Analyzer

Este proyecto consiste en una aplicación cliente-servidor que analiza expresiones matemáticas en tiempo real usando un autómata de pila (PDA). El frontend fue desarrollado con React + Vite y el backend en Go.
Este proyecto fue realizado en conjunto por los alumnos:
- César Morán Macías
- Ana Paola Jiménez Sedano
- Ana Elena Velasco García

---

## Requisitos previos

### Software necesario

#### Backend (Go)
- Go 1.21 o superior
- Conexión a internet para descargar dependencias (solo la primera vez)

#### Frontend (React + Vite)
- Node.js 18.x o superior
- npm (Node Package Manager)

---

## Pasos para correr el proyecto

### 1. Clonar o descargar el proyecto

```bash
git clone https://github.com/CodersSquad-Classes/ct-math-exp-analyzer-ui-cesaryana
cd ct-math-exp-analyzer-ui-cesaryana
```

## Iniciar el backend (Go)

Abrir una terminal en la carpeta `backend`.

Ejecutar:

```bash
go run math-exp-analyzer.go
```
Si da problemas el comando anterior, instalar las siguientes librerías:

```bash
go mod init
go get github.com/gin-gonic/gin
go get github.com/gin-contrib/cors
```

## Iniciar el frontend 

Abrir otra terminal en la carpeta `frontend` y posteriormente hacer otro cd a la carpeta dentro de la misma llamada `ui`  .

Ejecutar:

```bash
npm install
npm run dev
```
Abrir en el navegador la siguiente ruta:
```bash
http://localhost:5173
```

*La interfaz permitirá escribir expresiones que serán validadas en tiempo real.*


