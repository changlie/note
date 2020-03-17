1. 打开 “Edit”–> “Preferences”
2. 在Settings选项卡下面点击”Open Config Folder”即`.atom` 目录
3. 将下面复制到`.atom`文件夹下`style.less`文件保存按需要修改即可.
```css
/*
 * Your Stylesheet
 *
 * This stylesheet is loaded when Atom starts up and is reloaded automatically
 * when it is changed and saved.
 *
 * Add your own CSS or Less to fully customize Atom.
 * If you are unfamiliar with Less, you can read more about it here:
 * http://lesscss.org
 */


/*
 * Examples
 * (To see them, uncomment and save)
 */
 //顶部tab选项卡字体调节
 .tab-bar .tab {
  font-size: 12px;
  font-family: "Consolas";
}
//顶部选中时的tab选项卡字体
.tab.active .title {
 font-size: 14px;
 font-family: "Consolas";
}

//左边目录列表字体样式
.tree-view {
  font-size: 16px;
}

//设置界面字体样式调节
.settings-view .setting-description {
  font-size: 12px;
}
atom-pane {
    font-size: 14px;
}
//底部字体和窗框的样式
.status-bar  {
  font-size: 18px;
  height: 30px;
  line-height:30px;
}
```
