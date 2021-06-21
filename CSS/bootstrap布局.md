## bootstrap布局

### 一、container

- 居中布局

```html
<div class="container">
        <h5>你的第一个bootstrap布局</h5>
    </div>
    <div class="container-fluid">
        <h5>你的第二个bootstrap布局</h5>
    </div>
```



### 二、container-fluid

- 左右对齐布局
- col-x  x的取值范围是1-12! 且每一行的col值总和只能是12，否则就会排到下一行
- col-x  x的取值范围是1-12! 如果不指定的话,会等分
- offset 偏移量

```html
<body>
    <div class="container-fluid">
        <div class="row">
            <!-- col-x  x的取值范围是1-12! 且每一行的col值总和只能是12，否则就会排到下一行-->
            <div class="col-2">第一列</div>
            <div class="col-4">第二列</div>
            <div class="col-4">第三列</div>
        </div>
        <div class="row">
            <!-- col-x  x的取值范围是1-12! 如果不指定的话,会等分-->
            <div class="col">第一列</div>
            <div class="col">第二列</div>
            <div class="col">第三列</div>
            <!-- offset 偏移量 -->
            <div class="col offset-3">第四列</div>
        </div>
    </div>
</body>

```



### 三、列表

- 官方文档：[Tables · Bootstrap v4.5 (getbootstrap.com)](https://getbootstrap.com/docs/4.5/content/tables/)

```html
<table class="table table-striped table-bordered">
        <thead class="thead-dark">
            <tr>
                <th>ID</th>
                <th>名称</th>
                <th>岗位</th>
            </tr>
        </thead>
        <tbody>
            <tr class="table-primary">
                <td>01</td>
                <td>kk</td>
                <td>运维工程师</td>
            </tr>
            <tr>
                <td>02</td>
                <td>hh</td>
                <td>运维工程师</td>
            </tr>
            <tr>
                <td>03</td>
                <td>ll</td>
                <td>Java开发</td>
            </tr>
        </tbody>
    </table>
```



### 四、按钮

```html
<body>
    <div>
        <input type="submit" class="btn btn-primary" value="提交按钮" />
        <input type="button" class="btn btn-success" value="普通按钮" />

        <button class="btn btn-success">普通按钮</button>

        <a class="btn btn-success">新建</a>

        <form action="">
            <!-- 禁止登录 -->
            <input type="submit" class="btn btn-primary" value="提交数据" disabled="disabled" />
        </form>
    </div>

    <div class="btn-group">
        <!-- 按钮组 -->
        <button type="button" class="btn btn-primary">按钮1</button>
        <button type="button" class="btn btn-info">按钮2</button>
        <button type="submit" class="btn btn-danger">按钮3</button>
    </div>
</body>
```



### 五、列表

```html
<div class="row">
        <div class="col-3 offset-3">
            <ul class="list-group">
                <li class="list-group-item">洗衣服</li>
                <li class="list-group-item">做饭</li>
            </ul>
        </div>
    </div>
```

