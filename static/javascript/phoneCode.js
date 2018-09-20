function showHint(str)
{
  showtime(60);
  var xmlhttp;
  if (str.length==0)
  { 
    document.getElementById("txtHint").innerHTML="";
    return;
  }
  if (window.XMLHttpRequest)
  {
    // IE7+, Firefox, Chrome, Opera, Safari 浏览器执行代码
    xmlhttp=new XMLHttpRequest();
  }
  else
  {
    // IE6, IE5 浏览器执行代码
    xmlhttp=new ActiveXObject("Microsoft.XMLHTTP");
  }
  xmlhttp.onreadystatechange=function()
  {
    if (xmlhttp.readyState==4 && xmlhttp.status==200)
    {
      document.getElementById("txtHint").innerHTML=xmlhttp.responseText;
    }
  }
  xmlhttp.open("GET","/index.php/home/index/sms/s/"+str,true);
  xmlhttp.send();
}
function showtime(t){
  // document.myform[1].code.disabled=true;
  document.getElementById("code").disabled=true
  for(i=1;i<=t;i++) {
    window.setTimeout("update_p(" + i + ","+t+")", i * 1000);
  }

}

function update_p(num,t) {
  if(num == t) {
    // document.myform[1].code.value =" 重新发送 ";
	document.getElementById("code").innerHTML="重新发送";
    // document.myform[1].code.disabled=false;
	document.getElementById("code").disabled=false;
  }
  else {
    printnr = t-num;
    // document.myform[1].code.value = " (" + printnr +")秒后重新发送";
	document.getElementById("code").innerHTML=" (" + printnr +")秒";
	
  }
}

 
 