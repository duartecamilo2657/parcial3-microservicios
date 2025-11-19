Write-Host "========== INICIANDO PRUEBAS CRUD (MASCOTAS) ==========" -ForegroundColor Cyan

# BASE URLs para cada microservicio
$createUrl = "http://localhost:8071/mascotas"
$readUrl   = "http://localhost:8072/mascotas"
$updateUrl = "http://localhost:8073/mascotas"
$deleteUrl = "http://localhost:8074/mascotas"

# -------------------------------
# 1. CREATE
# -------------------------------
Write-Host "`n[CREATE] Creando mascota..." -ForegroundColor Yellow

$body = '{"name":"Azul","age":2}'
Invoke-RestMethod -Method POST -Uri $createUrl -ContentType "application/json" -Body $body

Write-Host "Mascota creada correctamente." -ForegroundColor Green

# -------------------------------
# 2. OBTENER ID DE LA ÚLTIMA MASCOTA
# -------------------------------
Write-Host "`n[INFO] Obteniendo ID de la mascota recién creada..." -ForegroundColor Yellow

$mascotas = Invoke-RestMethod -Method GET -Uri $readUrl

if ($mascotas.Length -eq 0) {
    throw "ERROR: No hay mascotas después del CREATE."
}

# Mongo retorna en orden → el último es el que acabamos de crear
$lastMascota = $mascotas[-1]
$id = $lastMascota._id

if (-not $id) {
    throw "ERROR: No se pudo obtener el ID de la última mascota."
}

Write-Host "ID obtenido = $id" -ForegroundColor Green

# -------------------------------
# 3. UPDATE
# -------------------------------
Write-Host "`n[UPDATE] Actualizando mascota..." -ForegroundColor Yellow

$updateBody = '{"name":"Lola","age":4}'
Invoke-RestMethod -Method PUT -Uri "$updateUrl/$id" -ContentType "application/json" -Body $updateBody

Write-Host "Mascota actualizada correctamente." -ForegroundColor Green

# -------------------------------
# 4. DELETE
# -------------------------------
Write-Host "`n[DELETE] Eliminando mascota..." -ForegroundColor Yellow

Invoke-RestMethod -Method DELETE -Uri "$deleteUrl/$id"

Write-Host "Mascota eliminada correctamente." -ForegroundColor Green

# -------------------------------
# 5. VERIFY DELETE
# -------------------------------
Write-Host "`n[VERIFY] Verificando eliminación..." -ForegroundColor Yellow

$mascotasAfter = Invoke-RestMethod -Method GET -Uri $readUrl

$exists = $mascotasAfter | Where-Object { $_._id -eq $id }

if ($exists) {
    throw "ERROR: La mascota con ID $id todavía existe."
}

Write-Host "Verificación exitosa: la mascota ya no existe." -ForegroundColor Green

Write-Host "========== TODAS LAS PRUEBAS CRUD (MASCOTAS) PASARON ==========" -ForegroundColor Cyan
