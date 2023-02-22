import axios from "axios";
import { useCallback, useState } from "react"
import { Folder } from "../types/api/folder";
import { useMessage } from "./use-message";

export const useAllFolders = () => {
    const [loading, setLoading] = useState<boolean>(false);
    const [folders, setFolders] = useState<Array<Folder>>([]);
    const { showMessage } = useMessage(); 

    const getFolders = useCallback(() => {
        setLoading(true);
        axios
        .get<Array<Folder>>("https://jsonplaceholder.typicode.com/todos")
        .then(res => setFolders(res.data))
        .catch(() => {
            showMessage({title: "フォルダー取得に失敗しました", status: "error"})
        })
        .finally(() => {
            setLoading(false);
        })
    }, []);
    
    return{getFolders, folders, loading}
}