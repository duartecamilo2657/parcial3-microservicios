Write-Host "========== INICIANDO PRUEBAS CRUD (PHONES) ==========" -ForegroundColor Cyan

# BASE URLs de cada microservicio
$createUrl = "http://localhost:8051/phones"
$readUrl   = "http://localhost:8052/phones"
$updateUrl = "http://localhost:8053/phones"
$deleteUrl = "http://localhost:8054/phones"

# -------------------------------
# 1. CREATE
# -------------------------------
Write-Host "`n[CREATE] Creando teléfono..." -ForegroundColor Yellow

$body = '{"brand":"Samsung","price":200}'
Invoke-RestMethod -Method POST -Uri $createUrl -ContentType "application/json" -Body $body

Write-Host "Teléfono creado correctamente." -ForegroundColor Green

# -------------------------------
# 2. OBTENER ID DEL ÚLTIMO PHONE
# -------------------------------
Write-Host "`n[INFO] Obteniendo ID del teléfono recién creado..." -ForegroundColor Yellow

$phones = Invoke-RestMethod -Method GET -Uri $readUrl

if ($phones.Length -eq 0) {
    throw "ERROR: No hay teléfonos después del CREATE."
}

# Mongo → el último suele ser el recién creado
$lastPhone = $phones[-1]
$id = $lastPhone._id

if (-not $id) {
    throw "ERROR: No se pudo obtener el ID del último teléfono."
}

Write-Host "ID obtenido = $id" -ForegroundColor Green

# -------------------------------
# 3. UPDATE
# -------------------------------
Write-Host "`n[UPDATE] Actualizando teléfono..." -ForegroundColor Yellow

$updateBody = '{"brand":"Samsung Pro","price":500}'
Invoke-RestMethod -Method PUT -Uri "$updateUrl/$id" -ContentType "application/json" -Body $updateBody

Write-Host "Teléfono actualizado correctamente." -ForegroundColor Green

# -------------------------------
# 4. DELETE
# -------------------------------
Write-Host "`n[DELETE] Eliminando teléfono..." -ForegroundColor Yellow

Invoke-RestMethod -Method DELETE -Uri "$deleteUrl/$id"

Write-Host "Teléfono eliminado correctamente." -ForegroundColor Green

# -------------------------------
# 5. VERIFY DELETE
# -------------------------------
Write-Host "`n[VERIFY] Verificando eliminación..." -ForegroundColor Yellow

$phonesAfter = Invoke-RestMethod -Method GET -Uri $readUrl

$exists = $phonesAfter | Where-Object { $_._id -eq $id }

if ($exists) {
    throw "ERROR: El teléfono con ID $id todavía existe."
}

Write-Host "Verificación exitosa: el teléfono ya no existe." -ForegroundColor Green

Write-Host "========== TODAS LAS PRUEBAS CRUD (PHONES) PASARON ==========" -ForegroundColor Cyan
