import { websocketAtom } from "../state/websocket";
import { messageListAtom } from "../state/messages";
import { useRecoilCallback, useRecoilValue } from "recoil";
import { Message } from "../models/message";

// 以下のmessageListAtomはグローバルなstateのことで、state/message.tsでRecoilを用いて
// 定義されているグローバルstate
// 参照は以下のようにまずconst messageList = useRecoilValue(messageListAtom);として
// stateを受け取ることができる
// useRecoilValueではなくuseRecoilStateを使えばstateの更新系のメソッドなども使うことができるようになる
// useRecoilValueは参照専用で、逆にuseSetRecoilStateというものは更新系のみを使うことを明示できるもの
// その二つを足し合わせたものが「useRecoilState」という感じ
export const useMessageList = (): Message[] => {
  const socket = useRecoilValue(websocketAtom);
  const messageList = useRecoilValue(messageListAtom);

  const updateMessageList = useRecoilCallback(
    ({ set }) =>
      (message: Message) => {
        set(messageListAtom, [...messageList, message]);
      }
  );
  socket.onmessage = (msg) => {
    const content = JSON.parse(msg.data as string);
    const message: Message = { content: content };
    updateMessageList(message);
  };

  return messageList;
};