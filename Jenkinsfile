@Library('CICD_Sonarqube-NonEP-lib') _
pipeline {
    agent { label 'golang'}
    stages {
        stage('Golang') {
            steps {
                sonarScan("internal, pkg, tests, main.go, coverage.out, test-report.out") 
            }
        }
        stage('Checkmarx') {
            steps {
                step([$class: 'SafeCheckmarxBuilder', failOn: 'low', additionalExcludes: 'vendor, bin'])
            }
        }
    }
}
