pipeline {
    agent any
    tools {
        go "go"
    }
    environment {
        GO_VERSION = '1.21'
        DOCKER_IMAGE = 'todo-list-api'
        DOCKER_TAG = "${BUILD_NUMBER}"
    }
    
    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        
        stage('Setup Go') {
            steps {
                sh '''
                    go version
                    go mod tidy
                '''
            }
        }
        
        stage('Test') {
            steps {
                sh '''
                    go test ./... -v -cover
                '''
            }
        }
        
        stage('Build') {
            steps {
                sh '''
                    go build -o bin/todo-list main.go
                '''
            }
        }
        
        stage('Docker Build') {
            steps {
                script {
                    docker.build("${DOCKER_IMAGE}:${DOCKER_TAG}")
                    docker.build("${DOCKER_IMAGE}:latest")
                }
            }
        }
        stage('Debug') {
            steps {
                sh 'which go'
                sh 'go version'
           }
        
        stage('Deploy') {
            when {
                branch 'main'
            }
            steps {
                sh '''
                    docker-compose down || true
                    docker-compose up -d --build
                '''
            }
        }
    }
    
    post {
        always {
            cleanWs()
        }
        success {
            echo 'Pipeline ejecutado exitosamente!'
        }
        failure {
            echo 'Pipeline falló. Revisar logs.'
        }
    }
}