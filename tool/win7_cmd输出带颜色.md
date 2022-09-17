
### 示例一
```bat
@echo off
set str=青天有月来几时 我今停杯一问之 人攀明月不可得 月行却与人相随
setlocal enabledelayedexpansion
set col=CAD9B
for /f "tokens=1-4 delims= " %%1 in ("%str%") do (echo.
for %%m in (%%1 %%2 %%3 %%4) do (set s=%%m
for /l %%a in (0,1,6) do (
call set b=0%%col:~!x!,1%%
set /a x+=1&if !x!==5 (set x=0)
set c=!s:~%%a,1!
set/p= <nul>!c! 
findstr /a:!b! .* "!c!*"
del !c!
ping /n 1 /w 500 127.1>nul&ping /n 1 /w 500 127.1>nul&ping /n 1 /w 500 127.1>nul)
echo.&echo.))
pause>nul&exit
```

### 示例二
```bat
@echo off
SETLOCAL EnableDelayedExpansion
for /F "tokens=1,2 delims=#" %%a in ('"prompt #$H#$E# & echo on & for %%b in (1) do rem"') 
do (
  set "DEL=%%a"
)
rem echo say the name of the colors, don't read
:start
rem cls
call :ColorText 0a "blue"
call :ColorText 0C "green"
call :ColorText 0b "red"
echo.
call :ColorText 19 "yellow"
call :ColorText 2F "black"
call :ColorText 4e "white"

echo.
echo aaaaaa

pause

goto :eof

:ColorText
echo off
<nul set /p ".=%DEL%" > "%~2"
findstr /v /a:%1 /R "^$" "%~2" nul
del "%~2" > nul 2>&1
goto :eof
```
使用注意事项：

1.你的代码放在 :start 和第一行 goto :eof之间 

2.换行使用 echo.  

3.最好使用双引号把内容括起来，不然遇到空格就会报拒绝访问的错误，内容如果包含特殊字符（例如："a：！aa" 等等，一般是win不允许我们用来命名文件的符号），也可能导致报错，或者字符不显示，所以最好不要使用特殊字符，这个会在当前目录生成一个临时文件，以你的内容为命名，如果包含了那些字符，就会产生那些文件，最好不删除，需要自己手动删除。

4.关于颜色，有背景色和字体色，第一个数字代表背景色，第二个数字代表字体色，至于颜色值，在 cmd 中输入 color /? 查看


### 示例三
```bat
@echo off
call :colortheword "将着色的字符" 37 "使用默认颜色显示"
goto :eof

rem 在参数中<>为必选项，[]为可选项，调用参数：<str1=将着色的字符> [str2=颜色设置] [str3=正
常显示字符] 
:colortheword <str1=将着色的字符> [str2=颜色设置] [str3=正常显示字符] 
set "objFile=%~1"
set "objColor=07"&if not "%~2."=="." set "objColor=%~2"
set "objMsg= "&if not "%~3."=="." set "objMsg=%~3"
for /F %%a in ('"prompt $h & for %%b in (1) do rem"')do set /p="%%a%objMsg%"<nul>"%objFile
%"
findstr /a:%objColor% .* "%objFile%" nul
del /q "%objFile%" >nul 2>nul
goto :eof
```


