import { getJWT } from "./auth";
const backend_url  = import.meta.env.VITE_BACKEND_URL_WS;

var socket = null;

let connect = (callback) => {

    if (socket && socket.readyState === WebSocket.OPEN) {
        console.log("Already connected");
        return;
    }
    const socketURL = `${backend_url}/ws?jwt=${encodeURIComponent(getJWT())}`;
    socket = new WebSocket(socketURL);

    socket.onopen = () => {
        // callback("Successfully connected")
    };

    socket.onmessage = msg => {
        callback(msg)
    };

    socket.onclose = event => {
        console.log("Socket Closed Connection: ", event);
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