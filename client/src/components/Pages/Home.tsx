import { Box, Center, Flex, Spacer, Spinner, Wrap, WrapItem } from "@chakra-ui/react";
import { memo, useEffect } from "react"
import { FolderIcon } from "../atoms/buttons/FileIcon";
import { useAllFolders } from "../../hooks/useAllFolders";


export const Home  = memo(() => {

    const { getFolders, loading, folders } = useAllFolders();
    useEffect(() => getFolders(), [getFolders])
    return(
        <>
            <p>ホームページです</p>
            {loading ? (
            <Center h="100vh">
                <Spinner />
            </Center>
            ) : (
            <Wrap p={{ base: 4, md:10 }} spacing='30px' justify="center" align='left'>
                {folders.map((folder) => (
                    <WrapItem key={folder.id} mx="auto">
                        <FolderIcon folderId={folder.id} title={folder.title}></FolderIcon>
                    </WrapItem>
                ))}
            </Wrap>
  
            )}
            
        </>
    )
});