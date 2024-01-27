
// registered functions to run from the backend
var callbacks;

window.onload = init;

function init(){
    callbacks = new Map();
    initForm();
    initWs();
}

var conn;  
function initWs(){
  
    if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + document.location.host + "/ws");
        conn.onclose = function (evt) {
            "<b>Application closed.</b>"
            document.getElementById("connState").innerHTML="<b>Application closed.</b>";
        };
        conn.onmessage = function (evt) {
            let message = JSON.parse(evt.data);

            callbacks.get(message.action)(message.data);
        };
    } else {
        document.getElementById("connState").innerHTML="<b>Brouser not support</b>";
    }
}

function exec(command, data){
    conn.send(JSON.stringify({action: command, data: data}));
}
