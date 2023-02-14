import { Link, useHistory } from "react-router-dom";
import axios from "axios";
import React, { useState, useContext } from 'react';
import { UserContext } from "../providers/UserProvider";

export const TrialReact = () => {
    // フォームの入力値を保持する state
  const [username, setUserName]: [string, React.Dispatch<React.SetStateAction<string>>] = useState("hogeter");
  const [email, setEmail]: [string, React.Dispatch<React.SetStateAction<string>>] = useState("hogeter@example.com");
  const [password, setPassword]: [string, React.Dispatch<React.SetStateAction<string>>] = useState("hogeh0ge");
  

  const handleEmailChange = (e: React.ChangeEvent<HTMLInputElement>) =>{
    setEmail(e.target.value);
  };
  const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) =>{
    setPassword(e.target.value);
  };
  const handleUserNameChange = (e: React.ChangeEvent<HTMLInputElement>) =>{
    setUserName(e.target.value);
  };

  // 以下はグローバルなstateであるuserInfo(UserProvider.tsx内)の状態を変化させる
  // setUserInfoメソッドをグローバルstateを扱うUseContextから取ってきているコード
  //stateはLintエラーが出てるけど一旦無視でおけ
  interface UserContextType {
    userInfo: any;
    setUserInfo: (info: any) => void;
  }
  
  const UserContext = React.createContext<UserContextType>({
    userInfo: {},
    setUserInfo: () => {},
  });

  const { userInfo, setUserInfo } = useContext(UserContext);

  interface UserData {
    username: string;
  }

  const onClickSignUp = (username: string, email: string, password: string) => {
    axios.post("http://localhost:8080/signup", {username, email, password})
    .then((result: any) => {
        console.log(userInfo);
        console.log(result.data);
        console.log(result.data.username);
        setUserName(result.data.username);
        setUserInfo({username: "signUped Hogeter"});
        console.log(result.data);
        console.log(userInfo.username);
    })
    .catch((err) => {console.log(err)});
};
  const handleSubmit = (e: any) =>{
    e.preventDefault();
    onClickSignUp(username, email, password)
  };

  //無意味な100件の配列
  const arr = [...Array(100).keys()];

  //素のjsでReactのルーティングに画面遷移するには以下のように書く
  //Linkでいう「to="/hoge"」のhogeの箇所を以下のpush内に記述する
  const history = useHistory();
  const onClickTrial500 = () => history.push("/trial/500");

  return (
   
    <>
        <div>
        <Link to="/trial/1">users</Link>
        <br />
        <Link to="/trial/2">user3</Link>
        <br />
        <Link to="/trial/500">URL Parameter</Link>
        <br />
        <Link to="/trial/500?name=hoge">Query Parameter</Link>
        <br />
        {/* 無意味なarrという配列（state）を渡してページ遷移 */}
        <Link to={{pathname: "/trial/500", state: arr}}>URL Parameter Page with State(無意味な配列)</Link>
        <br />
        {/* Linkを用いずに素のjsの遷移で「"/trial/500"」というリンクに遷移するには以下のように書く */}
        <button onClick={onClickTrial500}>trial/500</button>
        <form onSubmit={handleSubmit}>
            <div>
                <label htmlFor="user-name">UserName:</label>
                <input
                type="text"
                id="user-name"
                value={username}
                onChange={handleUserNameChange}
                />
            </div>
            <div>
                <label htmlFor="email">Email:</label>
                <input
                type="email"
                id="email"
                value={email}
                //onChangeの値（イベントハンドラ）に型をつけるには以下のように記述するといい
                onChange={(e: React.ChangeEvent<HTMLInputElement>) => handleEmailChange(e)}
                //onChange={handleEmailChange}
                />
            </div>
            <div>
                <label htmlFor="password">Password:</label>
                <input
                type="password"
                id="password"
                value={password}
                onChange={handlePasswordChange}
                />
            </div>
            <button type="submit">テストサインイン</button>
        </form>
        
        </div>
        <div>
            {username}
        </div>
    </>
  );
};

