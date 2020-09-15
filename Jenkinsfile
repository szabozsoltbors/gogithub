def gv

pipeline {
    agent {
        label 'oc_szs_agent'
    }

    stages {
        // INITIALIZATION
        stage("init_pipeline") {
            steps {
                script {
                    gv = load "Jenkins.groovy"
                    gv.initPipeline()
                }
            }
        }

        // BUILD
        stage("build_app") {
            steps {
                script {
                    gv.buildApp()
                }
            }
        }

        // TEST
        stage("test_app") {
            steps {
                script {
                    gv.testApp()
                }
            }
        }

        // BUILD DOCKER
        stage("build_docker") {
            steps {
                script {
                    gv.buildDocker()
                }
            }
        }

        // DEPLOY DOCKER
        stage("deploy_docker") {
            steps {
                script {
                    gv.deployDocker()
                }
            }
        }
    }

    post {
        success {
            echo "Pipeline succeeded!"
        }
        failure {
            echo "Pipeline failed!"
        }
    }
}
