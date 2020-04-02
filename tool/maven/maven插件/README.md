> 参考:  https://juejin.im/post/5c0fc3126fb9a049c2324878

maven插件其实也是一个maven项目

## 一、pom.xml 配置及依赖（打包类型及写自定义插件所需的依赖）
```xml
  <packaging>maven-plugin</packaging>

  <dependencies>
    <dependency>
      <groupId>org.apache.maven</groupId>
      <artifactId>maven-plugin-api</artifactId>
      <version>3.5.0</version>
    </dependency>

    <dependency>
      <groupId>org.apache.maven.plugin-tools</groupId>
      <artifactId>maven-plugin-annotations</artifactId>
      <version>3.5</version>
    </dependency>
  </dependencies>
```

## 二、业务逻辑要写在AbstractMojo的实现类的execute方法中
> 例子如下：
```java
// Mojo注解name对应配置文件的goal
@Mojo(name = "gen-source-handler", defaultPhase = LifecyclePhase.PROCESS_SOURCES)
public class MyMojo  extends AbstractMojo {

    // Parameter注解的属性defaultValue可以使用maven的内置属性，从而得到一个文件对象
    @Parameter(defaultValue = "${project.build.directory}", property = "outputDir", required = true)
    private File outputDirectory;

    @Parameter(defaultValue = "${project.build.sourceDirectory}", required = true)
    private File sourceDir;

    @Parameter
    private List<String> targetDirs;

    @Parameter
    private String packageName;

    public void execute() throws MojoExecutionException {
        // 具体业务逻辑
    }
}
```

## 三、插件使用时的配置
> 对应步骤二中的的例子
```xml
<plugin>
    <groupId>org.changlie</groupId>
    <artifactId>grpc-proto-source-handler</artifactId>
    <version>v7</version>
    <executions>
        <execution>
            <phase>process-sources</phase>
            <goals>
                <goal>gen-source-handler</goal>
            </goals>
        </execution>
    </executions>
    <configuration>
        <targetDirs>
            <dir>/generated-sources/protobuf/grpc-java</dir>
            <dir>/generated-sources/protobuf/java</dir>
        </targetDirs>
        <packageName>io.grpc.examples.helloworld</packageName>
    </configuration>
</plugin>
```

## 注意
1. 自定义maven插件使用时，要绑定到maven生命周期，要配置<executions>标签，<execution>标签的<phase>可以不配置，使用插件默认的，但<goals>一定要配置
2. Parameter注解的属性defaultValue可以使用maven的内置属性.
