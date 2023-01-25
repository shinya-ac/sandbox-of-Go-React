import React from "react";
import ReactDOM from "react-dom/client";
import { RecoilRoot } from "recoil";
import { App } from "./App";


ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <RecoilRoot>
      <App />
    </RecoilRoot>
  </React.StrictMode>
);
//2023/01/09
//Simple chatのページを開くところまでは実装できた
//ただバックエンド側の実装がメッセージ使用になりきっていないので、バックエンドの実装を変える必要がある
//からのその実装したバックエンドとこのSimpleChatのページとがwebsocketで連携できているか確認する
