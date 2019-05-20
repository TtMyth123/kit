!function (e, t) {
    if ("object" == typeof exports && "object" == typeof module) module.exports = t();
    else if ("function" == typeof define && define.amd) define([], t); else {
        var o = t();
        for (var n in o) ("object" == typeof exports ? exports : e)[n] = o[n]
    }
}(window, function () {
    return function (e) {
        var t = {};

        function o(n) {
            if (t[n]) return t[n].exports;
            var i = t[n] = {i: n, l: !1, exports: {}};
            return e[n].call(i.exports, i, i.exports, o), i.l = !0, i.exports
        }

        return o.m = e, o.c = t, o.d = function (e, t, n) {
            o.o(e, t) || Object.defineProperty(e, t, {enumerable: !0, get: n})
        }, o.r = function (e) {
            "undefined" != typeof Symbol && Symbol.toStringTag && Object.defineProperty(e, Symbol.toStringTag, {value: "Module"}), Object.defineProperty(e, "__esModule", {value: !0})
        }, o.t = function (e, t) {
            if (1 & t && (e = o(e)), 8 & t) return e;
            if (4 & t && "object" == typeof e && e && e.__esModule) return e;
            var n = Object.create(null);
            if (o.r(n), Object.defineProperty(n, "default", {
                enumerable: !0,
                value: e
            }), 2 & t && "string" != typeof e) for (var i in e) o.d(n, i, function (t) {
                return e[t]
            }.bind(null, i));
            return n
        }, o.n = function (e) {
            var t = e && e.__esModule ? function () {
                return e.default
            } : function () {
                return e
            };
            return o.d(t, "a", t), t
        }, o.o = function (e, t) {
            return Object.prototype.hasOwnProperty.call(e, t)
        }, o.p = "", o(o.s = 0)
    }([function (e, t, o) {
        "use strict";

        function n(e) {
            var t = e.url, o = e.pingTimeout, n = void 0 === o ? 15e3 : o, i = e.pongTimeout,
                r = void 0 === i ? 1e4 : i, c = e.reconnectTimeout, s = void 0 === c ? 2e3 : c, u = e.pingMsg,
                f = void 0 === u ? "heartbeat" : u;
            this.opts = {
                url: t,
                pingTimeout: n,
                pongTimeout: r,
                reconnectTimeout: s,
                pingMsg: f
            }, this.ws = null, this.onclose = function () {
            }, this.onerror = function () {
            }, this.onopen = function () {
            }, this.onmessage = function () {
            }, this.onreconnect = function () {
            }, this.createWebSocket()
        }

        Object.defineProperty(t, "__esModule", {value: !0}), n.prototype.createWebSocket = function () {
            try {
                this.ws = new WebSocket(this.opts.url), this.initEventHandle()
            } catch (e) {
                throw this.reconnect(), e
            }
        }, n.prototype.initEventHandle = function () {
            var e = this;
            this.ws.onclose = function () {
                e.onclose(), e.reconnect()
            }, this.ws.onerror = function () {
                e.onerror(), e.reconnect()
            }, this.ws.onopen = function () {
                e.onopen(), e.heartCheck()
            }, this.ws.onmessage = function (t) {
                e.onmessage(t), e.heartCheck()
            }
        }, n.prototype.reconnect = function () {
            var e = this;
            this.lockReconnect || this.forbidReconnect || (this.lockReconnect = !0, this.onreconnect(), setTimeout(function () {
                e.createWebSocket(), e.lockReconnect = !1
            }, this.opts.reconnectTimeout))
        }, n.prototype.send = function (e) {
            this.ws.send(e)
        }, n.prototype.heartCheck = function () {
            this.heartReset(), this.heartStart()
        }, n.prototype.heartStart = function () {
            var e = this;
            this.forbidReconnect || (this.pingTimeoutId = setTimeout(function () {
                e.ws.send(e.opts.pingMsg), e.pongTimeoutId = setTimeout(function () {
                    e.ws.close()
                }, e.opts.pongTimeout)
            }, this.opts.pingTimeout))
        }, n.prototype.heartReset = function () {
            clearTimeout(this.pingTimeoutId), clearTimeout(this.pongTimeoutId)
        }, n.prototype.close = function () {
            this.forbidReconnect = !0, this.heartReset(), this.ws.close()
        }, window && (window.WebsocketHeartbeatJs = n), t.default = n
    }])
});