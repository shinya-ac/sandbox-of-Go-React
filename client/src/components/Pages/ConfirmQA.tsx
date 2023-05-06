import { FormControl, FormLabel, Input, Stack, Checkbox, CheckboxGroup, Textarea, Flex, Box } from "@chakra-ui/react";
import axios from "axios";
import { ChangeEvent, memo, useEffect, useState } from "react"
import { useHistory, useLocation } from "react-router-dom";
import { Answer } from "../../types/api/answer";
import { Question } from "../../types/api/question";
import { PrimaryButton } from "../atoms/buttons/PrimaryButton";
import { useConfirmQA } from "../../hooks/useConfirmQA";


// TODO：レンダリングの最適化


export const ConfirmQA  = memo(() => {
    const { ConfirmQA } = useConfirmQA();
    const location = useLocation<{resQuestions: Question[], resAnswers: Answer[], folderId: number}>();
    const stateFromCreateQAPage = location.state;
    const [checkedItems, setCheckedItems] = useState<number[]>([]);
    const [resQuestions, setResQuestions] = useState<Question[]>(stateFromCreateQAPage.resQuestions);
    const [resAnswers, setResAnswers] = useState<Answer[]>(stateFromCreateQAPage.resAnswers);

    useEffect(() => {
        const allIndexes = stateFromCreateQAPage.resQuestions.map((_, index) => index);
        setCheckedItems(allIndexes);
      }, [stateFromCreateQAPage.resQuestions]);

    const handleCheckboxChange = (index: number, isChecked: boolean) => {
        setCheckedItems((prevCheckedItems) =>
          isChecked
            ? [...prevCheckedItems, index]
            : prevCheckedItems.filter((item) => item !== index),
        );
      };
    console.log(stateFromCreateQAPage.resQuestions);

    const handleInputChange = (
        index: number,
        type: "question" | "answer",
        newValue: string,
      ) => {
        if (type === "question") {
          setResQuestions((prevQuestions) => {
            const updatedQuestions = [...prevQuestions];
            updatedQuestions[index].Content = newValue;
            return updatedQuestions;
          });
        } else {
          setResAnswers((prevAnswers) => {
            const updatedAnswers = [...prevAnswers];
            updatedAnswers[index].Content = newValue;
            return updatedAnswers;
          });
        }
      };

    const onClickPostQAs = async () => {
        ConfirmQA(stateFromCreateQAPage, checkedItems)
    };

    return(<>
    <h1 style={{ fontSize: "2rem" }}>単語帳に追加する一問一答を選択してください</h1>
    <br />
    {resQuestions.map((question, index) => {
        const answer = resAnswers[index];
        return <>
        
        <Flex as="nav" color="black" align="center" justify="center" padding={{ base:3, md: 5 }}>
            <Stack spacing={4} direction={["column", "row"]}>
            <FormControl>
            <Box
                bg="white"
                w="md"
                p={4}
                borderRadius="md"
                shadow="md"
                _hover={{cursor: "pointer", opacity:0.8}}
            >
                <FormLabel display="flex" alignItems="center">
                <Checkbox
                    value="qAndA"
                    isChecked={checkedItems.includes(index)}
                    onChange={(e) => handleCheckboxChange(index, e.target.checked)}
                    >
                        生成された一問一答No.{index + 1}
                    </Checkbox>
                </FormLabel>
                <CheckboxGroup colorScheme="green" defaultValue={["qAndA"]}>
                    <br />
                    問題:
                    <Textarea
                    value={question.Content}
                    onChange={(e) => handleInputChange(index, "question", e.target.value)}
                    width="100%"
                    />
                    <br />
                    解答:
                    <Textarea
                    value={answer.Content}
                    onChange={(e) => handleInputChange(index, "answer", e.target.value)}
                    width="100%"
                    />
                </CheckboxGroup>
            </Box>
            </FormControl>
            </Stack>
        </Flex>
        
        </>
    })}
    <PrimaryButton onClick={onClickPostQAs}>選択した一問一答を追加する</PrimaryButton>
        
        </>);
});