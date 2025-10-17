pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                // Get some code from a GitHub repository
                git url: 'https://github.com/Stunner59/go-web-app-devops.git', branch: 'stage'
            }
        }

        // stage('Build'){
        //     steps{
        //         sh 'go build -o app .'
        //     }
        // }

        stage('Build the docker image'){
            steps{
                sh 'docker build -t go-app:$BUILD_NUMBER .'
            }
        }

        stage('Run the Image'){
            steps{
                sh '''
                // container_id=$(docker ps --filter ancestor=go-app:$((BUILD_NUMBER-1)) | awk 'NR==2 { print $1 }') || true
                // docker rm -f $container_id || true
                // docker rmi go-app:$((BUILD_NUMBER-1)) || true
                docker run -d -p 8040:8040 go-app:$BUILD_NUMBER
                '''
            }
        }
    }
}

