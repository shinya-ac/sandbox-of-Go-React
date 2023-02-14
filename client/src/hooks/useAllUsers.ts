import axios from "axios";
import { useCallback, useState } from "react"
import { User } from "../types/api/user";
import { useMessage } from "./use-message";

export const useAllUsers = () => {
    const [loading, setLoading] = useState<boolean>(false);
    const [users, setUsers] = useState<Array<User>>([]);
    const { showMessage } = useMessage(); 

    const getUsers = useCallback(() => {
        setLoading(true);
        axios
        .get<Array<User>>("https://jsonplaceholder.typicode.com/users")
        .then(res => setUsers(res.data))
        .catch(() => {
            showMessage({title: "ユーザー取得に失敗しました", status: "error"})
        })
        .finally(() => {
            setLoading(false);
        })
    }, []);
    
    return{getUsers, users, loading}
}