## CSS选择器

```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <title>CSS选择器</title>

    <style type="text/css">
        /* 标签选择器 */
        div {
            color: red;
        }

        /* ID 选择器 */
        #div2 {
            color: blue;
        }

        /* class选择器 */
        .green {
            color: blueviolet;
        }

        /* class复合选择器 */
        /* <div>下带有green的class */
        div.green {
            color: red;
        }

        /* 
            复合选择器:
                子孙选择器  selector selector
                子选择器    selector>selector
                属性选择器  []
         */
        [green] {
            color: magenta;
        }
    </style>

</head>

<body>
    <div class="green">DIV1</div>
    <div id="div2">DIV2</div>
    <div class="green">DIV3</div>
    <br />
    <span class="green">span1</span>
    <span id="span2">span2</span>
    <span class="green" green="green">span3</span>
</body>

</html>
```

