#根据protobfur 生成java代码
mvn protobuf:compile
mvn protobuf:compile-custom
#打包
mvn clean package -Dmaven.test.skip=true

