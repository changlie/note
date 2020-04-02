```xml
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/maven-v4_0_0.xsd">
	<modelVersion>4.0.0</modelVersion>
	<groupId>com.changlie</groupId>
	<artifactId>prototype1</artifactId>
	<packaging>war</packaging>
	<version>0.0.1-SNAPSHOT</version>
	<name>prototype1 Maven Webapp</name>
	<url>http://maven.apache.org</url>

	<properties>
		<junit.version>4.12</junit.version>
		<jsp.api.version>7.0.76</jsp.api.version>
		<jstl.version>1.2</jstl.version>
		<spring.version>4.3.6.RELEASE</spring.version>
		<aspectjweaver.version>1.8.9</aspectjweaver.version>



		<slf4j.version>1.7.21</slf4j.version>
		<log4j.version>1.2.17</log4j.version>
		<aliyun-sdk-oss.version>2.3.0</aliyun-sdk-oss.version>
		<commons-io.version>2.5</commons-io.version>
		<commons-fileupload.version>1.3.2</commons-fileupload.version>
		<commons-lang3.version>3.5</commons-lang3.version>


		<fastjson.version>1.2.24</fastjson.version>
		<shiro.version>1.3.2</shiro.version>

		<!-- Plugin的属性定义 -->
		<project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
		<java.version>1.8</java.version>

	</properties>

	<dependencies>

		<!-- junit单元测试 -->
		<dependency>
			<groupId>junit</groupId>
			<artifactId>junit</artifactId>
			<version>${junit.version}</version>
			<!-- 该jar包保留到测试 -->
			<scope>test</scope>
		</dependency>

		<!-- spring -->
		<dependency>
			<groupId>org.springframework</groupId>
			<artifactId>spring-test</artifactId>
			<version>${spring.version}</version>
			<scope>test</scope>
		</dependency>

		<dependency>
			<groupId>org.springframework</groupId>
			<artifactId>spring-webmvc</artifactId>
			<version>${spring.version}</version>
		</dependency>

		<dependency>
			<groupId>org.springframework</groupId>
			<artifactId>spring-jdbc</artifactId>
			<version>${spring.version}</version>
		</dependency>

		<dependency>
			<groupId>org.springframework</groupId>
			<artifactId>spring-tx</artifactId>
			<version>${spring.version}</version>
		</dependency>

		<!-- aspectjweaver -->
		<dependency>
			<groupId>org.aspectj</groupId>
			<artifactId>aspectjweaver</artifactId>
			<version>${aspectjweaver.version}</version>
		</dependency>

		<!-- jsp-api、servlet-api、el -->
		<dependency>
			<groupId>org.apache.tomcat</groupId>
			<artifactId>tomcat-jsp-api</artifactId>
			<version>${jsp.api.version}</version>
			<!-- 该jar包最终由Web容器提供 -->
			<scope>provided</scope>
		</dependency>




		<!-- slf4j -->
		<dependency>
			<groupId>org.slf4j</groupId>
			<artifactId>slf4j-jdk14</artifactId>
			<version>${slf4j.version}</version>
		</dependency>

		<!-- log4j -->
		<dependency>
			<groupId>log4j</groupId>
			<artifactId>log4j</artifactId>
			<version>${log4j.version}</version>
		</dependency>


		<!-- 数据库相关 -->
		<dependency>
			<groupId>com.alibaba</groupId>
			<artifactId>druid</artifactId>
			<version>1.0.29</version>
		</dependency>
		<dependency>
			<groupId>com.microsoft.sqlserver</groupId>
			<artifactId>sqljdbc4</artifactId>
			<version>4.0</version>
		</dependency>
		<dependency>
			<groupId>com.mchange</groupId>
			<artifactId>c3p0</artifactId>
			<version>0.9.5.2</version>
		</dependency>

		<!-- commons-io -->
		<dependency>
			<groupId>commons-io</groupId>
			<artifactId>commons-io</artifactId>
			<version>${commons-io.version}</version>
		</dependency>
		<dependency>
			<groupId>commons-collections</groupId>
			<artifactId>commons-collections</artifactId>
			<version>3.2.1</version>
		</dependency>
		<!-- commons-fileupload -->
		<dependency>
			<groupId>commons-fileupload</groupId>
			<artifactId>commons-fileupload</artifactId>
			<version>${commons-fileupload.version}</version>
		</dependency>

		<dependency>
			<groupId>org.apache.commons</groupId>
			<artifactId>commons-lang3</artifactId>
			<version>${commons-lang3.version}</version>
		</dependency>

		<!-- fastjson -->
		<dependency>
			<groupId>com.alibaba</groupId>
			<artifactId>fastjson</artifactId>
			<version>${fastjson.version}</version>
		</dependency>

		<!-- shiro -->
		<dependency>
			<groupId>org.apache.shiro</groupId>
			<artifactId>shiro-spring</artifactId>
			<version>${shiro.version}</version>
		</dependency>
		<dependency>
			<groupId>org.apache.shiro</groupId>
			<artifactId>shiro-ehcache</artifactId>
			<version>${shiro.version}</version>
		</dependency>
		<dependency>
			<groupId>org.apache.shiro</groupId>
			<artifactId>shiro-quartz</artifactId>
			<version>${shiro.version}</version>
		</dependency>

	</dependencies>

	<build>
		<!-- 打成war的最终名 -->
		<finalName>prototype</finalName>
		<!-- 指定java源代码存放的目录 -->
		<sourceDirectory>${basedir}/src/main/java</sourceDirectory>
		<!-- 指定资源文件存放的目录 -->
		<resources>
			<resource>
				<directory>${basedir}/src/main/resources</directory>
			</resource>
		</resources>
		<!-- 指定测试源代码存放的目录 -->
		<testSourceDirectory>${basedir}/src/test/java</testSourceDirectory>
		<!-- 指定测试资源文件存放的目录 -->
		<testResources>
			<testResource>
				<directory>${basedir}/src/test/resources</directory>
			</testResource>
		</testResources>

		<plugins>
			<!-- compiler插件, 设定JDK版本 -->
			<plugin>
				<groupId>org.apache.maven.plugins</groupId>
				<artifactId>maven-compiler-plugin</artifactId>
				<version>3.1</version>
				<configuration>
					<source>${java.version}</source>
					<target>${java.version}</target>
					<showWarnings>true</showWarnings>
				</configuration>
			</plugin>

			<!-- web容器插件 -->
			<plugin>
				<groupId>org.apache.tomcat.maven</groupId>
				<artifactId>tomcat7-maven-plugin</artifactId>
				<version>2.2</version>
				<configuration>
					<!-- 端口号 -->
					<port>8080</port>
					<!-- 支持热部署 -->
					<contextReloadable>true</contextReloadable>
				</configuration>
			</plugin>

			<!-- war打包插件, 设定war包名称不带版本号 -->
			<plugin>
				<groupId>org.apache.maven.plugins</groupId>
				<artifactId>maven-war-plugin</artifactId>
				<version>2.4</version>
				<configuration>
					<warName>${project.artifactId}</warName>
				</configuration>
			</plugin>
		</plugins>
	</build>
</project>



```
