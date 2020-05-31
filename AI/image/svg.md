### SVG概述   
SVG从1999年由一组加入W3C的公司启动开发，与之前1998年提交给W3C的标准（PGML、VML）构成竞争。SVG从这两个格式吸取了经验。[4]

SVG允许3种图形对象类型：矢量图形、栅格图像以及文本。图形对象——包括PNG、JPEG这些栅格图像——能够被编组、设计、转换及集成进先前的渲染对象中。文本可以在任何适用于应用程序的XML名字空间之内，从而提高SVG图形的搜索能力和无障碍性。SVG提供的功能集涵盖了嵌套转换、裁剪路径、Alpha通道、滤镜效果、模板对象以及可扩展性。

SVG严格遵从XML语法，并用文本格式的描述性语言来描述图像内容，因此是一种和图像分辨率无关的矢量图形格式。

<details>
  <summary> demo </summary>

```html
<html>

<title></title>
<body>


<svg width="150" height="100" viewBox="0 0 3 2">
  <rect width="1" height="2" x="0" fill="#008d46" />
  <rect width="1" height="2" x="1" fill="#ffffff" />
  <rect width="1" height="2" x="2" fill="#d2232c" />
</svg>

<br><br>
<svg version="1.1"
xmlns="http://www.w3.org/2000/svg">
<symbol id="pathSymbol">
    <path  class="path" stroke="#00adef"  d="M281.221,261.806c0,2.756-2.166,4.922-4.922,4.922l0,0h-33.964c-11.715-24.119-31.503-59.855-47.156-68.026
  c-15.751,7.974-35.637,43.907-47.451,68.026h-33.865l0,0c-2.756,0-4.922-2.166-4.922-4.922l0,0l0,0c0-0.295,0-0.689,0.098-0.984
  c0,0,14.078-69.109,79.15-129.161c-2.953-2.56-5.907-5.119-8.959-7.58c-1.87-1.575-2.166-4.233-0.591-6.104
  c1.575-1.772,4.43-2.166,6.497-0.689c3.347,2.461,6.694,5.218,9.746,8.073c3.15-2.953,6.497-5.71,10.041-8.368
  c2.067-1.378,4.922-1.083,6.497,0.689c1.575,1.87,1.28,4.529-0.591,6.104c-3.052,2.56-6.104,5.218-9.155,7.876
  c65.27,59.953,79.446,129.161,79.446,129.161C281.221,261.117,281.221,261.412,281.221,261.806L281.221,261.806L281.221,261.806z"/>
    <path  class="path" stroke="#00adef"  d="M194.589,212.583h0.984l0,0c19.886,28.451,31.503,54.145,31.503,54.145h-63.99C163.086,266.728,174.703,241.034,194.589,212.583
L194.589,212.583z"/>
</symbol>
<g>
  <use xlink:href="#pathSymbol"
    id="path1"></use>
    <use xlink:href="#pathSymbol"
      id="path2"></use>
</g>
</svg>

<svg width="320" height="320" xmlns="http://www.w3.org/2000/svg">
  <g>
    <text font-family="microsoft yahei" font-size="120" y="160" x="160">马</text>
    <animateTransform attributeName="transform" begin="0s" dur="10s" type="rotate" from="0 160 160" to="360 160 160" repeatCount="indefinite"/>
  </g>
</svg>

<svg width="360" height="200" xmlns="http://www.w3.org/2000/svg">
  <text font-family="microsoft yahei" font-size="40" x="0" y="0" fill="#cd0000">马
    <animateMotion path="M10,80 q100,120 120,20 q140,-50 160,0" begin="0s" dur="3s" rotate="auto" repeatCount="indefinite"/>
  </text>
  <path d="M10,80 q100,120 120,20 q140,-50 160,0" stroke="#cd0000" stroke-width="2" fill="none" />
</svg>

<br><br><br><br><br><br><br>

</body>
</html>

```

</details>
