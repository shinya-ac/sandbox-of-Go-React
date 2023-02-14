import { Box, Button, Divider, Flex, Heading, Input, Stack } from "@chakra-ui/react";
import { ChangeEvent, memo, useState } from "react"
import { useAuth } from "../../hooks/useAuth";
import { PrimaryButton } from "../atoms/buttons/PrimaryButton";


export const Login  = memo(() => {
    const {login, loading } = useAuth();
    const [userId, setUserId] = useState<string>('');
    const onChangeUserId = (e: ChangeEvent<HTMLInputElement>) => setUserId(e.target.value)

    const onClickLogin =  () => {login(userId)}
    return (
        <Flex align="center" justify="center" height="100vh">
            <Box bg="white" w="sm" p={4} borderRadius="md" shadow="md">
                <Heading as="h1" size="lg" textAlign="center">ユーザー管理アプリ</Heading>
                <Divider my={4} py={4} px={10} />
                <Stack spacing={6}>
                    <Input placeholder="ユーザーID" value={userId} onChange={onChangeUserId}/>
                    <PrimaryButton disabled={userId === ""} loading={loading} onClick={onClickLogin} >ログイン</PrimaryButton>
                </Stack>
            </Box>
        </Flex>
    )
});