import { memo, ReactNode, VFC } from "react"
import { Box, Button, IconButton, Stack, Image, Text} from "@chakra-ui/react"
import { HamburgerIcon } from "@chakra-ui/icons"

type Props = {
    id: number;
    answer_content: string;
    onClick: () => void;
}


export const AnswerCard: VFC<Props> = memo((props) => {
    const { id, answer_content, onClick} = props;
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
                <Text fontSize="lg" fontWeight="bold">{answer_content} </Text>
            </Box>
        </>
    )
});