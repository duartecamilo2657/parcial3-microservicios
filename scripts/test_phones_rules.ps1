Write-Host "========== PRUEBAS DE REGLAS DE NEGOCIO (PHONES) =========="

$createUrl = "http://localhost:8061/phones"

# -------------------------------------------------------------
# TEST 1: brand vacío
# -------------------------------------------------------------
Write-Host "`n[TEST 1] Validar que 'brand' no puede ser vacío..."
try {
    Invoke-RestMethod -Method POST -Uri $createUrl -ContentType "application/json" `
        -Body '{ "brand": "", "price": 300 }'

    Write-Host "❌ ERROR: El servicio aceptó un brand vacío (NO DEBERÍA)."
}
catch {
    Write-Host "✅ CORRECTO: El servicio rechazó un brand vacío."
}

# -------------------------------------------------------------
# TEST 2: price vacío
# -------------------------------------------------------------
Write-Host "`n[TEST 2] Validar que 'price' no puede ser vacío..."
try {
    Invoke-RestMethod -Method POST -Uri $createUrl -ContentType "application/json" `
        -Body '{ "brand": "Samsung", "price": "" }'

    Write-Host "❌ ERROR: El servicio aceptó un price vacío (NO DEBERÍA)."
}
catch {
    Write-Host "✅ CORRECTO: El servicio rechazó un price vacío."
}

# -------------------------------------------------------------
# TEST 3: price menor o igual a 0
# -------------------------------------------------------------
Write-Host "`n[TEST 3] Validar que 'price' debe ser mayor que 0..."
try {
    Invoke-RestMethod -Method POST -Uri $createUrl -ContentType "application/json" `
        -Body '{ "brand": "Samsung", "price": 0 }'

    Write-Host "❌ ERROR: El servicio aceptó price=0 (NO DEBERÍA)."
}
catch {
    Write-Host "✅ CORRECTO: El servicio rechazó un price <= 0."
}

# -------------------------------------------------------------
# TEST 4: CREATE válido (control)
# -------------------------------------------------------------
Write-Host "`n[TEST 4] Validar envío correcto para asegurar que CREATE funciona..."
$response = Invoke-RestMethod -Method POST -Uri $createUrl -ContentType "application/json" `
    -Body '{ "brand": "Motorola", "price": 350 }'

if ($response) {
    Write-Host "✅ CORRECTO: El servicio creó un teléfono válido."
} else {
    Write-Host "❌ ERROR: El servicio no aceptó un teléfono válido."
}

Write-Host "`n========== FIN DE PRUEBAS =========="
