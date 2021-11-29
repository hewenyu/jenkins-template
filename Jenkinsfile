pipeline {
    agent {
        docker {
            image 'golang:1.16'
//             args '-u root'
        }
    }
    environment {
        def BUILDVERSION = sh(script: "echo `date '+%Y-%m-%d-%H-%M-%S'`", returnStdout: true).trim()
    }
    stages {
        stage('check') {
            steps {
                checkout scm
            }
        }
        stage('build') {
            steps {
                sh "export GOPROXY=https://goproxy.io,direct && env GOOS=linux GOARCH=amd64 go build -o server"
                sh "mkdir -p archive"
                sh "mv server ./archive/"
                sh "mv README.md ./archive/"
                sh "mv *.yaml ./archive/"
            }
        }
        stage ('zip 打包') {
            steps {
                script{
                    zip archive: true, dir: "archive", glob: '', zipFile: "DEMO-${BUILDVERSION}.zip"
                }
            }
        }
    }
    post {
        cleanup {
            cleanWs()
        }
    }

}
