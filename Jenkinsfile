pipeline {
    agent any

    stages {

        stage('Checkout') {
            steps {
                git 'https://github.com/duartecamilo2657/parcial3-microservicios'
            }
        }

        stage('Build Docker Images') {
            steps {
                echo "Construyendo imágenes..."
                sh 'docker compose build'
            }
        }

        stage('Start Services') {
            steps {
                echo "Levantando servicios..."
                sh 'docker compose down -v || true'
                sh 'docker compose up -d'
                sh 'sleep 8'
                sh 'docker ps'
            }
        }

        stage('Run CRUD Tests') {
            steps {
                echo "Ejecutando pruebas CRUD automáticas..."
                powershell '''
                    ./scripts/test_crud.ps1
                '''
            }
        }

    }

    post {
        always {
            echo "Apagando contenedores..."
            sh 'docker compose down -v || true'
        }
    }
}
