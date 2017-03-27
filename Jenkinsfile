node 
{


    def appName = "demo-go-app"
    def dhubUser = "harshals"
    def imageTag = "${dhubUser}/${appName}:${env.BRANCH_NAME}.${env.BUILD_NUMBER}"
    checkout scm 
  stage 'Build image'
  sh("sudo docker build -t ${imageTag} .")

}

