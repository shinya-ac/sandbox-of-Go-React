import { Box, Center, Flex, Spacer, Spinner, Wrap, WrapItem } from "@chakra-ui/react";
import { memo, useEffect, useState } from "react"
import { useParams } from "react-router-dom";
import { useAllQAs } from "../../../../hooks/useAllQAs";
import { AnswerCard } from "./AnswerCard";
import { QuestionCard } from "./QuestionCard";
import ReactPaginate from 'react-paginate';
import { css } from "@emotion/react";
import { PaginationContainer } from "../../../molecules/Paginate";

const PAGE_SIZE = 1; // 1ページあたりのアイテム数

const pagingStyle = css`
    display: flex;
    align-items: center;
    justify-content: center;
    margin-bottom: 10px;
    gap: 20px 6px;

    display: inline-flex;
    align-items: center;
    border-radius: 30px;
    justify-content: center;
    font-weight: 700;
    font-size: 16px;
    height: 40px;
    width: 40px;
`;

export const QA  = memo(() => {
    // parameterとして受け渡しできるのはstring型のみ
    interface FolderParams {
        folderId: string;
    }
    const { folderId } = useParams<FolderParams>();
    const numFolderId = Number(folderId);// FolderIdはparamsとしてはstringだけど数値として扱いたいので型変換
    const { getQAs, loading, qas } = useAllQAs(numFolderId);
    useEffect(() => getQAs(), [getQAs])

    const [currentPage, setCurrentPage] = useState(0);
    const itemsOnPage = qas.slice(currentPage * PAGE_SIZE, (currentPage + 1) * PAGE_SIZE);
    const pageCount = Math.ceil(qas.length / PAGE_SIZE);

    const [showAnswer, setShowAnswer] = useState(false);
    const toggleAnswer = () => setShowAnswer(!showAnswer);
    const toggleToQuestion = () => setShowAnswer(false);
    
    return(
        <>
            <p>QAページです</p>
            {loading ? (
                <Center h="100vh">
                    <Spinner />
                </Center>
                ) : (
                <>
                <Wrap p={{ base: 4, md:10 }} spacing='30px' justify="center" align='left'>
                    {itemsOnPage.map((qa) => (
                        <WrapItem key={`${qa.question_content}-${qa.answer_content}`} mx="auto">
                            {showAnswer
                                ? <AnswerCard id={qa.aid} answer_content={qa.answer_content} onClick={toggleAnswer} />
                                : <QuestionCard id={qa.qid} question_content={qa.question_content} onClick={toggleAnswer} />}
                            {/* <QuestionCard id={qa.qid} question_content={qa.question_content}></QuestionCard>
                            <AnswerCard id={qa.aid} answer_content={qa.answer_content}></AnswerCard> */}
                        </WrapItem>
                    ))}
                </Wrap>
                <PaginationContainer>
                    <ReactPaginate
                    onClick={toggleToQuestion}
                    previousLabel="<<<前の単語"
                    nextLabel="次の単語>>>"
                    breakLabel="..."
                    pageCount={pageCount}
                    marginPagesDisplayed={2}
                    pageRangeDisplayed={5}
                    onPageChange={(data) => setCurrentPage(data.selected)}
                    containerClassName="pagination"
                    activeClassName="active"
                    pageClassName="page-item"
                    previousClassName="page-item"
                    nextClassName="page-item"
                    breakClassName="page-item"
                    pageLinkClassName="page-link"
                    previousLinkClassName="page-link"
                    nextLinkClassName="page-link"
                    breakLinkClassName="page-link"
                    disabledClassName='disabled'
                    />
                </PaginationContainer>
              </>
            )}
            
        </>
    )
});

// import { memo, ReactNode, VFC } from "react"
// import { Box, Button, IconButton, Stack, Image, Text} from "@chakra-ui/react"
// import { HamburgerIcon } from "@chakra-ui/icons"

// type Props = {
//     id: number;
//     imageUrl: string;
//     userName: string;
//     fullName:string;
//     onClick: (id: number) => void;
// }

// export const QA: VFC<Props> = memo((props) => {
//     const { id, imageUrl, userName, fullName , onClick} = props;
//     return (
//         <>
//             <Box 
//                 p={4} 
//                 w="260px" 
//                 height="260px" 
//                 bg="white" 
//                 borderRadius="10px" 
//                 shadow="md"
//                 _hover={{cursor: "pointer", opacity:0.8}}
//                 onClick={() => onClick(id)}
//                 >
//                 <Stack textAlign="center">
//                     <Image 
//                     boxSize="160px" 
//                     borderRadius="full" 
//                     alt={userName}
//                     m="auto"  
//                     src={imageUrl}
//                     />
//                     <Text fontSize="lg" fontWeight="bold">{userName}</Text>
//                     <Text fontSize="lg" color="gray"> {fullName}</Text>
//                 </Stack>
//             </Box>
//         </>
//     )
// });