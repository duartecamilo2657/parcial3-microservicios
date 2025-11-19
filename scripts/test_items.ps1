Write-Host "========== INICIANDO PRUEBAS CRUD ==========" -ForegroundColor Cyan

# BASE URLs
$createUrl = "http://localhost:8081/items"
$readUrl   = "http://localhost:8082/items"
$updateUrl = "http://localhost:8083/items"
$deleteUrl = "http://localhost:8084/items"

# -------------------------------
# 1. CREATE
# -------------------------------
Write-Host "`n[CREATE] Creando item..." -ForegroundColor Yellow

$body = '{"name":"Laptop Jenkins","value":2000}'
Invoke-RestMethod -Method POST -Uri $createUrl -ContentType "application/json" -Body $body

Write-Host "Item creado correctamente." -ForegroundColor Green

# -------------------------------
# OBTENER ID DEL ÚLTIMO ITEM
# -------------------------------
Write-Host "`n[INFO] Obteniendo ID del item recién creado..." -ForegroundColor Yellow

$items = Invoke-RestMethod -Method GET -Uri $readUrl

if ($items.Length -eq 0) {
    throw "ERROR: No hay items después del CREATE."
}

# Mongo retorna items en orden → el último es el recién creado
$lastItem = $items[-1]
$id = $lastItem._id

if (-not $id) {
    throw "ERROR: No se pudo obtener el ID del último item."
}

Write-Host "ID obtenido = $id" -ForegroundColor Green

# -------------------------------
# 3. UPDATE
# -------------------------------
Write-Host "`n[UPDATE] Actualizando item..." -ForegroundColor Yellow

$updateBody = '{"name":"Laptop Jenkins Pro","value":2500}'
Invoke-RestMethod -Method PUT -Uri "$updateUrl/$id" -ContentType "application/json" -Body $updateBody

Write-Host "Item actualizado correctamente." -ForegroundColor Green

# -------------------------------
# 4. DELETE
# -------------------------------
Write-Host "`n[DELETE] Eliminando item..." -ForegroundColor Yellow

Invoke-RestMethod -Method DELETE -Uri "$deleteUrl/$id"

Write-Host "Item eliminado correctamente." -ForegroundColor Green

# -------------------------------
# 5. VERIFY DELETE
# -------------------------------
Write-Host "`n[VERIFY] Verificando eliminación..." -ForegroundColor Yellow

$itemsAfter = Invoke-RestMethod -Method GET -Uri $readUrl

$exists = $itemsAfter | Where-Object { $_._id -eq $id }

if ($exists) {
    throw "ERROR: El item con ID $id todavía existe."
}

Write-Host "Verificación exitosa: el item ya no existe." -ForegroundColor Green

Write-Host "========== TODAS LAS PRUEBAS CRUD PASARON ==========" -ForegroundColor Cyan
