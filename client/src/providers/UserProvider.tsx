import React, { useState, createContext } from "react";

export const UserContext = createContext({});

interface Props {
    children: React.ReactNode;
  }

//UserContext.Providerで囲った配下全てでその値（value）を参照できるようになる
//propsはLintエラーが出てるけど一旦無視でおけ
export const UserProvider : React.FC<Props> = (props) => {
    const { children } = props;

    //useState + createContext(つまりUserContext)の合わせ技で読み込みも書き込みも双方向に
    //行えるグローバルなstateを以下のように扱えるようになる
    const [ userInfo, setUserInfo ] = useState<any>(null);

    const contextName = "Mr Hogeter Contexter"// ←固定値でグローバルなstateを使う場合はこんな感じ
    return(
        <UserContext.Provider value={{contextName, userInfo, setUserInfo}}>
            {children}
        </UserContext.Provider>
    )
}