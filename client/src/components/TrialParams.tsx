import { useParams, useLocation, useHistory } from "react-router-dom"
import React, { useContext } from "react"
import { UserContext } from "../providers/UserProvider";

export const TrialParams = () =>{
    //TrialRoutes.tsxで「/:id」という名前でurlパラメーターを受け取っているので
    //useParamsではidを受け取るような記述を書く
    const { id } = useParams();
    const location = useLocation();
    console.log(location)
    //useLocationはQueryパラメーター（URLの「?hogehoge」の部分）を扱えるもの
    //useLocationで取れるデータの中身は以下のconsole.logから確認できる（parameterとかstateとかSearchを確認できる）
    //searchの中にqueryパラメーターの値が入っている

    //ちなみにこのページへの遷移時にstateを渡して遷移した場合、
    //このlocationの中身のstateの中にそのstateが入っている

    //serchを直接受け取りたい場合は分割代入で以下のように受け取る
    const { search } = useLocation();
    console.log(search)

    //「<Link to={{pathname: "/trial/500", state: arr}}>URL Parameter Page with State(無意味な配列)</Link>」
    //のリンクから遷移した場合は上記のようにarrというstateも同時に渡されているので
    //以下のようにそのstate(arr)を扱うこともできる
    //例：APIから情報（ユーザー一覧など）を取得して、その一覧から詳細画面に飛ぶときなどに
    //このstateも渡すという方法で詳細画面に遷移してその配列データをもとに詳細ページを作ることもできる
    const { state } = useLocation();
    console.log(state)


    //そして上記で取得したQueryパラメーターをよしなに扱えるようにJSの以下のメソッドを利用する
    const query = new URLSearchParams(search);

    //useHistoryのgoBackを使えばブラウザの戻るボタンも実現できる
    const history = useHistory();
    const onClickBack = () => {history.goBack();}

    //以下はグローバルなstateを参照しているコード
    //下準備としてまずproviders配下にUserProvider.tsxを作成し、
    //その中でUserContextというユーザーの情報をグローバルに扱うコンテキストを定義した
    //そしてそれをApp.tsxの中で「全範囲で」そのコンテキスト（グローバルに参照できるstate）を
    //扱えるように設定した。その上で以下のように各コンポーネント内で参照することができる
    const context = useContext(UserContext);
    console.log(context);

    //以下のようにquery.get("name")とすると「?name=hogeter」のようにnameに一致する
    //Queryパラメーターを参照することができる
    return(
        <div>
            <h1>TrialParamsのページ</h1>
            <br/>
            <p>パラメーターは { id } です</p>
            <br/>
            <p>Queryパラメーターは { query.get("name") } です</p>
            <br/>
            <button onClick={onClickBack}>戻る</button>
        </div>
    )
}