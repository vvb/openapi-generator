lazy val root = (project in file(".")).
  settings(
    organization := "org.openapitools",
    name := "petstore-rest-assured",
    version := "1.0.0",
    scalaVersion := "2.11.4",
    scalacOptions ++= Seq("-feature"),
    javacOptions in compile ++= Seq("-Xlint:deprecation"),
    publishArtifact in (Compile, packageDoc) := false,
    resolvers += Resolver.mavenLocal,
    libraryDependencies ++= Seq(
      "io.swagger" % "swagger-annotations" % "1.5.21",
      "io.rest-assured" % "scala-support" % "4.3.0",
      "com.google.code.gson" % "gson" % "2.8.6",
      "io.gsonfire" % "gson-fire" % "1.8.4" % "compile",
      "org.threeten" % "threetenbp" % "1.4.3" % "compile",
      "com.squareup.okio" % "okio" % "1.17.5" % "compile",
      "javax.validation" % "validation-api" % "2.0.1.Final" % "compile",
      "org.hibernate" % "hibernate-validator" "6.0.19.Final" % "compile",
      "junit" % "junit" % "4.13" % "test",
      "com.novocode" % "junit-interface" % "0.10" % "test"
    )
  )
