Write-Host "========== PRUEBAS DE REGLAS DE NEGOCIO (ITEMS) =========="

$baseUrl = "http://localhost:8051/items"

function Test-Request($body) {
    return Invoke-WebRequest -Uri $baseUrl -Method POST -Body ($body | ConvertTo-Json) -ContentType "application/json" -SkipHttpErrorCheck
}

# TEST 1: name vacío
Write-Host "`n[TEST 1] Validar que 'name' no puede ser vacío..."
$response = Test-Request @{ name = ""; value = 10 }

if ($response.StatusCode -eq 400) {
    Write-Host "✔ CORRECTO: El servicio rechazó un name vacío."
} else {
    Write-Host "❌ ERROR: El servicio aceptó un name vacío (NO DEBERÍA)."
}

# TEST 2: value vacío (value=0)
Write-Host "`n[TEST 2] Validar que 'value' no puede ser vacío..."
$response = Test-Request @{ name = "Item X"; value = 0 }

if ($response.StatusCode -eq 400) {
    Write-Host "✔ CORRECTO: El servicio rechazó un value vacío."
} else {
    Write-Host "❌ ERROR: El servicio aceptó un value vacío (NO DEBERÍA)."
}

# TEST 3: value <= 0
Write-Host "`n[TEST 3] Validar que 'value' debe ser mayor que 0..."
$response = Test-Request @{ name = "Item X"; value = -5 }

if ($response.StatusCode -eq 400) {
    Write-Host "✔ CORRECTO: El servicio rechazó un value <= 0."
} else {
    Write-Host "❌ ERROR: El servicio aceptó un value <= 0 (NO DEBERÍA)."
}

# TEST 4: caso válido
Write-Host "`n[TEST 4] Validar creación correcta..."
$response = Test-Request @{ name = "Valid Item"; value = 99 }

if ($response.StatusCode -eq 201) {
    Write-Host "✔ CORRECTO: El servicio creó un item válido."
} else {
    Write-Host "❌ ERROR: El servicio NO aceptó un item válido."
}

Write-Host "`n========== FIN DE PRUEBAS =========="
