
```xml
<dependencies>
		<dependency>
			<groupId>junit</groupId>
			<artifactId>junit</artifactId>
			<version>4.12</version>
			<scope>test</scope>
		</dependency>
		<dependency>
			<groupId>io.vertx</groupId>
			<artifactId>vertx-unit</artifactId>
			<version>3.0.0</version>
			<scope>test</scope>
		</dependency>

		<dependency>
			<groupId>io.vertx</groupId>
			<artifactId>vertx-core</artifactId>
			<version>3.0.0</version>
		</dependency>


	</dependencies>

	<build>
		<plugins>
			<plugin>
				<groupId>org.apache.maven.plugins</groupId>
				<artifactId>maven-compiler-plugin</artifactId>
				<configuration>
					<defaultLibBundleDir>lib</defaultLibBundleDir>
					<source>1.8</source>
					<target>1.8</target>
					<encoding>UTF-8</encoding>
				</configuration>
			</plugin>
			<plugin>
				<groupId>org.apache.maven.plugins</groupId>
				<artifactId>maven-resources-plugin</artifactId>
				<configuration>
					<encoding>UTF-8</encoding>
					<useDefaultDelimiters>false</useDefaultDelimiters>
					<escapeString>\</escapeString>
					<delimiters>
						<delimiter>${*}</delimiter>
					</delimiters>
				</configuration>
			</plugin>

			<plugin>
				<groupId>org.apache.maven.plugins</groupId>
				<artifactId>maven-shade-plugin</artifactId>
				<version>2.3</version>
				<executions>
					<execution>
						<phase>package</phase>
						<goals>
							<goal>shade</goal>
						</goals>
						<configuration>
							<transformers>
								<transformer
									implementation="org.apache.maven.plugins.shade.resource.ManifestResourceTransformer">
									<manifestEntries>
										<Main-Class>io.vertx.core.Starter</Main-Class>
										<Main-Verticle>com.changlie.vertX1.MyFirstVerticle</Main-Verticle>
									</manifestEntries>
								</transformer>
							</transformers>
							<artifactSet />
							<outputFile>${project.build.directory}/${project.artifactId}.jar</outputFile>
						</configuration>
					</execution>
				</executions>
			</plugin>

		</plugins>
	</build>

```
