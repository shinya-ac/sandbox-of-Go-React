import axios from "axios";
import { useCallback } from "react";
import { useHistory } from "react-router-dom";
import { useMessage } from "./use-message";
import { Question } from "../types/api/question";
import { Answer } from "../types/api/answer";


export const useConfirmQA = () => {
    const history = useHistory();
    const { showMessage } = useMessage();
    const ConfirmQA = useCallback(async (stateFromCreateQAPage: {
        resQuestions: Question[];
        resAnswers: Answer[];
        folderId: number;
    }, checkedItems: number[])=>{
        // checkedItemsにはチェックされたQAのインデックスが格納されている
        // selectedQAsにそのチェックされたQAのペアだけを格納し、postでサーバーにそれらを送信している
        const selectedQAs = checkedItems.map((index) => ({
            question: stateFromCreateQAPage.resQuestions[index],
            answer: stateFromCreateQAPage.resAnswers[index],
          }));
        const config = {
            headers: {
                'FolderId': String(stateFromCreateQAPage.folderId)
            },
            withCredentials: true
        };
        try {
            const response = await axios.post('http://localhost:8080/registerQA', { selectedQAs }, config);
            console.log(response);
            showMessage({title:"QAを作成しました", status:"success"});
            history.push(`/folders/${stateFromCreateQAPage.folderId}`);
        } catch (error) {
            console.error(error);
        }
    }, [history, showMessage]);
    return { ConfirmQA }
}