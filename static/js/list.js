/*
@功能：列表页js
@作者：diamondwang
@时间：2013年11月13日
*/

$(function(){
	$(".child h3").click(function(){
		$(this).toggleClass("on").parent().find("ul").toggle();
	});
});