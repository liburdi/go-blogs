/*
@功能：首页面js
@作者：diamondwang
@时间：2013年11月13日
*/

/* 注意，要在页面中先引入jquery*/
$(function(){
	//右侧，话费、旅行、彩票、游戏切换效果
	$(".service h2 span").mouseover(function(){
		$(this).addClass("on").siblings().removeClass("on");
		$(".service_wrap div").hide().eq($(this).index()).show();

	});

	//导购区域切换效果，疯狂抢购，热卖商品、推荐商品、新品上架，猜您喜欢
	$(".guide_content h2 span").mouseover(function(){
		$(this).addClass("on").siblings().removeClass("on");
		$(".guide_wrap div").hide().eq($(this).index()).show();

	});

	//各楼层区域切换
	$(".goodslist h2 span").mouseover(function(){
		$(this).addClass("on").siblings().removeClass("on");
		$(".goodslist_wrap div").hide().eq($(this).index()).show();

	});

	//首页幻灯片效果
	var len = $(".slide_items li").size(); //获取图片的数目
	var init = 1; //设置图片切换初始值，从第二张开始
	var dt = null; //设置定时器
	
	//定义一个函数完成动画
	function slide(){
		dt = setInterval(function(){
			//大图切换
			$(".slide_items li").stop(true,false).fadeOut().eq(init).fadeIn();
			//数字索引切换
			$(".slide_controls li").removeClass("on").eq(init).addClass("on");
			init++;
			if (init >= len ){
				init = 0;
			}
		},5000)
	}

	// function slide(){
	// 	if (init >= len ){
	// 		init = 0;
	// 	}
	// 	//大图切换
	// 	$(".slide_items li").fadeOut().eq(init).fadeIn();
	// 	//数字索引切换
	// 	$(".slide_controls li").removeClass("on").eq(init).addClass("on");
	// 	init++;
		
	// 	setTimeout("slide()",2000);
	// }
	 
	//调用函数，实现动画
	slide();

	//鼠标放置在图片上则停止幻灯,离开则继续
	$(".slide_items li").mouseover(function(){
		clearInterval(dt);
	}).mouseout(function(){
		slide();
	});

	//鼠标放置到数字索引上时，立即切换到该图片上,并停止动画，离开则继续
	$(".slide_controls li").mouseover(function(){
		clearInterval(dt);
		init = $(this).index();
		$(".slide_items li").stop(true,false).fadeOut().eq(init).fadeIn();
		init++;
		$(this).addClass("on").siblings().removeClass("on");
	}).mouseout(function(){
		slide();
	});

});