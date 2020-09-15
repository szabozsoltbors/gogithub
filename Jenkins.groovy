// Initialization
def initPipeline() {
    echo "Pipeline initialization"
}

// Build
def buildApp() {
    echo "Build the application."
    sh "make build"
}

// Test
def testApp() {
    echo "Test the application."
    sh "make test"
}

def buildDocker() {
    echo "Build the docker image."
    sh "make docker_build"
}

// Deployment
def deployDocker() {
    echo "Deploy the docker image."
    sh "make docker_push"
}

// def deployKubernetes() {
//     echo "Deploying the app on kubernetes."
//     sh "oc login -u $OPENSHIFT_CREDENTIALS_USR -p $OPENSHIFT_CREDENTIALS_PSW"
//     sh "oc project microservices"

//     // Clean up
//     sh "oc delete -f $KUBERNETES_PATH -l module=backend"
//     sh "oc delete -f $KUBERNETES_PATH -l role=ingress"

//     // Create pods
//     sh "oc create -f $KUBERNETES_PATH -l module=backend"
//     if (params.ENVIRONMENT == "development") {
//         sh "oc create -f $KUBERNETES_PATH -l env=development"
//     } else {
//         sh "oc create -f $KUBERNETES_PATH -l env=production"
//     }
// }

return this