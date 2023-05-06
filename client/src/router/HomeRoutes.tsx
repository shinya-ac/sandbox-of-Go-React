import { CreateQA } from "../components/Pages/CreateQA";
import { ConfirmQA } from "../components/Pages/ConfirmQA"
import { Home } from "../components/Pages/Home";
import { Page404 } from "../components/Pages/Page404";
import { Setting } from "../components/Pages/Setting";
import { UserManagement } from "../components/Pages/UserManagement";

//Homeには「/home」「/home/hoge」などの複数ページを用意するのでこのファイルのように
//Homeのルーティング用のルーティング記述ファイルを作成してそこにルーティングを記載していく
export const HomeRoutes = [
        {path: "/",
        exact: true,
        children: <Home />},
        {path: "/create_qa/:folderId",
        exact: false,
        children: <CreateQA />},
        {
        path: "/confirm_qa/:folderId",
        exact: false,
        children: <ConfirmQA />
        },
        {path: "/user_management",
        exact: false,
        children: <UserManagement />},
        {path: "/setting",
        exact: false,
        children: <Setting />},
        {path: "*",
        exact: false,
        children: <Page404 />}
]