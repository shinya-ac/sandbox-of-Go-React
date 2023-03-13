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

    const [questionContent, setQuestionContent] = useState<string>('');

    interface folderParams {
        Id: string;
      }
    //TrialRoutes.tsxで「/:id」という名前でurlパラメーターを受け取っているので
    //useParamsではidを受け取るような記述を書く
    const { Id } = useParams<folderParams>();
    const location = useLocation();
    console.log(location)
    console.log(Id)

    const createQA = useCallback((Content: string) => {
        setLoading(true);
        console.log(`フォルダーID：${Id}`)
        const FolderId = parseInt(Id, 10)
        const config = {
            headers: {
                'FolderId': String(FolderId)
            },
            withCredentials: true
        };

        axios.post<QA>("http://localhost:8080/question", { Content }, config)
        .then((res) => {
            if (res.data){
                setQuestionContent(res.data.question_content)
                showMessage({title:"QAを作成しました", status:"success"});
                history.push(`/folders/${FolderId}`);
            } else {
                alert("QAを作成できませんでした");
                setLoading(false);
            }
        })
        .catch(() => {
            alert("QAを作成できませんでした")
            setLoading(false)
        })
        
    }, [history, showMessage, setQuestionContent]);
    return { createQA, loading }
}