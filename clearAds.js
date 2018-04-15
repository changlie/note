// ==UserScript==
// @name         ads clear
// @namespace    http://tampermonkey.net/
// @version      0.1
// @description  try to take over the world!
// @author       changlie
// @match        https://www.yiibai.com/*
// @match        https://www.linuxidc.com/*
// @match        http://www.runoob.com/*
// @match        http://www.w3school.com.cn/*
// @match        http://www.jb51.net/*
// @grant        none
// ==/UserScript==

(function() {
    'use strict';

    //var url = document.baseURI;
    var domain = document.domain;

    console.log('domain-> '+domain);
    var map = {
        'www.jb51.net':['.header'],
        //w3c
        'www.w3school.com.cn':['#ad_head', '#ad', '.adsbygoogle'],
        //linux公社
       'www.linuxidc.com':['#sitehead', '#li_all', '#printBody > a', '.adsbygoogle', '.linux902', '.linux903'],
        //易百
       'www.yiibai.com':['#adv-javalive', '.adsbygoogle'],
        //菜鸟教程
        'www.runoob.com':['#sidebar-right-ads']
    };

    var selectors = map[domain];

    if(!selectors){
      return;
    }

    if(typeof(jQuery) == "undefined" || typeof($) == "undefined"){
       //alert('no jquery!');
        //加载jquery,并执行回调函数
        loadScript("https://cdn.bootcss.com/jquery/3.3.1/jquery.min.js", function () {
            clear(selectors);
        });
    }


    function clear(arr){
      for(var i=0; i<arr.length;i++){
          $(arr[i]).remove();
      }
    }

    function loadScript(url, callback) {
        var script = document.createElement("script");
        script.type = "text/javascript";
        if(typeof(callback) != "undefined"){
            if (script.readyState) {
                script.onreadystatechange = function () {
                    if (script.readyState == "loaded" || script.readyState == "complete") {
                        script.onreadystatechange = null;
                        callback();
                    }
                };
            } else {
                script.onload = function () {
                    callback();
                };
            }
        }
        script.src = url;
        document.body.appendChild(script);
    }



    // Your code here...
})();
