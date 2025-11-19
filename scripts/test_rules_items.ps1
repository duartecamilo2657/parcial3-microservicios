Write-Host "========== PRUEBAS DE REGLAS DE NEGOCIO (ITEMS) =========="

$baseUrl = "http://localhost:8051/items"

function Test-Request($body) {

    $json = ConvertTo-Json $body -Depth 10

    try {
        return Invoke-WebRequest `
            -Uri $baseUrl `
            -Method POST `
            -Body $json `
            -ContentType "application/json" `
            -UseBasicParsing
    }
    catch {
        return $_.Exception.Response
    }
}


# TEST 1: name vacío
Write-Host "`n[TEST 1] Validar que 'name' no puede ser vacio..."
$response = Test-Request @{ name = ""; value = 10 }

if ($response.StatusCode -eq 400) {
    Write-Host "[OK] El servicio rechazo un name vacio."
} else {
    Write-Host "[ERROR] El servicio acepto un name vacio (NO DEBERIA)."
}

# TEST 2: value vacio (value=0)
Write-Host "`n[TEST 2] Validar que 'value' no puede ser vacio..."
$response = Test-Request @{ name = "Item X"; value = 0 }

if ($response.StatusCode -eq 400) {
    Write-Host "[OK] El servicio rechazo un value vacio."
} else {
    Write-Host "[ERROR] El servicio acepto un value vacio (NO DEBERIA)."
}

# TEST 3: value menor o igual a cero
Write-Host "`n[TEST 3] Validar que 'value' debe ser mayor que 0..."
$response = Test-Request @{ name = "Item X"; value = -5 }

if ($response.StatusCode -eq 400) {
    Write-Host "[OK] El servicio rechazo un value menor o igual a cero."
} else {
    Write-Host "[ERROR] El servicio acepto un value menor o igual a cero (NO DEBERIA)."
}

# TEST 4: caso valido
Write-Host "`n[TEST 4] Validar creacion correcta..."
$response = Test-Request @{ name = "Valid Item"; value = 99 }

if ($response.StatusCode -eq 201) {
    Write-Host "[OK] El servicio creo un item valido."
} else {
    Write-Host "[ERROR] El servicio NO acepto un item valido."
}

Write-Host "`n========== FIN DE PRUEBAS =========="
