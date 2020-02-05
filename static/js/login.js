window.onload = function(){	
    var name = document.getElementById('Uname');
    var pwd = document.getElementById('Upwd');
    var btn = document.getElementsByTagName('button')[0];
    
    var kpName = "Orianna";
    var kpPwd = 123456;
    
    btn.onclick = function(){
        if(name.value==kpName && pwd.value==kpPwd){
            document.f.action="https://www.baidu.com";
            document.f.submit();
        }else{
            alert("passwords error");
        }
    }   
}
