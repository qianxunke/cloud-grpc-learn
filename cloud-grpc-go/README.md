//生成proto文件 将*替换为文件名字  
protoc -I=../cloud-grpc-protos  --go_out=plugins=grpc:../cloud-grpc-go  ../cloud-grpc-protos/*.proto
    命令解释：-I 即扫描的路径 就是如果多个proto文件之间有互相依赖，生成某个proto文件时，需要import其他几个proto文件，这时候就要用-I来指定搜索目录。如果没有指定 –I 参数，则在当前目录进行搜索。此时protoc在此目录执行，所以输出目录要注意匹配
