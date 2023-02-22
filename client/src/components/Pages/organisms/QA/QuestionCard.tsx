import { Dispatch, memo, ReactNode, SetStateAction, VFC } from "react"
import { Box, Button, IconButton, Stack, Image, Text} from "@chakra-ui/react"
import { HamburgerIcon } from "@chakra-ui/icons"

type Props = {
    id: number;
    question_content: string;
    onClick: () => void;
}


export const QuestionCard: VFC<Props> = memo((props) => {
    const { id, question_content, onClick} = props;
    return (
        <>
            <Box
            onClick={onClick}
            cursor="pointer"
            borderWidth="1px"
            borderRadius="lg"
            overflow="hidden"
            p={3}
            bg="white"
            >
                <Text fontSize="lg" fontWeight="bold">{question_content} </Text>
            </Box>
        </>
    )
});