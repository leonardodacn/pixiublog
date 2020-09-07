<!DOCTYPE html>
<html lang="utf-8">
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
    <title>{{.config.siteName}} 网站地图</title>
    <meta name="author" content="SiteMapX.com"/>
    <meta name="robots" content="index,follow"/>
    <style type="text/css">
        body {
            color: #000000;
            background: #ffffff;
            margin: 20px;
            font-family: Verdana, Arial, Helvetica, sans-serif;
            font-size: 12px;
        }

        #myTable {
            list-style: none;
            margin: 10px 0px 10px 0px;
            padding: 0px;
            width: 100%;
            min-width: 804px;
        }

        #myTable li {
            list-style-type: none;
            width: 100%;
            min-width: 404px;
            height: 20px;
            line-height: 20px;
        }

        .pull-left{
            float: left!important;
        }
        .pull-right{
            float: right!important;
        }

        #myTable li .T1-h {
            font-weight: bold;
            min-width: 300px;
        }

        #myTable li .T2-h {
            width: 200px;
            font-weight: bold;
        }

        #myTable li .T3-h {
            width: 200px;
            font-weight: bold;
        }

        #myTable li .T4-h {
            width: 100px;
            font-weight: bold;
        }

        #myTable li .T1 {
            min-width: 300px;
        }

        #myTable li .T2 {
            width: 200px;
        }

        #myTable li .T3 {
            width: 200px;
        }

        #myTable li .T4 {
            width: 100px;
        }

        #footer {
            padding: 2px;
            margin: 0px;
            font-size: 8pt;
            color: gray;
            min-width: 900px;
        }

        #footer a {
            color: gray;
        }

        .myClear {
            clear: both;
        }

        #nav, #content, #footer {padding: 8px; border: 1px solid #EEEEEE; clear: both; width: 95%; margin: auto; margin-top: 10px;}

    </style>
</head>
<body>
<h2 style="text-align: center; margin-top: 20px">{{.config.siteName}} 网站地图 </h2>
<div id="content">
    <h3>最新文章</h3>
    <ul id="myTable">
        <li>
            <div class="T1-h pull-left">URL</div>
            <div class="T2-h pull-right">Last Change</div>
            <div class="T3-h pull-right">Change Frequency</div>
            <div class="T4-h pull-right">Priority</div>
        </li>
        <div class="myClear"></div>
        <li>
            <div class="T1 pull-left"><a href="{{.config.siteUrl}}" title="{{.config.siteName}}">{{.config.siteName}}</a></div>
            <div class="T2 pull-right">{{.currentDate}}</div>
            <div class="T3 pull-right">daily</div>
            <div class="T4 pull-right">1</div>
        </li>
        <div class="myClear"></div>
        {{range .blogs}}
        <li>
            <div class="T1 pull-left"><a href="{{$.config.siteUrl}}/blog/{{.Code}}" title="{{.Title}}">{{.Title}}</a></div>
            <div class="T2 pull-right">{{DateStr .CreateTime}}</div>
            <div class="T3 pull-right">daily</div>
            <div class="T4 pull-right">0.6</div>
        </li>
        <div class="myClear"></div>
        {{end}}
    </ul>
</div>
<div id="content">
    <h3>分类目录</h3>
    <ul id="myTable">
        <li>
            <div class="T1-h pull-left">URL</div>
            <div class="T2-h pull-right">Last Change</div>
            <div class="T3-h pull-right">Change Frequency</div>
            <div class="T4-h pull-right">Priority</div>
        </li>
        {{range .blogTypes}}    
        <li>
                <div class="T1 pull-left"><a href="{{$.config.siteUrl}}/type/{{.Id}}" title="{{.Name}}">{{.Name}}</a></div>
                <div class="T2 pull-right">{{DateStr .CreateTime}}</div>
                <div class="T3 pull-right">daily</div>
                <div class="T4 pull-right">0.6</div>
            </li>
            <div class="myClear"></div>
        {{end}}
    </ul>
</div>
<div id="content">
    <h3>标签目录</h3>
    <ul id="myTable">
        <li>
            <div class="T1-h pull-left">URL</div>
            <div class="T2-h pull-right">Last Change</div>
            <div class="T3-h pull-right">Change Frequency</div>
            <div class="T4-h pull-right">Priority</div>
        </li>
        {{range .blogTags}}    
        <li>
                <div class="T1 pull-left"><a href="{{$.config.siteUrl}}/tag/{{.Id}}" title="{{.Name}}">{{.Name}}</a></div>
                <div class="T2 pull-right">{{DateStr .CreateTime}}</div>
                <div class="T3 pull-right">daily</div>
                <div class="T4 pull-right">0.6</div>
            </li>
            <div class="myClear"></div>
        {{end}}
    </ul>
</div>
</body>
</html>