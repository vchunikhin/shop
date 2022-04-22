description="The pet project for golang researching"

task("createInitialMigration") {
    group = "migration"
    description = "Create schema folder with initial sql scripts"
    doLast {
        exec {
            executable("sh")
            args("migration/create_init_migration.sh")
        }
    }
}

task("configureDockerEnvironment") {
    group = "env"
    description = "Configure the docker environment"
    doLast {
        exec {
            executable("sh")
            args("env/env_local.sh")
        }
    }
}

task("dropDockerEnvironment") {
    group = "env"
    description = "Drop the docker environment"
    doLast {
        exec {
            executable("sh")
            args("env/env_local_drop.sh")
        }
    }
}

task("updateDependencies") {
    group = "go"
    description = "Download all dependencies based on go.mod file"
    doLast {
        exec {
            executable("go")
            args("mod", "tidy")
        }
    }
    doLast {
        exec {
            executable("go")
            args("mod", "download")
        }
    }
}

task("runMigration") {
    group = "migration"
    description = "Run migration"
    doLast {
        exec {
            executable("sh")
            args("migration/run_migration.sh")
        }
    }
}

task("runMigrationToDropTables") {
    group = "migration"
    description = "Run migration to drop the schema and tables"
    doLast {
        exec {
            executable("sh")
            args("migration/run_migration_to_drop_tables.sh")
        }
    }
}