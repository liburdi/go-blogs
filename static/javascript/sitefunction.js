/**
 * Created by Administrator on 2016/9/12.
 */

/*
 * tabs_name:
 * contents_name:
 * tabs_checked_style:
 * contents_checked_style:
 *
 * classList.toggle();
 *
 * */
function Tabs(tabs_name, contents_name, tabs_checked_style, contents_checked_style){
    var tabs = document.querySelectorAll(tabs_name),
        contents = document.querySelectorAll(contents_name);
		var e_mark=0;
    for (var i = 0, len = tabs.length; i < len; i++){
		tabs[i].num = i;
        tabs[i].onclick = function () {
			tabs[e_mark].classList.toggle(tabs_checked_style);
            tabs[this.num].classList.toggle(tabs_checked_style);
            contents[e_mark].classList.toggle(contents_checked_style);
            contents[this.num].classList.toggle(contents_checked_style);
            e_mark = this.num;
           
        };
    }
}
function Tabs_img(tabs_name, contents_name, tabs_checked_style, contents_checked_style){
    var tabs = document.querySelectorAll(tabs_name),
        contents = document.querySelectorAll(contents_name);
		var e_mark_img=0;
    for (var i = 0, len = tabs.length; i < len; i++){
		tabs[i].num = i;
        tabs[i].onclick = function () {
			tabs[e_mark_img].classList.toggle(tabs_checked_style);
            tabs[this.num].classList.toggle(tabs_checked_style);
            contents[e_mark_img].classList.toggle(contents_checked_style);
            contents[this.num].classList.toggle(contents_checked_style);
            e_mark_img = this.num;
           
        };
    }
}

function add_show_more() {
	var $this = $(this);
	var pHeight = 0;
	$this.find('p').each(function () {
		pHeight += $(this).height();
	});
	if (pHeight > $this.height()) {
		var show_more = $this.find('.show_more');
		if (show_more.length == 0) {
			$this.append('<div class="show_more"> &DownTeeArrow;  展开</div>')
		}
	}
 }
 function htmlEscape(str) {
    return String(str)
            .replace(/&/g, '&amp;')
            .replace(/"/g, '&quot;')
            .replace(/'/g, '&#39;')
            .replace(/</g, '&lt;')
            .replace(/>/g, '&gt;');
}
function GetUrlRelativePath(mode){
var url = document.location.toString();
var arrUrl = url.split("//");
var start = arrUrl[1].indexOf("/");
var relUrl = arrUrl[1].substring(start);
		if(mode==1){
		if(relUrl.indexOf("?") != -1){
		relUrl = '?'+relUrl.split("?")[1];
		}
		}else{
			if(relUrl.indexOf("?") != -1){
		relUrl =relUrl.split("?")[0];
	}
		}
return relUrl;
}
function tabchange(){
		var url=decodeURI(GetUrlRelativePath(1));
		var tab=$('.tab');
		
		
		tab.each(function(){
	
			console.log($(this).attr('href'));
			
			if($(this).attr('href')==url){
				
				$(this).addClass('tab_current');
			}else{
				$(this).removeClass('tab_current');
			}
		
		
		});
}
function GetQueryString(name)
{
     var reg = new RegExp("(^|&)"+ name +"=([^&]*)(&|$)");
     var r = window.location.search.substr(1).match(reg);
     if(r!=null)return  unescape(r[2]); return null;
}