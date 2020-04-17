`git log`用于查看提交历史
### 常用选项
|选项|说明|
|:------:|:----------:|
|`-<n> `              |仅显示最近的 n 条提交。|
|--date             | 与--pretty结合使用，常用值：`relative`,`local`,`short`,`human`,`unix`|
|--pretty|格式化输出，常用值: oneline, short, medium, full, fuller, reference, email, raw, `format:<string>` |

`git log --pretty=format`常用选项
|选项|说明|
|:------:|:----------:|
|%H|提交的完整哈希值|
|%h|提交的简写哈希值|
|%T|树的完整哈希值|
|%t|树的简写哈希值|
|%P|父提交的完整哈希值|
|%p|父提交的简写哈希值|
|%an|作者名字|
|%ae|作者的电子邮件地址|
|%ad|作者修订日期（可以用 --date=选项 来定制格式）|
|%ar|作者修订日期，按多久以前的方式显示|
|%cn|提交者的名字|
|%ce|提交者的电子邮件地址|
|%cd|提交日期|
|%cr|提交日期（距今多长时间）|
|%s|提交说明|




### 实例


查看所有提交记录
```bash
git log --pretty=format:"%h %s %cd" --date=short
```

查看最近6条提交记录
```bash
git log --pretty=format:"%h %s %cd" --date=short -6
```
