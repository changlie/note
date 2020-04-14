环境: ubuntu kylin 16.04

出错命令
```bash
# 将文件指定字符串替换成变量值
sed -i "s/destDir/${destDir}/" github-gitee-sync-conf
```

问题解决: 将分隔符'/' 换成 '#' 即可.
```bash
sed -i "s#destDir#${destDir}#" github-gitee-sync-conf
```

具体原因未明, 可能是bug
