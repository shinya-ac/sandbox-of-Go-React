import axios from "axios";

export const TrialA = () => {
    const onClickUsers = () =>{
        //axios.httpメソッド名("リクエストを送るエンドポイント").then((リクエストが返却された後の返り値、つまり「result」) => {返却された後の処理})
        axios.get("https://jsonplaceholder.typicode.com/users").then((result: any) => {
            console.log(result)
        }).catch((err) => {console.log(err)});;
        alert("users");
    }

    return  (
        <>
            <button onClick={onClickUsers}>users</button>
        </>
    );
}