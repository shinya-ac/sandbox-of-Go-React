import { TrialReact } from "../components/TrialReact"
import { TrialA } from "../components/Trial1";
import { TrialB } from "../components/Trial2";
import { TrialParams } from "../components/TrialParams";

export const TrialRoutes = [
    {
        path:"/",
        exact:true,
        children: < TrialReact />
    },
    {
        path:"/1",
        exact:false,
        children: < TrialA />
    },
    {
        path:"/2",
        exact:false,
        children: < TrialB />
    },
    {
        path:"/:id",
        exact:false,
        children:< TrialParams/>
    }
]