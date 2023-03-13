import { Box, Button, Divider, Flex, Heading, Input, Link, Stack, Textarea } from "@chakra-ui/react";
import { ChangeEvent, memo, useState } from "react"
import { useAuth } from "../../hooks/useAuth";
import { useCreateQA } from "../../hooks/useCreateQA";
import { PrimaryButton } from "../atoms/buttons/PrimaryButton";
import { WebCam } from "../WebCam";
import TextareaAutosize from 'react-textarea-autosize';


export const CreateQA  = memo(() => {
    const {createQA, loading } = useCreateQA();

    const [content, setContent] = useState<string>('');
    const onChangeContent =  (e: ChangeEvent<HTMLTextAreaElement>) => {setContent(e.target.value)}

    const onClickCreateQA = () => {createQA(content)}

    return (
        <>
            <Flex align="center" justify="center" height="100vh">
                <Box bg="white" w="sm" p={4} borderRadius="md" shadow="md">
                    <Heading as="h1" size="lg" textAlign="center">単語帳作成画面</Heading>
                    <Divider my={4} py={4} px={10} />
                    <Stack spacing={6}>
                        <TextareaAutosize
                            placeholder="この文章から一問一答を作成する"
                            value={content}
                            onChange={onChangeContent}
                            minRows={5}
                            maxRows={120}
                            style={{ overflow: 'hidden', resize:"none" }}
                            />
                        <PrimaryButton disabled={content === ""} loading={loading} onClick={onClickCreateQA} >一問一答を作成する</PrimaryButton>
                    </Stack>
                </Box>
            </Flex>

            <Flex align="center" justify="center" height="10vh">
                <Box pr={4}>
                    <WebCam setContent={setContent} content={content} />
                </Box>
            </Flex>
        </>
    )
});