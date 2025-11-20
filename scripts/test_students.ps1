Write-Host "========== INICIANDO PRUEBAS CRUD (STUDENTS) ==========" -ForegroundColor Cyan

# BASE URLs para cada microservicio
$createUrl = "http://localhost:8091/students"
$readUrl   = "http://localhost:8092/students"
$updateUrl = "http://localhost:8093/students"
$deleteUrl = "http://localhost:8094/students"

# -------------------------------
# 1. CREATE
# -------------------------------
Write-Host "`n[CREATE] Creando estudiante..." -ForegroundColor Yellow

$body = '{"age":20}'
Invoke-RestMethod -Method POST -Uri $createUrl -ContentType "application/json" -Body $body

Write-Host "Estudiante creado correctamente." -ForegroundColor Green

# -------------------------------
# 2. OBTENER ID DEL ÚLTIMO STUDENT
# -------------------------------
Write-Host "`n[INFO] Obteniendo ID del estudiante recién creado..." -ForegroundColor Yellow

$students = Invoke-RestMethod -Method GET -Uri $readUrl

if ($students.Length -eq 0) {
    throw "ERROR: No hay estudiantes después del CREATE."
}

# Mongo suele ordenar por fecha → el último es el recién creado
$lastStudent = $students[-1]
$id = $lastStudent._id

if (-not $id) {
    throw "ERROR: No se pudo obtener el ID del último estudiante."
}

Write-Host "ID obtenido = $id" -ForegroundColor Green

# -------------------------------
# 3. UPDATE
# -------------------------------
Write-Host "`n[UPDATE] Actualizando estudiante..." -ForegroundColor Yellow

$updateBody = '{"name":"Juan Pérez","age":21}'
Invoke-RestMethod -Method PUT -Uri "$updateUrl/$id" -ContentType "application/json" -Body $updateBody

Write-Host "Estudiante actualizado correctamente." -ForegroundColor Green

# -------------------------------
# 4. DELETE
# -------------------------------
Write-Host "`n[DELETE] Eliminando estudiante..." -ForegroundColor Yellow

Invoke-RestMethod -Method DELETE -Uri "$deleteUrl/$id"

Write-Host "Estudiante eliminado correctamente." -ForegroundColor Green

# -------------------------------
# 5. VERIFY DELETE
# -------------------------------
Write-Host "`n[VERIFY] Verificando eliminación..." -ForegroundColor Yellow

$studentsAfter = Invoke-RestMethod -Method GET -Uri $readUrl

$exists = $studentsAfter | Where-Object { $_._id -eq $id }

if ($exists) {
    throw "ERROR: El estudiante con ID $id todavía existe."
}

Write-Host "Verificación exitosa: el estudiante ya no existe." -ForegroundColor Green

Write-Host "========== TODAS LAS PRUEBAS CRUD PASARON ==========" -ForegroundColor Cyan
