# 🧩 Parcial 3 - Microservicios en Go con MongoDB y Docker

Este proyecto implementa un **CRUD distribuido en microservicios** usando **Golang**, **MongoDB**, **Docker Compose** y **GitHub Actions (CI/CD)**.

Cada operación del CRUD (Create, Read, Update, Delete) se maneja como un **servicio independiente**, comunicándose con una misma base de datos MongoDB.

---

## 📦 Estructura del Proyecto

parcial3-microservicios/
│
├── services/
│ ├── create/ # POST - Crear un item
│ ├── read/ # GET - Leer todos los items
│ ├── update/ # PUT - Actualizar un item
│ └── delete/ # DELETE - Eliminar un item
│
├── docker-compose.yml
├── .github/workflows/ci.yml # CI con GitHub Actions
└── README.md

yaml
Copiar código

---

## ⚙️ Tecnologías Usadas

| Tecnología | Uso |
|-------------|-----|
| **Golang 1.25** | Lógica principal de cada microservicio |
| **MongoDB 6** | Base de datos para persistencia |
| **Docker Compose** | Orquestación de todos los servicios |
| **GitHub Actions** | Integración Continua (CI) |
| **Postman / cURL** | Pruebas del CRUD |

---

## 🚀 Ejecución local

### 1️⃣ Clonar el repositorio

```
git clone https://github.com/tuusuario/parcial3-microservicios.git
cd parcial3-microservicios
```
### 2️⃣ Levantar todo el entorno con Docker
```
Copiar código
docker compose up -d
```
Esto inicia:

MongoDB (puerto 27017)

create_service (puerto 8081)

read_service (puerto 8082)

update_service (puerto 8083)

delete_service (puerto 8084)

Verifica con:

```
Copiar código
docker compose ps
```
## 🧪 Pruebas del CRUD (cURL)
### ➕ Crear un item
```
curl -X POST http://localhost:8081/items \
  -H "Content-Type: application/json" \
  -d '{"name":"Laptop","value":2000}'
 ```
### 📖 Leer todos los items
```
curl http://localhost:8082/items
```
### 🔁 Actualizar un item
```
curl -X PUT http://localhost:8083/items/<ID> \
  -H "Content-Type: application/json" \
  -d '{"name":"Laptop Pro","value":2500}'
  ```
### ❌ Eliminar un item
```
curl -X DELETE http://localhost:8084/items/<ID>
```
## 🧰 Pruebas con Postman
Puedes importar este archivo en Postman:

```
Parcial3-Microservicios.postman_collection.json
```
Incluye las 4 operaciones:

Operación	Método	URL
Create	POST	http://localhost:8081/items
Read	GET	http://localhost:8082/items
Update	PUT	http://localhost:8083/items/:id
Delete	DELETE	http://localhost:8084/items/:id

## ⚡ CI/CD con GitHub Actions
Archivo: .github/workflows/ci.yml

Este workflow:

1. Se ejecuta en cada push o pull request a main.

2. Crea un contenedor MongoDB para los tests.

3. Ejecuta los tests unitarios e integraciones para cada servicio:

    create

    read

    update

    delete

Si los tests pasan, construye todas las imágenes Docker.

## ✅ Pruebas unitarias locales
Ejecuta los tests de cada servicio:

```
cd services/create
go test ./... -v -cover
```
Repite para read, update, delete.

## 🧱 Construir manualmente las imágenes
```
docker compose build
```
O solo un servicio:

```
docker compose build create
```
Aquí está lo que falta corregido:

## 🧹 Detener y limpiar contenedores
```
docker compose down -v
```

---

## 📊 Resultados esperados

| Servicio | Puerto | Endpoint | Ejemplo de respuesta |
|----------|--------|----------|----------------------|
| Create | 8081 | POST /items | `{"name":"Laptop","value":2000}` |
| Read | 8082 | GET /items | `[{"_id":"...","name":"Laptop","value":2000}]` |
| Update | 8083 | PUT /items/:id | `{"status":"updated"}` |
| Delete | 8084 | DELETE /items/:id | `{"status":"deleted"}` |

---

## 🧾 Autor

**Camilo Duarte Rivera**  
Universidad EAM  
Parcial #3 — Sistemas Operativos  
Noviembre 2025
