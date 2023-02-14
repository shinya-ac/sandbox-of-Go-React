import axios from "axios";
import { useCallback, useState } from "react";
import { useHistory } from "react-router-dom";
import { User } from "../types/api/user";
import { useMessage } from "./use-message";
import { useLoginUser } from "../hooks/useLoginUser";


export const useAuth = () => {
    const history = useHistory();

    const [loading,  setLoading] = useState(false);
    const { showMessage } = useMessage();
    const { setLoginUser } = useLoginUser();

    const login = useCallback((id: string) => {
        setLoading(true);

        axios.get<User>(`https://jsonplaceholder.typicode.com/users/${id}`)
        .then((res) => {
            if (res.data){
                // 仮想的に管理者はidが10のユーザーとする
                const isAdmin = res.data.id === 10 ? true : false
                setLoginUser({...res.data, isAdmin})
                showMessage({title:"ログインしました", status:"success"});
                history.push("/home");
            } else {
                alert("ユーザーが見つかりません");
                setLoading(false);
            }
        })
        .catch(() => {
            alert("ログインできません")
            setLoading(false)
        })
        
    }, [history, showMessage, setLoginUser]);
    return { login, loading }
}