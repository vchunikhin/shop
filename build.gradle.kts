description="The pet project for golang researching"

task("createInitialMigration") {
    group = "build"
    description = "Create schema folder with initial sql scripts"
    doLast {
        exec {
            executable("sh")
            args("migration/create_init_migration.sh")
        }
    }
}

task("configureDockerEnvironment") {
    group = "build"
    description = "Configure the docker environment"
    doLast {
        exec {
            executable("sh")
            args("env/env_local.sh")
        }
    }
}

task("dropDockerEnvironment") {
    group = "drop"
    description = "Drop the docker environment"
    doLast {
        exec {
            executable("sh")
            args("env/env_local_drop.sh")
        }
    }
}

task("downloadDependencies") {
    description = "Download all dependencies based on go.mod file"
    doLast {
        exec {
            executable("go")
            args("mod", "download")
        }
    }
}


task("initBuild") {
    group = "build"
    description = "Build all initial files and the environment"
    dependsOn("createInitialMigration")
    dependsOn("configureDockerEnvironment")
    dependsOn("downloadDependencies")
}