pipeline {
    // Define environment variables
    environment {
        SERVICE_NAME     = 'my-service'
        CREDENTIAL_ID    = 'docker-registry'
        PROJECT_DOCKER   = 'ciao'
        DOCKER_APP_PATH  = '/root/'
        DOCKER_NETWORK   = 'ciao'
        SSH_ID           = 'cicd'
    }

    // Define the agent to run the pipeline
    agent any

    // Define stages of the pipeline
    stages {
        stage('Set Environment') {
            steps {
                echo 'Setup environment'
                script {
                    setupEnvironment()
                }
            }
        }

        stage('Build Image') {
            when { 
                expression { 
                    return isBuildRequired()
                }
            }          
            steps {
                buildDockerImage()
            }
        }

        stage('Docker Login tag and push') {
            when { 
                expression { 
                    return isBuildRequired()
                }
            }
            steps {
                dockerLoginTagAndPush()
            }
        }        

        stage('Set Image VM') {
            when { 
                expression { 
                    return isBuildRequired()
                }
            }
            steps {
                setImageVM()
            }
        }
    }        
}

// Function to set up environment variables based on branch
def setupEnvironment() {
    switch (env.BRANCH_NAME) {
        case 'development':
            env.FLAVOR         = 'development'
            env.SERVICE_FLAVOR = 'staging'
            env.SERVER_TARGET  = env.STAGING_SERVER
            env.APP_TAG        = "latest"
            env.USER_SSH       = env.USER_STAGING_SERVER
            break
        case 'master':
            env.FLAVOR         = 'master'
            env.SERVICE_FLAVOR = 'production'
            env.SERVER_TARGET  = env.PRD_SERVER
            env.APP_TAG        = env.BUILD_NUMBER
            env.USER_SSH       = env.USER_PRD_SERVER
            break
    }     
}

// Function to check if build is required based on branch
def isBuildRequired() {
    return env.BRANCH_NAME == 'development' || env.BRANCH_NAME == 'master'
}

// Function to build Docker image
def buildDockerImage() {
    echo 'Build Image ' + env.FLAVOR
    sh "sudo docker buildx build . -t ${SERVICE_NAME}:${BRANCH_NAME}-latest"
}

// Function to login, tag, and push Docker image to docker registry Harbor
def dockerLoginTagAndPush() {
    withCredentials([usernamePassword(credentialsId: CREDENTIAL_ID, usernameVariable: 'DOCKER_USERNAME', passwordVariable: 'DOCKER_PASSWORD')]) {
        echo 'Push docker image to docker registry Harbor'
        sh "echo $DOCKER_PASSWORD | docker login -u $DOCKER_USERNAME --password-stdin $URL_NON_PROTOCOL"
        sh "docker tag ${SERVICE_NAME}:${BRANCH_NAME}-latest ${URL_NON_PROTOCOL}/${PROJECT_DOCKER}/${SERVICE_NAME}:${BRANCH_NAME}-${APP_TAG}"
        sh "docker push ${URL_NON_PROTOCOL}/${PROJECT_DOCKER}/${SERVICE_NAME}:${BRANCH_NAME}-${APP_TAG}"
    }
}

// Function to set Docker image on VM
def setImageVM() {
    sh 'echo ssh connecting...'
    sshagent(credentials: [SSH_ID]) {
        // Check if the container exists and remove it if necessary
        sh "ssh -t -o StrictHostKeyChecking=no -l ${USER_SSH} ${SERVER_TARGET} -p 22 'pwd; if sudo docker inspect ${SERVICE_NAME}-${SERVICE_FLAVOR} &> /dev/null 2>&1; then sudo docker rm ${SERVICE_NAME}-${SERVICE_FLAVOR} -f; fi'"
        // Docker pull and run
        sh "ssh -t -o StrictHostKeyChecking=no -l ${USER_SSH} ${SERVER_TARGET} -p 22 'docker pull ${URL_NON_PROTOCOL}/${PROJECT_DOCKER}/${SERVICE_NAME}:${BRANCH_NAME}-${APP_TAG}; sudo docker run -d -p 9292:9292 --name ${SERVICE_NAME}-${SERVICE_FLAVOR} --network=${DOCKER_NETWORK} -v ${APPS_PATH_ENV}/${FLAVOR}/.env:${DOCKER_APP_PATH}.env ${URL_NON_PROTOCOL}/${PROJECT_DOCKER}/${SERVICE_NAME}:${BRANCH_NAME}-${APP_TAG}'"
    }
}