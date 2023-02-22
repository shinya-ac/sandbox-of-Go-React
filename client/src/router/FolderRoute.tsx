import { Page404 } from "../components/Pages/Page404";
import { QA } from "../components/Pages/organisms/QA/QA";

export const FolderRoutes = [
    {
        path:"/:folderId",
        exact:false,
        children:< QA/>
    },
    {path: "*",
        exact: false,
        children: <Page404 />
    }
]