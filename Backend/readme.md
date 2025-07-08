# 🚀 Backend API - Go

Un proyecto backend robusto desarrollado en Go que proporciona una API RESTful con conexión a base de datos PostgreSQL utilizando GORM.

## 👨‍💻 Desarrollador

<div align="center">
  <img src="https://media.licdn.com/dms/image/v2/D4E03AQHTsTzH2fM2iw/profile-displayphoto-shrink_400_400/profile-displayphoto-shrink_400_400/0/1718281288427?e=1757548800&v=beta&t=RFeU9SYo_SwcVVwuIQaYghk8Lx0yKtxMmpfsnGEjWU4" alt="Cesar Bejarano" width="150" style="border-radius: 50%;">
  
  ### Cesar Bejarano
  **Desarrollador de software**
  
  [![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/cesaraugustobejaranoparra/)
  [![GitHub](https://img.shields.io/badge/GitHub-100000?style=for-the-badge&logo=github&logoColor=white)](https://github.com/C3S4RBP)
  
  Desarrollador apasionado por crear soluciones escalables y eficientes.
</div>

---

## 📋 Tabla de Contenidos

- [Características](#-características)
- [Tecnologías](#-tecnologías)
- [Instalación](#-instalación)
- [Configuración](#-configuración)
- [Uso](#-uso)
- [Estructura del Proyecto](#-estructura-del-proyecto)
- [API Endpoints](#-api-endpoints)

## ✨ Características

- 🔧 **Go 1.24.4** - Lenguaje de programación moderno y eficiente
- 🗄️ **PostgreSQL** - Base de datos relacional robusta
- 🔌 **GORM** - ORM para Go con soporte completo para PostgreSQL
- 🔐 **Variables de Entorno** - Configuración segura con godotenv
- 🚀 **Arquitectura Modular** - Código organizado y mantenible
- 📦 **Gestión de Dependencias** - Go modules para dependencias

## 🛠️ Tecnologías

| Tecnología | Versión | Descripción |
|------------|---------|-------------|
| **Go** | 1.24.4 | Lenguaje de programación |
| **GORM** | 1.30.0 | ORM para Go |
| **PostgreSQL** | - | Base de datos |
| **godotenv** | 1.5.1 | Gestión de variables de entorno |

## 🚀 Instalación

### Prerrequisitos

- Go 1.24.4 o superior
- PostgreSQL instalado y ejecutándose
- Git

### Pasos de Instalación

1. **Clonar el repositorio**
   ```bash
   git clone https://github.com/C3S4RBP/recruitmentChallenge.git
   cd recruitmentChallenge/Backend
   ```

2. **Instalar dependencias**
   ```bash
   go mod download
   ```

3. **Configurar variables de entorno**
   ```bash
   cp .env-example .env
   # Editar el archivo .env con tus configuraciones
   ```

## ⚙️ Configuración

### Variables de Entorno

Crea un archivo `.env` basado en `.env-example` con las siguientes variables:

```env
PORT_SERVER=8080
PORT_DB=5432
SERVER_DB=localhost
USER_DB=tu_usuario
PASS_DB=tu_password
EXTERNAL_API=https://api.externa.com
TOKEN_EXTERNAL_API=tu_token
```

### Base de Datos

1. Asegúrate de que PostgreSQL esté ejecutándose
2. Crea una base de datos para el proyecto
3. Configura las credenciales en el archivo `.env`

## 🎯 Uso

### Ejecutar el Proyecto

```bash
# Desde el directorio Backend
go run cmd/main.go
```

### Construir el Proyecto

```bash
go build -o bin/backend cmd/main.go
```

### Ejecutar Tests

```bash
go test ./...
```

## 📁 Estructura del Proyecto

```
Backend/
├── cmd/
│   └── main.go          # Punto de entrada de la aplicación
├── db/
│   └── gorm.go          # Configuración de base de datos
├── go.mod               # Dependencias del proyecto
├── go.sum               # Checksums de dependencias
├── .env-example         # Ejemplo de variables de entorno
├── .gitignore           # Archivos ignorados por Git
└── README.md            # Este archivo
```

## 🔌 API Endpoints

> **Nota:** Los endpoints están en desarrollo. Esta sección se actualizará conforme se implementen.

### Endpoints Disponibles

| Método | Endpoint | Descripción | Estado |
|--------|----------|-------------|--------|
| `GET` | `/health` | Verificar estado del servidor | 🚧 En desarrollo |


<div align="center">
  <p>Desarrollado con ❤️ por <strong>Cesar Bejarano</strong></p>
</div>