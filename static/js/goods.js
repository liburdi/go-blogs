/*
@功能：商品页js
@作者：diamondwang
@时间：2013年11月13日
*/

$(function(){
	//商品缩略图左右移动效果
	//点击后退
	$("#backward").click(function(){
		var left = parseInt($(".smallpic_wrap ul").css("left")); //获取ul水平方向偏移量
		var offset = left + 62;
		if (offset <= 0){
			$(".smallpic_wrap ul").stop(true,false).animate({left:offset},"slow",'',function(){
				//动画完成之后，判断是否到了左边缘
				if ( parseInt($(".smallpic_wrap ul").css("left")) >= 0 ){
					$("#backward").removeClass("on").addClass("off");
				}
			});
			//开启右边的按钮
			$("#forward").removeClass("off").addClass("on");			
		}
		
		$(this).blur(); //去除ie 虚边框
	});

	//点击前进
	$("#forward").click(function(){
		var left = parseInt($(".smallpic_wrap ul").css("left")); //获取ul水平方向偏移量
		var len = $(".smallpic_wrap li").size() * 62; //获取图片的整体宽度(图片数 * 图片宽度)558
		var offset = left - 62;
		if (offset >= -(len - 62*5)){
			$(".smallpic_wrap ul").stop(true,false).animate({left:offset},"slow",'',function(){
				//判断是否到了右边缘
				if ( parseInt($(".smallpic_wrap ul").css("left")) <= -(len - 62*5) ){
					$("#forward").removeClass("on").addClass("off");
				}
			});
			//开启左边的按钮
			$("#backward").addClass("on").removeClass("off");
			
		}
		
		$(this).blur(); //去除ie 虚边框
	});

	//选择货品，如颜色、版本等
	$(".product a").click(function(){
		$(this).addClass("selected").siblings().removeClass("selected");
		$(this).find("input").attr({checked:"checked"});
		//去除虚边框
		$(this).blur();
	});


	//购买数量
	//减少
	$("#reduce_num").click(function(){
		if (parseInt($(".amount").val()) <= 1){
			alert("商品数量最少为1");
		} else{
			$(".amount").val(parseInt($(".amount").val()) - 1);
		}
	});

	//增加
	$("#add_num").click(function(){
		$(".amount").val(parseInt($(".amount").val()) + 1);
	});

	//直接输入
	$(".amount").blur(function(){
		if (parseInt($(".amount").val()) < 1){
			alert("商品数量最少为1");
			$(this).val(1);
		}
	});

	//商品详情效果
	$(".detail_hd li").click(function(){
		$(".detail_div").hide().eq($(this).index()).show();
		$(this).addClass("on").siblings().removeClass("on");
	});
});