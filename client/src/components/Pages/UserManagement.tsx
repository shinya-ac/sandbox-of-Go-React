import { Wrap, WrapItem, Image, Text, Spinner, Center, Modal, ModalOverlay, ModalContent, useDisclosure, ModalHeader, ModalCloseButton, ModalBody, FormControl, FormLabel, Input } from "@chakra-ui/react";
import { memo, useCallback, useEffect } from "react"
import { useAllUsers } from "../../hooks/useAllUsers";
import { useLoginUser } from "../../hooks/useLoginUser";
import { useSelectUser } from "../../hooks/useSelectUser";
import { UserCard } from "./organisms/user/UserCard";
import { UserDetailModal } from "./organisms/user/UserDetailModal";


export const UserManagement  = memo(() => {
    const { loginUser } = useLoginUser();
    const {isOpen, onOpen, onClose } = useDisclosure();
    const { getUsers, loading, users } = useAllUsers();
    const { onSelectUser, selectedUser } = useSelectUser();
    // useLoginUser（の中のloginUser）はグローバルなステート（useContext）
    // 定義もとはhooks内のuseLoginUser.tsに定義している（hook化することで自在にuseContextを呼び出して読み取ったり書き込んだりできるようにしている）
    // このloginUserに値を入れているのはログイン成功後の処理、つまりuseAuth内にある

    // useEffectの第二引数にから配列を渡すことで初期マウント時（最初ページが開かれたとき）の一回だけ
    // この処理を行えるという設定をできる
    useEffect(() => getUsers(), [getUsers])// このgetUsersでaxiosをしてUsersの一覧を取ってきている
    // propsとして渡していく関数（今回で言うとonClickUser）は毎回再作成されるとレンダリング
    // の効率が悪いのでuseCallbackで加工必要がある
    const onClickUser = useCallback((id: number) => {
        onSelectUser({id, users, onOpen});
    }, [users, onSelectUser, onOpen])
    // usersが変更されるたびにonSelectUserに渡す引数は設定し直してあげる必要があるので
    //（idは最初からonClickUserに渡しているので問題ないけどusersは更新する必要があるので）
    // 依存配列(useCallbackの第二引数、つまり[users, ...])と設定してあげる必要がある
    // もう少し丁寧に解説すると、もし上記のように設定しなかったら最初の一回だけgetUsersが実行される
    // その時タイミング（axiosに成功したタイミング）でusersが変更される（逆に言うとそれまではusersはnull）
    // けどこのuseCallbackの第二引数の配列に[users]を設定していないとusersの値が更新されても
    // onClickUserは最初に生成されタイミングの関数（つまりusersがnullの状態でのonClickUser）を
    // 利用してしまうのでユーザーをクリックしてもusersが空のようになってしまう
    // しかし第二引数の配列に[users]と、渡してやることでusersの値が変更されるたびに
    // onClickUserはそのusersの中身に応じたonClickUserの挙動を行うようになるので意図した動きになる
    // Udemyでも「なぜ、onClickUserの第二引数を指定しないとusersは空のままなのか？」と言うタイトルで質問がなされている
    return (
        <>
            {loading ? (
            <Center h="100vh">
                <Spinner />
            </Center>
            ) : (
            <Wrap p={{ base: 4, md:10 }}>
                
                    {users.map((user) => (
                        <WrapItem key={user.id} mx="auto">
                            <UserCard 
                            id={user.id}
                            imageUrl="https://source.unsplash.com/random" 
                            userName={user.username}
                            fullName={user.name}
                            onClick={onClickUser}
                            />
                        </WrapItem>
                    ))}
                
            </Wrap>
            )}
            <UserDetailModal user={selectedUser} isOpen={isOpen} onClose={onClose} isAdmin={loginUser?.isAdmin} />
        </>
    )
});