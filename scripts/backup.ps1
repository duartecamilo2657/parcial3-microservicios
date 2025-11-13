# backup.ps1
# Realiza un respaldo de la base de datos MongoDB que corre en Docker

# Fecha actual para nombrar el backup
$DATE = Get-Date -Format "yyyyMMdd-HHmm"

# Crear carpeta local si no existe
$backupDir = "backups\backup-$DATE"
New-Item -ItemType Directory -Force -Path $backupDir | Out-Null

# Ejecutar mongodump dentro del contenedor mongo
docker exec parcial_mongo mongodump --uri="mongodb://admin:admin123@mongo:27017" --out="/data/db/backup-$DATE"

# Copiar el backup del contenedor a tu máquina local
docker cp parcial_mongo:/data/db/backup-$DATE .\$backupDir

Write-Host "✅ Backup guardado en $backupDir"
