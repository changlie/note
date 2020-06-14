```bash
#!/bin/bash

pname=$1
mkdir $pname
mkdir $pname/header
mkdir $pname/impl

cat > $pname/${pname}.cpp << EOF
#include <stdio.h>

int main() {
    printf("hey! welcome.\n");
}
EOF

cat >> $pname/start << EOF
#!/bin/bash

# obtain header file name
EOF
echo "mf=\"${pname}\"" >> $pname/start
cat >> $pname/start << "EOF"
dir=`pwd`
fnames=""

# 当启用管道时，会生成一个subshell，while循环的代码在subshell中执行
# cat "${dir}/${mf}.cpp" | while read line
while read line
do
    prefix=`echo ${line} | awk '{print $1}'`
    if [[ ${prefix} != "#include" ]]; then
      continue
    fi

    fname=`echo ${line} | awk '{print $2}'`
    if [[ ${fname:0:1} == "<" ]]; then
      continue
    fi

    endIndex=${#fname}-4
    subName=${fname:1:endIndex}
    fnames=`echo "${fnames} ${subName}.h ${subName}.cpp"`
    # echo "$prefix > $fname"
done < "${dir}/${mf}.cpp"

# clean and move files
if [ ! -e gen ]; then
  mkdir gen
fi
rm -f gen/* &> /dev/null
cp -f `pwd`/header/* gen &> /dev/null
cp -f `pwd`/impl/* gen &> /dev/null
cp -f ${mf}.cpp gen

# echo ${fnames}
# /home/changlie/ws/cpp/cpp11demo
# g++ demo.cpp -include /header/share.h /impl/share.cpp -o bin/demo -std=c++11
cd gen
# compile
if [ ${#fnames} == 0 ]; then
  g++ "${mf}.cpp" -o ${mf} -std=c++11
else
  g++ "${mf}.cpp" -include ${fnames} -o ${mf} -std=c++11
fi

# start up
./${mf}
EOF

chmod +x $pname/start

```
