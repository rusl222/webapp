
function initForm(){

    // register functions to run from the backend here
    callbacks.set("printEcho",logText);
}


function logText(data){
    let div = document.createElement("div");
    div.textContent = data;
    div.classList.add("item");
    log.appendChild(div);
}

function onClickButton1(){
    exec("printEcho",edit1.value);
    edit1.value="";
}



