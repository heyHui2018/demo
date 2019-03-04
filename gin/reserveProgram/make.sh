#!/bin/sh
#需要安装govendor包管理工具体，go get github.com/kardianos/govendor
# 当前路径
SP=$(cd "$(dirname "$0")"; pwd)
echo "当前路径为 $SP"
echo "开始初始化依赖包..."
result=$(govendor init)
if [ $? != 0 ];then
echo "govendor init 初始化依赖包失败" $result
else
echo "初始化依赖包成功"
fi
echo "开始增加依赖包..."
result=$(govendor add +external)
if [ $? != 0 ];then
echo "govendor add +external 增加依赖包失败" $result
else
echo "合并依赖包成功"
fi
echo "开始更新依赖包..."
result=$(govendor add +external)
if [ $? != 0 ];then
echo "govendor add +external 更新依赖包失败" $result
else
echo "更新依赖包成功"
fi

cd $SP
echo "开始编译..."
result=$(govendor  build)
if [ $? != 0 ];then
echo "govendor build 编译失败" $result
else
echo "编译成功"
fi