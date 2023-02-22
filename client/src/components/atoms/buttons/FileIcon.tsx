/** @jsxRuntime classic */
/** @jsx jsx */
import { memo, ReactNode, VFC } from "react"
import { background, Box, Flex, HStack, IconButton, Stack, Text} from "@chakra-ui/react"
import { jsx, css } from "@emotion/react"
import { Link } from "react-router-dom"

type Props = {
    title: string;
    folderId: number;
}

const containerStyle = css`
    font-size: 100px;
    width: 2.6em;
    height: 1.7em;
`;

    export const FolderIcon: VFC<Props> = memo((props) => {
        const { folderId, title } = props;
    
        return (
                <div css={containerStyle}>
                    <Box
                        display="flex"
                        as='button'
                        fontSize= '100px'
                        position= 'relative'
                        width= '1.0em'
                        height= '0.6em'
                        backgroundColor= '#39a9d6'
                        borderRadius= '0.1em 0.1em 0 0'
                        px='8px'
                        _hover={{cursor: "pointer", opacity:0.8}}
                        _before={{
                            content: '""',
                            position: 'absolute',
                            top: '0.3em',
                            left: '0',
                            width: '1.8em',
                            height: '1.3em',
                            backgroundColor: '#39a9d6',
                            borderRadius: 
                            '0 0.1em 0.1em 0.1em',
                        }}
                        >
                         <Link to={`/folders/${folderId}`}>
                            <Stack textAlign="center" >
                                <Text noOfLines={1} fontSize="md" color="red">{title}</Text>
                            </Stack>
                        </Link>
                </Box>
                </div>
                
            
        )
    });