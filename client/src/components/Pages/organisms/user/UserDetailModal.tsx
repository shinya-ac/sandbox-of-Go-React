import { ChangeEvent, memo, ReactNode, useEffect, useState, VFC } from "react"
import { Box, Button, IconButton, Stack, Image, Text, Modal, ModalOverlay, ModalHeader, ModalContent, ModalCloseButton, ModalBody, FormControl, FormLabel, Input, ModalFooter} from "@chakra-ui/react"
import { HamburgerIcon } from "@chakra-ui/icons"
import { User } from "../../../../types/api/user";
import { useLoginUser } from "../../../../hooks/useLoginUser";
import { PrimaryButton } from "../../../atoms/buttons/PrimaryButton";


type Props = {
    user: User | null;
    isOpen: boolean;
    onClose: () => void;
    isAdmin?: boolean;
    
}


export const UserDetailModal: VFC<Props> = memo((props) => {
    const { user, isOpen, onClose, isAdmin = false } = props;
    const [username, setUsername] = useState("");
    const [name, setName] = useState("");
    const [email, setEmail] = useState("");
    const [phone, setPhone] = useState("");

    useEffect(() => {
        setUsername(user?.username ?? '')
        setName(user?.name ?? '')
        setEmail(user?.email ?? '')
        setPhone(user?.phone ?? '')
    }, [user])

    const onChangeUserName = (e: ChangeEvent<HTMLInputElement>) => {
        setUsername(e.target.value)
    }
    const onChangeName = (e: ChangeEvent<HTMLInputElement>) => {
        setName(e.target.value)
    }
    const onChangeEmail = (e: ChangeEvent<HTMLInputElement>) => {
        setEmail(e.target.value)
    }
    const onChangePhone = (e: ChangeEvent<HTMLInputElement>) => {
        setPhone(e.target.value)
    }

    const onClickUpdate = () => console.log("更新")
    
    return (
        <>
            <Modal isOpen={isOpen} onClose={onClose} autoFocus={false} motionPreset="slideInBottom">
                <ModalOverlay />
                <ModalContent pb={2}>
                    <ModalHeader>ユーザー詳細</ModalHeader>
                    <ModalCloseButton />
                    <ModalBody mx={4}>
                        <Stack spacing={4} >
                            <FormControl>
                                <FormLabel>名前</FormLabel>
                                <Input value={username} onChange={onChangeUserName} isReadOnly={!isAdmin} />
                            </FormControl>
                            <FormControl>
                                <FormLabel>フルネーム</FormLabel>
                                <Input value={name} onChange={onChangeName} isReadOnly={!isAdmin} />
                            </FormControl>
                            <FormControl>
                                <FormLabel>メール</FormLabel>
                                <Input value={email} onChange={onChangeEmail} isReadOnly={!isAdmin} />
                            </FormControl>
                            <FormControl>
                                <FormLabel>電話番号</FormLabel>
                                <Input value={phone} onChange={onChangePhone} isReadOnly={!isAdmin} />
                            </FormControl>
                        </Stack>
                    </ModalBody>
                    {isAdmin && (
                        <ModalFooter>
                        <PrimaryButton onClick={onClickUpdate}>更新</PrimaryButton>
                        </ModalFooter>
                    )}
                </ModalContent>
            </Modal>
        </>
    )
});