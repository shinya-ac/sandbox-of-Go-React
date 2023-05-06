import axios from "axios";
import { useCallback, useState } from "react";
import { useHistory, useLocation, useParams } from "react-router-dom";
import { User } from "../types/api/user";
import { useMessage } from "./use-message";
import { useLoginUser } from "../hooks/useLoginUser";
import { QA } from "../types/api/qa";


export const useCreateQA = () => {
    const history = useHistory();

    const [loading,  setLoading] = useState(false);
    const { showMessage } = useMessage();

    interface folderParams {
        folderId: string;
      }
    //TrialRoutes.tsxで「/:folderParams」という名前でurlパラメーターを受け取っているので
    //useParamsではfolderParamsを受け取るような記述を書く
    const { folderId } = useParams<folderParams>();
    const location = useLocation();
    console.log(location)
    console.log(folderId)

    const createQA = useCallback((Content: string) => {
        setLoading(true);
        console.log(`フォルダーID：${folderId}`)
        const FolderId = parseInt(folderId, 10)
        const config = {
            headers: {
                'FolderId': String(FolderId)
            },
            withCredentials: true
        };

        axios.post<any>("http://localhost:8080/question", { Content }, config)
        .then((res) => {
            if (res.data){
                console.log(res.data.questions[0])
                const resQuestions = res.data.questions;
                const resAnswers = res.data.answers;
                showMessage({title:"問題と解答が自動で生成されました", status:"success"});
                history.push({pathname: `/home/confirm_qa/${folderId}`,state: { resQuestions, resAnswers, folderId }});
            } else {
                alert("問題と解答を自動作成できませんでした");
                setLoading(false);
            }
        })
        .catch(() => {
            alert("問題と解答を自動作成できませんでした")
            setLoading(false)
        })
    }, [history, showMessage]);
    return { createQA, loading }
}