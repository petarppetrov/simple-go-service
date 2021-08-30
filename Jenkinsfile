def testFailed = false

pipeline {
    agent any

    tools {
        go 'go1.15.6'
    }
	
    stages {
        stage("CLEAN") {
            failFast true
            steps {
                cleanWs()
            }
        }
        stage('GitHub Pull') {
            steps {
                git branch: 'master', url: 'https://github.com/petarppetrov/simple-go-service.git'
            }
        }
        stage('Unit Tests') {
            steps {
                script {
                    try {
                        echo 'Performing Unit Testing'
                        powershell(script: 'go test ./...')
                    } catch (err) {
                        println(err.getMessage())
                        unstable("The Unit Tests have failed. Please check log.")
                        testFailed = true
                    }
                }
            }
        }
        stage('Build Docker Image') {
            when {
                expression {
                    testFailed == false
                }
            }
            steps {
                powershell(script: 'docker build -t simple-go-service:v$env:BUILD_NUMBER .')
            }
        }
    }
    
    post {
        success {
            echo 'Build is successful. Updating the service image in Kubernetes'
            powershell(script: 'kubectl set image deployments/simple-go-service application=simple-go-service:v$env:BUILD_NUMBER .')
        }
        unstable {
            echo 'There has been a problem. Kubernetes image will not be updated.'
        }
    }
}
