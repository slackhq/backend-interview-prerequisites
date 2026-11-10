plugins {
    kotlin("jvm") version "2.1.0"
    application
}
repositories { mavenCentral() }
dependencies {
    implementation("org.json:json:20240303")
    implementation("org.xerial:sqlite-jdbc:3.47.2.0")
}
kotlin { jvmToolchain(11) }
application { mainClass = "SetupKt" }
dependencyLocking { lockAllConfigurations() }
