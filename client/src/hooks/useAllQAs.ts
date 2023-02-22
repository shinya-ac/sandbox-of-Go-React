import axios from "axios";
import { useCallback, useState } from "react"
import { QA } from "../types/api/qa";
import { useMessage } from "./use-message";

export const useAllQAs = (folderId: number) => {
    const [loading, setLoading] = useState<boolean>(false);
    const [qas, setQAs] = useState<Array<QA>>([]);
    const { showMessage } = useMessage(); 

    const getQAs = useCallback(() => {
        setLoading(true);
        axios.defaults.withCredentials = true;
        axios
        .get<Array<QA>>(`http://localhost:8080/folders/${folderId}`)
        //.then(res => setQAs(res.data))
        //.then(res => console.log(res.data))
        .then((res) => {
            const qasData = res.data.map((data: any) => {
              const qa: QA = {
                qid: data.Id,
                aid: -1, // 初期値として-1を設定
                question_content: data.Content,
                answer_content: "これは解答です（決めうち文字列）",
              };
              return qa;
            });
            setQAs(qasData);
          })
        .catch((error) => {
            showMessage({title: "QA取得に失敗しました", status: "error"})
            console.log(error)
        })
        .finally(() => {
            setLoading(false);
        })
    }, []);
    
    return{getQAs, qas, loading}
}