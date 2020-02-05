function checkInput(){
    var uname = document.getElementById('uname');
    var upwd = document.getElementById('upwd');
    if(uname.value.length == 0){
        alert("Account number cannot be empty!");
        return false;
    }

    if(upwd.value.length == 0){
        alert("Password cannot be empty!");
        return false;
    }

    return true;
}