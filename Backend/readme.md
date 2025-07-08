# 🚀 Backend API - Go

Un proyecto backend robusto desarrollado en Go que proporciona una API RESTful con conexión a base de datos PostgreSQL utilizando GORM. Incluye funcionalidad para sincronizar datos de stocks desde APIs externas.

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
- [Troubleshooting](#-troubleshooting)

## ✨ Características

- 🔧 **Go 1.24.4** - Lenguaje de programación moderno y eficiente
- 🗄️ **PostgreSQL** - Base de datos relacional robusta
- 🔌 **GORM** - ORM para Go con soporte completo para PostgreSQL
- 🔐 **Variables de Entorno** - Configuración segura con godotenv
- 🚀 **Arquitectura Modular** - Código organizado y mantenible
- 📦 **Gestión de Dependencias** - Go modules para dependencias
- 🔄 **Sincronización de APIs Externas** - Integración con servicios externos
- 🏗️ **Patrón Repository** - Separación de lógica de negocio y datos
- 🛡️ **Manejo de Errores** - Gestión robusta de errores y excepciones
- ⚡ **Transacciones de Base de Datos** - Operaciones atómicas y seguras
- 🔍 **Health Check Avanzado** - Monitoreo de servicios y entorno

## 🛠️ Tecnologías

| Tecnología | Versión | Descripción |
|------------|---------|-------------|
| **Go** | 1.24.4 | Lenguaje de programación |
| **GORM** | 1.30.0 | ORM para Go |
| **PostgreSQL** | - | Base de datos |
| **godotenv** | 1.5.1 | Gestión de variables de entorno |
| **net/http** | - | Servidor HTTP estándar de Go |
| **pgx/v5** | 5.7.5 | Driver PostgreSQL nativo |

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
# Configuración del servidor
PORT_SERVER=8080

# Configuración de base de datos
PORT_DB=5432
SERVER_DB=localhost
USER_DB=tu_usuario
PASS_DB=tu_password
NAME_DB=nombre_base_datos

# Configuración de API externa
EXTERNAL_API=https://api.externa.com/stocks
TOKEN_EXTERNAL_API=tu_token_aqui

# Configuración del entorno
ENVIRONMENT=DEV|PRD
```

### Base de Datos

1. Asegúrate de que PostgreSQL esté ejecutándose
2. Crea una base de datos para el proyecto
3. Configura las credenciales en el archivo `.env`
4. Las tablas se crearán automáticamente al ejecutar la aplicación
5. Las migraciones se ejecutarán automáticamente

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

### Verificar Sintaxis

```bash
go vet ./...
```

## 📁 Estructura del Proyecto

```
Backend/
├── cmd/
│   └── main.go                    # Punto de entrada de la aplicación
├── internal/
│   ├── API_EXTERNAL/             # Módulo para integración con API externa
│   │   ├── model.go              # Modelos de datos para stocks
│   │   ├── dto.go                # Data Transfer Objects
│   │   ├── repository.go         # Operaciones de base de datos
│   │   └── handler.go            # Handlers HTTP para endpoints
│   └── health/                   # Módulo de health check
│       └── handler.go            # Handler de estado del servidor
├── Router/
│   └── router.go                 # Configuración de rutas HTTP
├── db/
│   ├── gorm.go                   # Configuración de base de datos
├── go.mod                        # Dependencias del proyecto
├── go.sum                        # Checksums de dependencias
├── .env-example                  # Ejemplo de variables de entorno
├── .gitignore                    # Archivos ignorados por Git
└── readme.md                     # Documentación del proyecto
```

## 🔌 API Endpoints

### Endpoints Disponibles

| Método | Endpoint | Descripción | Estado |
|--------|----------|-------------|--------|
 `GET` | `/health` | Verificar estado del servidor | ✅ Implementado |
| `POST` | `/sync-stocks` | Sincronizar stocks desde API externa | ✅ Implementado |
| `GET` | `/stocks` | Obtener listado de stocks | ✅ Implementado |
|


## 🔧 Comandos Útiles

```bash
# Ejecutar el servidor
go run cmd/main.go

# Construir binario
go build -o bin/backend cmd/main.go

# Limpiar dependencias
go mod tidy

# Ejecutar tests
go test ./...

# Verificar sintaxis
go vet ./...

# Ver dependencias
go list -m all

# Ejecutar con variables de entorno específicas
ENVIRONMENT=DEV go run cmd/main.go
```

## 🚨 Troubleshooting

### Problemas Comunes

**Error: "Error cargando archivo .env"**
```bash
# Solución: Crear archivo .env
cp .env-example .env
# Editar con tus configuraciones
```

**Error: "conexión a base de datos no inicializada"**
```bash
# Verificar variables de entorno
echo $SERVER_DB $USER_DB $PASS_DB $PORT_DB
# Asegurar que PostgreSQL esté ejecutándose
```

**Error: "Faltan variables de entorno para la API externa"**
```bash
# Configurar en .env
EXTERNAL_API=https://tu-api.com/stocks
TOKEN_EXTERNAL_API=tu_token_aqui
```

**Error: "Error de la API externa: 502"**
- Verificar conectividad a internet
- Validar URL y token de la API externa
- Revisar logs del servidor para más detalles

### Monitoreo

Usar el endpoint `/health` para:
- Verificar estado del servidor
- Monitorear conexión a base de datos
- Validar configuración de entorno
- Obtener tiempo de actividad

<div align="center">
  <p>Desarrollado con ❤️ por <strong>Cesar Bejarano</strong></p>
</div>