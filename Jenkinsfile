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
            }
        }

    }

    post {
        always {
            echo "Pipeline finalizado. Los contenedores quedan corriendo."
        }
    }
}
