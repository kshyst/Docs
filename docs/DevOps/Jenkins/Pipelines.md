# Jenkins Pipelines

Jenkins Pipeline is a suite of plugins which supports implementing and integrating **continuous delivery pipelines** into Jenkins.

## Pipeline-as-Code

The concept of **Pipeline-as-Code** means that the entire build process is defined in a text file called a `Jenkinsfile`. This file is stored in your version control system (e.g., Git), providing:
- **Version Control**: Track changes to the build process.
- **Audit Trail**: See who changed the pipeline and why.
- **Repeatability**: Ensure the build process is identical across different environments.

## Declarative Pipeline

The **Declarative Pipeline** is the modern and recommended syntax for most users. It provides a structured and opinionated way to define your pipeline.

```groovy
pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo 'Building...'
                sh './gradlew build'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing...'
                sh './gradlew test'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying...'
                sh './deploy.sh'
            }
        }
    }
    
    post {
        always {
            echo 'I will always run!'
        }
        success {
            echo 'I will run only if successful.'
        }
    }
}
```

## Scripted Pipeline

The **Scripted Pipeline** is the traditional syntax based on Groovy. It offers more flexibility and power but is more complex to write and maintain.

```groovy
node {
    stage('Build') {
        echo 'Building...'
        sh './gradlew build'
    }
    stage('Test') {
        echo 'Testing...'
        sh './gradlew test'
    }
}
```

## Jenkinsfile Structure

A typical `Jenkinsfile` for a Declarative Pipeline includes:
- **pipeline**: The top-level block.
- **agent**: Defines where the pipeline will execute (`any`, `none`, or a specific label).
- **stages**: Contains one or more **stage** blocks.
- **steps**: The actual commands to execute (e.g., `sh`, `echo`, `git`).
- **post**: Optional block for actions to take after the pipeline completes.

!!! info
    Always prefer **Declarative Pipelines** unless you have complex logic that strictly requires the flexibility of Groovy scripts.
