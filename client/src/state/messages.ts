import { Message } from "../models/message";
import { atom } from "recoil";

// グローバルなstateの管理方法で、useState+useContextよりも最新の方法がこのRecoil
// keyにはそのstateをグローバルで一意に示せる名前をつける
// defaultにはそのstateの初期値を入れる
// このRecoilの参照はhooks/use-message-list.tsで行っているのでそちらを参考に
// ただし前提条件としてmain.tsxで「RecoilRoot」で囲っているのでその範囲の中で利用することができる
// （今回は全範囲を囲っているのでどこからでも参照できるようになっている）
export const messageListAtom = atom<Message[]>({
  key: "messageList",
  default: [],
});