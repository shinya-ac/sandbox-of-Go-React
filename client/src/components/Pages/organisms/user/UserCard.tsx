import { memo, ReactNode, VFC } from "react"
import { Box, Button, IconButton, Stack, Image, Text} from "@chakra-ui/react"
import { HamburgerIcon } from "@chakra-ui/icons"

type Props = {
    id: number;
    imageUrl: string;
    userName: string;
    fullName:string;
    onClick: (id: number) => void;
}


export const UserCard: VFC<Props> = memo((props) => {
    const { id, imageUrl, userName, fullName , onClick} = props;
    return (
        <>
            <Box 
                p={4} 
                w="260px" 
                height="260px" 
                bg="white" 
                borderRadius="10px" 
                shadow="md"
                _hover={{cursor: "pointer", opacity:0.8}}
                onClick={() => onClick(id)}
                >
                <Stack textAlign="center">
                    <Image 
                    boxSize="160px" 
                    borderRadius="full" 
                    alt={userName}
                    m="auto"  
                    src={imageUrl}
                    />
                    <Text fontSize="lg" fontWeight="bold">{userName}</Text>
                    <Text fontSize="lg" color="gray"> {fullName}</Text>
                </Stack>
            </Box>
        </>
    )
});