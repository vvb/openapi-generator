import org.jetbrains.kotlin.gradle.tasks.KotlinCompile

buildscript {
    repositories {
        mavenCentral()
    }
    dependencies {
        classpath("org.springframework.boot:spring-boot-gradle-plugin:2.6.7")
    }
}

group = "org.openapitools"
version = "1.0.0"

repositories {
    mavenCentral()
}

tasks.withType<KotlinCompile> {
    kotlinOptions.jvmTarget = "1.8"
}

plugins {
    val kotlinVersion = "1.6.21"
    id("org.jetbrains.kotlin.jvm") version kotlinVersion
    id("org.jetbrains.kotlin.plugin.jpa") version kotlinVersion
    id("org.jetbrains.kotlin.plugin.spring") version kotlinVersion
    id("org.springframework.boot") version "2.6.7"
    id("io.spring.dependency-management") version "1.0.11.RELEASE"
}

dependencies {
    compile("org.jetbrains.kotlin:kotlin-stdlib-jdk8")
    compile("org.jetbrains.kotlin:kotlin-reflect")
    compile("org.springframework.boot:spring-boot-starter-web")
    compile("org.webjars:swagger-ui:4.10.3")
    compile("org.webjars:webjars-locator-core")
    compile("io.swagger.core.v3:swagger-annotations:2.2.0")

    compile("com.google.code.findbugs:jsr305:3.0.2")
    compile("com.fasterxml.jackson.dataformat:jackson-dataformat-yaml")
    compile("com.fasterxml.jackson.dataformat:jackson-dataformat-xml")
    compile("com.fasterxml.jackson.datatype:jackson-datatype-jsr310")
    compile("com.fasterxml.jackson.module:jackson-module-kotlin")
    compile("javax.validation:validation-api")
    compile("javax.annotation:javax.annotation-api:1.3.2")
    testCompile("org.jetbrains.kotlin:kotlin-test-junit5")
    testCompile("org.springframework.boot:spring-boot-starter-test") {
        exclude(module = "junit")
    }
}
