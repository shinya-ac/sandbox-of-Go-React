import { Switch, Route, } from "react-router-dom"
import { MessageInput } from "../components/MessageInput";
import { MessageList } from "../components/MessageList";
import { WebCam } from "../components/WebCam";
import { Page404 } from "../components/Page404"
import { TrialRoutes } from "./TrialRoutes";



export const Router = () => {
    return(
        <Switch>
      <Route exact path="/chat">
        <h1>Simple Chat</h1>
        <MessageInput />
        <MessageList />
      </Route>
      <Route exact path="/webcam">
        <WebCam />
      </Route>
      <Route 
      path="/trial"
      render={( { match: { url } } ) => (
        <Switch>
          { TrialRoutes.map((route) => (
            <Route
                key={route.path}
                exact={route.exact}
                path={`${url}${route.path}`}
            >
                {route.children}
            </Route>
          )) }
        </Switch>
      )
      }>
      </Route>
      {/* 「*」は全ての文字列に対して、を意味するので、上から順にルーティングが検索されて、
      そのどれにも当てはまらないけど何かしらのリクエストは来ているという場合は
      この*が反応して404ページに遷移するという仕組みになっている */}
      <Route path="*">
        <Page404 />
      </Route>
    </Switch>
    )
}

//以下のコードが元々このファイルのSwitch内に記載されていた
/* <Route exact path="/trial">
    <TrialReact />
    </Route>
    <Route exact path="/trial/1">
    <TrialA />
    </Route>
    <Route exact path="/trial/2">
    <TrialB />
</Route> */
//→TrialRouter.tsxというファイルに配列形式で切り離した
//→なのでこのファイルではその配列形式のやつをインポートしてきて展開している