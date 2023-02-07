import { atom, selector } from "recoil";
import * as WebSocket from "websocket";

const connect = (): Promise<WebSocket.w3cwebsocket> => {
  return new Promise((resolve, reject) => {
    const socket = new WebSocket.w3cwebsocket("ws://localhost:8080/ws");
    socket.onopen = () => {
      console.log("web socket connected by react");
      resolve(socket);
    };
    socket.onclose = () => {
      console.log("reconnecting...");
      connect();
    };
    socket.onerror = (err) => {
      console.log("connection error:", err);
      reject(err);
    };
  });
};

const connectWebsocketSelector = selector({
  key: "connectWebsocket",
  get: async (): Promise<WebSocket.w3cwebsocket> => {
    return await connect();
  },
});

// 以下はRecoilで「websocket」というグローバルなstateを定義しているコード
// 初期値には「connectWebsocketSelector」を入れている
export const websocketAtom = atom<WebSocket.w3cwebsocket>({
  key: "websocket",
  default: connectWebsocketSelector,
});