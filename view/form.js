
function initForm(){

    // register functions to run from the backend here
    callbacks.set("printEcho",logText);
    callbacks.set("printLog",logText);
}


function logText(data){
    let div = document.createElement("div");
    div.textContent = data;
    div.classList.add("item");
    log.insertBefore(div,log.firstChild);
}

function onClickButton1(){
    exec("printEcho",edit1.value);
    edit1.value="";
}

function setConfig(param, value){
    exec("setConfig", {param:param, value:value})
}

const Id       = "id"
const FuncCode = "funcCode"
const Start    = "start"
const Quantity = "quantity"




