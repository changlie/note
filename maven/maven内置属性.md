> 参考： https://www.cnblogs.com/whx7762/p/7927994.html

## maven属性
> maven预定义，用户可以直接使用的
```
${basedir}/${project.basedir}表示项目根目录，即包含pom.xml文件的目录
${version}表示项目版本
```

## POM属性
> 使用pom属性可以引用到pom.xml文件对应的元素的值
```
${project.build.sourceDirectory}:项目的主源码目录，默认为src/main/java/.
${project.build.testSourceDirectory}:项目的测试源码目录，默认为/src/test/java/.
${project.build.directory}:项目构建输出目录，默认为target/.
${project.outputDirectory}:项目主代码编译输出目录，默认为target/classes/.
${project.testOutputDirectory}:项目测试代码编译输出目录，默认为target/testclasses/.
${project.groupId}:项目的groupId.
${project.artifactId}:项目的artifactId.
${project.version}:项目的version,于${version}等价 
${project.build.finalName}:项目打包输出文件的名称，默认 为${project.artifactId}${project.version}.
```

## 自定义属性
> 在pom.xml文件的<properties>标签下定义的属性
```
<project>
  <properties>
      <spring.version>4.7.12-release</spring.version>
  </properties>

  <dependency>
      <groupId>org.springframework</groupId>
      <artifactId>spring-jdbc</artifactId>
      <version>${spring.version}</version>
   </dependency>
</project>
```

## setting.xml文件属性
> 与pom属性同理，用户可以用以settings.开头的属性引用setting.xml文件的XML元素值
```
${settings.localRepository}表示本地仓库的地址
```

## java系统属性
> 所有的java系统属性都可以用env,开头的maven属性引用,使用mvn help:system命令可查看所有环境变量
```
${env.JAVA_HOME}表示JAVA_HOME环境变量的值
```

