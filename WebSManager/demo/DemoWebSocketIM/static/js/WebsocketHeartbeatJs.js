/**
 * `WebsocketHeartbeatJs` constructor.
 *
 * @param {Object} opts
 * {
 *  url                  websocket链接地址
 *  pingTimeout 未收到消息多少秒之后发送ping请求，默认15000毫秒
    pongTimeout  发送ping之后，未收到消息超时时间，默认10000毫秒
    reconnectTimeout
 * }
 * @api public
 */

function WebsocketHeartbeatJs({
    url, 
    pingTimeout = 15000,
    pongTimeout = 10000,
    reconnectTimeout = 2000,
}){
    this.opts ={
        url: url,
        pingTimeout: pingTimeout,
        pongTimeout: pongTimeout,
        reconnectTimeout: reconnectTimeout
    };
    this.ws = null;//websocket实例

    //override hook function
    this.onclose = () => {};
    this.onerror = () => {};
    this.onopen = () => {};
    this.onmessage = () => {};
    this.onreconnect = () => {};

    //this.send = () => {};

    this.createWebSocket();
}
WebsocketHeartbeatJs.prototype.createWebSocket = function(){
    try {
        console.log("createWebSocket1");
        this.ws = new WebSocket(this.opts.url);
        console.log("createWebSocket2");
        this.initEventHandle();
        console.log("createWebSocket3");
    } catch (e) {
        this.reconnect();
        throw e;
    }     
};

WebsocketHeartbeatJs.prototype.initEventHandle = function(){
    this.ws.onclose = () => {
        this.onclose();
        this.reconnect();
    };
    this.ws.onerror = () => {
        this.onerror();
        this.reconnect();
    };
    this.ws.onopen = () => {
        this.onopen();
        //心跳检测重置
        this.heartCheck();
    };
    this.ws.onmessage = (event) => {
        this.onmessage(event);
        //如果获取到消息，心跳检测重置
        //拿到任何消息都说明当前连接是正常的
        this.heartCheck();
    };
};

WebsocketHeartbeatJs.prototype.reconnect = function(){
    if(this.lockReconnect || this.forbidReconnect) return;
    this.lockReconnect = true;
    console.log("reconnect1");
    this.onreconnect();
    console.log("reconnect2");
    //没连接上会一直重连，设置延迟避免请求过多
    setTimeout(() => {
        this.createWebSocket();
        this.lockReconnect = false;
    }, this.opts.reconnectTimeout);
};
WebsocketHeartbeatJs.prototype.send = function(msg){
    console.log("send1");
    this.ws.send(msg);
    console.log("send2");
};
//心跳检测
WebsocketHeartbeatJs.prototype.heartCheck = function(){
    console.log("heartCheck1");
    this.heartReset();
    this.heartStart();
    console.log("heartCheck2");
};
WebsocketHeartbeatJs.prototype.heartStart = function(){
    console.log("heartStart1");
    if(this.forbidReconnect) return;//不再重连就不再执行心跳
    console.log("heartStart2");
    this.pingTimeoutId = setTimeout(() => {
        console.log("heartStart3");
        //这里发送一个心跳，后端收到后，返回一个心跳消息，
        //onmessage拿到返回的心跳就说明连接正常
        var pingMsg = {t: 0, c: ""};

        var strData = JSON.stringify(pingMsg)
        this.ws.send(strData);
        console.log("heartStart3");
        //如果超过一定时间还没重置，说明后端主动断开了
        this.pongTimeoutId = setTimeout(() => {
            //如果onclose会执行reconnect，我们执行ws.close()就行了.如果直接执行reconnect 会触发onclose导致重连两次
            this.ws.close();
            console.log("heartStart4");
        }, this.opts.pongTimeout);
        console.log("heartStart5");
    }, this.opts.pingTimeout);
};
WebsocketHeartbeatJs.prototype.heartReset = function(){
    console.log("heartReset1");
    clearTimeout(this.pingTimeoutId);
    console.log("heartReset2");
    clearTimeout(this.pongTimeoutId);
    console.log("heartReset3");
};
WebsocketHeartbeatJs.prototype.close = function(){
    console.log("close1");
    //如果手动关闭连接，不再重连
    this.forbidReconnect = true;
    this.heartReset();
    console.log("close2");
    this.ws.close();
    console.log("close3");
};
if(window) window.WebsocketHeartbeatJs = WebsocketHeartbeatJs;