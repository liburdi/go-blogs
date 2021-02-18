$(function () {
    /*思路：
     *1.熟悉文件拖拽 目标元素 的四个事件,注意:ondragover、ondrop事件中阻止默认行为
     *2.拖拽放置后，获取到文件对象集合：e.dataTransfer.files
     *3.循环该集合中的每个文件对象，判断文件类型以及文件大小，是指定类型则进行相应的操作
     *4.读取文件信息对象：new FileReader()，它有读取文件对象为DataUrl等方法：readAsDataURL(文件对象)、读取成功之后触发的事件：onload事件等，this.result为读取到的数据
     *5.在FileReader对象中的几个事件中进行相应的逻辑处理
     *
     */

    //必须将jq对象转换为js对象，调用原生方法
    var oDiv = $(".flex_container").get(0);
    var oImg = $("#img").get(0);
    var oP = $(".text").get(0);
    var oMarkDown =$("#mk-introduce").get(0);
    var oType =$("#type").html();

    var maxWidth=0;
    if(oType==="comment"){
        maxWidth=200;
    }else{
        maxWidth=1000;
    }
    console.log(maxWidth)
    //进入
    oDiv.ondragenter = function () {
        oP.html('');
    };
    //移动，需要阻止默认行为，否则直接在本页面中显示文件
    oDiv.ondragover = function (e) {

        e.preventDefault();
    };
    //离开
    oDiv.onleave = function () {
        oP.html('');
    };
    //拖拽放置，也需要阻止默认行为
    oDiv.ondrop = function (e) {

        e.preventDefault();
        //获取拖拽过来的对象,文件对象集合
        var fs = e.dataTransfer.files;
        //若为表单域中的file标签选中的文件，则使用form[表单name].files[0]来获取文件对象集合
        //打印长度
        console.log(fs.length);
        //循环多文件拖拽上传
        for (var i = 0; i < fs.length; i++) {
            //文件类型
            var _type = fs[i].type;

            console.log(_type);
            //判断文件类型
            if (_type.indexOf('image') != -1) {
                //文件大小控制
                console.log(fs[i].size);
                if(fs[i].size>1024000){
                    alert("图片太大，")
                }
                //读取文件对象
                var reader = new FileReader();
                //读为DataUrl,无返回值
                reader.readAsDataURL(fs[i]);
                reader.onloadstart = function (e) {
                    //开始加载
                };
                // 这个事件在读取进行中定时触发
                reader.onprogress = function (e) {

                    $("#total").html(e.total);
                };

                //当读取成功时触发，this.result为读取的文件数据
                reader.onload = function () {
                    //文件数据
                    console.log(this.result);
                    //添加文件预览
                    console.log(oMarkDown)
                    console.log(oMarkDown.value)
                    oMarkDown.value=oMarkDown.value + "![avatar]("+this.result+")";
                    convert();
                    let oImgContent = $("<img style='height:40px;' src='' onclick='delImg()' />");
                    oImgContent.attr("src", this.result);
                    $(oImg).append(oImgContent); //oDiv转换为js对象调用方法
                    let fun=$("#result img").each(function(){
                        $(this).attr('width',maxWidth)
                    })
                    oMarkDown.scrollTop=oMarkDown.scrollHeight
                    oMarkDown.focus()
                    if(oType==="comment"){
                        setTimeout(fun,1000)
                    }

                };
                //无论成功与否都会触发
                reader.onloadend = function () {
                    if (reader.error) {
                        console.log(reader.error);
                    } else {
                        //上传没有错误，ajax发送文件，上传二进制文件
                    }
                }
            } else {
                alert('请上传图片文件！');
            }
        }

    };

    oMarkDown.onpaste = function (e) {

        e.preventDefault();

        if ( !(e.clipboardData && e.clipboardData.items) ) {
            return;
        }
        let f='';
        let isHaveFile=false;


        let promise = new Promise(function(resolve, reject){
            for (let i = 0, len = e.clipboardData.items.length; i < len; i++) {
                let item = e.clipboardData.items[i];
                if (item.kind === "file") {
                    console.log(1)
                    isHaveFile=true
                }
            }
            resolve()
        });
        let funA=function (){
            let len=2
            if(isHaveFile){
                len= e.clipboardData.items.length
            }
            for (let i = 0; i < len; i++) {
                let item = e.clipboardData.items[i];
                if (item.kind === "file") {
                    f= item.getAsFile();
                    return appendImage(f);
                }
                if (item.kind === "string" && i===1) {
                    item.getAsString(function (str) {
                        if(isHaveFile===true){
                            return;
                        }
                        console.log(str);
                        oMarkDown.value = oMarkDown.value.substring(0, oMarkDown.selectionStart) + str + oMarkDown.value.substring(oMarkDown.selectionEnd, oMarkDown.value.length);
                        convert();

                    })

                }
            }
        }

        let appendImage=function (f) {
            //获取拖拽过来的对象,文件对象集合
            var fs = [f];
            // console.log(fs)
            //若为表单域中的file标签选中的文件，则使用form[表单name].files[0]来获取文件对象集合
            //打印长度
            console.log(fs.length);
            //循环多文件拖拽上传
            for (var i = 0; i < fs.length; i++) {
                //文件类型
                var _type = fs[i].type;

                console.log(_type);
                //判断文件类型
                if (_type.indexOf('image') != -1) {
                    //文件大小控制
                    console.log(fs[i].size);
                    if(fs[i].size>1024000){
                        alert("图片太大，")
                    }
                    //读取文件对象
                    var reader = new FileReader();
                    //读为DataUrl,无返回值
                    reader.readAsDataURL(fs[i]);
                    reader.onloadstart = function (e) {
                        //开始加载
                    };
                    // 这个事件在读取进行中定时触发
                    reader.onprogress = function (e) {

                        $("#total").html(e.total);
                    };

                    //当读取成功时触发，this.result为读取的文件数据
                    reader.onload = function () {
                        //文件数据
                        console.log(this.result);
                        //添加文件预览
                        console.log(oMarkDown)
                        console.log(oMarkDown.value)

                        oMarkDown.value = oMarkDown.value.substring(0, oMarkDown.selectionStart) + "![avatar]("+this.result+")" + oMarkDown.value.substring(oMarkDown.selectionEnd, oMarkDown.value.length);
                        convert();
                        let oImgContent = $("<img style='height:40px;' src='' onclick='delImg()' />");
                        oImgContent.attr("src", this.result);
                        $(oImg).append(oImgContent); //oDiv转换为js对象调用方法

                        let fun=$("#result img").each(function(){
                            $(this).attr('width',maxWidth)
                        })
                        oMarkDown.scrollTop=oMarkDown.scrollHeight
                        oMarkDown.focus()
                        if(oType==="comment"){
                            setTimeout(fun,1000)
                        }

                        isHaveFile=false;
                    };
                    //无论成功与否都会触发
                    reader.onloadend = function () {
                        if (reader.error) {
                            console.log(reader.error);
                        } else {
                            //上传没有错误，ajax发送文件，上传二进制文件
                        }
                    }
                } else {
                    alert('请上传图片文件！');
                }
            }
        }
        promise.then(funA());


    };

    delImg = function () {
        let that=event.currentTarget;

        oMarkDown.value=oMarkDown.value.replace("![avatar]("+that.src+")",'');
        convert();
        $('#img img').attr("style",'display:none;')

    }
});
