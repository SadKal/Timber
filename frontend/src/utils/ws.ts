import { getJWT } from "./auth";

var socket = null;

let connect = (cb) => {
    console.log("Attempting Connection...");

    // Construct WebSocket URL with query parameter for username
    const socketURL = `ws://localhost:8080/ws?jwt=${encodeURIComponent(getJWT())}`;
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
        console.log("Socket Error: ", error);
    };
};

let sendMsg = msg => {
    // console.log("sending msg: ", msg);
    socket.send(msg);
};

export { connect, sendMsg };