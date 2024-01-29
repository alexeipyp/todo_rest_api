pipeline {
    agent { docker { image 'golang:1.21.6-alpine3.19' } }
    stages {
        stage('build') {
            steps {
                sh 'go version'
            }
        }
    }
}