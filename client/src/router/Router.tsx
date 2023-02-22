import { Switch, Route, } from "react-router-dom"
import { MessageInput } from "../components/MessageInput";
import { MessageList } from "../components/MessageList";
import { WebCam } from "../components/WebCam";
import { Page404 } from "../components/Pages/Page404"
import { TrialRoutes } from "./TrialRoutes";

import { memo } from "react";
import { Login } from "../components/Pages/Login";
import { HomeRoutes } from "./HomeRoutes";
import { HeaderLayout } from "../components/Pages/template/HeaderLayout";
import { LoginUserProvider } from "../providers/LoginUserProvider"
import { FolderRoutes } from "./FolderRoute";



export const Router = memo(() => {
    return(
      <Switch>
      <LoginUserProvider>
        <Route exact path="/">
          <Login />
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
          )}
      >
      </Route>
        <Route path="/home" render={({ match: { url } }) => (
          <Switch>
            {HomeRoutes.map((route) => (
              <Route
              key={route.path}
              exact={route.exact}
              path={`${url}${route.path}`}
              >
                <HeaderLayout>{route.children}</HeaderLayout>
              </Route>
            ))}
          </Switch>
        )}
        />

        <Route path="/folders" render={({ match: { url } }) => (
          <Switch>
            {FolderRoutes.map((route) => (
              <Route
              key={route.path}
              exact={route.exact}
              path={`${url}${route.path}`}
              >
                <HeaderLayout>{route.children}</HeaderLayout>
              </Route>
            ))}
          </Switch>
        )}
        />
        
        </LoginUserProvider>
        <Route path="*">
          <Page404 />
        </Route>

        
      </Switch>
    );
});

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