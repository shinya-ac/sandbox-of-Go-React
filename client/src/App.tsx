import { BrowserRouter, Link } from "react-router-dom"
import { UserProvider } from "./providers/UserProvider";
import React, {createContext} from "react"
import { Router } from "./router/Router";

//以下のUserProviderはproviders/UserProviderファイルの関数のことで、
//そこで利用したUserContext.Providerの値（contextName）を下で囲った範囲全てで
//使うことができる（<UserProvider> ~ </UserProvider>までの全てなので今回は全範囲で使えるようにしている）
export const App = () => {
  return (
    <UserProvider>
      <BrowserRouter>
      <div>
        <Link to="/chat">Chat</Link>
        <br />
        <Link to="/webcam">WebCam</Link>
        <br />
        <Link to="/trial">TrialReact</Link>
      </div>
      < Router />
      </BrowserRouter>
    </UserProvider>
  );
};