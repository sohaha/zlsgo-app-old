pipeline {
  agent any
  stages {
    stage('setup') {
      steps {
        checkout([
          $class: 'GitSCM',
          branches: [[name: env.GIT_BUILD_REF]],
          userRemoteConfigs: [[
            url: env.GIT_REPO_URL,
            credentialsId: env.CREDENTIALS_ID
          ]]])
        }
      }
      stage('Get dependencies') {
        steps {
          sh 'GOPROXY=https://goproxy.cn go mod tidy'
        }
      }
      stage('execute') {
        parallel {
          stage('Test') {
            steps {
              sh 'go test -v ./...'
            }
          }
          stage('Build') {
            steps {
              sh 'export NoChina=1 && curl -L https://raw.githubusercontent.com/sohaha/zzz/master/install.sh | bash'
              sh 'zzz build -P --os win,mac,linux --out build'
              sh 'ls'
              archiveArtifacts(artifacts: 'build/', allowEmptyArchive: true, onlyIfSuccessful: false)
            }
          }
        }
      }
      stage('End') {
        steps {
          echo 'succeed'
        }
      }
    }
  }