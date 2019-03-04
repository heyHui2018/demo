#!/bin/bash
#需要安装govendor包管理工具体，go get github.com/kardianos/govendor
app_name='reserveProgram'
echo "创建临时目录"
go_path="/tmp/gopath"
myfile="/src/$app_name"
if [ ! -d "$go_path" ];then
mkdir -p "$go_path"
fi
rm -rf $go_path$myfile
mkdir -p $go_path$myfile
echo "设置gopath临时目录"
export GOPATH=$go_path:$GOPATH
#开始拷贝文件

# 当前路径
SP=$(cd "$(dirname "$0")"; pwd)
echo "当前路径为 $SP"
cd $SP
cp -rf * $go_path$myfile
cd $go_path$myfile
echo "临时目录为 $go_path$myfile"

echo "开始编译..."
result=$(govendor  build)
if [ $? != 0 ];then
echo "govendor build 编译失败" $result
else
echo "编译成功,开始拷贝"
cp $app_name $SP/
echo "删除临时目录$go_path$myfile"
rm -rf $go_path$myfile
cd $SP

if [ ! -d target ];then
        mkdir -p target
fi
chmod u+x run.sh
tar -zcvf ./target/release.tar.gz --exclude=target $app_name conf/ run.sh
fi