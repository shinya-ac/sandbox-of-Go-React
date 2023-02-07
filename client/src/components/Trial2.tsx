import axios from "axios";

export const TrialB = () => {
    const onClickUser1 = () =>{
        axios.get("https://jsonplaceholder.typicode.com/users/3").then((result: any) => {
            console.log(result)
        }).catch((err) => {console.log(err)});;
        alert("user1");
    }

    return  (
        <>
            <button onClick={onClickUser1}>user3</button>
        </>
    );
}