import { getJWT } from "./auth";

var socket = null;

let connect = (cb) => {

    if (socket && socket.readyState === WebSocket.OPEN) {
        console.log("Already connected");
        return;
    }
    const socketURL = `ws://localhost:8080/ws?jwt=${encodeURIComponent(getJWT())}`;
    console.log("Connecting");
    socket = new WebSocket(socketURL);

    socket.onopen = () => {
        // cb("Successfully connected")
    };

    socket.onmessage = msg => {
        cb(msg)
    };

    socket.onclose = event => {
        console.log("Socket Closed Connection: ", event);
    };

    socket.onerror = error => {
        // console.log("Socket Error: ", error);
    };
};

const cleanupWebSocket = () => {
    if (socket) {
        socket.close();
        socket = null;
    }
};


let sendMsg = (message) => {
    socket.send(JSON.stringify(message));
};

window.addEventListener('beforeunload', cleanupWebSocket);

export { connect, sendMsg };