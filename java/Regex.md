```java
@Test
    public void test7(){
        String s = "{tyii}, {name}, ${userId}";
        String reg = "\\{(.*?)\\}|\\$\\{(.*?)\\}";
        Pattern p2 = Pattern.compile(reg);
        Matcher m2 = p2.matcher(s);
        System.out.println("groupCount: "+m2.groupCount());
        while (m2.find()){
            System.out.println("group(0):"+m2.group(0));
            System.out.println("group(1):"+m2.group(1));  // 组提取字符串
            System.out.println("group(1):"+m2.group(2));  // 组提取字符串
            System.out.println("------------------------------");
        }
    }


    @Test
    public void testGetSubStr6() {
        String reg = "\\D+(\\d+)\\D+(\\d+)";    //提取字符串末尾的数字：封妖塔守卫71 == >> 71
        String s = "封妖塔守卫89 == >> 71";
        Pattern p2 = Pattern.compile(reg);
        Matcher m2 = p2.matcher(s);
        while(m2.find()){
            System.out.println("group(0):"+m2.group(0));
            System.out.println("group(1):"+m2.group(1));  // 组提取字符串
            System.out.println("group(2):"+m2.group(2));  // 组提取字符串
        }
    }

    @Test public void testGetSubStr5() {
        String regex = "\\[{1}(.+)\\]{1}";
        String str = "短信中包含以下敏感字：[fuck,shit,bitch]并且号码中包含以下黑名单：张三-13701234567，李四-18701234567";
        System.out.print("敏感字如下：");
        System.out.println(getKeyWords(regex, str));
        System.out.print("黑名单号码如下：");
        System.out.println(getMobiles(str));
        System.out.println("字符替换：" + replaceStr("我喜欢红色"));
    }
    // 提取 中括号中关键字
    public static String getKeyWords(String regex, String str) {
        Pattern p = Pattern.compile(regex);
        Matcher m = p.matcher(str);
        if (m.find()) {
            System.out.println(m.group(0));
            return m.group(1);
        }
        return null;
    }
    // 提取字符串中的手机号码
    public static String getMobiles(String str) {
        Pattern p = Pattern.compile("(\\+86|)?(\\d{11})");
        Matcher m = p.matcher(str);
        StringBuilder sb = new StringBuilder();
        while (m.find()) {
            if (sb.length() > 0)
                sb.append(",");
            sb.append(m.group(2));
        }
		/*
		 * 不加"()"也能将手机号码输出 添加"()"是为了筛选数据添加上去的，
		 * 第一对"()"是为了获取字符串"+86"，代码是System.out.println(m.group(1));，
		 * 第二对"()"是获取11位纯数字电话号码， 本次的输出的手机号码中包含了"+86"，如果只要11位数字号码，
		 * 可将代码改为System.out.println(m.group(2));
		 */
        // System.out.println(m.groupCount());//
        // 该行代码是输出有几对"()"，即捕获组个数，本次输出结果是2，因为有两对"()"
        return sb.toString();
    }
    // 替换字符
    public static String replaceStr(String str) {
        String regex = "红";
        Pattern p = Pattern.compile(regex);
        Matcher m = p.matcher(str);
        return m.replaceAll("绿");
    }

    /**
     * 取出想取出内层的p标签 ：
     * <p>
     * 吃饭
     * </p>
     */
    @Test public void testGetSubStr4() {
        String p = "<p><p>吃饭</p></p>";

        Pattern pattern = Pattern.compile("\\<p\\>((?!\\<p\\>).*?)\\</p\\>");
        Matcher matcher = pattern.matcher(p);
        System.out.println("pattern: "+pattern.pattern());
        System.out.println("groupCount: "+matcher.groupCount());
        while (matcher.find()) {
            System.out.println("matcher.group(0): "+matcher.group(0));
            System.out.println("matcher.group(1): "+matcher.group(1));
//            System.out.println("matcher.group(2): "+matcher.group(2));
        }
    }

    @Test public void testGetSubStr3() {
        // 按指定模式在字符串查找
        String line = "This order was placed for QT3000abttr! OK?";
        String pattern = "(\\D*(p\\D*)QT(\\w))(\\d+)(.*r)(.*)";

        // 创建 Pattern 对象
        Pattern r = Pattern.compile(pattern);
        System.out.println(r.pattern());
        // 现在创建 matcher 对象
        Matcher m = r.matcher(line);
        System.out.println("groupCount:" + m.groupCount());
        int index = 0;
        if (m.find()) {
            do {
                System.out.println("index" + index + ": " + m.group(index++));
                System.out.println("index" + index + ": " + m.group(index++));
                System.out.println("index" + index + ": " + m.group(index++));
                System.out.println("index" + index + ": " + m.group(index++));
                System.out.println("index" + index + ": " + m.group(index++));
            } while (m.find());
        } else {
            System.out.println("NO MATCH");
        }
    }

    @Test public void testGetSubStr2() {
        String targetStr = " and id = {{userId  } and roleId in ({a  roles}), ({deps}) {aaaa} \\{tdt\\}}";
        String reg = "\\{(.*?)\\}";

        System.out.println("-------------");
        String replaceStr = "?";
        System.out.println("1: " + targetStr.indexOf("}"));
        System.out.println("2: " + targetStr.lastIndexOf("}"));
        Pattern pattern = Pattern.compile(reg);
        Matcher matcher = pattern.matcher(targetStr);
        System.out.println("groupCount: " + matcher.groupCount());
        StringBuffer sb = new StringBuffer();
        while (matcher.find()) {
            String group = matcher.group();
            System.out.println(group.length() + ", start:" + matcher.start() + ", end:" + matcher.end());
            System.out.println(matcher.group());
            matcher.appendReplacement(sb, replaceStr);
        }
        System.out.println("r1: " + sb);
        matcher.appendTail(sb);
        System.out.println("final:" + sb);
    }


    @Test public void testGetSubStr1() {
        String s = "A876X";
        // 把要匹配的字符串写成正则表达式，然后要提取的字符使用括号括起来
        // 在这里，我们要提取最后一个数字，正则规则就是“一个数字加上大于等于0个非数字再加上结束符”
        Pattern pattern = Pattern.compile("(\\d)[^\\d]*$");
        Matcher matcher = pattern.matcher(s);
        if (matcher.find())
            System.out.println(matcher.group(1));
    }

    @Test public void testGetSubStr0() {
        String str = "abc3443abcfgjhgabcgfjabc";
        String rgex = "abc(.*?)abc";

        Pattern pattern = Pattern.compile(rgex);// 匹配的模式
        Matcher m = pattern.matcher(str);
        while (m.find()) {
            System.out.println("m.group(0)-> "+m.group(0));
            System.out.println("m.group(1)-> "+m.group(1));
        }
    }

```
