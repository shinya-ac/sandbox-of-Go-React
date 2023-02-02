import { useMessageList } from "../hooks/use-message-list";
import axios from "axios";
import React, { useState } from 'react';

export const TrialReact = () => {
    const onClickUsers = () =>{
        //axios.httpメソッド名("リクエストを送るエンドポイント").then((リクエストが返却された後の返り値、つまり「result」) => {返却された後の処理})
        axios.get("https://jsonplaceholder.typicode.com/users").then((result: any) => {
            console.log(result)
        }).catch((err) => {console.log(err)});;
        alert("users");
    }

    const onClickUser1 = () =>{
        axios.get("https://jsonplaceholder.typicode.com/users/3").then((result: any) => {
            console.log(result)
        }).catch((err) => {console.log(err)});;
        alert("user1");
    }

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

  const onClickSignUp = (username: string, email: string, password: string) => {
    axios.post("http://localhost:8080/signup", {username, email, password})
    .then((result: any) => {
        console.log(result.data)
        console.log(result.data.username)
        setUserName(result.data.username);
        console.log(result.data)
    })
    .catch((err) => {console.log(err)});
};
  const handleSubmit = (e: any) =>{
    e.preventDefault();
    onClickSignUp(username, email, password)
  };

  return (
    <>
        <div>
        <button onClick={onClickUsers}>users</button>
        <button onClick={onClickUser1}>user1</button>
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

