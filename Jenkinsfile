pipeline {
    agent any

    stages {

        stage('Checkout') {
            steps {
                git branch: 'main',
                    url: 'https://github.com/duartecamilo2657/parcial3-microservicios'
            }
        }

        stage('Build Docker Images') {
            steps {
                echo "Construyendo imágenes..."
                powershell 'docker compose build'
            }
        }

        stage('Start Services') {
            steps {
                echo "Levantando servicios..."
                powershell 'docker compose down -v --remove-orphans'
                powershell 'docker compose up -d'
                powershell 'Start-Sleep -Seconds 10'
                powershell 'docker ps'
            }
        }

        stage('Run CRUD Tests') {
            steps {
                echo "Ejecutando pruebas CRUD automáticas (Students)..."
                powershell './scripts/test_students.ps1'

                echo "Ejecutando pruebas CRUD automáticas (Items)..."
                powershell './scripts/test_items.ps1'

                echo "Ejecutando pruebas CRUD automáticas (Mascotas)..."
                powershell './scripts/test_mascotas.ps1'

                echo "Ejecutando pruebas CRUD automáticas (Phones)..."
                powershell './scripts/test_phones.ps1'
            }
        }
        stage('Run Business Rules Tests (Phones)') {
            steps {
                echo "Validando reglas de negocio (Phones)..."
                powershell './scripts/test_phones_rules.ps1'
            }
        }
        stage('Run Business Rules Tests (Items)') {
            steps {
                echo "Validando reglas de negocio (Items)..."
                powershell './scripts/test_rules_items.ps1'
            }
        }

    }

    post {
        always {
            echo "Apagando contenedores..."
            powershell 'docker compose down -v'
        }
    }
}
