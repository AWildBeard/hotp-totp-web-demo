package main

const (
	setupOTPHTMLPage = `
<!DOCTYPE html>
<html><head><meta name="GCD" content="YTk3ODQ3ZWZhN2I4NzZmMzBkNTEwYjJld111d8e97c346b3a553e3f12a8b6b9ef"/>
  <meta charset="utf-8">
  <title>Setup!</title>
  <meta name="generator" content="Google Web Designer 5.0.4.0226">
  <style type="text/css" id="gwd-text-style">p{margin:0px}h1{margin:0px}h2{margin:0px}h3{margin:0px}</style>
  <style type="text/css">html,body{width:100%;height:100%;margin:0px}body{transform:matrix3d(1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1);perspective:1400px;transform-style:preserve-3d;background-image:none;background-color:#494dda}.gwd-div-13gl{position:absolute;background-image:none;background-color:#fff;border-radius:8px;width:69.33%;height:77.72%;transform-origin:417.016px 278.736px 0px;left:15.335%;top:11.14%}.gwd-span-3d7e{position:absolute;font-family:Verdana;left:0%;transform-origin:35.7344px 12.9219px 0px;height:100%;width:15.94%;top:0%}.gwd-span-zfvq{position:absolute;font-family:Verdana;text-align:right;height:100%;width:14.39%;transform-origin:32.7344px 16.5767px 0px;top:0%;left:85.61%}.gwd-div-13i1{position:absolute;transform-origin:234.109px 13.5188px 0px;width:56.23%;height:5.91%;top:0%;margin:3% 0%;left:21.885%}[data-gwd-group="qr-images"] .gwd-grp-1ihg.gwd-img-o0cf{position:absolute;top:50%;transform-style:preserve-3d;border-style:solid;border-width:4px;left:0%;transform:translate3d(0, -104px, 0);width:200px;height:200px}[data-gwd-group="qr-images"] .gwd-grp-1ihg.gwd-img-1aa2{position:absolute;width:200px;height:200px;top:50%;transform-style:preserve-3d;border-style:solid;border-width:4px;left:100%;transform:translate3d(-208px, -104px, 0)}[data-gwd-group="qr-images"]{width:834.031px;height:457px}.gwd-div-xwrp{position:absolute;width:77.69%;margin:10% 0%;height:37.63%;transform-origin:323.969px 85.4517px 0px;top:0%;left:11.155%}[data-gwd-group="io"] .gwd-grp-1jcx.gwd-input-1a1j{position:absolute;top:0%;width:35%;height:17%;padding:2px 6px;background-image:none;background-color:#fff;border-image-source:none;border-image-width:1;border-image-outset:0;border-image-repeat:stretch;border-color:#4294e8;border-style:solid;border-radius:3px;left:0%}[data-gwd-group="io"] .gwd-grp-1jcx.gwd-input-n9s8{position:absolute;transform-style:preserve-3d;top:0%;width:34.77%;height:17.24%;border-image-source:none;border-image-width:1;border-image-outset:0;border-image-repeat:stretch;border-color:#4294e8;border-style:solid;border-radius:3px;padding:2px 6px;transform-origin:97.0071px 13px 0px;background-image:none;background-color:#fff;left:65.23%;transform:translate3d(-16px, 0, 0)}[data-gwd-group="io"] .gwd-grp-1jcx.gwd-button-17uu{position:absolute;transform-style:preserve-3d;top:100%;border-style:solid;border-image-source:none;border-image-width:1;border-image-outset:0;border-image-repeat:stretch;border-color:#4294e8;border-radius:3px;width:18.3%;background-image:none;background-color:#fff;height:25%;left:81.7%;transform:translate3d(0, -24px, 0)}[data-gwd-group="io"] .gwd-grp-1jcx.gwd-button-l9za{position:absolute;left:0%;top:100%;transform:translate3d(0, -24px, 0);transform-style:preserve-3d;width:18.3%;background-image:none;background-color:#fff;border-style:solid;border-image-source:none;border-image-width:1;border-image-outset:0;border-image-repeat:stretch;border-color:#4294e8;border-radius:3px;height:25%}[data-gwd-group="io"]{width:546.469px;height:116.016px}.gwd-div-tkx9{position:absolute;transform-style:preserve-3d;top:64.87%;height:16.85%;width:77.69%;transform-origin:324.217px 46.9688px 0px;left:11.155%;transform:translate3d(0, -10px, 0)}.gwd-button-9j5c{position:absolute;transform-style:preserve-3d;top:93.9%;width:11.99%;height:4.3%;border-style:solid;border-image-source:none;border-image-width:1;border-image-outset:0;border-image-repeat:stretch;border-color:#4294e8;background-image:none;background-color:#fff;border-radius:3px;left:44.005%;transform:translate3d(0, -37px, 0)}input[type="text"]{outline:none}input[type="text"]:focus{box-shadow:#4294e8 0 0 5px}</style>
  
  <script type="text/javascript" gwd-gwdid-version="1.0">(function(){var c;var f=function(a){this.c="";this.a=[];this.b="";a&&("string"==typeof a?this.h(a):this.g(a))};c=f.prototype;c.i=function(a){for(var b=this.c==a.c&&this.b==a.b&&this.a.length==a.a.length,d=0;b&&d<this.a.length;d++)b=this.a[d]==a.a[d];return b};c.u=function(){return""==this.f()};c.clear=function(){this.c="";this.a=[];this.b=""};
c.g=function(a){this.clear();if(!a)return!1;var b=a.getAttribute("data-gwd-group-def")||"",d=b?[]:[a];for(a=a.parentElement;a&&!b;)a.hasAttribute("data-gwd-group")?d.push(a):b=a.getAttribute("data-gwd-group-def"),a=a.parentElement;d=d.reverse();a=!!b;var g=d[0];if(!a&&!g.id)return!1;for(var e=[],m=a?0:1;m<d.length;m++){var n=d[m].getAttribute("data-gwd-grp-id");if(!n)return!1;e.push(n)}this.c=a?"":g.id;this.a=e;this.b=b||"";return!0};
c.h=function(a){this.clear();a=a.split(" ");for(var b=""==a[0],d="",g=[],e=b?1:0;e<a.length;e++)if(a[e])d?g.push(a[e]):d=a[e];else return!1;if(!d)return!1;b?this.b=d:this.c=d;this.a=g;return!0};
c.j=function(a){if(this.b)a=a.querySelector('[data-gwd-group-def="'+this.b+'"]');else if(this.c){if(a=a.getElementById(this.c),!a){var b=document.querySelector("gwd-pagedeck")||document.querySelector("[is=gwd-pagedeck]");b&&"function"==typeof b.getElementById&&(a=b.getElementById(this.c))}}else return null;return h(this,a)};c.l=function(a,b){if(!this.b)return null;for(a=b;a;){var d=a.getAttribute("data-gwd-group");if(d==this.b)break;a=d&&a!=b?null:a.parentElement}return a?h(this,a):null};
c.f=function(){var a="";this.b?a=" "+this.b:this.c&&(a=this.c);for(var b=0;b<this.a.length;b++)a+=" "+this.a[b];return a};c.o=function(){return this.c};c.s=function(){return this.a};c.m=function(){return this.b};var h=function(a,b){for(var d=0;b&&d<a.a.length;d++){for(var g=b.querySelector('[data-gwd-grp-id="'+a.a[d]+'"]'),e=g?g.parentElement:null;g&&e&&e!=b;)e.hasAttribute("data-gwd-group")?g=null:e=e.parentElement;b=g}return b};var k=["gwd","GwdId"],l=this;k[0]in l||"undefined"==typeof l.execScript||l.execScript("var "+k[0]);for(var p;k.length&&(p=k.shift());)k.length||void 0===f?l=l[p]&&l[p]!==Object.prototype[p]?l[p]:l[p]={}:l[p]=f;f.GROUP_REFERENCE_ATTR="data-gwd-group";f.prototype.equals=f.prototype.i;f.prototype.clear=f.prototype.clear;f.prototype.isEmpty=f.prototype.u;f.prototype.setFromElement=f.prototype.g;f.prototype.setFromString=f.prototype.h;f.prototype.getElement=f.prototype.j;
f.prototype.getElementInInstance=f.prototype.l;f.prototype.getString=f.prototype.f;f.prototype.getLeadingId=f.prototype.o;f.prototype.getPseudoIds=f.prototype.s;f.prototype.getGroupName=f.prototype.m;}).call(this);
</script>
  <script type="text/javascript" gwd-events="support">var gwd=gwd||{};gwd.actions=gwd.actions||{};gwd.actions.events=gwd.actions.events||{};gwd.actions.events.getElementById=function(id){var gwdId=new gwd.GwdId(id);var retElement;var target=gwd.actions.events.currentTarget;if(!gwdId.getLeadingId()&&target){retElement=gwdId.getElementInInstance(document,gwd.actions.events.currentTarget)}else{retElement=gwdId.getElement(document)}if(!retElement){switch(id){case"document.body":retElement=document.body;break;case"document":retElement=document;break;case"window":retElement=window;break}}return retElement};gwd.actions.events.addHandler=function(eventTarget,eventName,eventHandler,useCapture){var gwdId=new gwd.GwdId(eventTarget);var groupName=gwdId.getGroupName();if(groupName!==""){var instances=document.querySelectorAll("["+gwd.GwdId.GROUP_REFERENCE_ATTR+" = "+groupName+"]");Array.prototype.forEach.call(instances,function(instance){var target=gwdId.getElementInInstance(document,instance);if(target){var actualHandlerProp=gwd.actions.events.actualHandlerProp;if(!eventHandler[actualHandlerProp]){eventHandler[actualHandlerProp]=gwd.actions.events.wrapHandler_.bind(null,eventHandler)}target.addEventListener(eventName,eventHandler[actualHandlerProp],useCapture)}})}else{var targetElement=gwd.actions.events.getElementById(eventTarget);if(targetElement){targetElement.addEventListener(eventName,eventHandler,useCapture)}}};gwd.actions.events.removeHandler=function(eventTarget,eventName,eventHandler,useCapture){var gwdId=new gwd.GwdId(eventTarget);var groupName=gwdId.getGroupName();if(groupName!==""){var instances=document.querySelectorAll("["+gwd.GwdId.GROUP_REFERENCE_ATTR+" = "+groupName+"]");Array.prototype.forEach.call(instances,function(instance){var target=gwdId.getElementInInstance(document,instance);if(target){var actualHandlerProp=gwd.actions.events.actualHandlerProp;if(eventHandler[actualHandlerProp]){target.removeEventListener(eventName,eventHandler[actualHandlerProp],useCapture)}}})}else{var targetElement=gwd.actions.events.getElementById(eventTarget);if(targetElement&&eventTarget[0]!=" "){targetElement.removeEventListener(eventName,eventHandler,useCapture)}}};gwd.actions.events.setInlineStyle=function(id,styles){var element=gwd.actions.events.getElementById(id);if(!element||!styles){return}var transitionProperty=element.style.transition!==undefined?"transition":"-webkit-transition";var prevTransition=element.style[transitionProperty];var splitStyles=styles.split(/\s*;\s*/);var nameValue;splitStyles.forEach(function(splitStyle){if(splitStyle){var regex=new RegExp("[:](?![/]{2})");nameValue=splitStyle.split(regex);nameValue[1]=nameValue[1]?nameValue[1].trim():null;if(!(nameValue[0]&&nameValue[1])){return}element.style.setProperty(nameValue[0],nameValue[1])}});function restoreTransition(event){var el=event.target;el.style.transition=prevTransition;el.removeEventListener(event.type,restoreTransition,false)}element.addEventListener("transitionend",restoreTransition,false);element.addEventListener("webkitTransitionEnd",restoreTransition,false)};gwd.actions.events.currentTarget=null;gwd.actions.events.actualHandlerProp="gwd_actualHandler";gwd.actions.events.wrapHandler_=function(handler,event){gwd.actions.events.currentTarget=event.target;handler.call(null,event)}</script>
  <script type="text/javascript" gwd-events="handlers">gwd.auto__Io_Test_hotp_buttonMouseover=function(event){gwd.actions.events.setInlineStyle(" io test-hotp-button","transition: all 0.2s ease; background-color: #4294e8;")};gwd.auto__Io_Test_hotp_buttonMouseleave=function(event){gwd.actions.events.setInlineStyle(" io test-hotp-button","transition: all 0.2s ease; background-color:  #ffffff;")};gwd.auto__Io_Test_totp_buttonMouseover=function(event){gwd.actions.events.setInlineStyle(" io test-totp-button","transition: all 0.2s ease; background-color:  #4294e8;")};gwd.auto__Io_Test_totp_buttonMouseleave=function(event){gwd.actions.events.setInlineStyle(" io test-totp-button","transition: all 0.2s ease; background-color:  #ffffff;")};gwd.auto_Finish_buttonMouseover=function(event){gwd.actions.events.setInlineStyle("finish_button","transition: all 0.2s ease; background-color:  #4294e8;")};gwd.auto_Finish_buttonMouseleave=function(event){gwd.actions.events.setInlineStyle("finish_button","transition: all 0.2s ease; background-color:  #ffffff;")}</script>
  <script type="text/javascript" gwd-events="registration">gwd.actions.events.registerEventHandlers=function(event){gwd.actions.events.addHandler(" io test-hotp-button","mouseover",gwd.auto__Io_Test_hotp_buttonMouseover,false);gwd.actions.events.addHandler(" io test-hotp-button","mouseleave",gwd.auto__Io_Test_hotp_buttonMouseleave,false);gwd.actions.events.addHandler(" io test-totp-button","mouseover",gwd.auto__Io_Test_totp_buttonMouseover,false);gwd.actions.events.addHandler(" io test-totp-button","mouseleave",gwd.auto__Io_Test_totp_buttonMouseleave,false);gwd.actions.events.addHandler("finish_button","mouseover",gwd.auto_Finish_buttonMouseover,false);gwd.actions.events.addHandler("finish_button","mouseleave",gwd.auto_Finish_buttonMouseleave,false)};gwd.actions.events.deregisterEventHandlers=function(event){gwd.actions.events.removeHandler(" io test-hotp-button","mouseover",gwd.auto__Io_Test_hotp_buttonMouseover,false);gwd.actions.events.removeHandler(" io test-hotp-button","mouseleave",gwd.auto__Io_Test_hotp_buttonMouseleave,false);gwd.actions.events.removeHandler(" io test-totp-button","mouseover",gwd.auto__Io_Test_totp_buttonMouseover,false);gwd.actions.events.removeHandler(" io test-totp-button","mouseleave",gwd.auto__Io_Test_totp_buttonMouseleave,false);gwd.actions.events.removeHandler("finish_button","mouseover",gwd.auto_Finish_buttonMouseover,false);gwd.actions.events.removeHandler("finish_button","mouseleave",gwd.auto_Finish_buttonMouseleave,false)};document.addEventListener("DOMContentLoaded",gwd.actions.events.registerEventHandlers);document.addEventListener("unload",gwd.actions.events.deregisterEventHandlers)</script>
  <script type="text/javascript">function testHotp() {
  var email = "{{ .Email }}";
  var key = document.getElementById("hotp-input-field").value;
  var json = {"email": email, key: key};
  var xhr = new XMLHttpRequest();
  
  xhr.open("post", "/testHotp", true);
  xhr.setRequestHeader('Content-Type', 'application/json');
  
  xhr.onreadystatechange = function() {
    if (xhr.readyState === 4 && xhr.status === 200) {
      var data = JSON.parse(xhr.responseText);
      console.log(data.status);
      if (data.status === "success") {
        alert("HOTP SUCCESS!")
      } else {
        alert("HOTP FAILURE!")
      }
    }
  }
  
  var data = JSON.stringify(json);
  console.log("Sending: " + data);
  xhr.send(data);
}

function testTotp() {
  var email = "{{ .Email }}";
  var key = document.getElementById("totp-input-field").value;
  var json = {"email": email, key: key};
  var xhr = new XMLHttpRequest();
  
  xhr.open("post", "/testTotp", true);
  xhr.setRequestHeader('Content-Type', 'application/json');
  
  xhr.onreadystatechange = function() {
    if (xhr.readyState === 4 && xhr.status === 200) {
      var data = JSON.parse(xhr.responseText);
      if (data.status === "success") {
        alert("TOTP SUCCESS!")
      } else {
        alert("TOTP FAILURE!")
      }
    } 
  }
  
  var data = JSON.stringify(json);
  console.log("Sending: " + data);
  xhr.send(data);
}
</script>
</head>

<body class="htmlNoPages">
  
  <div class="gwd-div-13gl">
    <div class="gwd-div-13i1"><span class="gwd-span-3d7e">HOTP<br class=""></span><span class="gwd-span-zfvq">TOTP<br></span>
    </div>
    <div class="gwd-div-xwrp" data-gwd-group="qr-images">
      <img src="{{ .HotpImagePath }}" class="gwd-img-o0cf gwd-grp-1ihg" alt="hotp-image">
      <img src="{{ .TotpImagePath }}" class="gwd-img-1aa2 gwd-grp-1ihg" alt="totp-image">
    </div>
    <div class="gwd-div-tkx9" data-gwd-group="io">
      <input type="text" id="hotp-input-field" class="gwd-input-1a1j gwd-grp-1jcx" data-gwd-grp-id="text_1" placeholder="HOTP Code" data-gwd-name="hotp-input-field">
      <input type="text" id="totp-input-field" class="gwd-input-n9s8 gwd-grp-1jcx" data-gwd-grp-id="totp-input-field" placeholder="TOTP Code" data-gwd-name="totp-input-field">
      <button id="test-totp-button" onclick="testTotp();" class="gwd-button-17uu gwd-grp-1jcx" data-gwd-grp-id="test-totp-button" data-gwd-name="test-totp-button">Test</button>
      <button id="test-hotp-button" onclick="testHotp();" class="gwd-button-l9za gwd-grp-1jcx" data-gwd-grp-id="test-hotp-button" data-gwd-name="test-hotp-button">Test</button>
    </div>
    <button id="finish_button" onclick="location.href='/logout';" class="gwd-button-9j5c" data-gwd-name="finished_button">Done!</button>
  </div>


</body></html>

`
)
