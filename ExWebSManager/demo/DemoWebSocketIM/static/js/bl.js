$(document).ready(function () {

    WebMsg_T_Ping = 0;//ping 消息
    WebMsg_T_Pong = 1;//pong 消息
    WebMsg_T_Jion = 10;//加入消息
    WebMsg_T_leave = 11;//离开消息

    WebMsg_T_redirect = 302;//跳转
    WebMsg_T_Chat = 1000;//聊天

    // var sendMsg = function (e) {
    //
    // }

    const url = 'ws://' + window.location.host + '/ws/join?uname=' + $('#UID').val() + '&sid=' + $('#SID').val();
    //socket = new WebSocket('ws://' + window.location.host + '/ws/join?uname=' + $('#uname').text());
    let firstHeartbeat = true;
    let websocketHeartbeatJs = new WebsocketHeartbeatJs({
        url: url
    });

    websocketHeartbeatJs.onopen = function () {
        // addLog('connect success', 'cadetblue');
        // addLog('send massage: test', 'cadetblue');
        //websocketHeartbeatJs.send('test');
        // setTimeout(() => {
        //     addLog(`wait ${websocketHeartbeatJs.opts.pingTimeout} ms will hava '${websocketHeartbeatJs.opts.pingMsg}'`, 'cadetblue');
        // }, 1500);
    }
    websocketHeartbeatJs.onmessage = function (e) {
        var data = JSON.parse(e.data);
        console.log("onmessage:" + data);
        if (data.T == WebMsg_T_Ping && firstHeartbeat) {
            //ping信息
            firstHeartbeat = false;
        } else if (data.T == WebMsg_T_Pong) {
            firstHeartbeat = false;
            //ping信息
        } else if (data.T == WebMsg_T_redirect) {
            alert("文本" + data.T);
            $('#host11').val(data.T + window.location.host);
            window.location.href = 'http://' + window.location.host;
            //ping信息
        } else if (data.T == WebMsg_T_Chat) {
            var li = document.createElement('li');
            var username = document.createElement('strong');
            var content = document.createElement('span');

            username.innerText = "DDD";
            content.innerText = data.C;
            li.appendChild(username);
            li.appendChild(document.createTextNode(': '));
            li.appendChild(content);
            $('#chatbox li').first().before(li);
        } 

        // if (e.data == websocketHeartbeatJs.opts.pingMsg && firstHeartbeat) {
        //     // setTimeout(() => {
        //     //     addLog(`Close your network, wait ${websocketHeartbeatJs.opts.pingTimeout + websocketHeartbeatJs.opts.pongTimeout}+ ms, websocket will reconnect`, 'cadetblue');
        //     // }, 1500);
        //     firstHeartbeat = false;
        // }
    }
    // websocketHeartbeatJs.onreconnect = function () {
    //     console.log("aa_onreconnect");
    //     // addLog(`reconnecting...`, 'chocolate');
    //     // addLog(`tips: if you network closing, please open network, reconnect will success`, 'brown');
    // }


    $('#sendbtn').click(function () {
        console.log("aaa1");
        var uname = $('#uname').text();
        var content = $('#sendbox').val();

        console.log("aaa2");

        var data = {T: WebMsg_T_Chat, C: content};
        var strData = JSON.stringify(data)

        console.log("click:" + strData);
        websocketHeartbeatJs.send(strData);
        $('#sendbox').val('');
        console.log("aaa3");

    });
});