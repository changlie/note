
Copy SOURCE to DEST, or multiple SOURCE(s) to DIRECTORY.

### 格式
```
       cp [OPTION]... [-T] SOURCE DEST
  or:  cp [OPTION]... SOURCE... DIRECTORY
  or:  cp [OPTION]... -t DIRECTORY SOURCE...
```

### 选项
```
  -a, --archive                same as -dR --preserve=all
      --attributes-only        don't copy the file data, just the attributes
      --backup[=CONTROL]       make a backup of each existing destination file
  -b                           like --backup but does not accept an argument
      --copy-contents          copy contents of special files when recursive
  -d                           same as --no-dereference --preserve=links
  
  -f, --force                  if an existing destination file cannot be
                                 opened, remove it and try again (this option
                                 is ignored when the -n option is also used)
  -i, --interactive            prompt before overwrite (overrides a previous -n
                                  option)
                                  
  -H                           follow command-line symbolic links in SOURCE
  -l, --link                   hard link files instead of copying
  -L, --dereference            always follow symbolic links in SOURCE
  -s, --symbolic-link          make symbolic links instead of copying
  
  -n, --no-clobber             do not overwrite an existing file (overrides
                                 a previous -i option)
  -P, --no-dereference         never follow symbolic links in SOURCE
  -p                           same as --preserve=mode,ownership,timestamps
      --preserve[=ATTR_LIST]   preserve the specified attributes (default:
                                 mode,ownership,timestamps), if possible
                                 additional attributes: context, links, xattr,
                                 all
      --no-preserve=ATTR_LIST  don't preserve the specified attributes
      --parents                use full source file name under DIRECTORY
      
  -R, -r, --recursive          copy directories recursively
      --reflink[=WHEN]         control clone/CoW copies. See below
      --remove-destination     remove each existing destination file before
                                 attempting to open it (contrast with --force)
      --sparse=WHEN            control creation of sparse files. See below
      --strip-trailing-slashes  remove any trailing slashes from each SOURCE
                                 argument
  

  -t, --target-directory=DIRECTORY  copy all SOURCE arguments into DIRECTORY
  -T, --no-target-directory    treat DEST as a normal file
  -u, --update                 copy only when the SOURCE file is newer
                                 than the destination file or when the
                                 destination file is missing
  -v, --verbose                explain what is being done(显示详情)
```

### 实例

1. 复制一个文件到指定位置: 复制cpDoc文件，到applog目录下
```bash
$cp cpDoc applog/cpDoc
```

2. 复制一个文件到指定位置, 并重命名：复制cpDoc文件到applog目录下，并重命名为cps
```bash
$cp cpDoc applog/cps
```

3. 复制一个文件到指定目录: 复制arg文件到applog目录下
```bash
$ cp arg applog
```

4. 复制一个目录下的所有文件到指定目录下
```bash
$ cp test/* test1    # test目录下包含目录会报错，但还是会复制所有文件到指定目录test1
$ cp -rT test test1  # test目录下所有文件及目录都会复制到test1
```

5. 复制一个目录到指定目录下: 复制完成后，test1目录会多一个test目录
```bash
$ cp -r test test1   
```

6. 版本备份: 如果是第一次备份，则会多出一个文件 ` demo.~1~ `
```bash
$ cp --force --backup=numbered demo demo
```


