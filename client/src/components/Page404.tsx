import { Link } from "react-router-dom"
//"react-router-dom"はインポートエラーが出てるけど一旦無視でおけ
export const Page404 = () =>{

    return(
        <div>
            <h1>404ページ</h1>
            <br/>
            <Link to="/Trial">TOPに戻る</Link>
        </div>
    )
}